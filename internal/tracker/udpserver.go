package tracker

import (
	"bytes"
	"encoding/json"
	"example.com/m/internal/pkg/utils"
	"github.com/xgfone/bt/metainfo"
	"github.com/xgfone/bt/tracker/udptracker"
	"net"
	"net/http"
)

var OtherTrackers []string

func UDPTrackerServer(sconn net.PacketConn) {
	server := udptracker.NewServer(sconn, testHandler{})
	defer server.Close()
	server.Run()
}

type testHandler struct{}

func (testHandler) OnConnect(raddr *net.UDPAddr) (err error) { return }

func (testHandler) OnAnnounce(raddr *net.UDPAddr, req udptracker.AnnounceRequest) (r udptracker.AnnounceResponse, err error) {
	go addTorrent(req.InfoHash.String())
	RedisAdd(req.InfoHash.String(), req.Left, raddr.IP, req.Port)
	r = udptracker.AnnounceResponse{
		Interval:  120, // 2 mins
		Leechers:  uint32(int(RedisCount(req.InfoHash.String() + ":incomplete"))),
		Seeders:   uint32(int(RedisCount(req.InfoHash.String() + ":complete"))),
		Addresses: utils.GetPeersFromTrackers(req.PeerID, req.InfoHash, OtherTrackers, RedisGet(req.InfoHash.String())),
	}
	return
}
func (testHandler) OnScrap(raddr *net.UDPAddr, infohashes []metainfo.Hash) (
	rs []udptracker.ScrapeResponse, err error) {
	rs = make([]udptracker.ScrapeResponse, len(infohashes))
	for i := range infohashes {
		rs[i] = udptracker.ScrapeResponse{
			Seeders:   uint32(int(RedisCount(infohashes[i].String() + ":complete"))),
			Leechers:  uint32(int(RedisCount(infohashes[i].String() + ":incomplete"))),
			Completed: uint32(int(RedisCount(infohashes[i].String() + ":complete"))),
		}
	}
	return
}

var APIServerURL string

func addTorrent(infohash string) {
	reqBody, _ := json.Marshal(map[string]string{
		"infohash": infohash,
	})

	_, _ = http.Post(APIServerURL+"/torrents", "application/json", bytes.NewBuffer(reqBody))
}
