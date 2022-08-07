package com.example.demo.services;

import org.libtorrent4j.AlertListener;
import org.libtorrent4j.SessionManager;
import org.libtorrent4j.SessionParams;
import org.libtorrent4j.SettingsPack;
import org.libtorrent4j.alerts.Alert;
import org.libtorrent4j.alerts.AlertType;
import org.libtorrent4j.alerts.DhtGetPeersAlert;
import org.libtorrent4j.swig.settings_pack;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.server.ServerResponse;
import reactor.core.publisher.Mono;
import reactor.netty.ByteBufFlux;
import reactor.netty.http.client.HttpClient;

import java.util.Timer;
import java.util.TimerTask;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

@Service
public class DHTSession {

    static Logger logger = LoggerFactory.getLogger(DHTSession.class);

    @Bean
    public static SessionManager prepareSession() throws InterruptedException {
        SessionManager s = new SessionManager();
        //final SessionManager s = new SessionManager(true);


        s.addListener(new AlertListener() {
            @Override
            public int[] types() {
                return null;
            }

            @Override
            public void alert(Alert<?> alert) {
                if (alert.type() == AlertType.DHT_GET_PEERS) {
                    String id = ((DhtGetPeersAlert) alert).infoHash().toString();
                    //System.out.println("Infohash: " + id);
                    sendToAPIServer(id);
                }
                //System.out.println(alert);
            }
        });

        SettingsPack sp = new SettingsPack();
        sp.setEnableDht(true);
        //sp.listenInterfaces("0.0.0.0:43567");
        //sp.listenInterfaces("[::]:43567");
        //sp.listenInterfaces("0.0.0.0:43567,[::]:43567");
        //sp.setString(settings_pack.string_types.dht_bootstrap_nodes.swigValue(), "router.silotis.us:6881");
        //sp.setString(settings_pack.string_types.dht_bootstrap_nodes.swigValue(), "router.bittorrent.com:6881");
        //sp.setString(settings_pack.string_types.dht_bootstrap_nodes.swigValue(), "dht.transmissionbt.com:6881");
        sp.setString(settings_pack.bool_types.dht_extended_routing_table.swigValue(), "true");

        SessionParams params = new SessionParams(sp);

        s.start(params);

        final CountDownLatch signal = new CountDownLatch(1);

        final Timer timer = new Timer();
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                long nodes = s.stats().dhtNodes();
                // wait for at least 10 nodes in the DHT.
                if (nodes >= 10) {
                    System.out.println("DHT contains " + nodes + " nodes");
                    signal.countDown();
                    timer.cancel();
                }
            }
        }, 0, 1000);

        System.out.println("Waiting for nodes in DHT (10 seconds)...");
        boolean r = signal.await(40, TimeUnit.SECONDS);
        if (!r) {
            System.out.println("DHT bootstrap timeout");
            System.exit(0);
        }
        return s;
    }

    public static void sendToAPIServer(String infohash) {
        Thread commandLineThread = new Thread(() -> {
            String rawData = "{\"infohash\":\""+infohash+"\"}";
            //String encodedData = URLEncoder.encode( rawData, StandardCharsets.UTF_8);

            HttpClient client = HttpClient.create();
            client.post()
                    .uri("http://apiserver:8080"+"/torrents")
                    .send(ByteBufFlux.fromString(Mono.just(rawData)))
                    .responseContent()
                    .onErrorStop()
                    .flatMap(s -> ServerResponse.ok()
                            .contentType(MediaType.TEXT_PLAIN)
                            .bodyValue(s));

        });
        commandLineThread.setDaemon(true);
        commandLineThread.start();
    }

}
