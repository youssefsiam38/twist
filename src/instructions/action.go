package instructions

import (
	"errors"
	"fmt"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/src/models"
)

const (
	STRINGPARAM = iota
	OBJECTPARAM
)

var Action map[string]func(ctx *models.Story, params interface{}) chromedp.Action

func init() {
	Action = map[string]func(ctx *models.Story, params interface{}) chromedp.Action{
		"click":        	click,
		"doubleClick":		doubleClick,
		"rightClick":		rightClick,
		"waitFor":      	waitFor,
		"waitUntilMissing":	waitUntilMissing,
		"assertPathIs": 	assertPathIs,
		"assertText": 		assertText,
		"do":           	do,
		"navigate":     	navigate,
		"saveScreenshotTo": saveScreenshotTo,
		"type":     		typeText,
		"submit":     		submit,
	}
}

func getSelector(params map[string]string) (*string, error) {
	if params["selector"] != "" {
		selector := fmt.Sprintf("%v", params["selector"])
		return &selector, nil
	} else {
		return nil, errors.New("ERROR: You didn't provide \"selector\" parameter in the \"waitFor\" instruction")
	}
}
