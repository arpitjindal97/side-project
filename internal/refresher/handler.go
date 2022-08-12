package refresher

import (
	"context"
	"example.com/m/internal/pkg/cassandra"
	"example.com/m/internal/pkg/elasticsearch"
	"example.com/m/internal/pkg/utils"
	"fmt"
	"github.com/xgfone/bt/metainfo"
	"github.com/xgfone/bt/tracker/udptracker"
	"github.com/xgfone/go-apiserver/http/reqresp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"net/http"
	"time"
)

var trackers = []string{
	"tracker.opentrackr.org:1337",
	"open.demonii.com:1337",
	"tracker.openbittorrent.com:6969",
}

var ActiveCount int

func UpdatePeers(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labeler, _ := otelhttp.LabelerFromContext(r.Context())
		labeler.Add(semconv.HTTPRouteKey.String(route))
		c := reqresp.GetContext(w, r)
		w.Header().Set("Content-Type", "application/json")
		id := fmt.Sprintf("%s", c.Data["id"])

		torrent, err := cassandra.FindTorrentByInfohash(id)
		if err != nil {
			w.WriteHeader(404)
			labeler.Add(semconv.HTTPStatusCodeKey.Int(404))
			_, _ = fmt.Fprintf(w, "%s", utils.JsonError(err))
			return
		}

		go getResult(torrent)

		_, _ = fmt.Fprintf(w, "%s", utils.JsonMessage("processed"))
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
		labeler.Add(semconv.HTTPStatusCodeKey.Int(200))
	}
}

func decreaseActiveCount() {
	ActiveCount--
}

func getResult(torrent cassandra.Torrent) {
	ActiveCount++
	defer decreaseActiveCount()
	peers := uint32(0)
	seeders := uint32(0)
	leechers := uint32(0)
	for _, tracker := range trackers {
		client, _ := udptracker.NewClientByDial("udp4", tracker)
		hs := []metainfo.Hash{metainfo.NewHashFromString(torrent.InfoHash)}
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
		rs, _ := client.Scrape(ctx, hs)
		for _, r := range rs {
			if peers < r.Seeders+r.Leechers {
				peers = r.Seeders + r.Leechers
				seeders = r.Seeders
				leechers = r.Leechers
			}
			//fmt.Println("Tracker: " + tracker)
			//fmt.Printf("Seeders: %d\n", r.Seeders)
			//fmt.Printf("Leechers: %d\n", r.Leechers)
			//fmt.Printf("Completed: %d\n", r.Completed)
		}
	}
	//fmt.Printf("Peers: %d\n", peers)
	//fmt.Printf("Seeders: %d\n", seeders)
	//fmt.Printf("Leechers: %d\n", leechers)
	torrent.Peers = int(peers)
	torrent.Seeders = int(seeders)
	torrent.Leechers = int(leechers)
	_ = cassandra.UpdateTorrentByInfohashPeers(torrent)
	elasticsearch.Update(torrent)
}
