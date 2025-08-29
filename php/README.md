# php-apache

## Instructions

This spins up a php application on port 8080.

The agent is containerize. Make sure that in your `~` directory, you have a file called sandbox.docker.env that contains:

```DD_API_KEY=<Your API Key>```

This is where the agent will read the API key.

Heads up to the fact that this sandbox was made on an M1 computer, so the architecture is `arm64`. If you are on an Intel computer, you will need to change the architecture to amd.

The architecture can be changed under `services.simple-dotnet.build.args.ARCH` (`arm64` for arm, and `amd64` for amd).

Launch with `./run.sh`. This runs the application in detached mode.

Then connect to:
[http://localhost:8080](http://localhost:8080)
![image](https://user-images.githubusercontent.com/65819327/215062150-fb0cba5b-3dad-4c66-8112-6cb64f8a223e.png)

The version of the agent in use is 7.41.1. It can be changed in the `docker-compose.yaml` file under `services.datadog-simple-php.image`.

The version of the tracer in use is 0.82.0, and the tracer used is for arm architecture. The version can be changed in the `docker-compose.yaml` file under `services.simple-php.build.args.VERSION`.

You can run an interactive shell on the container with:

```docker exec -it simple-php sh```


You can set the tracer to debug by changing `DD_TRACE_DEBUG` to `true` in the `docker-compose.yaml` under `simple-php.environment`. You can then access the container logs with:

```docker logs simple-php```

## Endpoints

Endpoints are defined in the `app` folder:
* `/`: this endpoints returns a hello world, and generate a trace with no custom instrumentation.
![image](https://user-images.githubusercontent.com/65819327/215062777-13a95611-7147-4cb1-bc25-8b25a4bd52eb.png)

* `/add-tag.php`: here, using custom instrumentation, we add a tag to the active span.
![image](https://user-images.githubusercontent.com/65819327/215062806-c6b8fb1c-5f2e-45ef-8b20-0cabff4b931e.png)

* `/exception.php`: here we throw an error and the error is added in the span.
![image](https://user-images.githubusercontent.com/65819327/215062837-04cb4a26-e1fa-4154-984d-83872f35a715.png)

* `/add-span.php`: we add a child span to the current one everytime a given function is called.
![image](https://user-images.githubusercontent.com/65819327/215062861-b3e6e99f-64df-4431-8a36-14b84aff2158.png)


## Tear down

Run `docker-compose down`
