package main
import (
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/dekcom/Repository/website/i-aor.com")))
	log.Println("Server starting")
	err := http.ListenAndServeTLS(":443", "/home/dekcom/Repository/HttpsServ/cert.pem", "/home/dekcom/Repository/HttpsServ/key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
