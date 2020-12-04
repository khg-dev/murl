FROM golang as runtime

COPY source source

WORKDIR source
EXPOSE 9100
CMD["bin/release","run"]