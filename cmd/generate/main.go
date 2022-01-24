package generate

import (
	"github.com/spf13/cobra"
)

var Generate = &cobra.Command{
	Use:   "generate",
	Short: "Generate various code, id, dummy data, etc",
	Long:  `generate various amount and/or type of code, id, dummy data, etc`,
	Args:  cobra.MinimumNArgs(1),
}
