/*
 * @ -*- Author: ayunwSky
 * @ -*- Date: 2022/8/14 20:40
 * @ -*- Desc:
 */

package main

import (
	"viperParseConfig/viperGetItem"
	"viperParseConfig/viperItemJson"
	"viperParseConfig/viperToml"
	"viperParseConfig/viperTomlAndJson"
	"viperParseConfig/viperYamlConf"
)

func main() {
	viperToml.ParseToml()

	viperTomlAndJson.ParseTomlAndJson()

	viperItemJson.ParseJsonItem()

	viperGetItem.ParseItems()

	viperYamlConf.ParseYamlConf()
}
