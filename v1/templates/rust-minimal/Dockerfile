FROM rust:alpine

WORKDIR /opt/app

ENV PORT=80

# daemon for cron jobs
RUN echo '' > /boot.sh
#RUN echo 'crond' > /boot.sh
# RUN echo 'crontab .openode.cron' >> /boot.sh

COPY . .

RUN cargo install --force
RUN cargo build --release

# start it!
CMD sh /boot.sh && cargo run --release