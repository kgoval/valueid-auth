package main

import (
	"fmt"
	"github.com/kgoval/envparser"
	"github.com/kgoval/erresp"
	"github.com/kgoval/valueid-auth/boot"
	"github.com/kgoval/valueid-auth/login"
	"github.com/kgoval/valueid-auth/servers"
	local "github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"strings"

	//"github.com/micro/go-micro"
)

func init()  {
	//-- registering error message
	erresp.Register(ErrorMessage)

	//-- parse config from env
	if err := envparser.Parse(boot.Config); err != nil{
		panic(err)
	}

}



func main(){

	fmt.Println("engine started..")
	var service micro.Service

	if strings.ToLower(boot.Config.Environment) == "local" {
		service = local.NewService()
	}else{
		service = k8s.NewService()
	}

	service.Init(
		micro.Name(boot.Config.ServiceName), )

	//handler.Register(service.Server())
	servers.Register(service.Client())

	fmt.Println("running..")
	//if err := service.Run(); err != nil {
	//	log.Fatal(err)
	//}

	l := login.Register(&servers.Servers)
	l.Token()

}
