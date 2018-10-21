FROM scratch
ADD ./assets/ca-certificates.crt /etc/ssl/certs/

ADD ./build/tasks /
CMD ["./tasks"]

EXPOSE 3000