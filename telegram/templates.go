package telegram

import (
  "io/ioutil"
  "fmt"
  "strings"
)

var templatesMap = map[string]string {}
type templateData map[string]interface{}

func removeExtension(filename string) string {
  return strings.Replace(filename, ".md", "", -1)
}

func readTemplates() {
  files, err := ioutil.ReadDir("./telegram/templates")
  check(err)
  for _, file := range files {
    dat, err := ioutil.ReadFile("./telegram/templates/" + file.Name())
    check(err)
    templatesMap[removeExtension(file.Name())] = string(dat)
  }
  fmt.Println(templatesMap)
}

func getResponseText(message string, data templateData) string {
  response := templatesMap[message]
  return response
}
