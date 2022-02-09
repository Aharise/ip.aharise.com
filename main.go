package main

import (
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getRemoteAddress(request *http.Request) string {
	if xRealIp := request.Header.Get("X-Real-IP"); 0 < len(xRealIp) {
		return xRealIp
	}

	return request.RemoteAddr
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		remoteAddress := getRemoteAddress(request)
		log.Print(remoteAddress)
		_, _ = w.Write([]byte("\n" + remoteAddress + "\n"))
	})

	port := getEnv("PORT", "21110")

	log.Println("http://localhost:" + port)
	_ = http.ListenAndServe(":"+port, nil)
}
