package com.example.springboot;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

// import thread to sleed when creating new spans
import java.lang.Thread;
// import the LocalDate class to add the date as a tag
import java.time.LocalDate; 

// Datadog tracer
import io.opentracing.util.GlobalTracer;
import io.opentracing.Span;
import io.opentracing.Scope;
import io.opentracing.tag.Tags;
import io.opentracing.log.Fields;
import io.opentracing.Tracer;
import java.util.Collections;
import datadog.trace.api.Trace;
import datadog.trace.api.DDTags;

// Making a class based on RestController
@RestController
public class HelloController {

	// Respond "Hello world!" on /
	@GetMapping(value="/")
	public String index() {
		return "Hello World!";
	}

	// Addind a tag to the current span
	@GetMapping(value="/add-tag")
	public String addTag() {
		// Getting the span
		final Span span = GlobalTracer.get().activeSpan();
		if (span != null) {
			// Create a date object
			LocalDate myDate = LocalDate.now(); 
			// Adding the date as tag to add extra information, but this is not required for spans in general since the tracer will add the timestamp to all traces.
			span.setTag("date", myDate.toString());
		}
		return "Adding a tag";
	}

	// Setting an error to the current span when a given exception happens
	@GetMapping(value="/java-set-error")
	public String setError() {
		// Getting the span
		final Span span = GlobalTracer.get().activeSpan();
		if (span != null) {
			// creating an error by accessing non existing element in list
			try{
				int[] smallArray = {1};
				System.out.println(smallArray[1]);
			} catch (Exception ex){
				// Since the error is catched, the http status code will be 200 and the span will have status Ok
				// So we use the following line to force the error on the span here.
				span.setTag(Tags.ERROR, true);
				span.log(Collections.singletonMap(Fields.ERROR_OBJECT, ex));
			}
		}
		return "Setting an error.";
	}

	// Following we use a function that we will annotate to trace it when it's called.
	// This is a waiting function.
	@Trace(operationName = "manual.span", resourceName = "Waiting")
	public void Waiting() {
		try{
			Thread.sleep(1000);
		} catch (Exception e){
			return;
		}
	}

	// We call the previous function and therefore create a span with trace annotation.
	@GetMapping(value="/trace-annotation")
	public String traceAnnotation() {
		Waiting();
		return "Creating a span with trace annotations";
	}

	// Manually creating a new span
	@GetMapping(value="/manual-span")
	public String manualSpan() {
		Tracer tracer = GlobalTracer.get();
		// Setting the spans service, resource and operation name
		Span span = tracer.buildSpan("manual.span")
            .withTag(DDTags.SERVICE_NAME, "java-manual")
            .withTag(DDTags.RESOURCE_NAME, "Manual")
            .start();

			try (Scope scope = tracer.activateSpan(span)) {
				// Alternatively, set tags after creation
				span.setTag("state", "crafted");
	
				try{
					Thread.sleep(1000);
				} catch (Exception e){}
	
			}

			span.finish();

		return "Manually creating a span";
	}
}
