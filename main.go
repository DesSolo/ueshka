package main

import (
	"context"
	"flag"
	"time"
	"ueshka/storage"
	"ueshka/ueshka"
)

var (
	configFile string
)

func genUniqID(date string, o ueshka.Operation) string {
	return date + o.Time
}

func runDaemon(cfg *AppConfig, c *ueshka.Client, repo *storage.Memory) {
	b := context.Background()
	ctx, close := context.WithCancel(b)
	defer close()

	cfg.Logger.Info("collecting data every", cfg.CheckInterval)

	today := time.Now().Format("2006-01-02")

	for {
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(1 * time.Hour):
			today = time.Now().Format("2006-01-02")

		case <-time.Tick(cfg.CheckInterval):
			cfg.Logger.Debug("starting collect after:", cfg.CheckInterval)

			stats, err := c.GetDailyStat(cfg.Ueshka.PupilID, today, today)
			if err != nil {
				cfg.Logger.Warning("fault get statistic err:", err)
				continue
			}

			for date, operations := range stats {
				for _, op := range operations {
					uid := genUniqID(date, op)
					if repo.IsExist(uid) {
						continue
					}

					cfg.Logger.Info(
						"found new operaton date:", date,
						"uid:", uid,
					)

					msg := cfg.gate.RenderMessage(&op)
					if err := cfg.gate.Send(msg); err != nil {
						cfg.Logger.Info("fault send message err:", err)
					}
					repo.Add(uid)
				}
			}

			cfg.Logger.Debug("end collecting")
		}
	}
}

func init() {
	flag.StringVar(&configFile, "f", "", "config file")
}

func main() {
	flag.Parse()

	cfg := &AppConfig{}
	if configFile == "" {
		cfg = NewAppConfigFromEnv()
	} else {
		cfg = NewAppConfigFromFile(configFile)
	}

	client := ueshka.NewClient(cfg.Ueshka.VersionAPI, cfg.Ueshka.Token)

	repo := storage.NewMemory()

	runDaemon(cfg, client, repo)
}
