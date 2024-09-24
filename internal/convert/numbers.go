package convert

import (
	"fmt"
	"math"
)

// IntToInt32 casts an int to an int64
func IntToInt32(i int) int32 {
	return int32(i)
}

// IntToInt64 casts an int to an int64.
func IntToInt64(i int) int64 {
	return int64(i)
}

// IntToUint32 casts an int to an uint64.
// It returns an error if i < 0.
func IntToUint32(i int) (uint32, error) {
	if i < 0 {
		return 0, fmt.Errorf("integer %d < 0, cannot cast to uint32", i)
	}
	if i > math.MaxUint32 {
		return 0, fmt.Errorf("integer %d > max uint32, cannot cast to uint32", i)
	}
	return uint32(i), nil
}

// IntToUint64 casts an int to an uint64.
// It returns an error if i < 0.
func IntToUint64(i int) (uint64, error) {
	if i < 0 {
		return 0, fmt.Errorf("integer %d < 0, cannot cast to uint64", i)
	}
	return uint64(i), nil
}

// Int64ToInt32 casts an int64 to an int32.
// It returns an error if i does not fit in the int32 range.
func Int64ToInt32(i int64) (int32, error) {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0, fmt.Errorf("converting int64 %d to int32 would cause overflow", i)
	}
	return int32(i), nil
}

// Int64ToInt converts an int64 to an int.
func Int64ToInt(i int64) int {
	return int(i)
}
