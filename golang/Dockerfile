FROM golang:1.22

WORKDIR /app

# Download Go modules
COPY golang/ .
RUN go install github.com/datadog/orchestrion@latest
RUN go mod download

# Build
RUN orchestrion go build -o hello-world

EXPOSE 8000

# Run
CMD [ "./hello-world" ]

# ####################  build stage  ####################
# # Orchestrion supports the two most-recent Go releases → use 1.22
# FROM golang:1.22 AS build

# WORKDIR /src

# # 2) copy module files and let orchestrion pin itself
# COPY golang/ .

# # 1) install the orchestrion CLI
# RUN go install github.com/datadog/orchestrion@v0.8

# RUN go mod download       # <— pin happens *inside* the image

# # 3) copy source and build through orchestrion
# # RUN orchestrion go build -ldflags="-s -w" -o /app/server ./main.go
# RUN orchestrion go build -o hello-world
# ####################  runtime stage  ##################
# FROM gcr.io/distroless/base-debian12
# WORKDIR /app
# COPY --from=build /hello-world .
# EXPOSE 8000
# USER nonroot:nonroot
# #ENTRYPOINT ["/app/server"]
# CMD [ "./hello-world" ]
