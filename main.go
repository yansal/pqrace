package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/lib/pq"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	pqconnector, err := pq.NewConnector(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	db := sql.OpenDB(pqconnector)

	const n = 2
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			var now time.Time
			if err := db.QueryRowContext(context.Background(), `select now()`).Scan(&now); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
}
