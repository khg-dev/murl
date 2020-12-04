FROM golang as runtime
WORKDIR bin
EXPOSE 9100
CMD["release","run"]