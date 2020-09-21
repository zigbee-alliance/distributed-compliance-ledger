// nolint: stylecheck
package range_helper

import "github.com/cosmos/cosmos-sdk/store/iavl"

type Context interface {
	QueryRange(startKey []byte, endKey []byte, limit int, storeName string) (iavl.RangeRes, int64, error)
	QueryStoreAtHeight(height int64, key []byte, storeName string) ([]byte, int64, error)
}
