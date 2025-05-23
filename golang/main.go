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

func helloService(w http.ResponseWriter, r *http.Request) {
    // 1) Figure out the dotnet service URL (with default)
    dotnetURL := os.Getenv("DOTNET_SERVICE_URL")
    if dotnetURL == "" {
        dotnetURL = "http://localhost:4567"
    }

    // 2) Call /dotnet
    endpoint := dotnetURL + "/dotnet"
    client := &http.Client{Timeout: 3 * time.Second}
    resp, err := client.Get(endpoint)
    if err != nil {
        http.Error(w, "failed to call dotnet service: "+err.Error(), http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    // 3) Propagate the dotnet status code
    w.WriteHeader(resp.StatusCode)
    // 4) Write our own greeting first
    w.Write([]byte("Hello World from Golang!\n"))
    // 5) Then stream the dotnet response body
    if _, err := io.Copy(w, resp.Body); err != nil {
        log.Printf("error copying dotnet response: %v", err)
    }
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", sayHello)
    mux.HandleFunc("/golang", helloService)
    mux.HandleFunc("/internal-work", internalWork)

    log.Printf("listening on :%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}


