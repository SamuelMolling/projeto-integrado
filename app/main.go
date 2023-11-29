package main

import "net/http"

// Path: projeto-integrado/app/main.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem vindo a LumenAI! 💡"))
	})

	http.ListenAndServe(":8000", nil)
}
