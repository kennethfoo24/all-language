// Adding System to add the date as tag in a span
using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Http;

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

// 1) /dotnet: greeting + upstream service call
app.MapGet("/dotnet", async (HttpContext context) =>
{
    var serviceUrl = Environment.GetEnvironmentVariable("A_SERVICE_URL");
    if (string.IsNullOrEmpty(serviceUrl))
    {
        context.Response.StatusCode = 500;
        await context.Response.WriteAsync("A_SERVICE_URL is not set");
        return;
    }

    using var client = new HttpClient { Timeout = TimeSpan.FromSeconds(5) };
    HttpResponseMessage upstream;
    try
    {
        upstream = await client.GetAsync(serviceUrl);
    }
    catch (Exception ex)
    {
        context.Response.StatusCode = 502;
        await context.Response.WriteAsync($"Failed to call upstream service: {ex.Message}");
        return;
    }

    var upstreamBody = await upstream.Content.ReadAsStringAsync();

    context.Response.StatusCode = (int)upstream.StatusCode;
    context.Response.ContentType = "text/plain";

    await context.Response.WriteAsync("Hello World from .NET!\n");
    await context.Response.WriteAsync(upstreamBody);
});


app.Run();