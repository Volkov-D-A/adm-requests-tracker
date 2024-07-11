FROM golang:latest as build
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -v ./cmd/adm-tsr

FROM alpine
COPY --from=build /app/apiserver /app/
WORKDIR /app
EXPOSE 8080