package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"baiaoC2/client/gui"
	"baiaoC2/server/teamserver"
)

// Carrega o perfil do C2 do JSON
func loadProfile(path string) (*teamserver.C2Profile, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var profile teamserver.C2Profile
	err = json.Unmarshal(file, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func main() {
	mode := flag.String("mode", "client", "Mode to run the application (client or server)")
	flag.Parse()

	if *mode == "server" {
		// Carregar o perfil antes de iniciar o servidor
		profile, err := loadProfile("config/c2profile.json")
		if err != nil {
			log.Fatalf("Erro ao carregar perfil C2: %v", err)
		}

		fmt.Println("🔥 BaiaoC2 Teamserver iniciado...")
		teamserver.Start(*profile)
	} else {
		gui.Start()
	}
}
