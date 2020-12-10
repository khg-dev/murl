FROM ubuntu

COPY bin/murl /usr/bin
RUN echo $'#!/bin/sh \n\
/usr/bin/murl' > run.sh
RUN chmod 777 run.sh /usr/bin/murl
RUN ls -la
EXPOSE 9100
CMD "ls -la /usr/"