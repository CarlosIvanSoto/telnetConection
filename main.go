package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Dirección IP y puerto del servidor Telnet
	server := "galesdtapps.ddns.net"
	port := "2638"

	// Número máximo de intentos
	maxAttempts := 5

	// Intentar conexión Telnet hasta que tenga éxito o se alcance el número máximo de intentos
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("Intento de conexión #%d\n", attempt)

		// Conexión TCP al servidor Telnet
		conn, err := net.Dial("tcp", server+":"+port)
		if err != nil {
			fmt.Println("Error de conexión:", err)
			fmt.Println("Intentando nuevamente en 5 segundos...")
			time.Sleep(5 * time.Second)
			continue
		}
		defer conn.Close()

		// Si la conexión fue exitosa, terminar el programa
		fmt.Println("Conexión exitosa.")
		return
	}

	fmt.Println("No se pudo establecer una conexión Telnet después de", maxAttempts, "intentos.")
}
