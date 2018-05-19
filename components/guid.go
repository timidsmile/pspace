package components

import (
	"errors"
	"time"
)

// global uniq uuid
// the core algorithm here was borrowed from:
// Blake Mizerany's `noeqd` https://github.com/bmizerany/noeqd
// and indirectly:
// Twitter's `snowflake` https://github.com/twitter/snowflake

// only minor cleanup and changes to introduce a type, combine the concept
// of workerID + datacenterId into a single identifier, and modify the
// behavior when sequences rollover for our specific implementation needs

const (
	workerIdBits       = uint64(5)                                      // 节点id长度
	datacenterIdBits   = uint64(5)                                      // 数据中心ID长度
	maxWorkerId        = int64(-1) ^ (int64(-1) << workerIdBits)        // 最大支持机器节点数0~31，一共32个
	maxDatacenterId    = int64(-1) ^ (int64(-1) << datacenterIdBits)    // 最大支持数据中心节点数0~31，一共32个
	sequenceBits       = uint64(12)                                     // 序列号12位
	workerIdShift      = sequenceBits                                   // 机器节点左移12位
	datacenterIdShift  = sequenceBits + workerIdBits                    // 数据中心节点左移17位
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits // 时间毫秒数左移22位
	sequenceMask       = int64(-1) ^ (int64(-1) << sequenceBits)        // 最大为4095
	//   1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111
	// ^ 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 0000 0000 0000
	// = 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 1111 1111 1111
	// = 4095

	// Tue, 21 Mar 2006 20:50:14.000 GMT
	twepoch = int64(1288834974657)
)

var ErrTimeBackwards = errors.New("time has gone backwards")
var ErrSequenceExpired = errors.New("sequence expired")
var ErrIDBackwards = errors.New("ID went backward")
var ErrInputIdInvalid = errors.New("datacenterID or workerID invalid, pls check!")

type Guid struct {
	sequence      int64
	lastTimestamp int64
	lastID        int64
}

func (f *Guid) NewGUID(datacenterID int64, workerID int64) (int64, error) {
	if datacenterID > maxDatacenterId || datacenterID < 0 {
		return 0, ErrInputIdInvalid
	}

	if workerID > maxWorkerId || workerID < 0 {
		return 0, ErrInputIdInvalid
	}

	// divide by 1048576, giving pseudo-milliseconds
	ts := f.getMillTime()

	if ts < f.lastTimestamp {
		return 0, ErrTimeBackwards
	}

	// 如果上次生成时间和当前时间相同(在同一毫秒内)
	if f.lastTimestamp == ts {
		// sequence自增，因为sequence只有12bit，所以和sequenceMask相与一下，去掉高位
		f.sequence = (f.sequence + 1) & sequenceMask
		// 判断是否溢出,每毫秒内最多支持4095。当为4096时，与sequenceMask相与，sequence就等于0
		if f.sequence == 0 {
			//自旋等待到下一毫秒
			ts = f.waitUntilNextMill(f.lastTimestamp)
		}
	} else {
		// 如果和上次生成时间不同,重置sequence. 从下一毫秒开始，sequence计数重新从0开始累加
		f.sequence = 0
	}

	// 更新时间
	f.lastTimestamp = ts

	// 最后按照规则拼出ID。
	// 000000000000000000000000000000000000000000  00000            00000       000000000000
	// time                                      datacenterId      workerId     sequence
	id := int64(((ts - twepoch) << timestampLeftShift) |
		(datacenterID << datacenterIdShift) |
		(workerID << workerIdShift) |
		f.sequence)

	// 如果id < 上次生成的id,错误
	if id <= f.lastID {
		return 0, ErrIDBackwards
	}

	// 更新 lastID
	f.lastID = id

	return id, nil
}

func (f *Guid) getMillTime() int64 {
	// divide by 1048576, giving pseudo-milliseconds
	return time.Now().UnixNano() >> 20
}

func (f *Guid) waitUntilNextMill(lastTimeMill int64) int64 {
	ts := f.getMillTime()
	for ts <= lastTimeMill {
		ts = f.getMillTime()
	}
	return ts
}
