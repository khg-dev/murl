FROM ubuntu as runtime

COPY bin/murl /usr/bin
WORKDIR /usr/bin
RUN chmod +x murl
EXPOSE 9100
CMD ["./murl"]