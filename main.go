package main

import (
  "fmt"
  "html/template"
  "net/http"

  "encoding/json"
  "io/ioutil"
  "os"
  
  "log"

  "github.com/gorilla/pat"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/google"
  "github.com/markbates/goth/providers/amazon"
)

var port = ":3000"
var hostName = "http://localhost" + port

type Config struct{
  Google Provider `json:"google"`
  Amazon Provider `json:"amazon"`
}

type Provider struct{
  ClientID string `json:"client_id"`
  Secret string `json:"secret"`
  Callback string `json:"callback"`
}

func readConfig(config *Config)  {
  raw, err := ioutil.ReadFile("./config.json")
  if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
  }

  json.Unmarshal(raw, &config)  
}

func indexHandlerfunc(res http.ResponseWriter, req *http.Request) {
  t, _ := template.ParseFiles("templates/index.html")
  t.Execute(res, false)
}

func authHandler(res http.ResponseWriter, req *http.Request) {
  var config Config
  readConfig(&config)

  provider := req.URL.Query().Get(":provider")
  switch provider {
    case "google":
      goth.UseProviders(
        google.New(config.Google.ClientID, config.Google.Secret, hostName+ "/auth/google/callback", "email", "profile"),
      )
    case "amazon":
      goth.UseProviders(
        amazon.New(config.Amazon.ClientID, config.Amazon.Secret, hostName+ "/auth/amazon/callback", "profile", "profile:user_id"),
      ) 

  }
  gothic.BeginAuthHandler(res, req)
}

func callBackHandler(res http.ResponseWriter, req *http.Request) {
  user, err := gothic.CompleteUserAuth(res, req)
  if err != nil {
    fmt.Fprintln(res, err)
    return
  }
  t, _ := template.ParseFiles("templates/success.html")
  t.Execute(res, user)
}

func main() {
  p := pat.New()
  p.Get("/auth/{provider}/callback", callBackHandler)
  p.Get("/auth/{provider}", authHandler)
  p.Get("/", indexHandlerfunc)

  log.Println("listening on " + hostName)
  log.Fatal(http.ListenAndServe(port, p))
}
