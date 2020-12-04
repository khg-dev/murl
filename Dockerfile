FROM golang as runtime

COPY . source

WORKDIR source
EXPOSE 9100
RUN ["bin/release"]