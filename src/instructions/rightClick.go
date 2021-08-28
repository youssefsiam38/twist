package instructions

import (
	"context"
	"errors"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func rightClick(story *models.Story, params interface{}) chromedp.Action {
	var selector *string
	var err error
	isString, s, m := utils.AssertTypeIsString(params)
	if isString {
		selector = s
	} else {
		selector, err = getSelector(*m)
		if err != nil {
			reporter.Write(err.Error(), reporter.INVALID_SELECTOR_ERROR)
			os.Exit(1)
		}
	}


	return chromedp.QueryAfter(*selector, func(ctx context.Context,r runtime.ExecutionContextID, nodes ...*cdp.Node) error {
        if len(nodes) < 1 {
            return errors.New("expected at least one node")
        }
       return chromedp.MouseClickNode(nodes[0], chromedp.ButtonRight).Do(ctx)
    })

}
