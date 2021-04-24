package models

import (
	"streamtec/util"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ConfigModel struct {
	UrlServer string `json:"url_server"`
	PathM3u8  string `json:"path_m3u8"`
}

func (M *ConfigModel) CreateUpdate(body ConfigModel) (err error) {
	db := util.GetDB()
	db, err = sjson.Set(db, "url_server", body.UrlServer)
	db, err = sjson.Set(db, "path_m3u8", body.PathM3u8)
	util.StoreDB(db)
	return
}

func (M *ConfigModel) Read() (data ConfigModel, err error) {
	db := util.GetDB()
	url_server := gjson.Get(db, "url_server").String()
	path_m3u8 := gjson.Get(db, "path_m3u8").String()
	data = ConfigModel{
		UrlServer: url_server,
		PathM3u8:  path_m3u8,
	}
	return data, err
}
