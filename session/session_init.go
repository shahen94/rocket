package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

func (s *Session) Init() *scs.SessionManager {
	var persist, secure bool

	minutes, err := strconv.Atoi(s.CookieLifetime)

	if err != nil {
		minutes = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	} else {
		secure = false
	}

	sess := scs.New()
	sess.Lifetime = time.Duration(minutes) * time.Minute
	sess.Cookie.Persist = persist
	sess.Cookie.Name = s.CookieName
	sess.Cookie.Secure = secure
	sess.Cookie.Domain = s.CookieDomain
	sess.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
	}

	return sess
}
