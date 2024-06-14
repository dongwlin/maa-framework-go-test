package main

import (
	"fmt"
	"github.com/MaaXYZ/maa-framework-go"
	"path/filepath"
)

func main() {
	maa.SetLogDir("debug")
	maa.SetSaveDraw(true)
	maa.SetStdoutLevel(maa.LoggingLevelInfo)

	testingDataSetDir := "TestingDataSet"
	testingPath := filepath.Join(testingDataSetDir, "PipelineSmoking", "MaaRecording.txt")
	resultPath := filepath.Join(testingDataSetDir, "debug")

	ctrl := maa.NewDbgController(testingPath, resultPath, maa.DbgControllerTypeReplayRecording, "{}", nil)
	defer ctrl.Destroy()
	ctrlId := ctrl.PostConnect()

	//res := maa.NewResource(nil)
	res := maa.NewResource(func(msg, detailsJson string) {
		fmt.Println("--------------------------------------")
		fmt.Println("msg: ", msg)
		fmt.Println("details: ", detailsJson)
		fmt.Println("--------------------------------------")
	})
	defer res.Destroy()
	resDir := filepath.Join(testingDataSetDir, "PipelineSmoking", "resource")
	resId := res.PostPath(resDir)

	ctrl.Wait(ctrlId)
	res.Wait(resId)

	inst := maa.New(nil)
	defer inst.Destroy()
	inst.BindResource(res)
	inst.BindController(ctrl)

	if !inst.Inited() {
		fmt.Println("failure")
		return
	}

	taskId := inst.PostTask("Wilderness", "{}")
	stats := inst.WaitTask(taskId)
	if stats == maa.Success {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}
}
