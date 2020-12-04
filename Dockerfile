FROM golang as runtime
WORKDIR source
EXPOSE 9100
CMD["bin/release","run"]