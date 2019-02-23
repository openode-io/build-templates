FROM amberframework/amber:latest

# !Important: make sure to change the project at the bottom of this file

ENV AMBER_ENV=production

WORKDIR /opt/app

RUN apt-get update && apt-get install -y libssl-dev libyaml-dev

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

# compile shards:
RUN echo 'shards install --production' >> /boot.sh

# uncomment if you want to create
# RUN echo 'amber db create' >> /boot.sh
RUN echo 'amber db migrate' >> /boot.sh

# IMPORTANT: change "YOUR-APP-NAME" to your project name (as in shard.yml)
RUN echo 'shards build YOUR-APP-NAME --release --static --no-debug' >> /boot.sh

CMD sh /boot.sh && bin/YOUR-APP-NAME
