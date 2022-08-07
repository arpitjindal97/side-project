package com.example.demo.entities;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.cassandra.core.cql.Ordering;
import org.springframework.data.cassandra.core.cql.PrimaryKeyType;
import org.springframework.data.cassandra.core.mapping.*;

import java.io.Serializable;
import java.util.Date;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
@Table(value = "torrent_by_infohash")
public class TorrentByInfohash {

    @PrimaryKey
    TorrentByInfohashKey key;

    String name;
    Date date;
    Integer peers;
    Integer seeders;
    Integer leechers;
    Long size;
    String magnet;
    String userid;
    String comment;
    String creator;

    @Column(value="num_files")
    Integer numFiles;

    @PrimaryKeyClass
    @Getter
    @Setter
    public static class TorrentByInfohashKey implements Serializable {

        @PrimaryKeyColumn(type = PrimaryKeyType.PARTITIONED,ordinal = 0)
        String infohash;

        @PrimaryKeyColumn(type = PrimaryKeyType.CLUSTERED,ordinal = 1,ordering = Ordering.ASCENDING)
        String category;

        @PrimaryKeyColumn(type = PrimaryKeyType.CLUSTERED, ordinal = 2,ordering = Ordering.ASCENDING)
        String subcategory;
    }
}
