package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/youssefsiam38/twist/src/utils"
)

var initDir string

func init() {
	initCmd.Flags().StringVarP(&initDir, "dir", "d", "twist", "The directory name you want to create")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short:   "Creates the twist folder with the defaults",
	Long:  `Creates the twist folder with the default values for you`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.FolderExist(initDir) {
			fmt.Printf("%s folder is already exist", initDir)
		} else {
			err := os.Mkdir(initDir, 0755)
			if err != nil {
				fmt.Println(err)
			}
			err = os.Mkdir(fmt.Sprintf("%s/stories",initDir), 0755)
			if err != nil {
				fmt.Println(err)
			}

			configBuf := []byte("execute: in order # or in parallel (not supported), if in parallel no need to specify the order\ntimeout: 70m\noutput: stdout # or filename (not supported)\norder: \n- blog")

			
			err = ioutil.WriteFile(fmt.Sprintf("%s/config.yml", initDir), configBuf, 0755)
			if err != nil {
				fmt.Println(err)
			}

			storyBuf := []byte("start: https://youssefsiam.me\nheadless: false\ntimeout: 1m\ninstructions:\n- waitFor: \n    selector: \".blog-post-card-wrapper\"\n- assertPathIs: https://youssefsiam.me")
			
			err = ioutil.WriteFile(fmt.Sprintf("%s/stories/blog.story.yml",initDir), storyBuf, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
