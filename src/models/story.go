package models

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"time"

	"github.com/youssefsiam38/twist/reporter"
	"github.com/youssefsiam38/twist/src/utils"
	"gopkg.in/yaml.v2"
)

type Story struct {
	StoryConfig  `yaml:",inline"`
	Instructions []Instruction `yaml:"instructions"`
	Assertions   []Assertion
	ActionTodoAfter []func()
}

type StoryConfig struct {
	Start    string `yaml:"start"`
	Headless bool   `yaml:"headless"`
	Timeout string   `yaml:"timeout"`
	TimeoutDur time.Duration
}

func ParseStories(config *Config) *[]Story {
	var stories []Story

	for _, v := range config.Order {
		story := parseStory(v)
		stories = append(stories, story)
	}

	return &stories
}

func parseStory(name string) Story {
	story := Story{StoryConfig: StoryConfig{Timeout: "5m"}}

	storyFile, err := ioutil.ReadFile(fmt.Sprintf("./twist/stories/%s.story.yml", name))
	if err != nil {
		reporter.Write(err.Error(), reporter.YAML_ERROR)
		os.Exit(1)
	}

	err = yaml.Unmarshal(storyFile, &story)
	if err != nil {
		reporter.Write(err.Error(), reporter.YAML_ERROR)
		os.Exit(1)
	}

	story.StoryConfig.TimeoutDur, err = time.ParseDuration(story.StoryConfig.Timeout) 
	if err != nil {
		reporter.Write(err.Error(), reporter.INPUT_ERROR)
		os.Exit(1)
	}

	return story
}

func (story Story) CheckAssertions(storyIndex int, config *Config) {
	i := storyIndex
	if len(story.Assertions) < 1 {
		return
	}
	reporter.Write(fmt.Sprintf(" %v Story ", config.Order[i]), reporter.STORY_HEADER)
	for j := 0; j < len(story.Assertions); j++ {
		assertion := story.Assertions[j]

		switch assertion.Instruction {
			
		case "assertPathIs":
			if utils.IsValidUrl(assertion.GetExpected()) {
				if assertion.GetExpected() == *assertion.Found {
					reporter.Write(assertion.Name, reporter.ASSERTION_SUCCESS)
				} else {
					reporter.Write(fmt.Sprintf("%s\n\t\tExpected: %s\n\t\tFound: %s", assertion.Name, assertion.GetExpected(), *assertion.Found), reporter.ASSERTION_FAILD)
				}
			} else {
				startUrl, _ := url.Parse(story.Start)

				fullExpected, _ := utils.JoinUrl(utils.CleanHost(startUrl), assertion.GetExpected())
				fullExpectedUrl, _ := url.Parse(*fullExpected)
				foundUrl, err := url.Parse(*assertion.Found)
				if err != nil {
					reporter.Write(fmt.Sprintf("Somthing wrong with the start url in %s", config.Order[i]), reporter.INPUT_ERROR)
				}
				if err != nil {
					reporter.Write(err.Error(), reporter.INPUT_ERROR)
				}

				if utils.Clean(fullExpectedUrl) == utils.Clean(foundUrl) {
					reporter.Write(assertion.Name, reporter.ASSERTION_SUCCESS)
					} else {
						reporter.Write(fmt.Sprintf("%s\n\t\tExpected: %s\n\t\tFound: %s", assertion.Name, utils.Clean(fullExpectedUrl), utils.Clean(foundUrl)), reporter.ASSERTION_FAILD)
					}
				}
		case "assertText":
			if assertion.GetExpected() == *assertion.Found {
				reporter.Write(assertion.Name, reporter.ASSERTION_SUCCESS)
			} else {
				reporter.Write(fmt.Sprintf("%s\n\t\tExpected: %s\n\t\tFound: %s", assertion.Name, assertion.GetExpected(), *assertion.Found), reporter.ASSERTION_FAILD)
			}
		}
		
		// reporter.Write(assertion.Instruction)
	}
}
