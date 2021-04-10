FROM golang:alpine

WORKDIR /opt/app

ENV PORT=80

# daemon for cron jobs
RUN echo 'crond' > /boot.sh
# RUN echo 'crontab .openode.cron' >> /boot.sh

# Bundle app source
COPY . .

CMD sh /boot.sh && go run server.go