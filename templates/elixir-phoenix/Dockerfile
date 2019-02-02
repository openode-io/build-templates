FROM elixir:1.8-alpine

ARG homedir="/opt/app"
WORKDIR $homedir

ENV PORT=80
ENV MIX_ENV=prod

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

RUN apk add nodejs npm make gcc erlang-dev libc-dev

# Local PostgreSQL (optional)
RUN apk add postgresql
RUN echo 'mkdir -p db' >> /boot.sh

RUN echo 'mkdir -p /run/postgresql' >> /boot.sh
RUN echo 'chmod -R 777 /run' >> /boot.sh
RUN echo 'chmod -R 700 /opt/app/db' >> /boot.sh
RUN echo 'chown -R postgres /opt/app/db' >> /boot.sh
RUN echo 'su -c "initdb /opt/app/db/ || true" postgres' >> /boot.sh
RUN echo 'su -c "postgres -D /opt/app/db/ &" postgres' >> /boot.sh
RUN echo 'sleep 5' >> /boot.sh

# Daemon for cron jobs (optional)
# RUN echo 'echo will install crond...' >> /boot.sh
# RUN echo 'crond' >> /boot.sh

# Main installation
RUN echo 'yes | mix local.hex' >> /boot.sh
RUN echo 'yes | mix local.rebar' >> /boot.sh
RUN echo 'mix deps.get' >> /boot.sh
RUN echo 'cd assets && npm install && npm run deploy' >> /boot.sh
RUN echo "cd $homedir" >> /boot.sh
RUN echo "mix ecto.create" >> /boot.sh
RUN echo "mix ecto.migrate" >> /boot.sh

# start it!
CMD sh /boot.sh && mix phx.server
