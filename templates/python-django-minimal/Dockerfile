FROM python:3.7-alpine

WORKDIR /opt/app

RUN touch /usr/bin/start.sh # this is the script which will run on start

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /usr/bin/start.sh

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /usr/bin/start.sh
RUN echo 'crond' >> /usr/bin/start.sh

RUN echo 'pip install -r requirements.txt' >> /usr/bin/start.sh

# run migrations
RUN echo 'python manage.py migrate' >> /usr/bin/start.sh

# start the server
RUN echo 'python manage.py runserver 0.0.0.0:80' >> /usr/bin/start.sh
