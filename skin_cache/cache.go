package skin_cache

import (
  "os"
  "image"
  "path"
  "regexp"
  "fmt"

  "github.com/hugopeixoto/minecraft/profiles"
  "github.com/hugopeixoto/minecraft/skins"
)

type Configuration struct {
  Filename string
  Convert  func(image.Image) image.Image
}

type SkinCache struct {
  Directory string
  Configurations []Configuration
}

var sizes = []uint{8, 16, 20, 24, 25, 32, 40, 48, 64}

func NewSkinCache(directory string) *SkinCache {
  sc := SkinCache{
    directory,
    []Configuration{Configuration{"full", skins.Full}},
  }

  for _, size := range sizes {
    sc.Configurations = append(
      sc.Configurations,
      Configuration{
        fmt.Sprintf("head-%v", size),
        func(src image.Image) image.Image { return skins.CoveredHead(src, size) }})
  }

  return &sc
}

func (sc SkinCache) ValidRequest(uuid string, configuration string) bool {
  ok, err := regexp.MatchString("^[a-z0-9]{32}$", uuid)
  if err != nil || !ok {
    return false
  }

  for _, conf := range sc.Configurations {
    if configuration == conf.Filename {
      return true
    }
  }

  return false
}

func (sc SkinCache) Cache(uuid string) error {
  uuidDirectory := path.Join(sc.Directory, uuid)

  _, err := os.Stat(uuidDirectory)
  if err == nil {
    return nil
  }

  profile, err := profiles.Fetch(uuid)
  if err != nil {
    return err
  }

  img, err := fetchHTTPPNG(profile.SkinURL)
  if err != nil {
    return err
  }

  err = os.Mkdir(uuidDirectory, 0755)
  if err != nil {
    return err
  }

  for _, config := range sc.Configurations {
    err = savePng(
      path.Join(uuidDirectory, config.Filename + ".png"),
      config.Convert(img))

    if err != nil {
      os.RemoveAll(uuidDirectory)
      return err
    }
  }

  return nil
}
