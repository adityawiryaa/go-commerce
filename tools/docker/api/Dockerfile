# argument for Go version
ARG GO_VERSION=1.18.4

# STAGE 1: building the executable
FROM golang:${GO_VERSION} AS builder

ENV APP_UID 1000
ENV APP_GID 1000
ENV APP_USER apps
ENV APP_SERVICE api
ENV APP_HOME /app

RUN groupadd -g $APP_GID $APP_USER
RUN useradd -u $APP_UID $APP_USER -g $APP_USER -M -l

USER $USER

COPY --chown=$APP_USER . .

RUN go mod init go-commerce && go mod tidy
RUN go mod vendor
RUN go mod verify

RUN go build $APP_HOME/services/$APP_SERVICE/cmd/main.go

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

ENV APP_HOME /app
WORKDIR $APP_HOME

COPY --from=builder $APP_HOME/main $APP_HOME/main

EXPOSE 80 443

CMD ["./main"]