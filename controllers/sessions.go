package controllers

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	_ = godotenv.Load()

	authKey := os.Getenv("SESSION_KEY")
	if authKey == "" {
		log.Fatal("A variável de ambiente 'SESSION_KEY' não foi definida. Por favor, defina-a.")
	}

	store = sessions.NewCookieStore([]byte(authKey))

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "app-session")
}

func SaveSession(w http.ResponseWriter, r *http.Request, session *sessions.Session) error {
	return session.Save(r, w)
}
