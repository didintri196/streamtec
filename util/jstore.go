package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var schema = `{"url_server":"http://localhost","path_m3u8":"","camera":[]}`

func GetDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
func StartDB() {
	// open input file
	dir := GetDir()
	_, err := os.Open(dir + "/streamtec.data")
	if err != nil {
		_ = ioutil.WriteFile(dir+"/streamtec.data", []byte(schema), 0644)
	}

}

func StoreDB(data string) {
	dir := GetDir()
	_ = ioutil.WriteFile(dir+"/streamtec.data", []byte(data), 0644)
}

func GetDB() (data string) {
	StartDB()
	dir := GetDir()
	content, err := ioutil.ReadFile(dir + "/streamtec.data")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
