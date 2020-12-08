FROM golang as runtime

COPY bin bin

WORKDIR bin
CMD chmod +x release
EXPOSE 9100
CMD ["./release"]