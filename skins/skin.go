package skins

import (
  "image"
  "image/draw"
  "github.com/nfnt/resize"
)

func Full(src image.Image) image.Image {
  return src
}

func Head(src image.Image, size uint) image.Image {
  head := crop(src, image.Rect(8, 8, 16, 16))

  return resize.Resize(size, size, head, resize.NearestNeighbor)
}

func CoveredHead(src image.Image, size uint) image.Image {
  head := crop(src, image.Rect(8, 8, 16, 16))
  hat  := crop(src, image.Rect(40, 8, 48, 16))

  draw.Draw(head, head.Bounds(), hat, hat.Bounds().Min, draw.Over)

  return resize.Resize(size, size, head, resize.NearestNeighbor)
}

func crop(src image.Image, r image.Rectangle) draw.Image {
  target := image.Rectangle{image.Point{0, 0}, r.Size()}

  cropped := image.NewRGBA(target)

  draw.Draw(cropped, target, src, r.Min, draw.Over)

  return cropped
}
