
FROM golang:1.21 as build 

WORKDIR /app

COPY . . 

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /auth ./cmd/auth


FROM scratch 

WORKDIR /app 

COPY --from=build /auth /auth 

EXPOSE 8080

ENTRYPOINT [ "/auth" ]