package session

import (
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func Test_InitSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "60",
		CookiePersist:  "true",
		CookieSecure:   "true",
		CookieName:     "rocket_session",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	sm := scs.New()

	ses := c.Init()

	var sessKind reflect.Kind
	var sessType reflect.Type

	rv := reflect.ValueOf(ses)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		sessKind = rv.Kind()
		sessType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Errorf("Invalid session object")
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Errorf("Invalid session object kind")
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Errorf("Invalid session object type")
	}
}
