package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("Enviroment")

	// Tenta carregar o .env apenas se estiver em desenvolvimento
	if env != "Development" && env != "" {
		log.Println("Rodando fora do ambiente local. Variáveis de sistema ativas.")
	} else {
		if err := godotenv.Load(); err != nil {
			log.Println("Aviso: Arquivo .env não encontrado. A aplicação usará as variáveis de sistema ou valores padrão.")
		} else {
			log.Println("Arquivo .env carregado com sucesso.")
		}
	}

	setDefaults()
}

func setDefaults() {

	defaults := map[string]string{
		"DB_USER":     "postgres",
		"DB_PASSWORD": "123456",
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_DBNAME":   "Gofre",
		"SECRET_KEY":  "devTest",
		"Enviroment":  "Development",
	}

	for key, defaultValue := range defaults {
		if os.Getenv(key) == "" {
			err := os.Setenv(key, defaultValue)
			if err != nil {
				log.Printf("Erro ao definir valor padrão para %s: %v", key, err)
			} else {
				log.Printf("Aviso: Variável %s não encontrada. Usando valor padrão.", key)
			}
		}
	}
}
