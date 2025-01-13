package solochain

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/snowfork/snowbridge/relayer/chain/ethereum"
	"github.com/snowfork/snowbridge/relayer/chain/relaychain"
	"github.com/snowfork/snowbridge/relayer/crypto/secp256k1"
	"github.com/snowfork/snowbridge/relayer/ofac"

	log "github.com/sirupsen/logrus"
)

type Relay struct {
	config                *Config
	relaychainConn        *relaychain.Connection
	ethereumConnWriter    *ethereum.Connection
	ethereumConnBeefy     *ethereum.Connection
	ethereumChannelWriter *EthereumWriter
	beefyListener         *BeefyListener
}

func NewRelay(config *Config, keypair *secp256k1.Keypair) (*Relay, error) {
	log.Info("Creating worker")

	relaychainConn := relaychain.NewConnection(config.Source.Polkadot.Endpoint)

	ethereumConnWriter := ethereum.NewConnection(&config.Sink.Ethereum, keypair)
	ethereumConnBeefy := ethereum.NewConnection(&config.Source.Ethereum, keypair)

	ofacClient := ofac.New(config.OFAC.Enabled, config.OFAC.ApiKey)

	// channel for messages from beefy listener to ethereum writer
	var tasks = make(chan *Task, 1)

	ethereumChannelWriter, err := NewEthereumWriter(
		&config.Sink,
		ethereumConnWriter,
		tasks,
	)
	if err != nil {
		return nil, err
	}

	beefyListener := NewBeefyListener(
		&config.Source,
		&config.Schedule,
		ethereumConnBeefy,
		relaychainConn,
		ofacClient,
		tasks,
	)

	return &Relay{
		config:                config,
		relaychainConn:        relaychainConn,
		ethereumConnWriter:    ethereumConnWriter,
		ethereumConnBeefy:     ethereumConnBeefy,
		ethereumChannelWriter: ethereumChannelWriter,
		beefyListener:         beefyListener,
	}, nil
}

func (relay *Relay) Start(ctx context.Context, eg *errgroup.Group) error {
	err := relay.ethereumConnWriter.Connect(ctx)
	if err != nil {
		return fmt.Errorf("unable to connect to ethereum: writer: %w", err)
	}

	err = relay.ethereumConnBeefy.Connect(ctx)
	if err != nil {
		return fmt.Errorf("unable to connect to ethereum: beefy: %w", err)
	}

	err = relay.relaychainConn.ConnectWithHeartBeat(ctx, 30*time.Second)
	if err != nil {
		return err
	}

	log.Info("Starting beefy listener")
	err = relay.beefyListener.Start(ctx, eg)
	if err != nil {
		return err
	}

	log.Info("Starting ethereum writer")
	err = relay.ethereumChannelWriter.Start(ctx, eg)
	if err != nil {
		return err
	}

	log.Info("Current relay's ID:", relay.config.Schedule.ID)

	return nil
}
