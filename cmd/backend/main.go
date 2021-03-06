package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bzdacs/config"
	"github.com/hardstylez72/bzdacs/generated"
	"github.com/hardstylez72/bzdacs/pkg/infra/logger"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

// @title BZDACS
// @version 1.0
// @description Access control system

// @contact.name Sergey Korotkov
// @contact.instruction ss

// @host localhost:8081
// @BasePath /api

// @securityDefinitions.basic BasicAuth

const (
	apiPathPrefix = "/api"
)

func main() {

	l, err := logger.New("")
	errCheck(err, "can't load config")
	defer func() {
		_ = l.Sync()
	}()

	err = NewServer(l).Run()
	errCheck(err, "can't run server")
}

func errCheck(err error, errorText string) {
	if err == nil {
		return
	}
	log.Fatal(errorText, ": ", err)
}

func NewServer(log *zap.SugaredLogger) *Server {
	return &Server{
		router: chi.NewRouter(),
		log:    log,
	}
}

func (s *Server) Run() error {

	configPath := flag.String("config", "/opt/cfg/config.yaml", "path to config file")
	flag.Parse()

	err := config.Load(*configPath)
	if err != nil {
		return err
	}
	//err = tracer.New(viper.GetString("tracer.jaeger.collectorEndpoint"), viper.GetString("tracer.jaeger.serviceName"))
	//if err != nil {
	//	return err
	//}

	generated.SwaggerInfo.Host = strings.ReplaceAll(config.GetHost(), "http://", "")

	done := make(chan struct{})
	ready := make(chan struct{})

	if viper.GetString("env") == "prod" {

		proxyPort := viper.GetString("proxy.port")
		staticServerHost := "http://localhost" + viper.GetString("ui.port")
		pathToStaticFiles := viper.GetString("ui.path")
		apiServerHost := "http://localhost" + viper.GetString("backend.port")
		go func() {
			err := startStaticServer(staticServerHost, pathToStaticFiles, done)
			if err != nil {
				done <- struct{}{}
				log.Fatal(err)
			}
		}()

		go func() {
			err = startProxyServer(apiServerHost, staticServerHost, proxyPort, done)
			if err != nil {
				done <- struct{}{}
				log.Fatal(err)
			}
		}()
	}

	go func() {
		err = s.startBackendServer(s.log, done, ready)
		if err != nil {
			done <- struct{}{}
			log.Fatal(err)
		}
	}()

	go func() {
		<-ready
		time.Sleep(time.Millisecond * 100)
	}()

	<-done
	return nil
}
