FROM node:10-alpine

WORKDIR /opt/app

ENV PORT=80

# Extras:
RUN apk --update add make python gcc g++

# Globally installed NPMs:
RUN npm install -g nodemon
RUN npm install -g pm2

RUN echo 'set -e' > /boot.sh # this is the script which will run on start

# Basic npm start verification
RUN echo 'nb=`cat package.json | grep start | wc -l` && if test "$nb" = "0" ; then echo "*** Boot issue: No start command found in your package.json in the scripts. See https://docs.npmjs.com/cli/start" ; exit 1 ; fi' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

# install packages
RUN echo 'npm install --production' >> /boot.sh

# logs by default are in logs
RUN mkdir -p logs

# npm start, make sure to have a start attribute in "scripts" in package.json
RUN echo 'pm2 --output logs/out.log --error logs/error.log start npm -- start' >> /boot.sh
RUN echo 'sh' >> /boot.sh

CMD sh /boot.sh
