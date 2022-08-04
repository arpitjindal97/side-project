package com.example.demo.entities;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.cassandra.core.mapping.PrimaryKey;
import org.springframework.data.cassandra.core.mapping.Table;

import java.util.Date;

@AllArgsConstructor
@Getter
@Setter
@Table(value = "queue_by_infohash")
public class QueueByInfohash {

    @PrimaryKey
    private String infoHash;

    private Date date;

    private Integer retry;

}