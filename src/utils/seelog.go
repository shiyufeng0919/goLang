package utils
import log "github.com/cihub/seelog"
/*
seelog是用GO语言实现的一个日志系统

Step1: go get -u github.com/cihub/seelog
*/

func SeelogDemo1(){
	defer log.Flush()
	log.Info("hello from seelog")
}
