FROM privetkakdela/devops_course:go_builder as builder
WORKDIR /go/src/app
COPY ./source/app_go.go .
RUN CGO_ENABLED=0 GOOS=linux go build app_go.go
RUN ls -l

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/src/app/app_go .
ENTRYPOINT ["./app_go"]
