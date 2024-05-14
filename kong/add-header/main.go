package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
	"log"
)

type Configuration struct {
	HeaderName string `json:"header_name"`
	Message    string `json:"message"`
}

func New() interface{} {
	return &Configuration{}
}

const Version = "1.0.0"
const Priority = 1

func (conf *Configuration) Access(kong *pdk.PDK) {
	err := kong.ServiceRequest.AddHeader(conf.HeaderName, conf.Message)
	if err != nil {
		log.Println(err)
		kong.Response.ExitStatus(500)
	}
}

func main() {
	err := server.StartServer(New, Version, Priority)
	if err != nil {
		panic(err)
	}
}
