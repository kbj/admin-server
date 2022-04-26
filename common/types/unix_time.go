package types

import (
	"fmt"
	"time"
)

// UnixTime 自定义时间类型，重写JSON序列化方法，使用unix时间戳形式
type UnixTime time.Time

// MarshalJSON implements json.Marshal.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).Unix())
	return []byte(stamp), nil
}
