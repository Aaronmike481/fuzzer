package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    "github.com/YuminosukeSato/pyproc/pkg/pyproc"
)

var pool *pyproc.Pool

func initPool() {
    var err error
    pool, err = pyproc.NewPool(pyproc.PoolOptions{
        Config: pyproc.PoolConfig{
            Workers:     4,
            MaxInFlight: 10,
        },
        WorkerConfig: pyproc.WorkerConfig{
            SocketPath:   "/tmp/fuzzer.sock",
            PythonExec:   "python3",
            WorkerScript: "../python-worker/worker.py",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    pool.Start(context.Background())
}

type FromPy struct {
    Keywords []string `json:"keywords"`
    Payload  string   `json:"payload"`
    TargetAI string   `json:"target_ai"`
}

func handle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    var req FromPy
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    input := map[string]interface{}{
        "keywords":  req.Keywords,
        "payload":   req.Payload,
        "target_ai": req.TargetAI,
        "count":     10,
    }

    var output map[string]interface{}
    if err := pool.Call(context.Background(), "run_campaign", input, &output); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(output)
}

func main() {
    initPool()
    defer pool.Shutdown(context.Background())
    http.HandleFunc("/scan", handle)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
