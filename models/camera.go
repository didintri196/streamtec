package models

import (
	"encoding/json"
	"streamtec/util"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CameraModel struct {
	Id   int    `json:"id"`
	Rtsp string `json:"rtsp"`
	Nama string `json:"nama"`
	Path string `json:"path"`
}

func (M *CameraModel) CreateUpdate(body []CameraModel) (err error) {
	db := util.GetDB()
	db, err = sjson.Set(db, "camera", body)
	util.StoreDB(db)
	return
}

func (M *CameraModel) Read() (host string, data []CameraModel, err error) {
	db := util.GetDB()
	camera := gjson.Get(db, "camera").String()
	host = gjson.Get(db, "url_server").String()
	json.Unmarshal([]byte(camera), &data)
	return host, data, err
}

func (M *CameraModel) StartStream() (err error) {
	db := util.GetDB()
	url_serverhost := gjson.Get(db, "url_server")
	path_m3u8 := gjson.Get(db, "path_m3u8")
	_, data_camera, err := M.Read()
	//MATIKAN SEMUA STREAM
	util.KillAllVlcWindows()
	//NYALAKAN KEMBALI
	for _, rowcam := range data_camera {
		util.StreamWindows(url_serverhost.String(), path_m3u8.String(), rowcam.Rtsp, rowcam.Path)
	}
	return err
}
