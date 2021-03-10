package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/louistaa/study-buddy/servers/gateway/handlers"
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

	studentStore, err := students.NewMySQLStore(DSN)

	if err != nil {
		log.Printf("Unable to create userStore")
	}

	handlerContext := &handlers.HandlerContext{
		SigningKey:   sessionKey,
		SessionStore: sessionStore,
		StudentStore: studentStore,
	}

	TLSKEY := os.Getenv("TLSKEY")
	TLSCERT := os.Getenv("TLSCERT")

	mux := http.NewServeMux()
	log.Printf("server is listening at %s...", addr)
	mux.HandleFunc("/students", handlerContext.StudentsHandler)
	mux.HandleFunc("/students/", handlerContext.SpecificStudentHandler)
	mux.HandleFunc("/sessions", handlerContext.SessionsHandler)
	mux.HandleFunc("/sessions/", handlerContext.SpecificSessionHanlder)

	corsMux := &handlers.CORS{Handler: mux}
	log.Fatal(http.ListenAndServeTLS(addr, TLSCERT, TLSKEY, corsMux))
}
