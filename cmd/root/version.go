package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Ting",
		Long:  `printing the current version of ting CLI`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("ting v0.1.0")
			return nil
		},
	}
}
