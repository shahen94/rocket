package rocket

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/shahen94/rocket/config"
	"github.com/shahen94/rocket/environment"
	"github.com/shahen94/rocket/helpers"
	"github.com/shahen94/rocket/logging"
	"github.com/shahen94/rocket/render"
	"github.com/shahen94/rocket/session"
)

const version = "0.0.1"

type Rocket struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config.Config
	Session  *scs.SessionManager
	Render   *render.Render
	JetViews *jet.Set
}

func (r *Rocket) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers",
			"migrations",
			"views",
			"data",
			"public",
			"tmp",
			"logs",
			"middlewares",
		},
	}

	err := r.Init(pathConfig)

	if err != nil {
		return err
	}

	err = environment.Touch(rootPath)

	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/" + ".env")
	if err != nil {
		return err
	}

	infoLog, errorLog := logging.New()

	r.InfoLog = infoLog
	r.ErrorLog = errorLog

	r.RootPath = rootPath
	r.AppName = os.Getenv("APP_NAME")
	r.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	r.Version = version
	r.Routes = r.routes().(*chi.Mux)

	r.config = config.Config{
		Port:     os.Getenv("PORT"),
		Renderer: os.Getenv("RENDERER"),
		Cookie: config.CookieConfig{
			Name:     os.Getenv("COOKIE_NAME"),
			Lifetime: os.Getenv("COOKIE_LIFETIME"),
			Persist:  os.Getenv("COOKIE_PERSIST"),
			Secure:   os.Getenv("COOKIE_SECURE"),
		},
		SessionType: os.Getenv("SESSION_TYPE"),
	}

	sess := session.Session{
		CookieLifetime: r.config.Cookie.Lifetime,
		CookiePersist:  r.config.Cookie.Persist,
		CookieSecure:   r.config.Cookie.Secure,
		CookieDomain:   r.config.Cookie.Domain,
		CookieName:     r.config.Cookie.Name,
		SessionType:    r.config.SessionType,
	}

	r.Session = sess.Init()

	views := jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	r.JetViews = views

	r.Render = render.New(r.config.Renderer, rootPath, r.config.Port, r.JetViews)

	return nil
}

func (r *Rocket) ListenAndServe() {
	r.InfoLog.Println("Server started on port " + r.config.Port)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     r.ErrorLog,
		Handler:      r.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		r.ErrorLog.Fatal(err)
	}
}

func (r *Rocket) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		err := helpers.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
