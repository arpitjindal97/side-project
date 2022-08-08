package com.example.demo.entities;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.cassandra.core.mapping.*;

import java.util.Date;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
@Table(value = "torrent_by_infohash")
public class TorrentByInfohash {

    @PrimaryKey
    String infohash;

    String category;
    String comment;
    String creator;
    Date date;
    Integer leechers;
    String magnet;
    String name;

    @Column(value="num_files")
    Integer numFiles;

    Integer peers;
    Integer seeders;
    Long size;
    String subcategory;
    String userid;

}
