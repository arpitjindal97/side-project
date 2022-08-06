package com.example.demo.daos;

import com.example.demo.documents.Torrent;
import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

public interface TorrentIndexRepository extends ElasticsearchRepository<Torrent, String> {
}
