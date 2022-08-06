package com.example.demo;

import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.Session;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.elasticsearch.client.elc.ElasticsearchTemplate;
import org.testcontainers.containers.CassandraContainer;
import org.testcontainers.containers.wait.strategy.Wait;
import org.testcontainers.elasticsearch.ElasticsearchContainer;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
@Testcontainers
class DemoApplicationTests {

	@Container
	public static CassandraContainer cassandra = (CassandraContainer) new CassandraContainer("cassandra:4.0.5")
			.withExposedPorts(9042);

	@Container
	private static ElasticsearchContainer esContainer = new ElasticsearchContainer("docker.elastic.co/elasticsearch/elasticsearch:8.3.3")
			.withExposedPorts(9200)
			.withEnv("ELASTIC_PASSWORD","password")
			.withEnv("xpack.security.http.ssl.enabled","false")
			.withEnv("discovery.type","single-node");

	@BeforeAll
	static void setupCassandraConnectionProperties() {
		System.setProperty("spring.data.cassandra.keyspace-name", "awesome");
		System.setProperty("spring.data.cassandra.contact-points", cassandra.getHost());
		System.setProperty("spring.data.cassandra.port", String.valueOf(cassandra.getMappedPort(9042)));
		createKeyspace(cassandra.getCluster());

		esContainer.start();
		System.setProperty("spring.elasticsearch.uris",esContainer.getHttpHostAddress());
	}

	private static void createKeyspace(Cluster cluster) {
		String query = "CREATE KEYSPACE IF NOT EXISTS awesome with REPLICATION = {'class':'SimpleStrategy','replication_factor':'1'};";
		try (Session session = cluster.connect()) {
			session.execute(query);
		}
	}

	@Test
	void givenCassandraContainer_whenSpringContextIsBootstrapped_thenContainerIsRunningWithNoExceptions() {
		assertTrue(cassandra.isRunning());
		assertTrue(esContainer.isRunning());
	}

	@Test
	void contextLoads() {
	}

}
