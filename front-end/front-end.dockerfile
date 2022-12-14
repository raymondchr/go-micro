FROM alpine:latest

RUN mdkir /app

COPY frontEndApp /app

CMD [ "/app/frontEndApp"]