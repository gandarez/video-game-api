package main

// import (
// 	"log/slog"
// 	"os"

// 	"github.com/gandarez/video-game-api/internal/exitcode"

// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// )

// func main() {
// 	v := viper.New()

// 	// Setup logger
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// 	cmd := &cobra.Command{
// 		Use:   "video-game-cli",
// 		Short: "Command line interface to query video-game service.",
// 		PreRun: func(cmd *cobra.Command, args []string) {
// 			verbose := v.GetBool("verbose")

// 			logger.
// 				log.SetOutput(os.Stdout)
// 			log.SetVerbose(verbose)
// 			log.SetJww(verbose, os.Stdout)
// 		},
// 		Run: func(cmd *cobra.Command, args []string) {
// 			Run(cmd, v, logger)
// 		},
// 	}

// 	setFlags(cmd, v, logger)
// }

// func setFlags(cmd *cobra.Command, v *viper.Viper, logger *slog.Logger) {
// 	flags := cmd.Flags()

// 	flags.String(
// 		"api-url",
// 		"",
// 		"API base url used to call video-game service.",
// 	)
// 	flags.Bool("verbose", false, "Turns on debug messages in output.")

// 	err := v.BindPFlags(flags)
// 	if err != nil {
// 		logger.Error("failed to bind cobra flags to viper", slog.Any("error", err))

// 		os.Exit(exitcode.ErrGeneric)
// 	}
// }
