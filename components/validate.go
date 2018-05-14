package components

import (
	"regexp"
)

const (
	emailRegular = "^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$"
	mobileRegular = "^1[0-9]{10}$"
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

func (obj EmailValidate) Validate (email interface{}) (bool) {
	emailStr, ok := email.(string)
	result := false;
	if ok == true {
		reg := regexp.MustCompile(emailRegular)
		result = reg.MatchString(emailStr)
	}
	return result
}


func (obj MobileValidate) Validate (mobile interface{}) (bool) {
	mobileStr,ok := mobile.(string)
	result := false;
	if ok == true {
		reg := regexp.MustCompile(mobileRegular)
		result = reg.MatchString(mobileStr)
	}
	return result
}