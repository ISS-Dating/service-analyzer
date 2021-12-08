package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/ISS-Dating/service-analyzer/repo"
	"github.com/ISS-Dating/service-analyzer/service"
	"github.com/ISS-Dating/service-analyzer/web"
)

var (
	localSetup = false
)

func main() {
	var db *sql.DB
	var err error
	if localSetup {
		db, err = sql.Open("postgres", "host=localhost user=postgres password=12345 port=5432 dbname=q_date sslmode=disable")
	} else {
		var stop bool
		for !stop {
			time.Sleep(time.Second * 8)
			db, _ = sql.Open("postgres", "host=postgres user=postgres password=12345 port=5432 dbname=postgres sslmode=disable")
			err = db.Ping()
			if err != nil {
				log.Println("Error connecting to db")
			} else {
				stop = true
			}
		}
	}
	if err != nil {
		log.Fatal("Can't connect to db: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Can't connect to db: ", err.Error())
	}

	service := service.NewService(repo.NewRepo(db))

	go web.Start()

	poller := &web.Poller{
		Service: service,
	}

	poller.Start()
}
