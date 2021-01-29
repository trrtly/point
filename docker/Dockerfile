FROM uhub.service.ucloud.cn/youyao/ubuntu:20.04
ENV GODEBUG netdns=go

ADD release/point-server /usr/bin

LABEL com.centurylinklabs.watchtower.stop-signal="SIGINT"

ENTRYPOINT ["/usr/bin/point-server"]
