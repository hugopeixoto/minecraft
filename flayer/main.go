package main

import (
  "flag"
  "log"
)

func main() {
  cache   := flag.String("cachedir", "cache", "directory to place cached skins (must be created beforehand)")
  address := flag.String("listen", ":9999", "listen address (0.0.0.0:9999)")

  flag.Parse()

  conf := Configuration{*cache, *address}

  ws := NewWebserver(conf)

  log.Fatalln(ws.Run())
}
