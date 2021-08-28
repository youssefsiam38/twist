package handle

import (
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	// "github.com/kr/pretty"
	stdContext "context"

	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/context"
	"github.com/youssefsiam38/twist/src/instructions"
	"github.com/youssefsiam38/twist/src/utils"
)

// i is the current story
var i int

func Handle() error {

	if !utils.FolderExist("./twist") {
		reporter.Write("There is no twist folder in this directory", reporter.FLOW_ERROR)
		os.Exit(1)
	}

	storiesCtx, err := context.NewStoriesContext()
	if err != nil {
		return err
	}

	// loop over the stories
	for i = 0; i < len(storiesCtx.Config.Order); i++ {
		actions := chromedp.Tasks{chromedp.Navigate(storiesCtx.Stories[i].Start)}
		story := storiesCtx.Stories[i]


		opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", story.Headless))
		ctx, _ := stdContext.WithTimeout(storiesCtx.Context, story.TimeoutDur)
		ctx, _ = chromedp.NewExecAllocator(ctx, opts...)
		ctx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(func(s string, i ...interface{}) {}))
		defer cancel()


		// loop over instructions
		for j := 0; j < len(story.Instructions); j++ {

			// get the key & value of the individual instruction
			for instruction, params := range story.Instructions[j] {

				if action, ok := instructions.Action[instruction]; ok {
					actions = append(actions, action(&story, params))
				} else {
					reporter.Write(fmt.Sprintf("Instruction (%s) not supported", instruction), reporter.INPUT_ERROR)
				}

			}
		}

		err := chromedp.Run(ctx, actions...)
		if err != nil {
			reporter.Write(err.Error(), reporter.FLOW_ERROR)
		}

		story.CheckAssertions(i, &storiesCtx.Config)
		
		for j := 0; j < len(story.ActionTodoAfter); j++ {
			story.ActionTodoAfter[j]()
		}
		
	}

	// pretty.Println(storiesCtx)

	return nil

}
