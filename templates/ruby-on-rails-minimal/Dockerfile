FROM ruby:2.5.1-alpine

ENV PORT=80
ENV RAILS_ENV=production
ENV RACK_ENV=production
ENV RAILS_ROOT /opt/app

WORKDIR /opt/app

RUN echo 'set -e' > /usr/bin/start.sh # this is the script which will run on start

RUN apk add --no-cache build-base tzdata git nodejs # sqlite-dev || postgressql-dev

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /usr/bin/start.sh

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /usr/bin/start.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /usr/bin/start.sh
RUN echo 'crond' >> /usr/bin/start.sh

RUN echo 'echo bundle install' >> /usr/bin/start.sh
RUN echo 'bundle install --jobs 20 --retry 5 --without development test' >> /usr/bin/start.sh

RUN echo 'echo assets:precompile' >> /usr/bin/start.sh
RUN echo 'bundle exec rake assets:precompile' >> /usr/bin/start.sh

# launch the application
RUN echo 'echo starting the application' >> /usr/bin/start.sh
RUN echo 'bundle exec puma -C config/puma.rb' >> /usr/bin/start.sh
