FROM mhart/alpine-node:8

WORKDIR /opt/app

RUN echo 'npm install --production' > /usr/bin/start.sh

# daemon for cron jobs
RUN echo 'crond' >> /usr/bin/start.sh

# npm start, make sure to have a start attribute in "scripts" in package.json
RUN echo 'npm start' >> /usr/bin/start.sh
