package main

import (
	"log"
	"time"
	"ueshka/storage"
	"ueshka/ueshka"
)

var (
	today = time.Now().Format("2006-01-02")
)

func genUniqID(date string, o ueshka.Operation) string {
	return date + o.Time
}

func runDaemon(d time.Duration, c *ueshka.Client, cfg *AppConfig, repo *storage.Memory) {
	log.Printf("collecting data every %s", d)
	for {
		select {
		case <-time.Tick(1 * time.Hour):
			today = time.Now().Format("2006-01-02")

		case <-time.Tick(d):
			log.Println("starting collect after:", d)

			stats, err := c.GetDailyStat(cfg.Ueshka.PupilID, today, today)
			if err != nil {
				log.Println("fault get statistic err:", err)
				continue
			}

			for date, operations := range stats {
				for _, op := range operations {
					uid := genUniqID(date, op)
					if repo.IsExist(uid) {
						continue
					}

					log.Printf("found new operaton date: %s uid: %s", date, uid)

					msg := cfg.gate.RenderMessage(&op)
					if err := cfg.gate.Send(msg); err != nil {
						log.Println("fault send message err:", err)
					}
					repo.Add(uid)
				}
			}

			log.Print("end collecting")
		}
	}
}

func main() {
	cfg := NewAppConfigFromEnv()

	client := ueshka.NewClient(cfg.Ueshka.VersionAPI, cfg.Ueshka.Token)

	repo := storage.NewMemory()

	runDaemon(
		cfg.CheckInterval,
		client,
		NewAppConfigFromEnv(),
		repo,
	)
}
