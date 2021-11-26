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
	xRealIp := request.Header.Get("X-Real-IP")
	if len(xRealIp) > 0 {
		return xRealIp
	}

	remoteAddr := request.RemoteAddr
	return remoteAddr
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		remoteAddress := getRemoteAddress(request)
		log.Print(remoteAddress)
		_, _ = w.Write([]byte(remoteAddress))
	})

	port := getEnv("PORT", "21110")

	log.Println("http://localhost:" + port)
	_ = http.ListenAndServe(":"+port, nil)
}
