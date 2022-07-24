package router

import (
	"fmt"
	"github.com/xgfone/bt/dht"
	"github.com/xgfone/bt/metainfo"
	"net"
	"strconv"
	"sync"
)

type testPeerManager struct {
	lock  sync.RWMutex
	peers map[metainfo.Hash][]metainfo.Address
}

func NewTestPeerManager() *testPeerManager {
	return &testPeerManager{peers: make(map[metainfo.Hash][]metainfo.Address)}
}

func (pm *testPeerManager) AddPeer(infohash metainfo.Hash, addr metainfo.Address) {
	pm.lock.Lock()
	var exist bool
	for _, orig := range pm.peers[infohash] {
		if orig.Equal(addr) {
			exist = true
			break
		}
	}
	if !exist {
		pm.peers[infohash] = append(pm.peers[infohash], addr)
	}
	pm.lock.Unlock()
}

func (pm *testPeerManager) GetPeers(infohash metainfo.Hash, maxnum int, ipv6 bool) (addrs []metainfo.Address) {
	// We only supports IPv4, so ignore the ipv6 argument.
	pm.lock.RLock()
	_addrs := pm.peers[infohash]
	if _len := len(_addrs); _len > 0 {
		if _len > maxnum {
			_len = maxnum
		}
		addrs = _addrs[:_len]
	}
	pm.lock.RUnlock()
	return
}

func onSearch(infohash string, ip net.IP, port uint16) {
	addr := net.JoinHostPort(ip.String(), strconv.FormatUint(uint64(port), 10))
	fmt.Printf("%s is searching %s\n", addr, infohash)
}

func onTorrent(infohash string, ip net.IP, port uint16) {
	addr := net.JoinHostPort(ip.String(), strconv.FormatUint(uint64(port), 10))
	fmt.Printf("%s has downloaded %s\n", addr, infohash)
}

func NewDHTServer(id metainfo.Hash, addr string, pm dht.PeerManager) (s *dht.Server, err error) {
	conn, err := net.ListenPacket("udp4", addr)
	if err == nil {
		c := dht.Config{ID: id, PeerManager: pm, OnSearch: onSearch, OnTorrent: onTorrent}
		s = dht.NewServer(conn, c)
	}
	return
}
