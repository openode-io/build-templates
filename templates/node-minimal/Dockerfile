FROM node:10-alpine

WORKDIR /opt/app

ENV PORT=80

COPY package*.json ./

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /boot.sh

# daemon for cron jobs
RUN echo 'crond' >> /boot.sh

RUN npm install --production

# npm start, make sure to have a start attribute in "scripts" in package.json
CMD sh /boot.sh && npm start
