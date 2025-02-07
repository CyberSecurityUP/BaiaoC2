package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// GetOS retorna o sistema operacional atual
func GetOS() string {
	return runtime.GOOS
}

// GetArch retorna a arquitetura do sistema
func GetArch() string {
	return runtime.GOARCH
}

// ExecuteCommand executa um comando no sistema
func ExecuteCommand(command string) (string, error) {
	var cmd *exec.Cmd
	if GetOS() == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// IsAdmin verifica se o processo está sendo executado como administrador
func IsAdmin() bool {
	if GetOS() == "windows" {
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		return err == nil
	}
	return os.Geteuid() == 0
}
