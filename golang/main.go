package main

import (
	"net/http"
	"os"
	"time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func say_hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func add_tag(w http.ResponseWriter, r *http.Request) {

	// Retrieve a span for a web request attached to a Go Context.
	if span, ok := tracer.SpanFromContext(r.Context()); ok {
		// Set tag
		span.SetTag("new", "tag")
	}

	w.Write([]byte("Adding a tag."))
}

func set_error(w http.ResponseWriter, r *http.Request) {

	// Retrieve a span for a web request attached to a Go Context.
	span, _ := tracer.StartSpanFromContext(r.Context(), "fileOpener")
	// Creating an error by opening a file that does not exist
	_, err := os.Open("filename.ext")
	// Set the error on the span
	span.Finish(tracer.WithError(err))

	w.Write([]byte("Setting an error on the span."))
}

func add_span(w http.ResponseWriter, r *http.Request) {
	// Create a span which will be the child of the span in the Context ctx, if there is a span in the context.
	parentSpan, _ := tracer.StartSpanFromContext(r.Context(), "parent.span", tracer.ResourceName("context-span"))
	// Creating a children to this new span
	span := tracer.StartSpan("waiting", tracer.ResourceName("With Childof"), tracer.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	span.Finish()
	parentSpan.Finish()

	w.Write([]byte("Adding a child span from context manually, then adding a child to this new span."))
}

func main() {
	tracer.Start(
		// Adding a global tag
		tracer.WithGlobalTag("team", "go_sandbox"),
	)
	defer tracer.Stop()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := httptrace.NewServeMux()

	mux.HandleFunc("/", say_hello)

	mux.HandleFunc("/add-tag", add_tag)

	mux.HandleFunc("/set-error", set_error)

	mux.HandleFunc("/add-span", add_span)

	http.ListenAndServe(":"+port, mux)
}
