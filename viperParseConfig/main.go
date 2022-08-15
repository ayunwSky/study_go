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
	"github.com/ayunwSky/study_go/viperParseConfig/viperJsonItem"
)

func main() {
	//fmt.Println("I amd main function...---------------------------")
	viperToml.ParseToml()

	//fmt.Println("I amd main function...---------------------------")
	viperTomlAndJson.ParseTomlAndJson()
	
	//fmt.Println("I amd main function...---------------------------")
	viperGetItem.ParseItems()
	
	//fmt.Println("I amd main function...---------------------------")
	viperJsonItem.ParseJsonItem()
}
