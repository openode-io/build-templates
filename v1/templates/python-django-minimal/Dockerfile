FROM python:3

WORKDIR /opt/app

RUN touch /boot.sh # this is the script which will run on start

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

# run migrations
RUN echo 'python manage.py migrate' >> /boot.sh

# start the server
CMD sh /boot.sh && python manage.py runserver 0.0.0.0:80
