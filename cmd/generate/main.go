package generate

import (
	"github.com/spf13/cobra"
)

var Total uint
var Version string

var Generate = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate various code, id, dummy data, etc",
	Long:    `generate various amount and/or type of code, id, dummy data, etc`,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	Generate.PersistentFlags().UintVarP(&Total, "total", "t", 1, "Total data needs to be generated")
	GenerateUuid.PersistentFlags().StringVarP(&Version, "uuidVersion", "", "v4", "Version of UUID (supported: v1 and v4)")
}
