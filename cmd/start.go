package cmd

import (
	"os"
	"os/signal"

	"github.com/ingvaar/indeks-api/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	flagPort     = "port"
	flagTimeout  = "timeout"
	flagMongoURI = "mongo-uri"
	flagDevMode = "dev"
)

// NewStartCmd creates a new instance of the start command.
func NewStartCmd() *cobra.Command {
	var config internal.Config

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the application",
		Run: func(cmd *cobra.Command, args []string) {
			server, err := internal.New(config)
			if err != nil {
				log.Fatal().Err(err).Msg("Cannot initialize the server")
			}

			if err := server.Start(); err != nil {
				log.Fatal().Err(err).Msg("Cannot start the server")
			}

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig

			if err := server.Stop(); err != nil {
				log.Fatal().Err(err).Msg("Cannot stop the server gracefully")
			}
		},
	}

	startCmd.PersistentFlags().UintVar(&config.Port, flagPort, 8181, "Application port")
	startCmd.PersistentFlags().UintVar(&config.Timeout, flagTimeout, 5, "Time in sec before timeout")
	startCmd.PersistentFlags().StringVar(&config.MongoURI,
		flagMongoURI,
		"mongodb://localhost:27017",
		"MongoDB connection URI",
	)
	startCmd.PersistentFlags().BoolVarP(&config.DevMode, flagDevMode, "d", false, "Enable dev mode with additional debug")

	return startCmd
}
