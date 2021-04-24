package main

import (
	"log"
	"streamtec/middleware"
	"streamtec/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("TEST")
	// cmd_path := "vlc"
	// cmd_instance := exec.Command(cmd_path, "-I", "dummy", "rtsp://admin:12345hik@192.168.7.227/Streaming/Channels/1", `:sout=#std{access=livehttp{seglen=5,delsegs=true,numsegs=10,index=camera.m3u8,index-url=http://36.89.38.193/streamer/camera-########.ts},mux=ts{use-key-frames},dst=camera-########.ts}}
	// `)
	// cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// cmd_output, _ := cmd_instance.Output()
	// fmt.Println(cmd_output)

	util.StartDB()
	gin.SetMode(gin.DebugMode)
	log.Print("Starting the pagu app")
	middleware.StartEngine()
	middleware.Middleware()

}
