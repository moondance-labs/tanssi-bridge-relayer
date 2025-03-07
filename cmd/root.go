package cmd

import (
	"os"

	"github.com/moondance-labs/tanssi-bridge-relayer/cmd/run"
	"github.com/spf13/cobra"
)

var dataDir string
var configFile string

var rootCmd = &cobra.Command{
	Use:          "snowbridge-relay",
	Short:        "Snowbridge Relay is a bridge between Ethereum and Polkadot",
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(run.Command())
	rootCmd.AddCommand(GenerateBeaconCheckpointCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
