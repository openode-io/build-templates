FROM rust:1.23.0

WORKDIR /opt/app

ENV PORT=80

RUN echo 'set -e' > /usr/bin/start.sh # this is the script which will run on start

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
# RUN echo 'echo will install crond...' >> /usr/bin/start.sh
# RUN echo 'crond' >> /usr/bin/start.sh

RUN echo 'cargo install --force' >> /usr/bin/start.sh
RUN echo 'cargo build --release' >> /usr/bin/start.sh

# start it!
RUN echo 'target/release/YOUR-BINARY-FILE' >> /usr/bin/start.sh
