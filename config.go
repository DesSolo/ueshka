package main

import (
	"log"
	"os"
	"time"
	"ueshka/gate"
)

// AppConfig ...
type AppConfig struct {
	Ueshka struct {
		VersionAPI string
		Token      string
		PupilID    string
	}
	GateType      string
	CheckInterval time.Duration

	gate gate.Sender
}

// NewAppConfigFromEnv ...
func NewAppConfigFromEnv() *AppConfig {
	cfg := &AppConfig{}
	cfg.Ueshka.VersionAPI = getEnvOrString("UESHKA_API_VERSION", "LK/1.8.12")
	cfg.Ueshka.Token = getEnvOrError("UESHKA_TOKEN")
	cfg.Ueshka.PupilID = getEnvOrError("UESHKA_PUPIL_ID")

	gateType := getEnvOrError("GATE_TYPE")

	switch gateType {
	case "telegram":
		cfg.gate = gate.NewTelegram(
			getEnvOrError("TELEGRAM_TOKEN"),
			getEnvOrError("TELEGRAM_CHAT_ID"),
		)
	default:
		log.Fatalf("gate type \"%s\" not supported", gateType)
	}

	cfg.CheckInterval = time.Duration(getEnvOrInt("CHECK_INTERVAL", 10)) * time.Second

	return cfg
}

func getEnvOrError(value string) string {
	v := os.Getenv(value)
	if v == "" {
		log.Fatalf("env value \"%s\" not set", value)
	}
	return v
}

func getEnvOrString(value, blank string) string {
	v := os.Getenv(value)
	if v == "" {
		return blank
	}
	return v
}

func getEnvOrInt(value string, blank int) int {
	v := os.Getenv(value)
	if v == "" {
		return blank
	}
	return 2
}
