package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	contador int
	mu       sync.Mutex
)

func main() {
	// Cargar el contador desde un almacenamiento persistente (opcional)
	contador = 0

	// Iniciar una goroutine para aumentar el contador cada 6 horas
	go func() {
		ticker := time.NewTicker(6 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			mu.Lock()
			contador++
			mu.Unlock()
			fmt.Println("Contador incrementado:", contador)
		}
	}()

	http.HandleFunc("/contador", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		// Responder con el valor del contador en JSON
		response := map[string]int{"contador": contador}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Servidor en ejecución en el puerto 9090...")
	http.ListenAndServe(":9090", nil)
}
