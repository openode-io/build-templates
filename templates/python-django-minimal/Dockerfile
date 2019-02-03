FROM python:3.7-alpine

WORKDIR /opt/app

RUN touch /boot.sh # this is the script which will run on start

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /boot.sh

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

RUN echo 'pip install -r requirements.txt' >> /boot.sh

# run migrations
RUN echo 'python manage.py migrate' >> /boot.sh

# start the server
CMD sh /boot.sh && python manage.py runserver 0.0.0.0:80
