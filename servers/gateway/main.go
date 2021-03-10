package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louistaa/study-buddy/servers/gateway/handlers"
	"github.com/louistaa/study-buddy/servers/gateway/models/users"
	"github.com/louistaa/study-buddy/servers/gateway/sessions"
)

//main is the main entry point for the server
func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	sessionKey := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	DSN := os.Getenv("DSN")

	if len(redisAddr) == 0 {
		redisAddr = "127.0.0.1:6379"
	}

	redisDB := redis.NewClient(&redis.Options{Addr: redisAddr})

	sessionStore := sessions.NewRedisStore(redisDB, time.Hour)

	userStore, err := users.NewMySQLStore(DSN)

	if err != nil {
		log.Printf("Unable to create userStore")
	}

	handlerContext := &handlers.HandlerContext{
		SigningKey:   sessionKey,
		SessionStore: sessionStore,
		UserStore:    userStore,
	}

	TLSKEY := os.Getenv("TLSKEY")
	TLSCERT := os.Getenv("TLSCERT")

	mux := http.NewServeMux()
	log.Printf("server is listening at %s...", addr)
	mux.HandleFunc("/v1/summary/", handlers.SummaryHandler)
	mux.HandleFunc("/v1/users", handlerContext.UsersHandler)
	mux.HandleFunc("/v1/users/", handlerContext.SpecificUserHandler)
	mux.HandleFunc("/v1/sessions", handlerContext.SessionsHandler)
	mux.HandleFunc("/v1/sessions/", handlerContext.SpecificSessionHanlder)

	corsMux := &handlers.CORS{Handler: mux}
	log.Fatal(http.ListenAndServeTLS(addr, TLSCERT, TLSKEY, corsMux))
}
