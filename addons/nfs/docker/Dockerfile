# alpine:3.19.1
FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

RUN apk update && \
    apk add --no-cache python3 \
    nfs-utils \
    jq

COPY fluid_config_init.py /
COPY entrypoint.sh /usr/local/bin

RUN chmod u+x /usr/local/bin/entrypoint.sh

CMD ["/usr/local/bin/entrypoint.sh"]
