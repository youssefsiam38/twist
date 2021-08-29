package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authorCmd = &cobra.Command{
	Use:   "author",
	Short:   "Prints the author name",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("© Youssef Siam (https://github.com/youssefsiam38)")
	},
}
