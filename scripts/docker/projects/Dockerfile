FROM scratch
ADD ./assets/ca-certificates.crt /etc/ssl/certs/

ADD ./build/projects /
CMD ["./projects"]

EXPOSE 3000