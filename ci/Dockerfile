FROM ubuntu as runtime

COPY bin/release /usr/bin

RUN chmod +x /usr/bin/release
EXPOSE 9100

ENTRYPOINT ["/usr/bin/release"]
 