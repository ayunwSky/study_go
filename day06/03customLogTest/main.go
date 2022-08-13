/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:56
 * @Desc:
 */

package main

import (
	"github.com/ayunwSky/study_go/day06/custom_log"
	"time"
)

// 测试自己定义的日志库
func main() {
	id := 1101
	name := "ayunwSky"
	log := custom_log.NewLog("INFO")

	for {
		log.Debug("这是一条 Debug 级别的测试日志")
		log.Info("这是一条 Info 级别的测试日志")
		log.Warning("这是一条 Warning 级别的测试日志，id: %d, name: %s", id, name)
		log.Error("这是一条 Error 级别的测试日志，id: %d, name: %s", id, name)
		log.Fatal("这是一条 Fatal 级别的测试日志\n")

		time.Sleep(3 * time.Second)
	}
}
