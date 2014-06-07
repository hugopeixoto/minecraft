package profiles

import (
  "net/http"
  "encoding/json"
  "encoding/base64"
  "io"
  "strings"
  "errors"
)

func parseJson(reader io.Reader, object interface{}) error {
  return json.NewDecoder(reader).Decode(object)
}

type profileJson struct {
  Id   string `json:"id"`
  Name string `json:"name"`
  Properties []struct {
    Name      string `json:"name"`
    Signature string `json:"signature"`
    Value     string `json:"value"`
  } `json:"properties"`
}

type texturesJson struct {
  Textures struct{
    Skin struct {
      URL string `json:"url"`
    } `json:"SKIN"`
  } `json:"textures"`
}

var NotFoundError = errors.New("Profile not found")

func Fetch (id string) (*Profile, error) {
  response, err := http.Get(
    "https://sessionserver.mojang.com/session/minecraft/profile/" + id)

  if err != nil {
    return nil, err
  }

  defer response.Body.Close()

  if response.StatusCode != 200 {
    return nil, NotFoundError
  }

  var profile profileJson
  err = parseJson(response.Body, &profile)
  if err != nil {
    return nil, err
  }

  propertyReader := strings.NewReader(profile.Properties[0].Value)

  var textures texturesJson
  err = parseJson(base64.NewDecoder(base64.StdEncoding, propertyReader), &textures)
  if err != nil {
    return nil, err
  }

  return &Profile{profile.Id, profile.Name, textures.Textures.Skin.URL}, nil
}
