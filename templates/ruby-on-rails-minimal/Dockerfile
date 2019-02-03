FROM ruby:2.5.1-alpine

ENV PORT=80
ENV RAILS_ENV=production
ENV RACK_ENV=production
ENV RAILS_ROOT /opt/app

WORKDIR /opt/app

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

RUN apk add --no-cache build-base tzdata git nodejs # sqlite-dev || postgressql-dev

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /boot.sh

# if you need redis, uncomment the lines below
# RUN apk --update add redis
# RUN echo 'redis-server &' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

RUN echo 'echo bundle install' >> /boot.sh
RUN echo 'bundle install --jobs 20 --retry 5 --without development test' >> /boot.sh

RUN echo 'echo assets:precompile' >> /boot.sh
RUN echo 'bundle exec rake assets:precompile' >> /boot.sh

# launch the application
RUN echo 'echo starting the application' >> /boot.sh

CMD sh /boot.sh && bundle exec puma -C config/puma.rb
