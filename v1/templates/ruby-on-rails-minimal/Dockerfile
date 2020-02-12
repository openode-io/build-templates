FROM ruby:2.7.0-alpine

ENV PORT=80
ENV RAILS_ENV=production
ENV RACK_ENV=production
ENV RAILS_ROOT /opt/app

WORKDIR /opt/app

RUN mkdir -p /opt/app
RUN touch /boot.sh # this is the script which will run on start

RUN apk add --no-cache build-base tzdata git sqlite-dev nodejs # sqlite-dev || postgressql-dev

# if you need a build script, uncomment the line below
# RUN echo 'sh mybuild.sh' >> /boot.sh

# daemon for cron jobs
RUN echo 'echo will install crond...' >> /boot.sh
RUN echo 'crond' >> /boot.sh

COPY Gemfile* ./
RUN bundle install --jobs 20 --retry 5 --without development test

COPY . .

# compile assets
RUN bundle exec rake assets:precompile

RUN echo 'echo db:migrate' >> /boot.sh
RUN echo 'bin/rails db:migrate' >> /boot.sh

# launch the application
RUN echo 'echo starting the application' >> /boot.sh
CMD sh /boot.sh && bundle exec puma -C config/puma.rb
