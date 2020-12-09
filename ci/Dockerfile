FROM ubuntu as runtime

COPY bin/murl /usr/bin
RUN chmod +x /usr/bin/murl.1

EXPOSE 9100
CMD ["/usr/bin/murl.1"]