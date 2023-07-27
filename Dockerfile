FROM alpine

WORKDIR /root

COPY bin/IceFireDB-Proxy /root/IceFireDB-Proxy
COPY config/config.yaml /root/config/config.yaml


CMD /root/IceFireDB-Proxy -c /root/config/config.yaml
