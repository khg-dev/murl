FROM ubuntu as runtime
ADD https://murl-releases.s3.us-east-2.amazonaws.com/releases/murl /usr/bin
WORKDIR /usr/bin
RUN ls -la
RUN chmod +777 murl
EXPOSE 9100
CMD ./murl