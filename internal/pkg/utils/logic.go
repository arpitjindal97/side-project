package utils

import (
	"context"
	"github.com/xgfone/bt/metainfo"
	"github.com/xgfone/bt/tracker"
	"time"
)

func GetPeersFromTrackers(id, infohash metainfo.Hash, trackers []string, peers []metainfo.Address) []metainfo.Address {
	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if len(peers) >= 20 {
		return peers
	}

	resp := tracker.GetPeers(c, id, infohash, trackers)
	for _, r := range resp {
		for _, addr := range r.Resp.Addresses {
			peers = append(peers, addr)
			if len(peers) == 25 {
				return peers
			}
		}
	}
	return peers
}
