package instructions

import (
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func typeText(story *models.Story, params interface{}) chromedp.Action {

	isString, _, m := utils.AssertTypeIsString(params)
	if isString || (*m)["text"] == "" || (*m)["selector"] == ""{
		reporter.Write("You must specify the selector and the text in the \"type\" instruction", reporter.INPUT_ERROR)
		os.Exit(1)
	}

	return chromedp.SendKeys((*m)["selector"], (*m)["text"])
}
	
