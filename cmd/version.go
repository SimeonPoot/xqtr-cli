package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of xqtr",
	Long:  `All software has versions. This is xqtr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("xqtr version -- commit %s HEAD \n", commit)
		// fmt.Println("inside version: ", author)
	},
}
