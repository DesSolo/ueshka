package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"ueshka/gate"
	"ueshka/logging"
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
	Logger        *logging.Logger

	gate gate.Sender
}

// NewAppConfigFromFile ...
func NewAppConfigFromFile(file string) *AppConfig {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("fault read config file err:", err)
	}

	for _, line := range strings.Split(string(body), "\n") {
		l := strings.Split(line, "=")
		if len(l) != 2 {
			continue
		}
		os.Setenv(l[0], strings.Trim(l[1], "\""))
	}

	return NewAppConfigFromEnv()
}

// NewAppConfigFromEnv ...
func NewAppConfigFromEnv() *AppConfig {
	cfg := &AppConfig{}

	cfg.Logger = &logging.Logger{}

	ll := getEnvOrString("LOG_LEVEL", "info")

	switch ll {
	case "info":
		cfg.Logger.Level = logging.Info
	case "debug":
		cfg.Logger.Level = logging.Debug
	default:
		log.Fatalf("logging type \"%s\" not supported", ll)
	}

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
