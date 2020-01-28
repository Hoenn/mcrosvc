package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/hoenn/mcrosvc/proto"
	"github.com/hoenn/mcrosvc/udb/pkg/db"
	"github.com/hoenn/mcrosvc/udb/pkg/server"

	"google.golang.org/grpc"
)

func main() {
	dbConn := &DBConn{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDRESS"),
		Name:     os.Getenv("DB_NAME"),
	}

	fmt.Println("Setting up DB")

	d, err := sql.Open("mysql", dbConn.Format())
	if err != nil {
		panic(err.Error())
	}
	err = d.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to DB")

	udb := &server.UDBServer{
		DB: db.NewUserAPI(d),
	}

	fmt.Println("Starting GRPC server")
	startGRPC(udb)
}

func startGRPC(udb *server.UDBServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUDBAPIServer(grpcServer, udb)
	grpcServer.Serve(lis)
}

type DBConn struct {
	Username string
	Password string
	Addr     string
	Name     string
}

func (d *DBConn) Format() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		d.Username,
		d.Password,
		d.Addr,
		d.Name)
}
