package bridge

import (
	"github.com/SuperconsensusMatrixchain/matrixcore/kernel/contract"
	"github.com/SuperconsensusMatrixchain/matrixcore/protos"
)

func eventsResourceUsed(events []*protos.ContractEvent) contract.Limits {
	var size int64
	for _, event := range events {
		size += int64(len(event.Contract) + len(event.Name) + len(event.Body))
	}
	return contract.Limits{
		Disk: size,
	}
}
