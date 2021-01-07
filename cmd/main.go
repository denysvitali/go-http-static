package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var args struct {
	Path string `arg:"positional,required"`
	Port string `arg:"-p,--port" default:"8080"`
	ListenAddress string `arg:"-l,--listen" default:""`
	TLS *bool `arg:"-t,--tls"`
	CertFile *string `arg:"-c,--certificate"`
	CertKey *string `arg:"-k,--key"`
}

func main(){
	arg.MustParse(&args)
	router := gin.Default()
	router.StaticFS("/", http.Dir(args.Path))

	addr := fmt.Sprintf("%s:%s", args.ListenAddress, args.Port)
	
	if args.TLS != nil {
		if args.CertFile == nil {
			logrus.Fatal("you need to provide a certificate when using TLS")
		}
		
		if args.CertKey == nil {
			logrus.Fatal("you need to provide a key when using TLS")
		}

		_ = router.RunTLS(addr, *args.CertFile, *args.CertKey)
	} else {
		_ = router.Run(addr)
	}
}
