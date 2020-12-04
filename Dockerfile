FROM golang as build

COPY . /source
RUN EXPORT GOPATH="/source"
WORKDIR /source
RUN go get -d ./...
RUN go build -o ../output/release

FROM golang as runtime

COPY --from=build /output/release .
EXPOSE 9100
CMD["main","run"]