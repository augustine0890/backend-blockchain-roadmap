package nft

import (
	"encoding/json"
	"sub-service/pkg/models"
	"sync"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
)

type NFT struct {
	nftMtx sync.Mutex

	// storage
	// search
	Msg chan *sarama.ConsumerMessage
}

const (
	MINT     = 1
	BURN     = 2
	TRANSFER = 3
)

func NewNFT() *NFT {
	return &NFT{
		Msg: make(chan *sarama.ConsumerMessage),
	}
}

func (n *NFT) ProcessMsg() {
	for {
		m := <-n.Msg
		value := m.Value
		data := models.Messages{}
		json.Unmarshal(value, &data)
		switch data.TRANSACTIONTYPE {
		case MINT:
			log.Info("Msg: ", data)
		case BURN:
			log.Info("Msg: ", data)
		case TRANSFER:
			log.Info("Msg: ", data)
		default:
		}
	}
}
