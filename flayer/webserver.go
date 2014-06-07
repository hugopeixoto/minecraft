package main

import (
  "github.com/hugopeixoto/minecraft/skin_cache"
	"github.com/gorilla/mux"
  "net/http"
  "path"
)

type Configuration struct {
  CachePath     string
  ListenAddress string
}

type Webserver struct {
  Config *Configuration
  Cache *skin_cache.SkinCache
}

func NewWebserver(config Configuration) *Webserver {
  return &Webserver{
    &config,
    skin_cache.NewSkinCache(config.CachePath),
  }
}

func (ws *Webserver) Handle(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)

  uuid := params["uuid"]
  configuration := params["configuration"]

  if !ws.Cache.ValidRequest(uuid, configuration) {
    http.NotFound(w, r)
    return
  }

  ws.Cache.Cache(uuid)

  http.ServeFile(w, r, path.Join(ws.Config.CachePath, uuid, configuration + ".png"))
}

func (ws *Webserver) Run() error {
	r := mux.NewRouter()

  r.HandleFunc("/skins/{uuid}/{configuration}.png", ws.Handle)
	http.Handle("/", r)

  return http.ListenAndServe(ws.Config.ListenAddress, nil)
}
