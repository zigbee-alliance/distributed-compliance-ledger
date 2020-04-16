package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
)

type ReadResult struct {
	Result json.RawMessage `json:"result"`
	Height int64           `json:"height"`
}

func NewReadResult(cdc *codec.Codec, data []byte, height int64) ReadResult {
	var value json.RawMessage
	cdc.MustUnmarshalJSON(data, &value)

	return ReadResult{
		Result: value,
		Height: height,
	}
}

// Implement fmt.Stringer
func (n ReadResult) String() string {
	res, err := json.Marshal(n)

	if err != nil {
		panic(err)
	}

	return string(res)
}
