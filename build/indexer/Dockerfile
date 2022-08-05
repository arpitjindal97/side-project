FROM eclipse-temurin:latest

WORKDIR /work
COPY java-src/build/libs/*SHOT.jar indexer.jar
COPY java-src/lib/x86_64/libtorrent4j.so libtorrent4j.so
EXPOSE 8082

CMD ["java","-Djava.library.path=/work","-jar","indexer.jar"]
