package com.example.demo.services;

import com.example.demo.daos.FilesByInfohashRepository;
import com.example.demo.daos.QueueByInfohashRepository;
import com.example.demo.daos.TorrentByInfohashRepository;
import com.example.demo.daos.TorrentIndexRepository;
import com.example.demo.documents.Torrent;
import com.example.demo.entities.QueueByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import com.example.demo.utils.Convert;
import org.libtorrent4j.SessionManager;
import org.libtorrent4j.TorrentInfo;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.Calendar;
import java.util.Date;
import java.util.Optional;

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

    @Scheduled(fixedRate = 1000*60)
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
        queueByInfohashRepository.save(queue);

        TorrentInfo torrentInfo = fetchInfo(queue.getInfoHash());
        if (torrentInfo == null) {
            queue.setRetry(queue.getRetry()+1);
            queueByInfohashRepository.save(queue);
            logger.error("Infohash "+queue.getInfoHash()+" failed");
            return;
        }
        logger.info("Successfully fetched "+queue.getInfoHash());

        TorrentByInfohash duplicate = torrentByInfohashRepository.
                findByInfohashEquals(torrentInfo.infoHash().toString());
        Optional<Torrent> doc = torrentIndexRepository.findById(torrentInfo.infoHash().toString());
        if (duplicate != null && doc.isPresent()) {
            logger.error("Infohash "+queue.getInfoHash()+" already exists on database");
            queueByInfohashRepository.delete(queue);
        } else {
            torrentByInfohashRepository.save(Convert.getTorrentByInfohash(torrentInfo));
            filesByInfohashRepository.save(Convert.getFilesByInfohash(torrentInfo));
            torrentIndexRepository.save(Convert.getTorrentDocument(torrentInfo));
            queueByInfohashRepository.delete(queue);
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