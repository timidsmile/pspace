package components

import (
	"fmt"
	"regexp"
)

const (
	emailRegular    = "^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$"
	mobileRegular   = "^1[0-9]{10}$"
	userNameRegular = "^[a-zA-Z0-9_]{4,32}$"
)

type ValidateFactory interface {
	Validate()
}

type EmailValidate struct {
	input interface{}
}

type MobileValidate struct {
	input interface{}
}

// 校验用户名的合法性
type UserNameValidate struct {
	input interface{}
}

func (obj EmailValidate) Validate(email interface{}) bool {
	emailStr, ok := email.(string)
	result := false
	if ok == true {
		reg := regexp.MustCompile(emailRegular)
		result = reg.MatchString(emailStr)
	}
	return result
}

func (obj MobileValidate) Validate(mobile interface{}) bool {
	mobileStr, ok := mobile.(string)
	result := false
	if ok == true {
		reg := regexp.MustCompile(mobileRegular)
		result = reg.MatchString(mobileStr)
	}
	return result
}

// 由字母、数字、下划线组成，最少4位，最多32位
func (obj UserNameValidate) Validate(userName interface{}) bool {
	userNameStr, ok := userName.(string)
	result := false
	if ok == true {
		reg := regexp.MustCompile(userNameRegular)
		result = reg.MatchString(userNameStr)
	}

	fmt.Println("user name is not ok!")
	return result
}
