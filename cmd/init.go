package cmd

import (
	"os"

	"github.com/hisamafahri/ting/cmd/convert"
	"github.com/hisamafahri/ting/cmd/generate"
	"github.com/hisamafahri/ting/cmd/root"
	"github.com/spf13/cobra"
)

/*
Init the root command

*/
var rootCmd = &cobra.Command{
	Use:   "ting",
	Short: "A CLI for tools developer use everyday",
	Long: `Ting is a CLI tools to automate your everyday
task as a developer.`,
}

/*
Init the root function

*/
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/*
Init all of the available commands

*/
func init() {
	rootCmd.AddCommand(root.VersionCmd())

	// Generate
	rootCmd.AddCommand(generate.Generate)
	generate.Generate.AddCommand(generate.GenerateUuid)
	generate.Generate.AddCommand(generate.GenerateJwt)

	// Convert
	rootCmd.AddCommand(convert.Convert)
	convert.Convert.AddCommand(convert.ConvertColor)
}
