package models

import (
	"fmt"
	"streamtec/util"
	"testing"
)

var M CameraModel

func TestAddCamera(T *testing.T) {
	var data []CameraModel
	data = append(data, CameraModel{
		Id:   1,
		Rtsp: "rtsp://admin:12345hik@192.168.7.227/Streaming/Channels/1",
		Nama: "Streaming 1",
		Path: "streaming-1",
	})
	M.CreateUpdate(data)
	data_now := util.GetDB()
	fmt.Println(data_now)
}

func TestReadCamera(T *testing.T) {
	_, data_now, _ := M.Read()
	fmt.Println(data_now)
}
