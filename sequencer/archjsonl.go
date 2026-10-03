package sequencer

import (
	"fmt"
	"krc20d/protowire"
	"krc20d/storage"
)

// //////////////////////////////
func archJsonlInit(url string) error {

	// ...

	return nil
}

// //////////////////////////////
func archJsonlGetVspcTxDataList(vspcList []storage.DataVspcType) (bool, uint64, uint64, []storage.DataVspcType, []storage.DataTransactionType, error) {
	return false, 0, 0, nil, nil, fmt.Errorf("disabled")

	// ...

}

// //////////////////////////////
func archJsonlGetTxDataMap(txDataList []storage.DataTransactionType) (map[string]*protowire.RpcTransaction, int64, error) {
	txDataMap := map[string]*protowire.RpcTransaction{}
	return txDataMap, 0, nil
}

// //////////////////////////////
func archJsonlGetNodeArchiveVspcTxDataList(daaScore string) (string, string, []string, map[string]string, error) {
	return "", "", nil, nil, fmt.Errorf("disabled")
}

// //////////////////////////////
func archJsonlGetNodeArchiveTxData(txId string) (string, error) {
	return "", fmt.Errorf("disabled")
}
