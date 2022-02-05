package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
Init the ersion checker command

*/

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Ting",
		Long:  `printing the current version of ting CLI`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "ting v0.1.0\n")
			return nil
		},
	}
}
