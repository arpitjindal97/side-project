package com.example.demo.controllers;

import org.libtorrent4j.SessionManager;
import org.springframework.scheduling.concurrent.ThreadPoolTaskScheduler;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/metrics")
public class Metrics {

    private final ThreadPoolTaskScheduler threadPoolTaskScheduler;
    private final SessionManager sessionManager;

    public Metrics(ThreadPoolTaskScheduler threadPoolTaskScheduler,
                   SessionManager sessionManager) {
        this.threadPoolTaskScheduler = threadPoolTaskScheduler;
        this.sessionManager = sessionManager;
    }

    @GetMapping(value = "/activeCount")
    public int metrics() {
        return this.threadPoolTaskScheduler.getActiveCount();
    }

    @GetMapping(value = "/dhtNodes")
    public long dhtNodes() {
        return sessionManager.stats().dhtNodes();
    }
}
