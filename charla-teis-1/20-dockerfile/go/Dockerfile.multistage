##
## Build
##
FROM golang:1.17-bullseye AS build

WORKDIR /build

COPY . .

RUN go build -o /app

##
## Run
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app /app

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/app"]