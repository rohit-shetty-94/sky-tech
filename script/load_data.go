package script

import (
	"fmt"
	"log"
	"math/rand"
	"sky-tech/entity"
	"time"

	"github.com/go-pg/pg/v10"
)

func Ingest_metrics(db *pg.DB) {
	// Find latest timestamp in DB
	var lastTimestamp int64

	_, err := db.QueryOne(
		pg.Scan(&lastTimestamp),
		`SELECT MAX(timestamp) FROM metrics`,
	)
	if err != nil {
		log.Fatalf("Failed to query max timestamp: %v", err)
	}

	now := time.Now().Unix()
	fiveMinutesAgo := now - 300 // last 5 minutes in seconds

	// Only ingest data that doesn't exist
	var metrics []entity.Metric
	for t := fiveMinutesAgo; t < now; t += 60 {
		if t > lastTimestamp {
			metrics = append(metrics, entity.Metric{
				Timestamp:   t,
				CPULoad:     rand.Float64() * 100,
				Concurrency: rand.Intn(500001),
			})
		}
	}

	if len(metrics) == 0 {
		fmt.Println("No new data to ingest. Already up to date.")
		return
	}

	// Insert new metrics
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Transaction start failed:", err)
	}

	stmt, err := tx.Prepare(`INSERT INTO metrics (timestamp, cpu_load, concurrency) VALUES ($1, $2, $3)`)
	if err != nil {
		log.Fatal("Statement preparation failed:", err)
	}
	defer stmt.Close()

	for _, m := range metrics {
		_, err := stmt.Exec(m.Timestamp, m.CPULoad, m.Concurrency)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Insert failed at timestamp %d: %v", m.Timestamp, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Transaction commit failed:", err)
	}

	fmt.Printf("Inserted %d new metrics.\n", len(metrics))
}
