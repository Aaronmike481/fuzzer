package main

import(
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

func initPool(){
	var err error

	pool, err := pyproc.NewPool(pyproc.PoolOptions{
		Config: pyproc.PoolConfig{
			Workers: 4,
			MaxInFlight,
		},
		WorkerConfig: pyproc.WorkerConfig{
			SocketPath: "/tmp/fuzzer.sock",
			PythonExec: "python3",
			WorkerScript: "fuzzer.py",
		}
	},
	nil 
)

if err != nil{
	log.Fatal(err)
}

pool.Start(context.Background())
}
