# ------------------------------- DOWNLOAD go DEPENDENCIES
FROM golang:1.12-stretch AS dependencies

WORKDIR /pod-vis
COPY go.mod go.sum ./
RUN go mod download

# ------------------------------- BUILD AND EXECUTE TESTS
FROM dependencies AS builder

# copy all sources to the docker image
COPY . .

RUN gofmt -l ./
RUN test -z $(gofmt -l ./)

ARG commit_sha_arg
ARG timestamp_arg

ENV COMMIT_SHA=$commit_sha_arg
ENV TIMESTAMP=$timestamp_arg

RUN CGO_ENABLED=0 go build -o ./pod-visualizer ./internal/main.go

# ------------------------------ BUILD THE REAL CONTAINER
FROM alpine:latest

RUN addgroup -S -g 1001 gopher && adduser -S -G gopher -u 1001 gopher

# copy golang binary
COPY --from=builder --chown=gopher:gopher /pod-vis/pod-visualizer /home/gopher
RUN chmod +x /home/gopher/pod-visualizer

USER gopher

ENTRYPOINT /home/gopher/pod-visualizer
