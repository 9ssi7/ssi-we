package app

import (
	"context"
	"log"

	cnf "github.com/ssibrahimbas/ssi-we/pkg/config"
	"github.com/ssibrahimbas/ssi-we/pkg/db"
	"github.com/ssibrahimbas/ssi-we/pkg/env"
	"github.com/ssibrahimbas/ssi-we/pkg/helper"
	"github.com/ssibrahimbas/ssi-we/pkg/http"
	"github.com/ssibrahimbas/ssi-we/src/config"
	"github.com/ssibrahimbas/ssi-we/src/internal"
)

type App struct {
	Cnf   *config.App
	Mongo *db.MongoDB
	Http  *http.Client
	Repo  *internal.Repo
	Srv   *internal.Srv
	Hnd   *internal.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Prepare() *App {
	a.loadConfig()
	a.loadMongo()
	a.loadHttp()
	a.loadInternal()
	return a
}

func (a *App) Listen() {
	log.Fatal(a.Http.Listen(a.Cnf.Host + ":" + a.Cnf.Port))

}

func (a *App) loadConfig() {
	cnfg := config.App{}
	cnf.LoadConfig(&cnfg)
	env.LoadEnv("./", &cnfg)
	a.Cnf = &cnfg
}

func (a *App) loadMongo() {
	uri := db.CalcMongoUri(db.UriParams{
		Host: a.Cnf.MongoHost,
		Port: a.Cnf.MongoPort,
		User: a.Cnf.MongoUsername,
		Pass: a.Cnf.MongoPassword,
		Db:   a.Cnf.MongoName,
	})
	d, err := db.NewMongo(uri, a.Cnf.MongoName)
	helper.CheckErr(err)
	d.Ping()
	a.Mongo = d
}

func (a *App) loadHttp() {
	a.Http = http.New()
}

func (a *App) loadInternal() {
	a.Repo = internal.NewRepo(&internal.RepoParams{
		Collection: a.Mongo.GetCollection(a.Cnf.MongoCollection),
		DB:         a.Mongo,
		Ctx:        context.TODO(),
	})
	a.Srv = internal.NewSrv(&internal.SrvParams{
		Repo: a.Repo,
	})
	a.Hnd = internal.NewHandler(&internal.HandlerParams{
		Srv:  a.Srv,
		Cnf:  a.Cnf,
		Http: a.Http,
	})
	a.Hnd.Init()
}
