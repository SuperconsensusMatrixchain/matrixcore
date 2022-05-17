/*
Copyright Baidu Inc. All Rights Reserved.
*/

package gm

import (
	"github.com/SuperconsensusMatrixchain/matrixcore/lib/crypto/client/base"
	"github.com/xuperchain/crypto/client/service/gm"
)

// make sure this plugin implemented the interface
var _ base.CryptoClient = (*GmCryptoClient)(nil)

type GmCryptoClient struct {
	gm.GmCryptoClient
}

func GetInstance() base.CryptoClient {
	gmCryptoClient := GmCryptoClient{}
	return &gmCryptoClient
}
