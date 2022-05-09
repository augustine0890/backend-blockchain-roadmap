package nft

import (
	"sub-service/pkg/models"
	"time"

	"github.com/Shopify/sarama"
)

type NFT struct {
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

func (n *NFT) processMsg() {
	processReceivedTicker := time.NewTicker(10 * time.Millisecond)
	doProcessCh := make(chan struct{}, 1)

	for {
		select {
		case <-processReceivedTicker.C:
			select {
			case doProcessCh <- struct{}{}:
			default:
			}
		case <-doProcessCh:
			for {
				m := <-n.Msg
				value := m.Value

				data := models.Messages{}
				json.Unmarshal(value, &data)
				switch data.TRANSACTIONTYPE {
				case MINT:
					log.Info("Msg: ", data)
				case BURN:
				case TRANSFER:
					}
				}
			n.Quit()
		}
		case <- n.Quit():
			break
		}
	}
}
