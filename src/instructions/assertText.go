package instructions

import (
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func assertText(story *models.Story, params interface{}) chromedp.Action {
	var expected, found, selector string
	name := "assertText"
	isString, _, m := utils.AssertTypeIsString(params)

	if isString || (*m)["expect"] == "" {
		reporter.Write("You must specify the selector and the expect paramaeters in the \"assertText\" instruction", reporter.INPUT_ERROR)
		os.Exit(1)
	} else {

		selector = (*m)["selector"]
		expected = (*m)["expect"]

		if (*m)["name"] != "" {
			name = (*m)["name"]
		}
	}

	assertion := models.Assertion{
		Instruction: "assertText",
		Name:        name,
		Expected:    expected,
		Found:       &found,
	}
	story.Assertions = append(story.Assertions, assertion)
	return chromedp.Text(selector, &found)
}
