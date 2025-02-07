package teamserver

import (
	"fmt"
	"log"
	"net"
)

// Start inicia todos os listeners do C2
func Start(profile C2Profile) {
	log.Println("[+] Iniciando BaiaoC2 Teamserver...")

	// Iniciar listeners configurados
	for _, listener := range profile.Listeners {
		go StartListener(listener)
	}

	// Manter o servidor rodando
	select {}
}

// StartListener inicia um listener individual
func StartListener(listener Listener) {
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
	}
}
