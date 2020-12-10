FROM ubuntu as runtime

COPY bin/murl /usr/bin
WORKDIR /usr/bin
RUN chmod 777 murl
RUN ls -la
EXPOSE 9100
CMD murl