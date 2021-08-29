package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/youssefsiam38/twist/src/handle"
)

var dir string

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(authorCmd)
	rootCmd.Flags().StringVarP(&dir, "dir", "d", "twist", "The Twist directory name")
}

var rootCmd = &cobra.Command{
	Use:   "twist",
	Short: "Twist is a Simple UI testing tool",
	Long:  `A fast easy to read and write UI automation testing tool.Complete documentation is available at https://github.com/youssefsiam38/twist`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handle.Handle(dir); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
