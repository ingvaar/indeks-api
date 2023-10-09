package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "indeks <start> [arguments]",
	Short: "Indeks API",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: cmd.OutOrStdout()})
		return nil
	},
}

// Execute the commands.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(NewStartCmd())
}
