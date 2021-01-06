FROM alpine:3.12 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3.12
ENV GODEBUG netdns=go

ADD release/point-server /bin/

RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

LABEL com.centurylinklabs.watchtower.stop-signal="SIGINT"

ENTRYPOINT ["/bin/point-server"]
