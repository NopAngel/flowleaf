package main

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer(sourceFile string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sourceCode, err := os.ReadFile(sourceFile)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error for the read file: %s", err), http.StatusInternalServerError)
			return
		}

		tokens := Tokenizer(string(sourceCode))

		htmlCode, err := ParseTokens(tokens)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error for the compile HTML: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(htmlCode))
		if err != nil {
			fmt.Println("Error al escribir la respuesta:", err)
		}
	})

	fmt.Println("Server ruuning in: http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error: An error occurred while starting the server. Verify that port 3000 is free on your computer, or that permissions and authorization to access the local network have been granted.", err)
	}
}
