package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short:  `Print the version of Twist`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Twist v0.1")
	},
}
