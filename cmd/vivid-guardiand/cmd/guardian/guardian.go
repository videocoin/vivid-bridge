package guardian

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/pubsub"

	gossipv1 "vivid-bridge/pkg/proto/gossip/v1/gossip.pb.go"
)

var GuardianCmd = &cobra.Command{
	Use:   "guardiand",
	Short: "Run the Vivid-Bridge Guardian Node",
	Run:   runGuardian,
}

type (
	ChainID uint16
)

const (
	ChainIDUnset    ChainID = 0
	ChainIDVivid    ChainID = 1
	ChainIDEthereum ChainID = 2
)

func (c ChainID) String() string {
	switch c {
	case ChainIDUnset:
		return "unset"
	case ChainIDVivid:
		return "vivid"
	case ChainIDEthereum:
		return "ethereum"
	default:
		return fmt.Sprintf("Unknow chain ID: %d", c)
	}
}

type MessagePublication struct {
	TxHash         common.Hash
	Sequence       uint64
	TimeStamp      time.Time
	EmitterChain   ChainID
	EmitterAddress common.Address
	Payload        []byte
}

type Watcher struct {
	url              string
	contract         common.Address
	networkName      string
	chainID          ChainID
	msgChan          chan *MessagePublication
	minConfirmations uint64
}

var (
	rootCtx       context.Context
	rootCtxCancel context.CancelFunc
)

func NewWatcher(
	url string,
	contract common.Address,
	networkName string,
	chainID ChainID,
	msgChan chan *MessagePublication,
	minConfirmations uint64,
) *Watcher {
	return &Watcher{
		url:              url,
		contract:         contract,
		networkName:      networkName,
		chainID:          chainID,
		msgChan:          msgChan,
		minConfirmations: minConfirmations,
	}
}

func (w *Watcher) Run(ctx context.Context) {
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	errC := make(chan error)
	//messageC := make(chan *ethabi.AbiLogMessagePublished, 2)
	messageC := make(chan *int, 2) // TODO
	var currentBlockNumber uint64

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ev <- messageC:
				message := &MessagePublication{}
				w.msgChan <- message
			}
		}
	}()
}

func runGuardian(cmd *cobra.Command, args []string) {

	rootCtx, rootCtxCancel = context.WithCancel(context.Background())
	defer rootCtxCancel()
	// Initialize logger

	lockC := make(chan *MessagePublication)

	signedInC := make(chan gossipv1.SigendEventWIthQuorum, 10)
	// Run server on port 6060

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exiting runGuardian...")
	time.Sleep(2)
}
