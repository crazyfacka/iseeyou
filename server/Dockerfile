FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
ADD app /
ADD confs.toml /

EXPOSE 9000

CMD ["/app"]
