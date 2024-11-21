package main

import (
	"image/png"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"example/hello/iplogger"
	"example/hello/utils"
)


func main() {
	db, err := iplogger.InitializeDB()
	if err != nil { log.Print("error initializing DB: ", err) }
	var service iplogger.IPLoggerService = iplogger.SqliteLoggerSerivce{Db: db}

	var geoip *utils.GeoIPParser;

	geoip, err = utils.NewGeoIPParser("GeoLite2-Country.mmdb")
	if err != nil { log.Println("Couldn't initialize the geoip service: ", err) }

	/*var ips []string*/
	
	//var service iplogger.IPLoggerService = &iplogger.InmemoryLoggerService{Ips: ips}

	r := mux.NewRouter()

	RegisterIPLoggerHandlers(r, service, geoip)

	http.ListenAndServe(":80", r)
}

func RegisterIPLoggerHandlers (router *mux.Router, service iplogger.IPLoggerService, geoip *utils.GeoIPParser) {
	router.HandleFunc("/{page}.png", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		log.Println(r.Header.Get("X-Real-Ip"))
		realip := strings.Fields(r.Header.Get("X-Real-Ip"))[1]
		//host, _, _ := net.SplitHostPort(r.Header.Get("X-Real-Ip"))
		err := service.InsertIP(realip)
		if err != nil { log.Println("Error while inserting new IP: ", err) }
		w.Header().Set("Content-Type", "image/png")
		
		list, err := service.GetIPList(10)
		if err != nil { log.Println("Error while getting new IP: ", err) }

		img := utils.GenerateImage(list, geoip)
		png.Encode(w, img);
	})
}