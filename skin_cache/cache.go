package skin_cache

import (
  "os"
  "path"
  "path/filepath"

  "github.com/hugopeixoto/minecraft/profiles"
)

type SkinCache struct {
  Directory string
  Configurations []Configuration
}

var sizes = []uint{8, 16, 20, 24, 25, 32, 40, 48, 64}

func NewSkinCache(directory string) *SkinCache {
  absolutePath, err := filepath.Abs(directory)
  if err != nil {
    return nil
  }

  sc := SkinCache{absolutePath, []Configuration{FullSkinConfiguration()}}

  for _, size := range sizes {
    sc.Configurations = append(sc.Configurations, CoveredHeadConfiguration(size))
  }

  if !sc.IsCached("steve") {
    sc.cacheURL(sc.SteveDirectory(), "https://minecraft.net/images/char.png")
  }

  return &sc
}

func (sc SkinCache) UserDirectory(uuid string) string {
  return path.Join(sc.Directory, uuid)
}

func (sc SkinCache) SteveDirectory() string {
  return sc.UserDirectory("steve")
}

func (sc SkinCache) ValidRequest(uuid string, configuration string) bool {
  if !profiles.ValidIdentifier(uuid) {
    return false
  }

  for _, conf := range sc.Configurations {
    if configuration == conf.Basename {
      return true
    }
  }

  return false
}

func (sc SkinCache) IsCached(uuid string) bool {
  _, err := os.Stat(sc.UserDirectory(uuid))

  return err == nil
}

func (sc SkinCache) ClearCache(uuid string) error {
  directory := sc.UserDirectory(uuid)
  fi, err := os.Stat(directory)
  if err != nil {
    return err
  }

  if fi.IsDir() {
    return os.RemoveAll(directory)
  } else {
    // probably a symlink to steve.
    // Don't remove its contents, just the link.
    return os.Remove(directory)
  }
}

func (sc SkinCache) Cache(uuid string) error {
  if sc.IsCached(uuid) {
    return nil
  }

  profile, err := profiles.Fetch(uuid)
  if err != nil && err != profiles.NotFoundError {
    return err
  }

  if err == profiles.NotFoundError {
    err = sc.linkToSteve(uuid)
  } else {
    err = sc.cacheURL(sc.UserDirectory(uuid), profile.SkinURL)
  }

  if err != nil {
    sc.ClearCache(uuid)
  }

  return err
}

func (sc SkinCache) cacheURL(directory, url string) error {
  img, err := fetchHTTPPNG(url)
  if err != nil {
    return err
  }

  err = os.Mkdir(directory, 0755)
  if err != nil && !os.IsExist(err) {
    return err
  }

  for _, config := range sc.Configurations {
    err = config.Convert(directory, img)
    if err != nil {
      return err
    }
  }

  return nil
}

func (sc SkinCache) linkToSteve(uuid string) error {
  return os.Symlink(sc.SteveDirectory(), sc.UserDirectory(uuid))
}
