package com.example.demo.entities;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;
import org.springframework.data.cassandra.core.mapping.PrimaryKey;
import org.springframework.data.cassandra.core.mapping.Table;

import java.util.List;

@AllArgsConstructor
@Getter
@Setter
@Table
public class Todo {

    @PrimaryKey
    private String infoHash;

    private List<String> tasks;

}