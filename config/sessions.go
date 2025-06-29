package config

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"twiteer/utils"
)

var Store *sessions.CookieStore

func init() {

	utils.LoadEnv()

	// Initialize the session store with the secret key (from environment variable)
	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY"))) // Use a secure secret key
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,  // Session expires after 1 hour
		HttpOnly: true,  // Prevent JavaScript access to the cookie
		Secure:   false, // Set to true if using HTTPS
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	session, err := Store.Get(r, "users")
	if err != nil {
		log.Printf("Error getting session: %v", err)
		return nil, err
	}

	return session, nil
}
