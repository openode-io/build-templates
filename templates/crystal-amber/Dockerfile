FROM amberframework/amber:latest

# !Important: make sure to change the project at the bottom of this file

ENV AMBER_ENV=production

WORKDIR /opt/app

RUN apt-get update && apt-get install -y libssl-dev libyaml-dev

RUN touch /usr/bin/start.sh # this is the script which will run on start

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /usr/bin/start.sh

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /usr/bin/start.sh
RUN echo 'crond' >> /usr/bin/start.sh

# compile shards:
RUN echo 'shards install --production' >> /usr/bin/start.sh

# uncomment if you want to create
# RUN echo 'amber db create' >> /usr/bin/start.sh
RUN echo 'amber db migrate' >> /usr/bin/start.sh

# IMPORTANT: change "YOUR-APP-NAME" to your project name (as in shard.yml)
RUN echo 'shards build YOUR-APP-NAME --release' >> /usr/bin/start.sh
RUN echo 'bin/YOUR-APP-NAME' >> /usr/bin/start.sh
