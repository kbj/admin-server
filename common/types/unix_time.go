package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UnixTime 自定义时间类型，重写JSON序列化方法，使用unix时间戳形式
type UnixTime time.Time

// MarshalJSON implements json.Marshal.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).UnixMilli())
	return []byte(stamp), nil
}

// UnmarshalJSON 字符串时间戳转为时间类型
func (t *UnixTime) UnmarshalJSON(b []byte) error {
	timeString := string(b)
	if timeString != "" && strings.ContainsRune(timeString, '"') {
		timeString = strings.ReplaceAll(timeString, "\"", "")
	}
	millis, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		return err
	}
	*t = UnixTime(time.UnixMilli(millis))
	return nil
}

func (t UnixTime) ToTime() time.Time {
	return time.Time(t)
}
