package bluetang

func Name(field string) (string, string) {
  if len(field) < 4 {
    return field, "Nome deve conter pelo menos quatro letras"
  } else {
    return field, ""
  }
}
