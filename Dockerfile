FROM golang as runtime

COPY . source

WORKDIR source
CMD chmod +x release
EXPOSE 9100
CMD ["./release"]