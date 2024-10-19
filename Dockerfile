# https://www.programonaut.com/pocketbase-as-a-framework-deploy-with-front-and-backend/

FROM golang:latest as builder
WORKDIR /build
COPY ./ ./
RUN go mod download
RUN go build -o ./main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main ./main
EXPOSE 8090
ENTRYPOINT ["./main"]
CMD ["serve", "--http=0.0.0.0:8090"]