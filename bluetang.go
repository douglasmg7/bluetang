package bluetang

import (
  "regexp"
  "strings"
)

func Name(s string) (string, string) {
  var msg string
  if len(strings.TrimSpace(s)) < 4 {
    msg = "Nome deve conter pelo menos quatro letras"
  }
  // clean string
  s = strings.Join(strings.Fields(strings.Title(strings.ToLower(s))), " ")
  return s, msg
}

func Email(s string) (string, string) {
  var msg string
  reg := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  if !reg.MatchString(s) {
    msg = "E-mail inválido"
  }
  // clean string
  s = strings.TrimSpace(strings.ToLower(s))
  return s, msg
}

func Mobile(s string) (string, string) {
  var msg string
  // remove all non-numeric character sequence
  reg := regexp.MustCompile("[^0-9]+")
  s = reg.ReplaceAllString(s, "")
  // test valid mobile number
  reg = regexp.MustCompile("^[0-9]{10,11}$")
  if !reg.MatchString(s) {
    msg = "Número de celular inválido"
  } else {
    if len(s) == 10 {
      s = "(" + s[:2] + ") " + s[2:6] + "-" + s[6:]
    } else {
      s = "(" + s[:2] + ") " + s[2:3] + " " + s[3:7] + "-" + s[7:]
    }
  }
  // clean string
  return s, msg
}
