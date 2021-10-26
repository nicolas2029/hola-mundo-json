package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func holaMundo(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hola Mundo"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/api/v1/hola", holaMundo)
	fmt.Println("Server iniciado en el puerto 3000")
	http.ListenAndServe(":3000", nil)
}
