package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/youssefsiam38/twist/src/handle"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringVarP(&Verbose, "verbose", "v", "", "verbose output")
	// rootCmd.MarkFlagRequired("verbose")
}

var Verbose string

var rootCmd = &cobra.Command{
	Use:   "twist",
	Short: "Twist is a convenient automated UI testing tool",
	Long:  `A fast easy to read and write automation testing tool.Complete documentation is available at https://github.com/youssefsiam38/twist`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := handle.Handle(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
