package instructions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
	"github.com/youssefsiam38/twist/src/utils"
)

func saveScreenshotTo(story *models.Story, params interface{}) chromedp.Action {

	var ( 
		buf []byte
		filename string
	)

	isString, s, m := utils.AssertTypeIsString(params)

	if isString {
		filename = *s
	} else {
		if (*m)["file"] == "" {
			reporter.Write("You must specify file parameter to the \"saveScreenshotTo\" instruction", reporter.INPUT_ERROR)
			os.Exit(1)
		}

		filename = (*m)["file"]
	}

	story.ActionTodoAfter = append(story.ActionTodoAfter, func() {
		if !utils.FolderExist("./twist/report") {
			err := os.Mkdir("./twist/report", 0o777)
			if err != nil {
				reporter.Write("Can't make report folder", reporter.FLOW_ERROR)
			}
		}
		if err := ioutil.WriteFile(path.Join("./twist/report/", fmt.Sprintf("%s.png", filename)), buf, 0o777); err != nil {
			reporter.Write(fmt.Sprintf("Can't save screenshot: %v", err), reporter.FLOW_ERROR)
		}
	})

	return fullScreenshot(90, &buf)
}
func fullScreenshot(quality int, res *[]byte) chromedp.Action {
	return	chromedp.FullScreenshot(res, quality)
}