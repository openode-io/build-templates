FROM openjdk:8-jdk-alpine

WORKDIR /opt/app

# add maven for building the application
RUN apk add maven

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

# Build with maven
RUN echo 'mvn package -DskipTests' >> /boot.sh

# IMPORTANT: change "YOUR-JAR-FILE"
CMD sh /boot.sh && java -jar target/YOUR-JAR-FILE.jar
