package instructions

import (
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func waitUntilMissing(story *models.Story, params interface{}) chromedp.Action {
	isString, s, m := utils.AssertTypeIsString(params)
	if isString {
		return chromedp.WaitNotPresent(*s)
	}

	selector, err := getSelector(*m)
	if err != nil {
		reporter.Write(err.Error(), reporter.INVALID_SELECTOR_ERROR)
		os.Exit(1)
	}
	

	return chromedp.WaitNotPresent(*selector)
}
