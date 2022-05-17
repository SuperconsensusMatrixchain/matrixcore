package cmd

import (
	chaincmd "github.com/SuperconsensusMatrixchain/matrixcore/example/xchain/cmd/client/cmd/chain"
	"github.com/SuperconsensusMatrixchain/matrixcore/example/xchain/cmd/client/common/global"
	xdef "github.com/SuperconsensusMatrixchain/matrixcore/example/xchain/common/def"

	"github.com/spf13/cobra"
)

type ChainCmd struct {
	global.BaseCmd
}

func GetChainCmd() *ChainCmd {
	chainCmdIns := new(ChainCmd)

	chainCmdIns.Cmd = &cobra.Command{
		Use:           "chain",
		Short:         "query chain info.",
		Example:       xdef.CmdLineName + " chain status",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// query chain status
	chainCmdIns.Cmd.AddCommand(chaincmd.GetChainStatusCmd().GetCmd())
	// create chain
	chainCmdIns.Cmd.AddCommand(chaincmd.GetCreateChainCmd().GetCmd())

	return chainCmdIns
}
