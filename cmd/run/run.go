package run

import (
	"github.com/moondance-labs/tanssi-bridge-relayer/cmd/run/beefy"
	"github.com/moondance-labs/tanssi-bridge-relayer/cmd/run/solochain"
	"github.com/snowfork/snowbridge/relayer/cmd/run/beacon"
	"github.com/snowfork/snowbridge/relayer/cmd/run/execution"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start a relay service",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(beefy.Command())
	cmd.AddCommand(beacon.Command())
	cmd.AddCommand(execution.Command())
	cmd.AddCommand(solochain.Command())

	return cmd
}
