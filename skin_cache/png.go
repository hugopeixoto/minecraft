package skin_cache

import (
  "image"
  "image/png"
  "net/http"
  "os"
)

func savePng(filename string, image image.Image) error {
  writer, err := os.Create(filename)
  if err != nil {
    return err
  }

  defer writer.Close()

  return png.Encode(writer, image)
}

func fetchHTTPPNG(url string) (image.Image, error) {
  response, err := http.Get(url)
  if err != nil {
    return nil, err
  }

  defer response.Body.Close()

  return png.Decode(response.Body)
}
