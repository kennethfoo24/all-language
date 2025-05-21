// package main

// import (
// 	"net/http"
// 	"os"
// 	"time"

// 	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
// 	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
// )

// func say_hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World!"))
// }

// func add_tag(w http.ResponseWriter, r *http.Request) {

// 	// Retrieve a span for a web request attached to a Go Context.
// 	if span, ok := tracer.SpanFromContext(r.Context()); ok {
// 		// Set tag
// 		span.SetTag("new", "tag")
// 	}

// 	w.Write([]byte("Adding a tag."))
// }

// func set_error(w http.ResponseWriter, r *http.Request) {

// 	// Retrieve a span for a web request attached to a Go Context.
// 	span, _ := tracer.StartSpanFromContext(r.Context(), "fileOpener")
// 	// Creating an error by opening a file that does not exist
// 	_, err := os.Open("filename.ext")
// 	// Set the error on the span
// 	span.Finish(tracer.WithError(err))

// 	w.Write([]byte("Setting an error on the span."))
// }

// func add_span(w http.ResponseWriter, r *http.Request) {
// 	// Create a span which will be the child of the span in the Context ctx, if there is a span in the context.
// 	parentSpan, _ := tracer.StartSpanFromContext(r.Context(), "parent.span", tracer.ResourceName("context-span"))
// 	// Creating a children to this new span
// 	span := tracer.StartSpan("waiting", tracer.ResourceName("With Childof"), tracer.ChildOf(parentSpan.Context()))
// 	time.Sleep(1 * time.Second)
// 	span.Finish()
// 	parentSpan.Finish()

// 	w.Write([]byte("Adding a child span from context manually, then adding a child to this new span."))
// }

// func main() {
// 	tracer.Start(
// 		// Adding a global tag
// 		tracer.WithGlobalTag("team", "go_sandbox"),
// 	)
// 	defer tracer.Stop()

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8000"
// 	}

// 	mux := httptrace.NewServeMux()

// 	mux.HandleFunc("/", say_hello)

// 	mux.HandleFunc("/add-tag", add_tag)

// 	mux.HandleFunc("/set-error", set_error)

// 	mux.HandleFunc("/add-span", add_span)

// 	http.ListenAndServe(":"+port, mux)
// }
//
//
// package main

// import (
//     "io"
//     "net/http"
//     "os"
//     "time"

//     httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
//     "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
// )

// // sayHello is the public handler a client will hit.
// // It now fan‑outs to an internal endpoint and does some extra local work so
// // the Datadog trace shows multiple spans.
// func sayHello(w http.ResponseWriter, r *http.Request) {
//     // ➜ child span that groups everything we do in this handler
//     span, ctx := tracer.StartSpanFromContext(r.Context(), "say_hello.handler", tracer.ResourceName("say_hello"))
//     defer span.Finish()

//     // --- downstream HTTP call (still inside the same binary) -------------
//     client := httptrace.WrapClient(&http.Client{Timeout: 2 * time.Second})

//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8000"
//     }
//     req, _ := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:"+port+"/internal-work", nil)

//     // propagate a synthetic request‑id just to prove headers flow
//     req.Header.Set("X-Request-ID", time.Now().Format(time.RFC3339Nano))

//     resp, err := client.Do(req)
//     if err != nil {
//         span.SetTag("error", err)
//     } else {
//         io.Copy(io.Discard, resp.Body)
//         resp.Body.Close()
//     }

//     // --- extra local span ------------------------------------------------
//     local := tracer.StartSpan("local.computation", tracer.ChildOf(span.Context()))
//     time.Sleep(150 * time.Millisecond) // pretend to work
//     local.Finish()

//     w.Write([]byte("Hello World with fan‑out call!"))
// }

// // internalWork is hit only by the internal HTTP client in sayHello.
// // It adds its own nested span so you get a multi‑service‑looking trace.
// func internalWork(w http.ResponseWriter, r *http.Request) {
//     // simulate business logic
//     time.Sleep(100 * time.Millisecond)

//     if parent, ok := tracer.SpanFromContext(r.Context()); ok {
//         sub := tracer.StartSpan("internal.calculation", tracer.ChildOf(parent.Context()))
//         time.Sleep(50 * time.Millisecond)
//         sub.Finish()
//     }

//     w.Write([]byte("Internal work done"))
// }

// // addTag keeps the original logic but renames for Go style.
// func addTag(w http.ResponseWriter, r *http.Request) {
//     if span, ok := tracer.SpanFromContext(r.Context()); ok {
//         span.SetTag("new", "tag")
//     }
//     w.Write([]byte("Adding a tag."))
// }

// func setError(w http.ResponseWriter, r *http.Request) {
//     span, _ := tracer.StartSpanFromContext(r.Context(), "fileOpener")
//     _, err := os.Open("filename.ext")
//     span.Finish(tracer.WithError(err))
//     w.Write([]byte("Setting an error on the span."))
// }

// func addSpan(w http.ResponseWriter, r *http.Request) {
//     parentSpan, _ := tracer.StartSpanFromContext(r.Context(), "parent.span", tracer.ResourceName("context-span"))
//     child := tracer.StartSpan("waiting", tracer.ResourceName("With Childof"), tracer.ChildOf(parentSpan.Context()))
//     time.Sleep(1 * time.Second)
//     child.Finish()
//     parentSpan.Finish()
//     w.Write([]byte("Adding a child span from context manually, then adding a child to this new span."))
// }

// func main() {
//     tracer.Start(tracer.WithGlobalTag("team", "go_sandbox"))
//     defer tracer.Stop()

//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8000"
//     }

//     mux := httptrace.NewServeMux()
//     mux.HandleFunc("/", sayHello)
//     mux.HandleFunc("/internal-work", internalWork)
//     mux.HandleFunc("/add-tag", addTag)
//     mux.HandleFunc("/set-error", setError)
//     mux.HandleFunc("/add-span", addSpan)

//     http.ListenAndServe(":"+port, mux)
// }

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
    time.Sleep(150 * time.Millisecond)

    w.Write([]byte("Hello World with fan-out call!"))
}

// internalWork is reached only by the internal HTTP client in sayHello.
// It performs some fake CPU work to illustrate nested spans (if tracing
// is added back later).
func internalWork(w http.ResponseWriter, r *http.Request) {
    // simulate business logic
    time.Sleep(100 * time.Millisecond)

    // nested pretend calculation
    time.Sleep(50 * time.Millisecond)

    w.Write([]byte("Internal work done"))
}

func addTag(w http.ResponseWriter, r *http.Request) {
    // placeholder – no tracing, just respond
    w.Write([]byte("Adding a tag."))
}

func setError(w http.ResponseWriter, r *http.Request) {
    // trigger an error by opening a non‑existent file so we can see logs
    if _, err := os.Open("filename.ext"); err != nil {
        log.Printf("setError simulated failure: %v", err)
    }
    w.Write([]byte("Setting an error (simulated)."))
}

func addSpan(w http.ResponseWriter, r *http.Request) {
    // Simulate a blocking operation
    time.Sleep(1 * time.Second)
    w.Write([]byte("Simulated additional span work."))
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
    mux.HandleFunc("/add-tag", addTag)
    mux.HandleFunc("/set-error", setError)
    mux.HandleFunc("/add-span", addSpan)

    log.Printf("listening on :%s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}


