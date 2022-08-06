package com.example.demo.documents;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.data.annotation.Id;
import org.springframework.data.elasticsearch.annotations.Document;
import org.springframework.data.elasticsearch.annotations.Field;
import org.springframework.data.elasticsearch.annotations.FieldType;

import java.util.Date;

@Document(indexName = "torrents")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class Torrent {

    @Id
    private String infohash;

    @Field(type = FieldType.Keyword)
    private String category;

    @Field(type = FieldType.Keyword)
    private String subcategory;

    @Field(type = FieldType.Text)
    private String comment;

    @Field(type = FieldType.Text)
    private String creator;

    @Field(type = FieldType.Date)
    private Date date;

    @Field(type = FieldType.Integer)
    private Integer leechers;

    @Field(type = FieldType.Text)
    private String magnet;

    @Field(type = FieldType.Text)
    private String name;

    @Field(type = FieldType.Integer, name = "num_files")
    private Integer numFiles;

    @Field(type = FieldType.Integer)
    private Integer peers;

    @Field(type = FieldType.Integer)
    private Integer seeders;

    @Field(type = FieldType.Long)
    private String size;

    @Field(type = FieldType.Keyword)
    private String userid;

}