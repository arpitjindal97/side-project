package com.example.demo;

import com.example.demo.daos.TorrentIndexRepository;
import com.example.demo.documents.Torrent;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;

import java.util.Date;

@Component
class ElasticRepositoryTest {

	@Autowired
	TorrentIndexRepository torrentIndexRepository;

	@Bean
	ElasticRepositoryTest getInstance() {
		return new ElasticRepositoryTest();
	}

	void crudOperationElastic() {
		Torrent torrent = new Torrent();
		torrent.setInfohash("abcd");
		torrent.setCategory("Others");
		torrent.setSubcategory("Others");
		torrent.setName("Sample Name");
		torrent.setNumFiles(15);
		torrent.setSize(1024L);
		torrent.setComment("Test Comment");
		torrent.setCreator("Anonymous");
		torrent.setSeeders(0);
		torrent.setLeechers(0);
		torrent.setPeers(0);
		torrent.setMagnet("magnet:blah&blah");
		torrent.setUserid("Anonymous");
		torrent.setDate(new Date());
		torrentIndexRepository.save(torrent);
	}

}
