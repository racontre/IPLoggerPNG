package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/oschwald/geoip2-golang"
)

type Datas struct {
	Status string 		`json:"status"`
	CountryCode string 	`json:"countryCode"`
	City string 		`json:"city"`
}

type GeoIPParser struct {
	db *geoip2.Reader
}

func NewGeoIPParser(path string) (*GeoIPParser, error) {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {return nil, err}
	parser := &GeoIPParser{
		db: db,
	}
	return parser, nil
}

func (parser *GeoIPParser) GetCountry_DB(ip string) (string, error) {
	ipNet := net.ParseIP(ip)
	record, err := parser.db.City(ipNet)
	if err != nil {return "Unknown", err}
	return record.Country.IsoCode, nil
}

func GetCountry_API(ip string) (string, error) {
	requestURL := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,countryCode,city", ip)
	res, err := http.Get(requestURL)
	if err != nil {return "Unknown", err}
	body, err := io.ReadAll(res.Body)
	if err != nil {return "Unknown", err}
	bodyString := string(body);
	log.Println(requestURL, bodyString)
	var data Datas	
	json.Unmarshal([]byte(bodyString), &data)
	if data.Status == "fail" {return "Unknown", nil}
	return data.CountryCode, nil
}