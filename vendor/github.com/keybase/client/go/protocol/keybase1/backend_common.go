// Auto-generated by avdl-compiler v1.3.11 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/backend_common.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type BlockIdCombo struct {
	BlockHash string `codec:"blockHash" json:"blockHash"`
	ChargedTo UID    `codec:"chargedTo" json:"chargedTo"`
}

type ChallengeInfo struct {
	Now       int64  `codec:"now" json:"now"`
	Challenge string `codec:"challenge" json:"challenge"`
}

type BackendCommonInterface interface {
}

func BackendCommonProtocol(i BackendCommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.backendCommon",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type BackendCommonClient struct {
	Cli rpc.GenericClient
}
