package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Twist",
	Long:  `All software has versions. This is Twist's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Twist automation testing tool v0.1")
	},
}
