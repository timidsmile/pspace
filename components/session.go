package components

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/timidsmile/pspace/consts"
)

const TTL_SESSION = 86400

type Session struct {
	UserID    int64 // user ID
	LoginTime int64 // 登陆时间
}

func (s *Session) Save() (string, error) {
	accessToken := s.genAccessToken()

	sessionVal := strconv.FormatInt(s.UserID, 10) + "-" + strconv.FormatInt(s.LoginTime, 10)

	fmt.Println("session key: ", accessToken)
	fmt.Println("session value: ", sessionVal)

	err := SetEx(accessToken, sessionVal, TTL_SESSION)
	if err != nil {
		fmt.Println("save token failed!")
		return "", err
	}

	return accessToken, nil
}

func (s *Session) Get(accessToken string) *Session {
	token, err := Get(accessToken)
	if err != nil {
		fmt.Println("get token failed!")
		return nil
	}

	arr := strings.Split(token, "-")

	if len(arr) != 2 {
		fmt.Println("invalid token!")
		return nil
	}

	userID, _ := strconv.ParseInt(arr[0], 10, 64)
	loginTime, _ := strconv.ParseInt(arr[1], 10, 64)

	userToken := &Session{
		UserID:    userID,
		LoginTime: loginTime,
	}

	return userToken
}

func (s *Session) genAccessToken() string {
	// accessToken 生成规则： userID_uniqID

	var str string

	guid := Guid{}
	uuid, _ := guid.NewGUID(consts.DataCenterID_session, consts.WorkID_session)

	str = strconv.FormatInt(uuid, 10)

	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}
