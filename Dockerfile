# https://www.programonaut.com/pocketbase-as-a-framework-deploy-with-front-and-backend/

FROM golang:latest AS builder
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o /pocketbase

FROM alpine:latest
WORKDIR /app
COPY --from=builder /pocketbase /app/pocketbase

#EXPOSE 8090
#CMD ["/app/pocketbase", "serve", "--https=testing.snippy.info:8090"]

EXPOSE 3000
#CMD ["/app/pocketbase", "serve", "--http=testing.snippy.info:3000"]

CMD ["/app/pocketbase", "serve", "--http=0.0.0.0:8090"]
