package instructions

import (
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func do(story *models.Story, params interface{}) chromedp.Action {
	
	isString, s, _ := utils.AssertTypeIsString(params)
	
	if !isString || *s != "refresh" {
		reporter.Write("Issue with the \"do\" instruction", reporter.INPUT_ERROR)
		os.Exit(1)
	}
	
	return chromedp.Reload()
}
