FROM golang:1.23.4 AS build
LABEL maintainer="rbennum"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o service-acc .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/.env .
COPY --from=build /app/service-acc .
COPY --from=build /app/database/migrate database/migrate
ENTRYPOINT [ "/app/service-acc" ]