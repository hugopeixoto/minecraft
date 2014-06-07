package main

import (
  "flag"
  "log"
  "time"
)

func main() {
  cache   := flag.String("cachedir", "cache", "directory to place cached skins (must be created beforehand)")
  address := flag.String("listen", ":9999", "listen address (0.0.0.0:9999)")
  sweep   := flag.Int("interval", 6, "cache sweep interval, in hours")

  flag.Parse()

  conf := Configuration{*cache, *address, time.Duration(*sweep) * time.Hour}

  ws := NewWebserver(conf)

  log.Fatalln(ws.Run())
}
