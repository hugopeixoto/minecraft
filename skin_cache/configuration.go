package skin_cache

import (
  "path"
  "image"
  "fmt"

  "github.com/hugopeixoto/minecraft/skins"
)

type Configuration struct {
  Basename  string
  ConvertFn func(image.Image) image.Image
}

func (c Configuration) Filename(directory string) string {
  return path.Join(directory, c.Basename + ".png")
}

func (c Configuration) Convert(directory string, img image.Image) error {
  return savePng(c.Filename(directory), c.ConvertFn(img))
}

func FullSkinConfiguration() Configuration {
  return Configuration{"full", skins.Full}
}

func CoveredHeadConfiguration(size uint) Configuration {
  return Configuration{
    fmt.Sprintf("head-%v", size),
    func(src image.Image) image.Image {
      return skins.CoveredHead(src, size)
    }}
}
