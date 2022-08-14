/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 20:40
 * @Desc:
 */

package main

import (
	"github.com/ayunwSky/study_go/viperParseConfig/viperGetItem"
	"github.com/ayunwSky/study_go/viperParseConfig/viperToml"
	"github.com/ayunwSky/study_go/viperParseConfig/viperTomlAndJson"
)

func main() {
	viperToml.ParseToml()
	//fmt.Println("I amd main function...---------------------------")
	viperTomlAndJson.ParseTomlAndJson()
	//fmt.Println("I amd main function...---------------------------")
	viperGetItem.ParseItems()
}
