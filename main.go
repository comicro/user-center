package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"user-center/conf"
	"user-center/db"
	"user-center/handler"
	"user-center/subscriber"

	user "user-center/proto/user"
)

func main() {
	// load config
	conf.Load()
	// load db
	db.Load()
	defer db.Orm().Close()

	db.Orm().AutoMigrate(&user.User{})

	// New Service
	service := micro.NewService(
		micro.Name("comicro.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserServiceHandler(service.Server(), handler.UserHandle())

	// Register Struct as Subscriber
	micro.RegisterSubscriber("comicro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("comicro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
