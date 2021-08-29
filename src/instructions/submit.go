package instructions

import (
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func submit(story *models.Story, params interface{}) chromedp.Action {

	isString, s, m := utils.AssertTypeIsString(params)
	if isString {
		return chromedp.Submit(*s)
	}

	if (*m)["selector"] == "" {
		reporter.Write("You must specify selector in the \"submit\" instruction", reporter.INPUT_ERROR)
		os.Exit(1)
	}

	return chromedp.Submit((*m)["selector"])
}
	
