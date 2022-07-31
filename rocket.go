package rocket

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/shahen94/rocket/render"
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
	config   config
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

	err = r.checkDotEnv(rootPath)

	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/" + ".env")
	if err != nil {
		return err
	}

	infoLog, errorLog := r.startLoggers()

	r.InfoLog = infoLog
	r.ErrorLog = errorLog

	r.RootPath = rootPath
	r.AppName = os.Getenv("APP_NAME")
	r.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	r.Version = version
	r.Routes = r.routes().(*chi.Mux)

	r.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	views := jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	r.JetViews = views

	r.Render = render.New(r.config.renderer, rootPath, r.config.port, r.JetViews)

	return nil
}

func (r *Rocket) ListenAndServe() {
	r.InfoLog.Println("Server started on port " + r.config.port)

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
		err := r.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
