FROM crystallang/crystal:latest-alpine

WORKDIR /opt/app

ENV PORT=80

# daemon for cron jobs
RUN echo 'crond' > /boot.sh
# RUN echo 'crontab .openode.cron' >> /boot.sh

# Bundle app source
COPY . .

RUN if test -e shard.yml; then shards install; fi
RUN crystal build server.cr --release

CMD sh /boot.sh && ./server