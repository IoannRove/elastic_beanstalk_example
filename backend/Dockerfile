##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY app/go.mod ./
COPY app/go.sum ./
RUN go mod download

COPY app ./
RUN go build -o /backend

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /backend /backend

USER nonroot:nonroot

ENTRYPOINT ["/backend"]