package cassandra

import (
	"time"
)

type Torrent struct {
	InfoHash    string    `json:"infohash"`
	Category    string    `json:"category"`
	SubCategory string    `json:"subcategory"`
	Comment     string    `json:"comment"`
	Creator     string    `json:"creator"`
	Date        time.Time `json:"date"`
	Leechers    int       `json:"leechers"`
	Magnet      string    `json:"magnet"`
	Name        string    `json:"name"`
	NumFiles    int       `json:"num_files"`
	Peers       int       `json:"peers"`
	Seeders     int       `json:"seeders"`
	Size        int64     `json:"size"`
	User        string    `json:"userid"`
}

type User struct {
	Username string `json:"userid"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Suggestion struct {
	Username string `json:"userid"`
	InfoHash string `json:"infohash"`
	OldJson  string `json:"old_json"`
	NewJson  string `json:"new_json"`
}

type Queue struct {
	InfoHash string    `json:"infohash"`
	Date     time.Time `json:"date"`
	Retry    int       `json:"retry"`
}
