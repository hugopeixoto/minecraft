package skin_cache

import (
  "os"

  "github.com/hugopeixoto/minecraft/profiles"
)

func Sweep(cache *SkinCache) error {
  dir, err := os.Open(cache.Directory)
  if err != nil {
    return err
  }

  defer dir.Close()

  uuids, err := dir.Readdirnames(0)
  if err != nil {
    return err
  }

  for _, uuid := range uuids {
    if profiles.ValidIdentifier(uuid) {
      if hasUpdate(cache, uuid) {
        cache.ClearCache(uuid)
        cache.Cache(uuid)
      }
    }
  }

  return nil
}

func hasUpdate(cache *SkinCache, uuid string) bool {
  profile, err := profiles.Fetch(uuid)
  if err != nil {
    return false
  }

  url, err := cache.CachedURL(uuid)
  if err != nil {
    return true
  }

  return profile.SkinURL != url
}
