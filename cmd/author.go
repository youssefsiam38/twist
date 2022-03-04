package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authorsCmd = &cobra.Command{
	Use:   "authors",
	Short:   "Prints the authors names",
	// Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("© Youssef Siam")
		fmt.Println("© Youssef Attia")
		fmt.Println("© Hla Shaheen")
		fmt.Println("© Youssef Byoumy")
	},
}
