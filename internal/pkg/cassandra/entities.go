package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type Cluster struct {
	URL      []string
	KeySpace string
	Session  *gocql.Session
}

var Conn Cluster

type Torrent struct {
	InfoHash    string `json:"infohash"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Date        string `json:"date"`
	User        string `json:"userid"`
	Magnet      string `json:"magnet"`
	Size        string `json:"size"`
	Peers       int    `json:"peers"`
	Seeders     int    `json:"seeders"`
	Leechers    int    `json:"leechers"`
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
	Date     string `json:"date"`
	InfoHash string `json:"infohash"`
	Retry    int    `json:"retry"`
}

func Init() {
	fmt.Println("Initializing Cassandra")
	cluster := gocql.NewCluster(Conn.URL...)
	cluster.Keyspace = Conn.KeySpace
	cluster.Consistency = gocql.Quorum
	Conn.Session, _ = cluster.CreateSession()

	CreateTables()
	fmt.Println("Cassandra Session is ready")
}

func CreateTables() {

	fmt.Println("Creating Table torrent_by_infohash")
	var err error
	if err = Conn.Session.Query(create_torrent_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table queue_by_infohash")
	if err = Conn.Session.Query(create_queue_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_userid")
	if err = Conn.Session.Query(create_suggestion_by_userid).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_infohash")
	if err = Conn.Session.Query(create_suggestion_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_date")
	if err = Conn.Session.Query(create_suggestion_by_date).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_userid")
	if err = Conn.Session.Query(create_user_by_userid).Exec(); err != nil {
		log.Fatal(err)
	}

}

/*


func TorrentFindAll(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get(+"/torrents")
		defer resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", err)
			return
		}
		var infohashDTO InfoHash
		err = json.Unmarshal(body, &infohashDTO)
		if err != nil {
			_, _ = fmt.Fprintf(w, "%s", err)
			return
		}
		_, _ = fmt.Fprintf(w, "%s", infohashDTO)
		//_, _ = fmt.Fprintf(w, "%s: %s", route, infohash
	}
}

*/
