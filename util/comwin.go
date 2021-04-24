package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func StreamWindows(host, path_m3u8, rtsp, path string) error {
	fmt.Println("START STREAM")
	cmd_path := "vlc"
	cmd := exec.Command(cmd_path, "-I", "dummy", rtsp, "--http-reconnect", `:sout=#std{access=livehttp{seglen=5,delsegs=true,numsegs=10,index=`+path_m3u8+path+`.m3u8,index-url=`+host+path+`-########.ts},mux=ts{use-key-frames},dst=`+path_m3u8+path+`-########.ts}}`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	return nil
}

func KillAllVlcWindows() {
	fmt.Println("KILL STREAM")
	cmd_path := "TASKKILL"
	cmd_instance := exec.Command(cmd_path, "/IM", "VLC.EXE", "/F")
	cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd_output, _ := cmd_instance.Output()
	fmt.Println(cmd_output)
}
