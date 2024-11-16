package utils

import (
	"database/sql"
	"encoding/binary"
	"errors"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() (*sql.DB, error) {
	const file string = "visits.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	const create string = `
		CREATE TABLE IF NOT EXISTS IP (
		id INTEGER NOT NULL PRIMARY KEY,	
		IPAddress INTEGER
		);`
	
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return db, nil
}

// selecting 10 latest ips
func GetIP(db *sql.DB) [10]string {
	row, err := db.Query("SELECT * FROM IP ORDER BY id DESC LIMIT 10")
	if err != nil {
		log.Fatal("Error during query: ", err)
	}
	defer row.Close()

	var data [10]string
	index := 0
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var IPAddress uint32
		row.Scan(&id, &IPAddress)
		address := Long2ip(IPAddress)
		data[index] = address
		index++
	}
	return data
}

func InsertIP(db *sql.DB, address string) {
	log.Println("Inserting IP record ...")
	insertIPSQL := `INSERT INTO IP(IPAddress) VALUES (?)`
	statement, err := db.Prepare(insertIPSQL) // Prepare statement. 
                                                   // This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	addressLong, err := Ip2long(address)
	if err != nil {
		log.Println("Failed to convert ", address, " to longint: ", err)
	}

	_, err = statement.Exec(addressLong)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// the IP is stored as a long int
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