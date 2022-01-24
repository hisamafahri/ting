package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ting",
	Long:  `printing the current version of ting CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ting v0.1.0")
	},
}
