package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// UnixMillis 커스텀 타입 정의
type UnixMillis int64

// Scan - sql.Scanner 구현
func (um *UnixMillis) Scan(value interface{}) error {
	if value == nil {
		*um = 0
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*um = UnixMillis(v.UnixMilli())
		return nil
	case []byte:
		// []byte를 string으로 변환
		strVal := string(v)
		// string을 time.Time으로 파싱
		parsedTime, err := time.Parse("2006-01-02 15:04:05.000", strVal)
		if err != nil {
			return err
		}
		*um = UnixMillis(parsedTime.UnixMilli())
		return nil
	default:
		return fmt.Errorf("cannot convert %T to UnixMillis", value)
	}
}

// Value - driver.Valuer 구현
func (um UnixMillis) Value() (driver.Value, error) {
	// 밀리세컨드 단위의 유닉스 타임스탬프를 time.Time으로 변환
	return time.UnixMilli(int64(um)), nil
}
