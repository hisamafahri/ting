package generate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var GenerateUuid = &cobra.Command{
	Use:   "uuid [OPTIONS]",
	Short: "Generate universally unique identifier (UUID)",
	Long:  `generate various amount and/or version of universally unique identifier (UUID)`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < int(Total); i++ {
			uuid := uuid.New()
			fmt.Println(uuid)
		}
	},
}
