package ibft

import (
	"github.com/bloxapp/ssv/ibft/proto"
	"github.com/bloxapp/ssv/ibft/sync"
	"github.com/bloxapp/ssv/network"
	"go.uber.org/zap"
)

// ProcessDecidedMessage is responsible for processing an incoming decided message.
// If the decided message is known or belong to the current executing instance, do nothing.
// Else perform a sync operation
/* From https://arxiv.org/pdf/2002.03613.pdf
We can omit this if we assume some mechanism external to the consensus algorithm that ensures
synchronization of decided values.
upon receiving a valid hROUND-CHANGE, λi, −, −, −i message from pj ∧ pi has decided
by calling Decide(λi,− , Qcommit) do
	send Qcommit to process pj
*/
func (i *ibftImpl) ProcessDecidedMessage(msg *proto.SignedMessage) {
	// TODO - validate msg

	i.logger.Info("received highest decided", zap.Uint64("seq number", msg.Message.SeqNumber), zap.Uint64s("node ids", msg.SignerIds))

	// if we already have this in storage, pass
	found, err := i.storage.GetDecided(msg.Message.ValidatorPk, msg.Message.SeqNumber)
	if err != nil {
		i.logger.Error("could not get decided instance from storage", zap.Error(err))
		return
	}
	if found != nil {
		return
	}

	// If received decided for current instance, let that instance play out.
	// otherwise sync
	if i.currentInstance != nil || i.currentInstance.State.SeqNumber < msg.Message.SeqNumber {
		// stop current instance
		i.currentInstance.Stop()

		// sync
		s := sync.NewHistorySync(i.network)
		go s.Start()
	} else {
		// if no current instance and highest decided is lower than msg seq number, sync
		highest, err := i.storage.GetHighestDecidedInstance(msg.Message.ValidatorPk)
		if err != nil {
			i.logger.Error("could not get highest decided instance from storage", zap.Error(err))
			return
		}
		if highest.Message.SeqNumber < msg.Message.SeqNumber {
			// sync
			s := sync.NewHistorySync(i.network)
			go s.Start()
		}
	}

}

func (i *ibftImpl) ProcessSyncMessage(msg *network.SyncChanObj) {
	s := sync.NewReqHandler(i.logger, i.network, i.storage)
	go s.Process(msg)
}
