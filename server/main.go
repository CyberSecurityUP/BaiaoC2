package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"baiaoC2/server/teamserver"
)

// C2Profile define a estrutura do perfil do C2
type C2Profile struct {
	Name      string     `json:"name"`
	Version   string     `json:"version"`
	Listeners []teamserver.Listener `json:"listeners"`
	Implants  Implants   `json:"implants"`
	Logging   Logging    `json:"logging"`
}

// Implants define a configuração dos implantes
type Implants struct {
	SleepTime     int     `json:"sleep_time"`
	Jitter        float64 `json:"jitter"`
	MaxRetries    int     `json:"max_retries"`
}

// Logging define os logs do teamserver
type Logging struct {
	Level string `json:"level"`
	File  string `json:"file"`
}

// Carrega o perfil do C2 do JSON
func loadProfile(path string) (*C2Profile, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profile C2Profile
	err = json.Unmarshal(file, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// Função principal do teamserver
func main() {
	profile, err := loadProfile("config/c2profile.json")
	if err != nil {
		log.Fatalf("Erro ao carregar perfil C2: %v", err)
	}

	fmt.Println("🔥 BaiaoC2 Teamserver iniciado...")
	teamserver.Start(profile)
}
