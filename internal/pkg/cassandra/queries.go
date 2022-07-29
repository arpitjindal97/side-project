package cassandra

const create_keyspace = `create keyspace if not exists awesome with replication = {'class':'SimpleStrategy', 'replication_factor':1};`

const create_torrent_by_infohash = `CREATE TABLE IF NOT EXISTS torrent_by_infohash (
    infohash text,
    date timestamp,
    category text,
    peers int,
    seeders int,
    leechers int,
    size text,
    subcategory text,
    title text,
    magnet text,
    userid text,
	PRIMARY KEY ((infohash), category, subcategory)
	) WITH CLUSTERING ORDER BY (category ASC, subcategory ASC);`

const create_queue_by_infohash = `CREATE TABLE IF NOT EXISTS queue_by_infohash (
    date timestamp,
    infohash text,
    retry int,
	PRIMARY KEY (infohash));`

const create_suggestion_by_userid = `CREATE TABLE IF NOT EXISTS suggestion_by_userid (
    infohash text,
    userid   text,
	date     timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY ((userid), date) ) WITH CLUSTERING ORDER BY (date DESC);`

const create_suggestion_by_infohash = `CREATE TABLE IF NOT EXISTS suggestion_by_infohash (
    infohash text,
    userid   text,
	date     timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY (infohash) );`

const create_suggestion_by_date = `CREATE TABLE IF NOT EXISTS suggestion_by_date (
    infohash text,
    userid   text,
	date	 timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY (date));`

const create_user_by_userid = `CREATE TABLE IF NOT EXISTS user_by_userid (
    userid   text,
	password text,
	email    text,
	PRIMARY KEY (userid) );`

const find_torrent_by_infohash = `SELECT * FROM torrent_by_infohash where infohash = ?`

const insert_queue_by_infohash = `INSERT INTO queue_by_infohash (infohash,date,retry) values(?,?,?)`

const find_queue_by_infohash = `SELECT * FROM queue_by_infohash (infohash,data,retry) values(?,?,?)`
