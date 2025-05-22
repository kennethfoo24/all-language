// Datadog tracer
using Datadog.Trace;

// Adding System to add the date as tag in a span
using System;

using Serilog;
using Serilog.Context;
using Serilog.Formatting.Json;

// Initialize a new instance of the WebApplication builder
var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

// Regardless of the output layout, your LoggerConfiguration must be
// enriched from the LogContext to extract the Datadog
// properties that are automatically injected by the .NET tracer
//
// Additions to LoggerConfiguration:
// - .Enrich.FromLogContext()
var loggerConfiguration = new LoggerConfiguration()
                              .Enrich.FromLogContext()
                              .MinimumLevel.Is(Serilog.Events.LogEventLevel.Information);

// configure serilog to output to console so the datadog docker agent can pick it up

// raw version
// loggerConfiguration = loggerConfiguration
//                           .WriteTo.Console(
//                               outputTemplate: "{Timestamp:yyyy-MM-dd HH:mm:ss.fff zzz} [{Level:u3}] {Properties} {Message:lj} {NewLine}{Exception}");

// json version
loggerConfiguration = loggerConfiguration
                          .WriteTo.Console(
                              new JsonFormatter());

// Main procedure
var log = loggerConfiguration.CreateLogger();

// Respond "Hello world!" on /
app.MapGet("/", () => "Hello World!");

// Endpoint where we add a tag
app.MapGet("/add-tag", () =>
{

  log.Information("Adding a tag");

  // Getting the active span
  var scope = Tracer.Instance.ActiveScope;

  if (scope != null)
  {
    // Adding a custom span tag for use in the Datadog UI. 
    // The tag added here is the current date. The UI already allows scoping by date, this is just for example purposes.

    var currentDateTime = DateTime.Today.ToString();

    scope.Span.SetTag("date", currentDateTime);
    log.Information("Date tag has been set to current time {CurrentTime}", currentDateTime);
  }

  return "Here we add a tag";
});

// Endpoint where we throw an exception
app.MapGet("/exception", () =>
{
  // Getting the active span
  var scope = Tracer.Instance.ActiveScope;

  if (scope != null)
  {
    // Throwing an exception by dividing by 0
    int numer = 1;
    int denom = 0;
    int result;
    try
    {
      result = numer / denom;
    }
    catch (Exception e)
    {
      log.Error(e, "Exception has been thrown!");
      scope.Span.SetException(e);
    }
  }

  return "Throwing an exeption!";
});

// Endpoint where we add a custom span
app.MapGet("/manual-span", () =>
{
  log.Information("Adding a custom span");
  using (var parentScope = Tracer.Instance.StartActive("manual.sortorders"))
  {
    parentScope.Span.ResourceName = "Parent Resource";
    using (var childScope = Tracer.Instance.StartActive("manual.sortorders.child"))
    {
      // Nest using statements around the code to trace
      childScope.Span.ResourceName = "Child Resource";
    }
  }

  return "Adding a span manually";
});

app.Run();