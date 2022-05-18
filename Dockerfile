FROM golang:1.18-alpine as builder
WORKDIR /build 
COPY . .
RUN go mod download
RUN go build -o xd_working_trial
ENTRYPOINT ["./xd_working_trial"]


FROM golang:1.18-alpine

RUN addgroup -g 1000 xduser \
    && adduser -u 1000 -G xduser -s /bin/sh -D xduser \
    && mkdir -p /app \
    && chown -R xduser:xduser /app 
COPY --from=builder /build/statics /app/statics
COPY --from=builder /build/xd_working_trial /app/xd_working_trial

WORKDIR /app
USER xduser

EXPOSE 8080

ENTRYPOINT ["./xd_working_trial"]