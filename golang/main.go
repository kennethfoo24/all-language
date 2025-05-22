package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

// sayHello is hit by external callers. It makes an internal HTTP
// round‑trip to /internal-work and performs some local work so you can
// still see multiple spans if you eventually re‑enable tracing tools.
func sayHello(w http.ResponseWriter, r *http.Request) {
    // downstream call to our own server ---------------------------------
    client := &http.Client{Timeout: 2 * time.Second}

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    req, _ := http.NewRequestWithContext(r.Context(), "GET", "http://127.0.0.1:"+port+"/internal-work", nil)

    // propagate a synthetic request‑id just to prove headers flow
    req.Header.Set("X-Request-ID", time.Now().Format(time.RFC3339Nano))

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("error calling internal-work: %v", err)
    } else {
        io.Copy(io.Discard, resp.Body)
        resp.Body.Close()
    }

    // extra local work ---------------------------------------------------
    time.Sleep(10 * time.Millisecond)

    w.Write([]byte("Hello World with fan-out call!"))
}

// internalWork is reached only by the internal HTTP client in sayHello.
// It performs some fake CPU work to illustrate nested spans (if tracing
// is added back later).
func internalWork(w http.ResponseWriter, r *http.Request) {
    // simulate business logic
    time.Sleep(100 * time.Millisecond)

    w.Write([]byte("Internal work done"))
}

func say_hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", sayHello)
    mux.HandleFunc("/hello", say_hello)
    mux.HandleFunc("/internal-work", internalWork)

    log.Printf("listening on :%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}


