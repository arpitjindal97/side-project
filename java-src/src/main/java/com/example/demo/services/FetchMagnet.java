package com.example.demo.services;

import com.example.demo.daos.QueueByInfohashRepository;
import com.example.demo.daos.TorrentByInfohashRepository;
import com.example.demo.daos.TorrentIndexRepository;
import com.example.demo.documents.Torrent;
import com.example.demo.entities.QueueByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import org.libtorrent4j.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.*;

@Service
public class FetchMagnet {

    final
    QueueByInfohashRepository queueByInfohashRepository;
    SessionManager sessionManager;
    TorrentByInfohashRepository torrentByInfohashRepository;
    TorrentIndexRepository torrentIndexRepository;

    Logger logger = LoggerFactory.getLogger(FetchMagnet.class);

    public FetchMagnet(QueueByInfohashRepository queueByInfohashRepository,
                       TorrentByInfohashRepository torrentByInfohashRepository,
                       TorrentIndexRepository torrentIndexRepository,
                       SessionManager s) {
        this.queueByInfohashRepository = queueByInfohashRepository;
        this.sessionManager = s;
        this.torrentByInfohashRepository = torrentByInfohashRepository;
        this.torrentIndexRepository = torrentIndexRepository;
    }

    /*
    @Scheduled(fixedRate = 1000*60*1)
    public void task() throws InterruptedException {
        Random random = new Random();
        int a = random.nextInt();
        System.out.println("Started :"+a);
        Thread.sleep(1000*60*2);
        System.out.println("Done "+a);
    }
     */

    @Scheduled(fixedRate = 1000*60*1)
    public void initTask() {

        Calendar calendar = Calendar.getInstance();
        calendar.setTime(new Date());
        calendar.add(Calendar.MINUTE, -5);

        QueueByInfohash queue = queueByInfohashRepository.findFirstByDateLessThanAndRetryLessThan(calendar.getTime(),5);
        if (queue == null) {
            //logger.info("Nothing in queue to process");
            return;
        }
        logger.info("Processing Infohash: "+queue.getInfoHash());
        queue.setDate(new Date());
        //queue.setRetry(queue.getRetry()+1);
        queueByInfohashRepository.save(queue);

        TorrentByInfohash torrent = fetchInfo(queue.getInfoHash());
        if (torrent == null) {
            queue.setRetry(queue.getRetry()+1);
            if (queue.getRetry() == 5) {
                queueByInfohashRepository.delete(queue);
            } else {
                queueByInfohashRepository.save(queue);
            }
            logger.error("Infohash "+queue.getInfoHash()+" failed");
            return;
        }
        logger.info("Successfully fetched "+queue.getInfoHash());

        TorrentByInfohash duplicate = this.torrentByInfohashRepository.findByInfohashEquals(torrent.getKey().getInfohash());

        if (duplicate != null) {
            logger.error("Infohash "+queue.getInfoHash()+" already exists on database");
        } else {
            torrentByInfohashRepository.save(torrent);
            queueByInfohashRepository.delete(queue);
            saveToElasticSearch(torrent);
            logger.info("Saved into Database");
        }
    }


    public TorrentByInfohash fetchInfo(String infohash) {

        String uri = "magnet:?xt=urn:btih:"+infohash+"&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce";

        //System.out.println("Fetching the magnet uri, please wait...");
        try {
            byte[] data = sessionManager.fetchMagnet(uri, 180, new File("/tmp"));

            TorrentByInfohash torrent = new TorrentByInfohash();

            if (data != null) {
                TorrentInfo torrentInfo = new TorrentInfo(data);
                torrentInfo.addTracker("udp://tracker.opentrackr.org:1337/announce",0);
                torrentInfo.addTracker("udp://tracker.openbittorrent.com:6969",1);
                torrentInfo.addTracker("udp://tracker.coppersurfer.tk:6969/announce",2);

                TorrentByInfohash.TorrentByInfohashKey key = new TorrentByInfohash.TorrentByInfohashKey();
                key.setInfohash(torrentInfo.infoHash().toString());
                key.setCategory("Others");
                key.setSubcategory("Others");
                torrent.setKey(key);
                torrent.setName(torrentInfo.name());
                torrent.setNumFiles(torrentInfo.numFiles());
                torrent.setSize(torrentInfo.totalSize()+"");
                torrent.setComment(torrentInfo.comment());
                torrent.setCreator(torrentInfo.creator());
                torrent.setSeeders(0);
                torrent.setLeechers(0);
                torrent.setPeers(0);
                torrent.setMagnet(torrentInfo.makeMagnetUri());
                torrent.setUserid("Anonymous");
                torrent.setDate(new Date());
                return torrent;
            }
        } catch(IllegalArgumentException e) {
            logger.error(e.getMessage() + infohash);
        }
        return null;
    }

    public void saveToElasticSearch(TorrentByInfohash torrentByInfohash) {
        Torrent document = new Torrent();

        document.setInfohash(torrentByInfohash.getKey().getInfohash());
        document.setCategory(torrentByInfohash.getKey().getCategory());
        document.setSubcategory(torrentByInfohash.getKey().getSubcategory());

        document.setComment(torrentByInfohash.getKey().getSubcategory());
        document.setCreator(torrentByInfohash.getKey().getSubcategory());
        document.setDate(torrentByInfohash.getDate());
        document.setLeechers(torrentByInfohash.getLeechers());
        document.setMagnet(torrentByInfohash.getMagnet());
        document.setName(torrentByInfohash.getName());
        document.setNumFiles(torrentByInfohash.getNumFiles());
        document.setPeers(torrentByInfohash.getPeers());
        document.setSeeders(torrentByInfohash.getSeeders());
        document.setSize(torrentByInfohash.getSize());
        document.setUserid(torrentByInfohash.getUserid());

        torrentIndexRepository.save(document);
    }
}