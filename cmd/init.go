package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/youssefsiam38/twist/src/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short:   "Creates the twist folder with the defaults",
	Long:  `Creates the twist folder with the default values for you`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.FolderExist("twist") {
			fmt.Println("twist folder is already exist")
		} else {
			err := os.Mkdir("twist", 0755)
			if err != nil {
				fmt.Println(err)
			}
			err = os.Mkdir("twist/stories", 0755)
			if err != nil {
				fmt.Println(err)
			}

			configBuf := []byte("execute: in order # or in parallel, if in parallel no need to specify the order\ntimeout: 70m\noutput: stdout # or filename\norder: \n- blog")

			
			err = ioutil.WriteFile("twist/config.yml", configBuf, 0755)
			if err != nil {
				fmt.Println(err)
			}

			storyBuf := []byte(`execute: in order # or in parallel, if in parallel no need to specify the order
			timeout: 70m
			output: stdout # or filename
			order: 
			- blog
			`)
			
			err = ioutil.WriteFile("twist/stories/blog.story.yml", storyBuf, 0755)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
