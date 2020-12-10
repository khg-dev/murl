FROM ubuntu as runtime

COPY bin/murl /usr/bin
WORKDIR /usr/bin
RUN ls -la
RUN chmod +x /usr/bin/murl1

EXPOSE 9100
CMD ["/usr/bin/murl"]