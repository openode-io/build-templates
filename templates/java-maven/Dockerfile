FROM openjdk:8-jdk-alpine

WORKDIR /opt/app

# add maven for building the application
RUN apk add maven

RUN touch /usr/bin/start.sh # this is the script which will run on start

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /usr/bin/start.sh
RUN echo 'crond' >> /usr/bin/start.sh

# Build with maven
RUN echo 'mvn package -DskipTests' >> /usr/bin/start.sh

# IMPORTANT: change "YOUR-JAR-FILE"
RUN echo 'java -jar target/YOUR-JAR-FILE.jar' >> /usr/bin/start.sh
