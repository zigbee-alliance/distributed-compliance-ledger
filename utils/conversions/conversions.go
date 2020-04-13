package conversions

import "strconv"

func ParseInt16FromString(str string) (int16, error) {
	val, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(val), nil
}
