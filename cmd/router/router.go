package main

import (
	"example.com/m/internal/router"
	"fmt"
	"github.com/xgfone/bt/metainfo"
)

func main() {
	pm := router.NewTestPeerManager()
	server, err := router.NewDHTServer(metainfo.NewRandomHash(), "0.0.0.0:9001", pm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer server.Close()
	server.Bootstrap([]string{"router.bittorrent.com:6881"})
	server.Run()
}
