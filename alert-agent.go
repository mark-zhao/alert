package main

import (
	//"github.com/golang/glog"
	"glog"
	"google.golang.org/grpc"
	"net"
	"pb"
)

const (
	port = ":1234"
)


// init
func init() {
	//@TODO
}


func main() {
	defer glog.Flush()
	//flag.Parse()

	glog.Info("alert agent start")
	listen, err := net.Listen("tcp", port)
	if err != nil{
		glog.Error("net Listen failed; message: ", err)
		return
	}

	service := grpc.NewServer()
	pb.RegisterCephServiceServer(service, &pb.Ceph{})
	if err := service.Serve(listen); err != nil{
		glog.Error("grpc failed to serve: ", err)
	}
}
