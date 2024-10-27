package main

import (
	"github.com/saintbyte/postgresURItoDSN"
	"log/slog"
	"os"
)

func main() {
	var database_url string = os.Getenv("DATABASE_URL")
	if len(os.Args) > 1 {
		database_url = os.Args[1]
	} else if database_url == "" {
		database_url = "postgresql://user_111:passwordssf@" +
			"qy-blue-block-65767118.eu-central-1.aws.neon.tech/neondb?sslmode=require"
	}
	slog.Info("database_url: ", database_url)
	dsn, err := postgresURItoDSN.UriToDSN(database_url)
	if err != nil {
		slog.Error("dsn error: ", err)
		return
	}
	slog.Info("dsn: ", dsn)
}
