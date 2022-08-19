package com.example.demo.services;

import com.example.demo.daos.FilesByInfohashRepository;
import com.example.demo.daos.QueueByInfohashRepository;
import com.example.demo.daos.TorrentByInfohashRepository;
import com.example.demo.daos.TorrentIndexRepository;
import com.example.demo.documents.Torrent;
import com.example.demo.entities.FilesByInfohash;
import com.example.demo.entities.QueueByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import com.example.demo.utils.Convert;
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
    FilesByInfohashRepository filesByInfohashRepository;

    Logger logger = LoggerFactory.getLogger(FetchMagnet.class);

    public FetchMagnet(QueueByInfohashRepository queueByInfohashRepository,
                       TorrentByInfohashRepository torrentByInfohashRepository,
                       TorrentIndexRepository torrentIndexRepository,
                       FilesByInfohashRepository filesByInfohashRepository,
                       SessionManager s) {
        this.queueByInfohashRepository = queueByInfohashRepository;
        this.sessionManager = s;
        this.torrentByInfohashRepository = torrentByInfohashRepository;
        this.torrentIndexRepository = torrentIndexRepository;
        this.filesByInfohashRepository = filesByInfohashRepository;
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

        TorrentInfo torrentInfo = fetchInfo(queue.getInfoHash());
        if (torrentInfo == null) {
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

        TorrentByInfohash duplicate = this.torrentByInfohashRepository.findByInfohashEquals(torrentInfo.infoHash().toString());

        if (duplicate != null) {
            logger.error("Infohash "+queue.getInfoHash()+" already exists on database");
            queueByInfohashRepository.delete(queue);
        } else {
            torrentByInfohashRepository.save(Convert.getTorrentByInfohash(torrentInfo));
            filesByInfohashRepository.save(Convert.getFilesByInfohash(torrentInfo));
            queueByInfohashRepository.delete(queue);
            torrentIndexRepository.save(Convert.getTorrentDocument(torrentInfo));
            logger.info("Saved into Database");
        }
    }

    public TorrentInfo fetchInfo(String infohash) {

        String uri = "magnet:?xt=urn:btih:"+infohash+"&tr=udp%3a%2f%2ftracker.opentrackr.org%3a1337%2fannounce";

        //System.out.println("Fetching the magnet uri, please wait...");
        try {
            byte[] data = sessionManager.fetchMagnet(uri, 180, new File("/tmp"));

            if (data != null) {
                return new TorrentInfo(data);
            }
        } catch(IllegalArgumentException e) {
            logger.error(e.getMessage() + infohash);
        }
        return null;
    }

}