FROM golang:1.23-alpine AS backend

WORKDIR /pb

RUN GOOS=linux go build -o snippy.exe .

FROM alpine:latest AS production
COPY --from=backend /pb .

EXPOSE 8090

CMD ["./snippy.exe", "serve", "--http=0.0.0.0:8080"]

# https://www.programonaut.com/pocketbase-as-a-framework-deploy-with-front-and-backend/