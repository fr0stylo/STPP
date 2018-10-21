FROM scratch
ADD ./assets/ca-certificates.crt /etc/ssl/certs/

ADD ./build/time-entries /
CMD ["./time-entries"]

EXPOSE 3000