FROM ubuntu as runtime
ADD https://murl-releases.s3.us-east-2.amazonaws.com/releases/murl /usr/bin
RUN chmod +777 /usr/bin/murl
EXPOSE 9100
CMD /usr/bin/murl