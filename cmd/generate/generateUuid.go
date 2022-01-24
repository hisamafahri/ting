package generate

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var GenerateUuid = &cobra.Command{
	Use:   "uuid",
	Short: "Generate universally unique identifier (UUID)",
	Long:  `generate various amount and/or version of universally unique identifier (UUID)`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < int(Total); i++ {
			/*
				Implementation of v2, v3, and v5 is really confusing at this moment.
				probably will be add in the future

			*/
			if Version == "v1" {
				uuid := uuid.NewV1() // v1
				fmt.Println(uuid)
			} else if Version == "v4" {
				uuid := uuid.NewV4() // v4
				fmt.Println(uuid)
			} else {
				log.Fatalf("UUID version '%v' is not valid/supported", Version)
			}
		}
	},
}
