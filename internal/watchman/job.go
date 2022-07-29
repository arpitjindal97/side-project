package watchman

import (
	"example.com/m/internal/apiserver"
	"example.com/m/internal/pkg/cassandra"
	"fmt"
	"sync"
)

var JobStartTime string

func StartJob() {
	iteration := 0
	for apiserver.RedisCount(JobStartTime) > 0 {
		iteration++
		fmt.Printf("Iteration : %d\n", iteration)
		var wg sync.WaitGroup

		wg.Add(0)
		go thread(&wg, 0, 100, 0)

		wg.Add(1)
		go thread(&wg, 100, 100, 1)

		wg.Add(2)
		go thread(&wg, 200, 100, 2)

		// waiting for all threads to get complete
		wg.Wait()
	}
	fmt.Println("Job Completed")
}

func thread(wg *sync.WaitGroup, offset, count int64, threadId int) {
	defer wg.Done()
	infohashes := apiserver.RedisGet(JobStartTime, offset, count)
	//var torrent cassandra.Torrent
	//var queue cassandra.Queue
	var err error
	for _, infohash := range infohashes {

		// select * from cassandra.indexed table where infohash = infohash
		// if this is already present in cassandra
		// update ttl of this entry in cassandra
		// else
		// insert into cassandra.queue NX

		_, err = cassandra.FindTorrentByInfohash(infohash)

		if err == nil {
			// it was already indexed and present in cassandra
			// send to peer_updater

		} else {
			// it is new torrent, needs to be queued
			// check if it's already present in queue
			_, err = cassandra.FindQueueByInfohash(infohash)
			if err != nil {
				// not present in queue
				fmt.Println("Inserting into queue: " + infohash)
				_ = cassandra.InsertQueueByInfohash(infohash)
			}
		}

		//time.Sleep(time.Second * 1)
		fmt.Printf("Thread %d evaluated %s\n", threadId, infohash)

		// finally
		apiserver.RedisRemove(infohash)
	}
}
