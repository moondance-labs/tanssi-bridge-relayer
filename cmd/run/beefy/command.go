package beefy

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/moondance-labs/tanssi-bridge-relayer/relays/beefy"
	"github.com/sirupsen/logrus"
	"github.com/snowfork/snowbridge/relayer/chain/ethereum"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

var (
	configFile     string
	privateKey     string
	privateKeyFile string
	privateKeyID   string
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "beefy",
		Short: "Start the beefy relay",
		Args:  cobra.ExactArgs(0),
		RunE:  run,
	}

	cmd.Flags().StringVar(&configFile, "config", "", "Path to configuration file")
	cmd.MarkFlagRequired("config")

	cmd.Flags().StringVar(&privateKey, "ethereum.private-key", "", "Ethereum private key")
	cmd.Flags().StringVar(&privateKeyFile, "ethereum.private-key-file", "", "The file from which to read the private key")
	cmd.Flags().StringVar(&privateKeyID, "ethereum.private-key-id", "", "The secret id to lookup the private key in AWS Secrets Manager")

	return cmd
}

func run(_ *cobra.Command, _ []string) error {
	log.SetOutput(logrus.WithFields(logrus.Fields{"logger": "stdlib"}).WriterLevel(logrus.InfoLevel))
	logrus.SetLevel(logrus.DebugLevel)

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var config beefy.Config
	err := viper.UnmarshalExact(&config)
	if err != nil {
		return err
	}

	err = config.Validate()
	if err != nil {
		return fmt.Errorf("config file validation failed: %w", err)
	}

	keypair, err := ethereum.ResolvePrivateKey(privateKey, privateKeyFile, privateKeyID)
	if err != nil {
		return err
	}

	relay, err := beefy.NewRelay(&config, keypair)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	// Ensure clean termination upon SIGINT, SIGTERM
	eg.Go(func() error {
		notify := make(chan os.Signal, 1)
		signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-notify:
			logrus.WithField("signal", sig.String()).Info("Received signal")
			cancel()
		}

		return nil
	})

	err = relay.Start(ctx, eg)
	if err != nil {
		logrus.WithError(err).Fatal("Unhandled error")
		cancel()
		return err
	}

	err = eg.Wait()
	if err != nil {
		logrus.WithError(err).Fatal("Unhandled error")
		return err
	}

	return nil
}
