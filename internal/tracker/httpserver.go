package tracker

import (
	"fmt"
	"github.com/xgfone/bt/metainfo"
	"github.com/xgfone/bt/tracker/httptracker"
	"log"
	"net/http"
	"net/url"
)

func httptrackerserver() {
	http.HandleFunc("/announce", func(w http.ResponseWriter, r *http.Request) {
		m, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Fatal(err)
		}
		announceRequest := &httptracker.AnnounceRequest{}
		announceRequest.FromQuery(m)
		address, err := metainfo.NewAddressFromString(r.RemoteAddr)
		peer := httptracker.Peer{
			ID:   announceRequest.PeerID.String(),
			IP:   address.IP.String(),
			Port: address.Port}
		fmt.Println(peer)
		fmt.Println(announceRequest)
		fmt.Fprintf(w, "Welcome to new server!")
	})
	http.ListenAndServe(":8080", nil)
}
