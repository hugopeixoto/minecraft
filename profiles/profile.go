package profiles

import "regexp"

type Profile struct {
  Id   string
  Name string

  SkinURL string
}

func ValidIdentifier(uuid string) bool {
  ok, err := regexp.MatchString("^[a-z0-9]{32}$", uuid)

  return err == nil && ok
}
