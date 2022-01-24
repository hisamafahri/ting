package generate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var GenerateUuid = &cobra.Command{
	Use:   "uuid",
	Short: "Generate universally unique identifier (UUID)",
	Long:  `generate various amount and/or version of universally unique identifier (UUID)`,
	Run: func(cmd *cobra.Command, args []string) {
		uuid := uuid.New()
		fmt.Println(uuid)
	},
}