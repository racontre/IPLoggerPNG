package iplogger

import (
	"database/sql"
	"encoding/binary"
	"errors"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteLoggerSerivce struct {
	Db *sql.DB
}

// Creates necessary tables if they don't exist (maybe force create visits.db if it doesn't exist)
func InitializeDB() (*sql.DB, error) {
	const file string = "visits.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil { return nil, err }

	// Perhaps add a datetime field eventually?
	const create string = `
		CREATE TABLE IF NOT EXISTS IP (
		id INTEGER NOT NULL PRIMARY KEY,	
		IPAddress INTEGER
		);`
	
	if _, err := db.Exec(create); err != nil { return nil, err }

	return db, nil
}

func (s SqliteLoggerSerivce) InsertIP(ip string) error {
	if ip == "127.0.0.1" {
		log.Println("Did not add 127.0.0.1 to IP list")
		return nil
	}

	log.Println("Inserting IP record ...")
	insertIPSQL := `INSERT INTO IP(IPAddress) VALUES (?)`
	statement, err := s.Db.Prepare(insertIPSQL)
                                                
	if err != nil { return err }

	addressLong, err := Ip2long(ip)
	if err != nil {	return err }

	_, err = statement.Exec(addressLong)
	if err != nil {	return err }

	return nil
}

func (s SqliteLoggerSerivce) GetIPList(num int) ([]string, error) {
	row, err := s.Db.Query("SELECT * FROM IP ORDER BY id DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var data []string
	for row.Next() {
		var id int
		var IPAddress uint32
		row.Scan(&id, &IPAddress)
		address := Long2ip(IPAddress)
		data = append(data, address)
	}
	return data, nil
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func Ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong IP Address format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}