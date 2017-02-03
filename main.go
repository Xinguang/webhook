package main

import (
	"fmt"
	"log"
	"net/http"

	mflag "github.com/starboychina/webhook/mflag"
)

var (
	// version
	version = "0.2.0"
)
var (
	// --version
	cmdVersion = mflag.Bool([]string{"v", "-version"}, false, "")
	// --help
	cmdHelp = mflag.Bool([]string{"h", "-help"}, false, "helps you out when in dire need of information")
	// hook setting
	cmdIP         = mflag.String([]string{"ip"}, "0.0.0.0", "ip the webhook should serve hooks on")
	cmdPort       = mflag.Int([]string{"p", "-port"}, 9000, "port to listen on")
	cmdConfigFile = mflag.String([]string{"c", "-config"}, "./config.json", "config")
	cmdLogFile    = mflag.String([]string{"l", "-log"}, "/var/log/webhook.log", "log file")
	// https
	cmdHTTPS = mflag.Bool([]string{"https"}, false, "use HTTPS instead of HTTP")
	cmdCert  = mflag.String([]string{"cert"}, "cert/cert.pem", "path to the HTTPS certificate pem file")
	cmdKey   = mflag.String([]string{"key"}, "cert/key.pem", "path to the HTTPS certificate private key pem file")
)
var config Config

func init() {

	mflag.Usage = func() {
		fmt.Println("Usage: webhook [OPTIONS] \n       webhook [ --help | -v | --version ]")
		mflag.PrintDefaults()
	}
	mflag.Parse()
}

func main() {

	if *cmdVersion {
		fmt.Println("webhook version", version)
		return
	}

	if *cmdHelp {
		mflag.Usage()
		return
	}

	setLog()
	loadConfig()

	if *cmdHTTPS {
		log.Printf("serving hooks on https://%s:%d/", *cmdIP, *cmdPort)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", *cmdIP, *cmdPort), *cmdCert, *cmdKey, nil))
	} else {
		log.Printf("serving hooks on http://%s:%d", *cmdIP, *cmdPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *cmdIP, *cmdPort), nil))
	}
}
