package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

var Session *gocql.Session

func Init(uris []string, username, password, keyspace string) {
	log.Println("Initializing Cassandra")
	cluster := gocql.NewCluster(uris...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Second * 3
	cluster.WriteTimeout = time.Second * 3
	cluster.ConnectTimeout = time.Second * 3
	Session, _ = cluster.CreateSession()

	CreateTables()
	log.Println("Cassandra Session is ready")
}

func CreateTables() {

	fmt.Println("Creating Table torrent_by_infohash")
	var err error
	if err = Session.Query(create_torrent_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table queue_by_infohash")
	if err = Session.Query(create_queue_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_userid")
	if err = Session.Query(create_suggestion_by_userid).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_infohash")
	if err = Session.Query(create_suggestion_by_infohash).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_date")
	if err = Session.Query(create_suggestion_by_date).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Table suggestion_by_userid")
	if err = Session.Query(create_user_by_userid).Exec(); err != nil {
		log.Fatal(err)
	}

}
