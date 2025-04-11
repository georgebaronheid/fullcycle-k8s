package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {

	fs := http.FileServer(http.Dir("content"))
	http.Handle("/content/", http.StripPrefix("/content/", fs))

	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secrets", Secrets)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt).Seconds()

	if duration < 10 { // it'll take 10s for my application to start
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Health check failed after %v seconds", duration)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Health check passed after %v seconds", duration)
}

func Secrets(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	pwd := os.Getenv("PWD")

	fmt.Fprintf(w, "User: %s, Password: %s", user, pwd)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	fmt.Fprintf(w, "My family: %s", string(data))
}

func Hello(w http.ResponseWriter, r *http.Request) {

	_ = os.Getenv("NAME")
	_ = os.Getenv("AGE")

	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Baronheid</title>
	<link rel="shortcut icon" href="content/fav-icon.png" type="image/png">
</head>
<body>
    <div style="text-align: center;">
        <img src="content/em-dev-branco.png" alt="Em desenvolvimento">
    </div>
</body>
</html>
`) // w.Write([]byte("<h1> Hello, FullCycle!!! </h1>"))
}
