package convert

import (
	"github.com/spf13/cobra"
)

var Convert = &cobra.Command{
	Use:     "convert",
	Aliases: []string{"c"},
	Short:   "Convert various type of data into another data",
	Long:    `convert various types or format of data into another type or format of data`,
	Args:    cobra.MinimumNArgs(1),
}
