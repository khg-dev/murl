FROM golang as runtime

COPY source source

WORKDIR source
EXPOSE 9100
RUN ["bin/release"]