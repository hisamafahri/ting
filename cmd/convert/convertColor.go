package convert

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConvertColor = &cobra.Command{
	Use:   "color",
	Short: "Convert color into all popular format",
	Long: `convert one color format into all other 
popular color format such as Hex, RGB, HSL, HSV, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("color")
	},
}
