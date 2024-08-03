package config

import (
	"context"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickmn/go-cache"
)

var Dbpool *pgxpool.Pool
var err error
var Cache *cache.Cache

func Connect(r *chi.Mux) {
	Dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		println(err)
		println("unable to connect")
	}
	Cache = cache.New(5*time.Minute, 10*time.Minute)
}
