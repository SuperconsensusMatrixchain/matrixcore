package p2p

import (
	xctx "github.com/SuperconsensusMatrixchain/matrixcore/kernel/common/xcontext"
	nctx "github.com/SuperconsensusMatrixchain/matrixcore/kernel/network/context"
	pb "github.com/SuperconsensusMatrixchain/matrixcore/protos"
)

// P2P is the p2p server interface
type Server interface {
	Init(*nctx.NetCtx) error
	Start()
	Stop()

	NewSubscriber(pb.XuperMessage_MessageType, interface{}, ...SubscriberOption) Subscriber
	Register(Subscriber) error
	UnRegister(Subscriber) error

	SendMessage(xctx.XContext, *pb.XuperMessage, ...OptionFunc) error
	SendMessageWithResponse(xctx.XContext, *pb.XuperMessage, ...OptionFunc) ([]*pb.XuperMessage, error)

	Context() *nctx.NetCtx

	PeerInfo() pb.PeerInfo
}
