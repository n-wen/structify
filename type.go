package main

import (
	"fmt"
	"regexp"
)

var TypeMysqlDicMp = map[string]string{
	"smallint":            "int",
	"smallint unsigned":   "int",
	"int":                 "int",
	"int unsigned":        "int",
	"bigint":              "int64",
	"bigint unsigned":     "int64",
	"varchar":             "string",
	"char":                "string",
	"date":                "time.Time",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int",
	"tinyint unsigned":    "int",
	"tinyint(1)":          "int",
	"tinyint(1) unsigned": "int",
	"json":                "string",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
}

// TypeMysqlMatchMp Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchMp = map[string]string{
	`^(tinyint)[(]\d+[)]`:            "int",
	`^(tinyint)[(]\d+[)] unsigned`:   "int",
	`^(smallint)[(]\d+[)]`:           "int",
	`^(int)[(]\d+[)]`:                "int",
	`^(bigint)[(]\d+[)]`:             "int64",
	`^(char)[(]\d+[)]`:               "string",
	`^(enum)[(](.)+[)]`:              "string",
	`^(varchar)[(]\d+[)]`:            "string",
	`^(varbinary)[(]\d+[)]`:          "[]byte",
	`^(binary)[(]\d+[)]`:             "[]byte",
	`^(decimal)[(]\d+,\d+[)]`:        "float64",
	`^(mediumint)[(]\d+[)]`:          "string",
	`^(double)[(]\d+,\d+[)]`:         "float64",
	`^(float)[(]\d+,\d+[)]`:          "float64",
	`^(float)[(]\d+,\d+[)] unsigned`: "float64",
	`^(datetime)[(]\d+[)]`:           "time.Time",
}

// getTypeName Type acquisition filtering.类型获取过滤
func getTypeName(name string, isNull bool) string {
	// Precise matching first.先精确匹配
	if v, ok := TypeMysqlDicMp[name]; ok {
		return v
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for k, v := range TypeMysqlMatchMp {
		if ok, _ := regexp.MatchString(k, name); ok {
			return v
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.maybe need to add", name))
}