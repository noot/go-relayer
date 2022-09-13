package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/AthanorLabs/go-relayer/common"
	"github.com/AthanorLabs/go-relayer/contracts"
	"github.com/AthanorLabs/go-relayer/relayer"
	"github.com/AthanorLabs/go-relayer/rpc"
	"github.com/athanorlabs/atomic-swap/ethereum/block"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"

	logging "github.com/ipfs/go-log"
)

const (
	flagEndpoint         = "endpoint"
	flagForwarderAddress = "forwarder-address"
	flagKey              = "key"
	flagRPCPort          = "rpc-port"
	flagDev              = "dev"
	flagLog              = "log"

	defaultGanacheKey = "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
)

var (
	log = logging.Logger("cmd")

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:  flagEndpoint,
			Value: "http://localhost:8545",
			Usage: "Ethereum RPC endpoint",
		},
		&cli.StringFlag{
			Name:  flagForwarderAddress,
			Value: "",
			Usage: "Forwarder contract address",
		},
		&cli.StringFlag{
			Name:  flagKey,
			Value: "eth.key",
			Usage: "Path to file containing Ethereum private key",
		},
		&cli.UintFlag{
			Name:  flagRPCPort,
			Value: 9545,
			Usage: "Relayer RPC server port",
		},
		&cli.BoolFlag{
			Name:  flagDev,
			Usage: "Use development configuration and deploy forwarder contract",
		},
		&cli.StringFlag{
			Name:  flagLog,
			Value: "info",
			Usage: "Set log level: one of [error|warn|info|debug]",
		},
	}

	errInvalidAddress       = errors.New("invalid forwarder address")
	errNoEthereumPrivateKey = errors.New("must provide ethereum private key with --key")
)

func main() {
	app := &cli.App{
		Name:   "relayer",
		Usage:  "Ethereum transaction relayer",
		Flags:  flags,
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func setLogLevels(c *cli.Context) error {
	const (
		levelError = "error"
		levelWarn  = "warn"
		levelInfo  = "info"
		levelDebug = "debug"
	)

	level := c.String(flagLog)
	switch level {
	case levelError, levelWarn, levelInfo, levelDebug:
	default:
		return fmt.Errorf("invalid log level %q", level)
	}

	_ = logging.SetLogLevel("cmd", level)
	_ = logging.SetLogLevel("rpc", level)
	return nil
}

func run(c *cli.Context) error {
	err := setLogLevels(c)
	if err != nil {
		return err
	}

	port := uint16(c.Uint(flagRPCPort))
	dev := c.Bool(flagDev)

	endpoint := c.String(flagEndpoint)
	ec, err := ethclient.Dial(endpoint)
	if err != nil {
		return err
	}

	chainID, err := ec.ChainID(context.Background())
	if err != nil {
		return err
	}

	log.Infof("starting relayer with ethereum endpoint %s and chain ID %s", endpoint, chainID)

	key, err := getPrivateKey(c.String(flagKey), dev)
	if err != nil {
		return err
	}

	forwarder, err := deployOrGetForwarder(
		c.String(flagForwarderAddress),
		ec,
		key,
		chainID,
	)
	if err != nil {
		return err
	}

	cfg := &relayer.Config{
		Ctx:                   context.Background(),
		EthClient:             ec,
		Forwarder:             contracts.NewIMinimalForwarderWrapped(forwarder),
		Key:                   key,
		ChainID:               chainID,
		NewForwardRequestFunc: contracts.NewIMinimalForwarderForwardRequest,
	}

	r, err := relayer.NewRelayer(cfg)
	if err != nil {
		return err
	}

	rpcCfg := &rpc.Config{
		Port:    port,
		Relayer: r,
	}

	server, err := rpc.NewServer(rpcCfg)
	if err != nil {
		return err
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)

	errCh := server.Start()
	select {
	case err := <-errCh:
		return err
	case <-sigc:
		return nil
	}
}

func deployOrGetForwarder(
	addressString string,
	ec *ethclient.Client,
	key *common.Key,
	chainID *big.Int,
) (*contracts.IMinimalForwarder, error) { // TODO: change to interface
	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey(), chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to make transactor: %w", err)
	}

	if addressString == "" {
		address, tx, _, err := contracts.DeployMinimalForwarder(txOpts, ec)
		if err != nil {
			return nil, err
		}

		_, err = block.WaitForReceipt(context.Background(), ec, tx.Hash())
		if err != nil {
			return nil, err
		}

		return contracts.NewIMinimalForwarder(address, ec)
	}

	ok := ethcommon.IsHexAddress(addressString)
	if !ok {
		return nil, errInvalidAddress
	}

	return contracts.NewIMinimalForwarder(ethcommon.HexToAddress(addressString), ec)
}

func getPrivateKey(keyFile string, dev bool) (*common.Key, error) {
	if dev {
		return common.NewKeyFromPrivateKeyString(defaultGanacheKey)
	}

	if keyFile != "" {
		fileData, err := os.ReadFile(filepath.Clean(keyFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read private key file: %w", err)
		}
		keyHex := strings.TrimSpace(string(fileData))
		return common.NewKeyFromPrivateKeyString(keyHex)
	}

	return nil, errNoEthereumPrivateKey
}
