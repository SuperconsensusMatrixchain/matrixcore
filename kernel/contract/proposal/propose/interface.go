package propose

import (
	pb "github.com/SuperconsensusMatrixchain/matrixcore/protos"
)

type ProposeManager interface {
	GetProposalByID(proposalID string) (*pb.Proposal, error)
}
