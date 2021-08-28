package instructions

import (
	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func assertPathIs(story *models.Story, params interface{}) chromedp.Action {
	var expected, found string
	name := "assertPathIs"
	isString, s, m := utils.AssertTypeIsString(params)

	if isString {
		expected = *s
	} else {
		if (*m)["expect"] == "" {
			panic("you must provide expect")
		}
		expected = (*m)["expect"]

		if (*m)["name"] != "" {
			name = (*m)["name"]
		}
	}

	assertion := models.Assertion{
		Instruction: "assertPathIs",
		Name:        name,
		Expected:    expected,
		Found:       &found,
	}
	story.Assertions = append(story.Assertions, assertion)
	return chromedp.Location(&found)
}
