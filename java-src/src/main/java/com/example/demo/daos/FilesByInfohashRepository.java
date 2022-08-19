package com.example.demo.daos;

import com.example.demo.entities.FilesByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.data.cassandra.repository.Query;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface FilesByInfohashRepository extends CassandraRepository<FilesByInfohash, String> {

}
