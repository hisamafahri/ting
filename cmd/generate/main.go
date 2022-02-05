package generate

import (
	"github.com/spf13/cobra"
)

var Total uint
var Version string
var IsExpired bool

var Generate = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate various code, id, dummy data, etc",
	Long:    `generate various amount and/or type of code, id, dummy data, etc`,
	Args:    cobra.MinimumNArgs(1),
}

func init() {
	// Parent flag
	Generate.PersistentFlags().UintVarP(&Total, "total", "t", 1, "Total data needs to be generated")

	// Child flag
	GenerateUuid.PersistentFlags().StringVarP(&Version, "version", "v", "v4", "Version of UUID (supported: v1 and v4)")
	GenerateJwt.PersistentFlags().BoolVarP(&IsExpired, "expired", "e", false, "Is JWT token valid or expire")
}
