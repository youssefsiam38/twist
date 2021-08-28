package models

import "github.com/youssefsiam38/twist/src/utils"

// import (
// 	"fmt"
// 	"io/ioutil"

// 	// "github.com/chromedp/chromedp"
// )

type Assertion struct {
	Instruction string
	Name        string
	Expected    interface{}
	Found       *string
}

func (a Assertion) GetExpected() string {
	isString, s, m := utils.AssertTypeIsString(a.Expected)

	if isString {
		return *s
	}
	return (*m)["expect"]
}