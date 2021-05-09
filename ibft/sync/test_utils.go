package sync

//type testNetwork struct {
//	t                      *testing.T
//	highestDecidedReceived *proto.SignedMessage
//}
//
//// newTestNetwork testnet
//func newTestNetwork(t *testing.T) *testNetwork {
//	return &testNetwork{t: t}
//}
//
//// Broadcast implementation
//func (n *testNetwork) Broadcast(_ *proto.SignedMessage) error {
//	return nil
//}
//
//// ReceivedMsgChan implementation
//func (n *testNetwork) ReceivedMsgChan() <-chan *proto.SignedMessage {
//	return nil
//}
//
//// BroadcastSignature implementation
//func (n *testNetwork) BroadcastSignature(_ *proto.SignedMessage) error {
//	return nil
//}
//
//// ReceivedSignatureChan implementation
//func (n *testNetwork) ReceivedSignatureChan() <-chan *proto.SignedMessage {
//	return nil
//}
//
//// BroadcastDecided implementation
//func (n *testNetwork) BroadcastDecided(_ *proto.SignedMessage) error {
//	return nil
//}
//
//// ReceivedDecidedChan implementation
//func (n *testNetwork) ReceivedDecidedChan() <-chan *proto.SignedMessage {
//	return nil
//}
//
//// GetHighestDecidedInstance implementation
//func (n *testNetwork) GetHighestDecidedInstance(_ []peer.ID, _ *network.SyncMessage) (*network.Message, error) {
//	return nil, nil
//}
//
//// RespondToHighestDecidedInstance implementation
//func (n *testNetwork) RespondToHighestDecidedInstance(_ network.SyncStream, msg *network.SyncMessage) error {
//	n.highestDecidedReceived = msg.SignedMessages[0]
//	return nil
//}
//
//// ReceivedSyncMsgChan implementation
//func (n *testNetwork) ReceivedSyncMsgChan() <-chan *network.SyncChanObj {
//	return nil
//}

//type testStorage struct {
//	highestDecided *proto.SignedMessage
//}
//
//// newTestStorage test
//func newTestStorage(highestDecided *proto.SignedMessage) *testStorage {
//	return &testStorage{highestDecided: highestDecided}
//}
//
//// SaveCurrentInstance implementation
//func (s *testStorage) SaveCurrentInstance(_ *proto.State) error {
//	return nil
//}
//
//// GetCurrentInstance implementation
//func (s *testStorage) GetCurrentInstance(_ []byte) (*proto.State, error) {
//	return nil, nil
//}
//
//// SaveDecided implementation
//func (s *testStorage) SaveDecided(_ *proto.SignedMessage) error {
//	return nil
//}
//
//// GetDecided implementation
//func (s *testStorage) GetDecided(_ []byte, _ uint64) (*proto.SignedMessage, error) {
//	return nil, nil
//}
//
//// SaveHighestDecidedInstance implementation
//func (s *testStorage) SaveHighestDecidedInstance(_ *proto.SignedMessage) error {
//	return nil
//}
//
//// GetHighestDecidedInstance implementation
//func (s *testStorage) GetHighestDecidedInstance(_ []byte) (*proto.SignedMessage, error) {
//	return s.highestDecided, nil
//}
//
//type testSyncStream struct {
//}
//
//// NewTestStream test
//func NewTestStream() network.SyncStream {
//	return &testSyncStream{}
//}
//
//func (s *testSyncStream) Read(p []byte) (n int, err error) {
//	return 0, nil
//}
//
//func (s *testSyncStream) Write(p []byte) (n int, err error) {
//	return 0, nil
//}
//
//func (s *testSyncStream) Close() error {
//	return nil
//}
//
//func (s *testSyncStream) CloseWrite() error {
//	return nil
//}
//
//func (s *testSyncStream) RemotePeer() string {
//	return ""
//}
