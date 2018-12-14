FROM elixir:1.7.4-alpine

ARG homedir="/opt/app"
WORKDIR $homedir

ENV PORT=80
ENV MIX_ENV=prod

RUN echo 'set -e' > /usr/bin/start.sh # this is the script which will run on start

RUN apk add nodejs npm make gcc erlang-dev libc-dev

# Local PostgreSQL (optional)
RUN apk add postgresql
RUN echo 'mkdir -p db' >> /usr/bin/start.sh

RUN echo 'mkdir -p /run/postgresql' >> /usr/bin/start.sh
RUN echo 'chmod -R 777 /run' >> /usr/bin/start.sh
RUN echo 'chmod -R 700 /opt/app/db' >> /usr/bin/start.sh
RUN echo 'chown -R postgres /opt/app/db' >> /usr/bin/start.sh
RUN echo 'su -c "initdb /opt/app/db/ || true" postgres' >> /usr/bin/start.sh
RUN echo 'su -c "postgres -D /opt/app/db/ &" postgres' >> /usr/bin/start.sh
RUN echo 'sleep 5' >> /usr/bin/start.sh

# Daemon for cron jobs (optional)
# RUN echo 'echo will install crond...' >> /usr/bin/start.sh
# RUN echo 'crond' >> /usr/bin/start.sh

# Main installation
RUN echo 'yes | mix local.hex' >> /usr/bin/start.sh
RUN echo 'yes | mix local.rebar' >> /usr/bin/start.sh
RUN echo 'mix deps.get' >> /usr/bin/start.sh
RUN echo 'cd assets && npm install && npm run deploy' >> /usr/bin/start.sh
RUN echo "cd $homedir" >> /usr/bin/start.sh
RUN echo "mix ecto.create" >> /usr/bin/start.sh
RUN echo "mix ecto.migrate" >> /usr/bin/start.sh

# start it!
RUN echo "mix phx.server" >> /usr/bin/start.sh
