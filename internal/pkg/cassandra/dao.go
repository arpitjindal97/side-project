package cassandra

import (
	"github.com/gocql/gocql"
	"time"
)

func FindTorrentByInfohash(id string) (Torrent, error) {
	var torrent Torrent
	err := Conn.Session.Query(find_torrent_by_infohash, id).Consistency(gocql.One).Scan(
		&torrent.InfoHash,
		&torrent.Category,
		&torrent.SubCategory,
		&torrent.Comment,
		&torrent.Creator,
		&torrent.Date,
		&torrent.Leechers,
		&torrent.Magnet,
		&torrent.Name,
		&torrent.NumFiles,
		&torrent.Peers,
		&torrent.Seeders,
		&torrent.Size,
		&torrent.User,
	)
	return torrent, err
}

func InsertQueueByInfohash(infohash string) error {
	return Conn.Session.Query(insert_queue_by_infohash, infohash, time.Now(), 0).Exec()
}

func FindQueueByInfohash(id string) (queue Queue, err error) {
	err = Conn.Session.Query(find_queue_by_infohash, id).Consistency(gocql.One).Scan(
		&queue.InfoHash,
		&queue.Date,
		&queue.Retry,
	)
	return
}

func UpdateTorrentByInfohashPeers(torrent Torrent) error {
	return Conn.Session.Query(update_torrent_by_infohash,
		torrent.Peers,
		torrent.Seeders,
		torrent.Leechers,
		torrent.InfoHash,
		torrent.Category,
		torrent.SubCategory).Exec()
}
