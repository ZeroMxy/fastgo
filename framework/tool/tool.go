package tool

import (
	"encoding/json"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Tool struct {}

// 搜索数组中是否存在指定的值
// 存在 => true / 不存在 => false
func In_array (array []string, target string) bool {

	// 有序字符串数组
	sort.Strings(array)

	// 通过二分查找
	// index 为目标字符串在排序后的列表中第一次出现的索引
	// 若不存在，返回数组中最后一个元素的索引
	var index = sort.SearchStrings(array, target)

	// 小于等于最后一个元素的索引和值相等的情况下存在
	if index <= len(array) && array[index] == target {
		return true
	}

	return false

}

// 检查文件或文件夹是否存在
// 存在 => true / 不存在 => false
func Path_is_exist (path string) bool {

	var _, path_err = os.Stat(path)

	if path_err != nil {
		return false
	}

	return true

}

// 正则
const (
	// 邮箱
	EMAIL = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 手机号
	PHONE = "^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$"
	// 身份证号
	ID_CARD = "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"
)

// 正则匹配
// 匹配成功 => true / 匹配失败 => false
func Match (role, context string) bool {

	var matched, _ = regexp.MatchString(role, context)

	return matched

}

// 日期
const (
	YMDHIS 	= "2006-01-02 15:04:05"
	YMD 	= "2006-01-02"
	HIS 	= "15:04:05"
)

// 日期格式化
func Format_date_time (layout, date_time string) string {

	if date_time == "" {
		date_time = time.Now().Format(layout)
	} else {
		var time_value, _ = time.Parse(layout, date_time)
		date_time = time_value.Format(layout)
	}

	return date_time

}

// 隐藏部分字符串
func Hide_string (data, symbol string, start, length int) string {

	// 符号赋默认值
	if symbol == "" {
		symbol = "*"
	}

	// 字符串空返回拼接隐藏符号
	if data == "" {
		return symbol + symbol + symbol + symbol
	}

	if start < 0 || length <= 0 {
		return data
	}

	var string_value = ""
	// 字符串转数组并循环
	for index, value := range strings.Split(data, "") {

		if index >= start && index <= length - 1 {
			string_value += symbol
		} else {
			string_value += value
		}

	}

	return string_value

}

// json 转 map
func Json_to_map (json_data []byte) map[string] interface {} {

	// 初始化 map 类型
	var data map[string] interface {}

	var err = json.Unmarshal(json_data, &data)

	if err != nil {
		return nil
	}

	return data
	
}

// struct 转 map
func Struct_to_map (struct_data interface {}, tag_name string) map[string] interface {} {

	// 初始化 map 类型
	var map_data map[string] interface {}

	var value = reflect.ValueOf(struct_data)

	// 非结构体返回空
	if value.Kind() != reflect.Struct {
		return nil
	}

	// 比对类别取元素
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	var type_value = value.Type()

	// 遍历结构体 tag_name 为 key 值为 value
	for i := 0; i < value.NumField(); i++ {

		var field_index = type_value.Field(i)
		var tag_value 	= field_index.Tag.Get(tag_name)

		if tag_value != "" {
			map_data[tag_value] = value.Field(i).Interface()
		}

	}

	return map_data

}
