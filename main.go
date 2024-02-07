package  main

import (
	"log"
	"net/http"
	"time"
)
func rootPage (w http.ResponseWriter, r * http.Request){
	w.Write([]byte("Hello World!"))
}
func main(){
	http.HandleFunc("/", rootPage)

	server := http.Server{
		Addr: ":8000",
		WriteTimeout: time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		IdleTimeout: time.Second * 60,
	}
	log.Println("The server is listening on port: 8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}