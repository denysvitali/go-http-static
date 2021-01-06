package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/gin-gonic/gin"
	"net/http"
)

var args struct {
	Path string `arg:"positional,required"`
	Port string `arg:"-p,--port" default:"8080"`
	ListenAddress string `arg:"-l,--listen" default:""`
}

func main(){
	arg.MustParse(&args)
	router := gin.Default()
	router.StaticFS("/", http.Dir(args.Path))

	addr := fmt.Sprintf("%s:%s", args.ListenAddress, args.Port)
	_ = router.Run(addr)
}
