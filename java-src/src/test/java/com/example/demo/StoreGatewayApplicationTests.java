package com.example.demo;

import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.Session;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.testcontainers.containers.CassandraContainer;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

import static org.junit.Assert.*;

@SpringBootTest
@Testcontainers
class StoreGatewayApplicationTests {

	@Container
	public static final CassandraContainer cassandra
			= (CassandraContainer) new CassandraContainer("cassandra:4.0.5").withExposedPorts(9042);

	@BeforeAll
	static void setupCassandraConnectionProperties() {
		System.setProperty("spring.data.cassandra.keyspace-name", "testKeySpace");
		System.setProperty("spring.data.cassandra.contact-points", cassandra.getHost());
		System.setProperty("spring.data.cassandra.port", String.valueOf(cassandra.getMappedPort(9042)));
		createKeyspace(cassandra.getCluster());
	}

	private static void createKeyspace(Cluster cluster) {
		String query = "CREATE KEYSPACE IF NOT EXISTS testkeyspace with REPLICATION = {'class':'SimpleStrategy','replication_factor':'1'};";
		try (Session session = cluster.connect()) {
			session.execute(query);
		}
	}

	@Test
	void givenCassandraContainer_whenSpringContextIsBootstrapped_thenContainerIsRunningWithNoExceptions() {
		assertTrue(cassandra.isRunning());
	}

	@Test
	void contextLoads() {
	}

}
