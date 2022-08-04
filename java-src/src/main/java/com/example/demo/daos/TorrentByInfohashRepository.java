package com.example.demo.daos;

import com.example.demo.entities.QueueByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.data.cassandra.repository.Query;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface TorrentByInfohashRepository extends CassandraRepository<TorrentByInfohash, String> {

    @Query("SELECT * FROM torrent_by_infohash WHERE infohash=?0 LIMIT 1")
    TorrentByInfohash findByInfohashEquals(String infohash);
}
