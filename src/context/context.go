package context

import (
	"context"
	// "github.com/kr/pretty"
	// "github.com/chromedp/chromedp"
	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/models"
)

type StoriesContext struct {
	context.Context
	Config  models.Config
	Stories []models.Story
}

func NewStoriesContext(dirName string) (*StoriesContext, error) {

	config, err := models.ParseConfig(dirName)
	if err != nil {
		return nil, err
	}

	// opts := append(chromedp.DefaultExecAllocatorOptions[:])
	// ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := context.WithTimeout(context.Background(), config.TimeoutDur)

	// The author said to do this to disable logs: https://github.com/chromedp/chromedp/issues/143
	// ctx, _ = chromedp.NewContext(ctx, chromedp.WithLogf(func(s string, i ...interface{}) {}))

	stories := models.ParseStories(config, dirName)

	if config.Output == "stdout" || config.Output == "" {
		reporter.SetConsoleWriter()
	} else {
		reporter.SetFileWriter(config.Output)
	}

	return &StoriesContext{
		Context: ctx,
		Config:  *config,
		Stories: *stories,
	}, nil
}
