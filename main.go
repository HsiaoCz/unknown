package main

import (
	"flag"
	"go-hello/api"
	"go-hello/storage"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9091", "set hello server listen Addr")
	flag.Parse()

	log.Println("the service is running on port", *listenAddr)

	mysql_store := storage.NewMysqlStorage()

	srv := api.NewServer(*listenAddr, mysql_store)

	log.Fatal(srv.Start())
}
