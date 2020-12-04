FROM golang as runtime

COPY . source

WORKDIR source/bin
CMD chmod +x release
EXPOSE 9100
RUN ["./release"]