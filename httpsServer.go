package main
import (
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("i-aor.com")))
	log.Println("Server starting")
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
