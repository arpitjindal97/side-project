package cassandra

const create_keyspace = `create keyspace if not exists awesome with replication = {'class':'SimpleStrategy', 'replication_factor':1};`

const create_torrent_by_infohash = `CREATE TABLE IF NOT EXISTS torrent_by_infohash (
    infohash text,
    category int,
	comment text,
	creator text,
    date timestamp,
    leechers int,
    magnet text,
    name text,
	num_files int,
    peers int,
    seeders int,
    size bigint,
    userid text,
	PRIMARY KEY (infohash));`

const create_files_by_infohash = `CREATE TABLE IF NOT EXISTS files_by_infohash (
	infohash text,
    filepath list<text>,
    size list<bigint>,
	PRIMARY KEY (infohash));`

const create_queue_by_infohash = `CREATE TABLE IF NOT EXISTS queue_by_infohash (
	infohash text,
	date timestamp,
	retry int,
	PRIMARY KEY (infohash) ); `

const create_suggestion_by_userid = `CREATE TABLE IF NOT EXISTS suggestion_by_userid (
	infohash text,
	userid   text,
	date     timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY ((userid), date) ) WITH CLUSTERING ORDER BY (date DESC); `

const create_suggestion_by_infohash = `CREATE TABLE IF NOT EXISTS suggestion_by_infohash (
	infohash text,
	userid   text,
	date     timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY (infohash) ); `

const create_suggestion_by_date = `CREATE TABLE IF NOT EXISTS suggestion_by_date (
	infohash text,
	userid   text,
	date     timestamp,
	old_json text,
	new_json text,
	PRIMARY KEY (date)); `

const create_user_by_userid = `CREATE TABLE IF NOT EXISTS user_by_userid (
	userid   text,
	password text,
	email    text,
	PRIMARY KEY (userid) ); `

const find_torrent_by_infohash = `SELECT * FROM torrent_by_infohash where infohash = ?`

const insert_queue_by_infohash = `INSERT INTO queue_by_infohash (infohash, date, retry) values(?, ?, ?)`

const find_queue_by_infohash = `SELECT * FROM queue_by_infohash where infohash = ?`

const update_torrent_by_infohash = `UPDATE torrent_by_infohash set peers=?, seeders=?, leechers=? where infohash=?`

const find_files_by_infohash = `SELECT * FROM files_by_infohash where infohash = ?`

const delete_files_by_infohash = `DELETE from files_by_infohash where infohash = ?`

const delete_torrent_by_infohash = `DELETE from torrent_by_infohash where infohash = ?`
