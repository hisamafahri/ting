package cmd

import (
	"os"

	"github.com/hisamafahri/piranti/cmd/generate"
	"github.com/hisamafahri/piranti/cmd/root"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "piranti",
	Short: "A CLI for tools developer use everyday",
	Long: `Piranti is a CLI tools to automate your everyday
task as a developer.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(root.VersionCmd)

	// Generate
	rootCmd.AddCommand(generate.Generate)
	generate.Generate.AddCommand(generate.GenerateUuid)
}
