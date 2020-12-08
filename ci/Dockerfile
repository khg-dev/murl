FROM ubuntu as runtime

COPY bin/murl /usr/bin
RUN chmod +x /usr/bin/murl

EXPOSE 9100
CMD ["/usr/bin/murl"]
 