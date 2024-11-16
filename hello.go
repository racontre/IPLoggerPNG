package main

import (
	//"fmt"
	"image/png"
	"net/http"

	//"os"
	"log"
	"net"

	//"io"
	"github.com/gorilla/mux"

	"example/hello/utils"
)


func main() {
	db, err := utils.InitializeDB()
	if err != nil {
		log.Fatal("Error during db init: ", err)
	}
    r := mux.NewRouter()

    r.HandleFunc("/{page}.png", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        page := vars["page"]
		log.Print(page)
		host, _, _ := net.SplitHostPort(r.RemoteAddr) // move SplitHostPort to InsertDB
		utils.InsertIP(db, host) 

		/*img, err := os.Open("data.png")
    	if err != nil {
        	log.Fatal(err) // perhaps handle this nicer
    	}
    	defer img.Close()*/
		
    	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header

		//ips_test := [10]string{"127.0.0.1", "127.0.0.1", "127.0.0.1", "127.0.0.1", "127.0.0.1",
		//"127.0.0.1", "127.0.0.1", "127.0.0.1", "127.0.0.1", "127.0.0.1"}
		img := utils.GenerateImage(utils.GetIP(db))

    	//io.Copy(w, img)
		png.Encode(w, img);
    })

    http.ListenAndServe(":80", r)
}