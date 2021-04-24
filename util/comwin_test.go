package util

import (
	"testing"
)

func TestKillVlcWindows(t *testing.T) {
	KillAllVlcWindows()
}

func TestStreamWindows(t *testing.T) {
	host := "http://localhost/stream/"
	path_m3u8 := `C:/Users/Administrator/go/src/streamtec/stream/`
	rtsp := "rtsp://admin:12345hik@192.168.7.227/Streaming/Channels/1"
	path := "streaming-1"
	StreamWindows(host, path_m3u8, rtsp, path)
}
