FROM hseeberger/scala-sbt:8u181_2.12.7_1.2.6

WORKDIR /opt/app

ENV PORT=80

RUN touch /usr/bin/start.sh # this is the script which will run on start

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
# RUN echo 'echo will install crond...' >> /usr/bin/start.sh
# RUN echo 'crond' >> /usr/bin/start.sh

# start it!
RUN echo 'sbt "run 80"' >> /usr/bin/start.sh
