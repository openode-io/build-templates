FROM clojure:lein-2.8.1-alpine

WORKDIR /opt/app

ENV PORT=80

RUN echo 'set -e' > /usr/bin/start.sh # this is the script which will run on start

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
# RUN echo 'echo will install crond...' >> /usr/bin/start.sh
# RUN echo 'crond' >> /usr/bin/start.sh

RUN echo 'lein uberjar' >> /usr/bin/start.sh

# start it!
RUN echo 'java -jar target/uberjar/YOUR-JAR-NAME.jar' >> /usr/bin/start.sh
