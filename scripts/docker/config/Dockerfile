FROM scratch

ADD ./assets/ca-certificates.crt /etc/ssl/certs/

ADD ./build/config /
ADD ./assets/config.toml /

CMD ["./config"]

EXPOSE 3000