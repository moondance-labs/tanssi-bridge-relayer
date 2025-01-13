package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/snowfork/go-substrate-rpc-client/v4/types"
	"github.com/snowfork/snowbridge/relayer/relays/beacon/header/syncer"
	"github.com/snowfork/snowbridge/relayer/relays/beacon/header/syncer/api"
	"github.com/snowfork/snowbridge/relayer/relays/beacon/header/syncer/scale"
	"github.com/snowfork/snowbridge/relayer/relays/beacon/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	beaconConf "github.com/snowfork/snowbridge/relayer/relays/beacon/config"
	"github.com/snowfork/snowbridge/relayer/relays/beacon/protocol"

	log "github.com/sirupsen/logrus"
)

func GenerateBeaconCheckpointCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-beacon-checkpoint",
		Short: "Generate beacon checkpoint.",
		Args:  cobra.ExactArgs(0),
		RunE:  generateBeaconCheckpoint,
	}

	cmd.Flags().String("config", "/tmp/snowbridge/beacon-relay.json", "Path to the beacon relay config")
	cmd.Flags().Uint64("finalized-slot", 0, "Optional finalized slot to create checkpoint at")
	cmd.Flags().Bool("export-json", false, "Export Json")

	return cmd
}

func generateBeaconCheckpoint(cmd *cobra.Command, _ []string) error {
	err := func() error {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}
		finalizedSlot, _ := cmd.Flags().GetUint64("finalized-slot")

		viper.SetConfigFile(config)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		var conf beaconConf.Config
		err = viper.Unmarshal(&conf)
		if err != nil {
			return err
		}

		p := protocol.New(conf.Source.Beacon.Spec, conf.Sink.Parachain.HeaderRedundancy)
		store := store.New(conf.Source.Beacon.DataStore.Location, conf.Source.Beacon.DataStore.MaxEntries, *p)
		store.Connect()
		defer store.Close()

		client := api.NewBeaconClient(conf.Source.Beacon.Endpoint, conf.Source.Beacon.StateEndpoint)
		s := syncer.New(client, &store, p)

		var checkPointScale scale.BeaconCheckpoint
		if finalizedSlot == 0 {
			checkPointScale, err = s.GetCheckpoint()
		} else {
			checkPointScale, err = s.GetCheckpointAtSlot(finalizedSlot)
		}

		if err != nil {
			return fmt.Errorf("get initial sync: %w", err)
		}
		exportJson, err := cmd.Flags().GetBool("export-json")
		if err != nil {
			return fmt.Errorf("get export-json flag: %w", err)
		}
		if exportJson {
			initialSync := checkPointScale.ToJSON()
			err = writeJSONToFile(initialSync, "dump-initial-checkpoint.json")
			if err != nil {
				return fmt.Errorf("write initial sync to file: %w", err)
			}
		}
		checkPointCallBytes, _ := types.EncodeToBytes(checkPointScale)
		checkPointCallHex := hex.EncodeToString(checkPointCallBytes)
		fmt.Println(checkPointCallHex)
		return nil
	}()
	if err != nil {
		log.WithError(err).Error("error generating beacon checkpoint")
	}

	return nil
}

func writeJSONToFile(data interface{}, path string) error {
	file, _ := json.MarshalIndent(data, "", "  ")

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer f.Close()

	_, err = f.Write(file)

	if err != nil {
		return fmt.Errorf("write to file: %w", err)
	}

	return nil
}
