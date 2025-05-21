# FROM golang:1.18

# WORKDIR /app

# # Download Go modules
# COPY golang/* .
# RUN go mod download

# # Build
# RUN go build -o hello-world

# EXPOSE 8000

# # Run
# CMD [ "./hello-world" ]

####################  build stage  ####################
# Orchestrion supports the two most-recent Go releases → use 1.22
FROM golang:1.22-bullseye-slim AS build

# 1) install the orchestrion CLI
RUN go install github.com/datadog/orchestrion@v0.8.0

WORKDIR /src

# 2) copy module files and let orchestrion pin itself
COPY golang/* .
RUN orchestrion pin && go mod download       # <— pin happens *inside* the image

# 3) copy source and build through orchestrion
RUN orchestrion go build -ldflags="-s -w" -o /app/server ./main.go

####################  runtime stage  ##################
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8000
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
