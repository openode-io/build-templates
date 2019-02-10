FROM clojure:lein-2.8.1-alpine

WORKDIR /opt/app

ENV PORT=80

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /boot.sh

# daemon for cron jobs
# RUN echo 'echo will install crond...' >> /boot.sh
# RUN echo 'crond' >> /boot.sh

RUN echo 'lein uberjar' >> /boot.sh

# start it!
CMD sh /boot.sh && java -jar target/uberjar/YOUR-JAR-NAME.jar
