package com.example.demo.daos;

import com.example.demo.entities.QueueByInfohash;
import org.springframework.data.cassandra.repository.AllowFiltering;
import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.Date;

@RepositoryRestResource
public interface QueueByInfohashRepository extends CassandraRepository<QueueByInfohash,String> {

    @AllowFiltering
    QueueByInfohash findFirstByDateLessThanAndRetryLessThan(Date date,int retry);

}