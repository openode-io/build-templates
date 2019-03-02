FROM rust:1.23.0

WORKDIR /opt/app

ENV PORT=80

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

# daemon for cron jobs
# RUN echo 'echo will install crond...' >> /boot.sh
# RUN echo 'crond' >> /boot.sh
# RUN echo 'crontab .openode.cron' >> /boot.sh

RUN echo 'cargo install --force' >> /boot.sh
RUN echo 'cargo build --release' >> /boot.sh

# start it!
CMD sh /boot.sh && target/release/YOUR-BINARY-FILE
