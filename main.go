package main

import (
	"fmt"
	"math/rand"
	"strconv"

	modules "github.com/AmirMirzayi/arvancloud-interview/pkg"
)

var processingQueue, userList, userRemainingQuotaInMinute, userRemainingQuotaInMonth *modules.DataMap

func init() {
	processingQueue = modules.NewDataMap()
	userList = modules.NewDataMap()
	userRemainingQuotaInMinute = modules.NewDataMap()
	userRemainingQuotaInMonth = modules.NewDataMap()
}

func main() {

	processingQueue.SetValue("amir", 30)

	fmt.Println(processingQueue.GetValue("amir"))
	fmt.Println(processingQueue.ExistedKey("amirs"))
	fmt.Println(processingQueue.GetData())

	go modules.CronJob(modules.GetNextMinute, func() {
		userRemainingQuotaInMinute = modules.NewDataMap()
	})
	go modules.CronJob(modules.GetNextMonth, func() {
		userRemainingQuotaInMonth = modules.NewDataMap()
	})

}

func checkQuotaAndProduce(dataId string, userId int) {
	if !userRemainingQuotaInMonth.ExistedKey(strconv.Itoa(userId)) {
		userRemainingQuotaInMonth.SetValue(strconv.Itoa(userId), rand.Int())
	}

	if userRemainingQuotaInMinute.ExistedKey(strconv.Itoa(userId)) {
		userRemainingQuotaInMinute.SetValue(strconv.Itoa(userId), rand.Int())
	}

	if userRemainingQuotaInMonth.GetValue(strconv.Itoa(userId)) < 1 ||
		userRemainingQuotaInMinute.GetValue(strconv.Itoa(userId)) < 1 {
		return
	}
	
	userRemainingQuotaInMonth.SetValue(strconv.Itoa(userId), userRemainingQuotaInMonth.GetValue(strconv.Itoa(userId))-1)
	userRemainingQuotaInMinute.SetValue(strconv.Itoa(userId), userRemainingQuotaInMinute.GetValue(strconv.Itoa(userId))-1)
	produce(dataId, userId)
}

func produce(dataId string, userId int) {
	if !processingQueue.ExistedKey(dataId) {
		return
	}
	processingQueue.SetValue(dataId, userId)
}
