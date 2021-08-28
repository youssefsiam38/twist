package reporter

import (
	"fmt"
	"io"
	"os"
	"sync"

	c "github.com/logrusorgru/aurora/v3"
)

const (
	NO_ERROR = iota
	INPUT_ERROR
	FLOW_ERROR
	INVALID_SELECTOR_ERROR
	YAML_ERROR

	// non error statuses
	STORY_HEADER
	ASSERTION_SUCCESS
	ASSERTION_FAILD
)
var writer io.Writer
func init() {
	writer = os.Stdout
}


func SetFileWriter(filename string) {
	var (
		once sync.Once
		err error
	)
	once.Do(func() {
		writer, err = os.Create(fmt.Sprintf("./twist/report/%s", filename))
		if err != nil {
			panic(err)
		}
	})
}
func SetConsoleWriter() {
	writer = os.Stdout
}

func Write(s string, writeStatus int) (int, error) {
	switch writeStatus {
	case STORY_HEADER:
		s = fmt.Sprint(c.Gray(0, s).BgYellow().Bold())
	case ASSERTION_SUCCESS:
		s = fmt.Sprintf("\t✅ %s", s)
	case ASSERTION_FAILD:
		s = fmt.Sprintf("\t❌ %s", s)
	}
	return writer.Write([]byte(fmt.Sprint(s, "\n")))
}
