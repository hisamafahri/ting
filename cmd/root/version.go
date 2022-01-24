package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Piranti",
	Long:  `printing the current version of piranti CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piranti v0.1.0")
	},
}
