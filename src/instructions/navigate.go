package instructions

import (
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func navigate(story *models.Story, params interface{}) chromedp.Action {
	isString, s, m := utils.AssertTypeIsString(params)

	if isString {
		chromedpAction, err := handleNavigateArg(*s, STRINGPARAM)
		if err != nil {
			reporter.Write(err.Error(), reporter.INPUT_ERROR)
			os.Exit(1)
		}
		return chromedpAction
	}

	if (*m)["url"] == "" {
		reporter.Write("You must specify valid url as an argument to the url parameter in the \"navigate\" instruction", reporter.INPUT_ERROR)
		os.Exit(1)
	}

	chromedpAction, err := handleNavigateArg((*m)["url"], OBJECTPARAM)
	if err != nil {
		reporter.Write(err.Error(), reporter.INPUT_ERROR)
		os.Exit(1)
	}
	return chromedpAction
}

func handleNavigateArg(url string, paramType int) (chromedp.Action, error) {

	if utils.IsValidUrl(url) {
		return chromedp.Navigate(url), nil
	}

	if paramType == STRINGPARAM {
		switch url {
		case "back":
			return chromedp.NavigateBack(), nil
		case "forward":
			return chromedp.NavigateForward(), nil
		}

		return nil, fmt.Errorf("YOU MUST SPECIFY VALID URL AS A PARAMETER IN THE \"navigate\" INSTRUCTION\nYOU WROTE: %v", url)
	}

	return nil, fmt.Errorf("YOU MUST SPECIFY VALID URL AS A PARAMETER IN THE \"navigate\" INSTRUCTION OR WRITE IT AS STRING PARAMETER\nYOU WROTE: %v", url)
}
