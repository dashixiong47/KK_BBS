package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type JSONSlice []interface{}
type JSONMap []interface{}
type IntArray []int64

func (j *JSONSlice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}
func (j *JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}

// Value 实现 driver.Valuer 接口
func (a *IntArray) Value() (driver.Value, error) {
	if len(*a) == 0 {
		return "{}", nil
	}

	strs := make([]string, len(*a))
	for i, v := range *a {
		strs[i] = fmt.Sprint(v)
	}

	return "{" + strings.Join(strs, ",") + "}", nil
}

// Scan 实现 sql.Scanner 接口
func (a *IntArray) Scan(src interface{}) error {
	asString, ok := src.(string)
	if !ok {
		return errors.New("not_byte")
	}
	// 检查是否是空数组
	if asString == "{}" {
		*a = make(IntArray, 0)
		return nil
	}
	items := strings.Split(asString[1:len(asString)-1], ",")
	*a = make(IntArray, len(items))
	for i, v := range items {
		_, err := fmt.Sscan(v, &(*a)[i])
		if err != nil {
			return errors.New("unknown")
		}
	}

	return nil
}
