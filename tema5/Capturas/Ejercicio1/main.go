package main

import (  
    "fmt"
    "context"
//    "log"
    "go.etcd.io/etcd/clientv3"
    "time"
//    "strconv"
)

func main() {
   	cli, err := clientv3.New(clientv3.Config{
	Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	resp, err := cli.Get(ctx, "clave")
	cancel()
	if err != nil {
		// handle error!
	}else{
	// use the response
	fmt.Print(resp)
	fmt.Print("\r")
	}
	
	defer cli.Close()
}
