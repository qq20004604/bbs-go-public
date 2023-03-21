package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

// YYYY-MM-DD hh:mm:ss 格式
type DateTime struct {
	time.Time
}

// 2. 为 DateTime 重写 MarshaJSON 和 UnmarshalJSON 方法，在此方法中实现自定义格式的转换；
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = DateTime{t1}
	return err
}

// 自定义格式转为 JSON
func (t DateTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	if output == `"0001-01-01 00:00:00"` {
		output = `""`
	}
	return []byte(output), nil
}

// 3. 为 DateTime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
	//return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 4. 为 DateTime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
