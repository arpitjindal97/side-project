package com.example.demo.utils;

import com.example.demo.documents.Torrent;
import com.example.demo.entities.FilesByInfohash;
import com.example.demo.entities.TorrentByInfohash;
import org.libtorrent4j.TorrentInfo;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Date;

public class Convert {

    public static TorrentByInfohash getTorrentByInfohash(TorrentInfo torrentInfo) {
        TorrentByInfohash torrent = new TorrentByInfohash();

        torrentInfo.addTracker("udp://tracker.opentrackr.org:1337/announce",0);
        torrentInfo.addTracker("udp://tracker.openbittorrent.com:6969",1);
        torrentInfo.addTracker("udp://tracker.coppersurfer.tk:6969/announce",2);

        torrent.setInfohash(torrentInfo.infoHash().toString());
        torrent.setCategory("Others");
        torrent.setSubcategory("Others");
        torrent.setName(torrentInfo.name());
        torrent.setNumFiles(torrentInfo.numFiles());
        torrent.setSize(torrentInfo.totalSize());
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

    public static FilesByInfohash getFilesByInfohash(TorrentInfo torrentInfo){
        FilesByInfohash files = new FilesByInfohash();
        files.setInfohash(torrentInfo.infoHash().toString());
        String[] filePath = new String[torrentInfo.numFiles()];
        Long[] size = new Long[torrentInfo.numFiles()];
        for(int x=0; x<torrentInfo.numFiles();x++) {
            filePath[x] = torrentInfo.files().filePath(x);
            size[x] = torrentInfo.files().fileSize(x);
        }
        files.setFilepath(Arrays.asList(filePath));
        files.setSize(Arrays.asList(size));
        return files;
    }

    public static Torrent getTorrentDocument(TorrentInfo torrentInfo) {
        TorrentByInfohash torrentByInfohash = getTorrentByInfohash(torrentInfo);
        Torrent document = new Torrent();

        document.setInfohash(torrentByInfohash.getInfohash());
        document.setCategory(torrentByInfohash.getCategory());
        document.setSubcategory(torrentByInfohash.getSubcategory());
        document.setComment(torrentByInfohash.getComment());
        document.setCreator(torrentByInfohash.getCreator());
        document.setDate(torrentByInfohash.getDate());
        document.setLeechers(torrentByInfohash.getLeechers());
        document.setMagnet(torrentByInfohash.getMagnet());
        document.setName(torrentByInfohash.getName());
        document.setNumFiles(torrentByInfohash.getNumFiles());
        document.setPeers(torrentByInfohash.getPeers());
        document.setSeeders(torrentByInfohash.getSeeders());
        document.setSize(torrentByInfohash.getSize());
        document.setUserid(torrentByInfohash.getUserid());
        return document;
    }

}
