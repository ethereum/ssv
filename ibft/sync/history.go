package sync

import "github.com/bloxapp/ssv/network"

// HistorySync is responsible for syncing and iBFT instance when needed by
// fetching decided messages from the network
type HistorySync struct {
	network network.Network
}

// NewHistorySync returns a new instance of HistorySync
func NewHistorySync(network network.Network) *HistorySync {
	return &HistorySync{network: network}
}

// Start the sync
func (sync *HistorySync) Start() {
	panic("implement HistorySync")
}

// FindHighestInstance returns the highest found instance identifier found by asking the P2P network
func (sync *HistorySync) FindHighestInstance() []byte {

	return nil
}

// FetchValidateAndSaveInstances fetches, validates and saves decided messages from the P2P network.
func (sync *HistorySync) FetchValidateAndSaveInstances(startID []byte, endID []byte) {

}
