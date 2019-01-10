FROM google/dart
# debian-based image

WORKDIR /opt/app

RUN echo 'set -e' > /usr/bin/start.sh # this is the script which will run on start

RUN echo 'pub get' >> /usr/bin/start.sh
RUN echo 'pub get --offline' >> /usr/bin/start.sh

# activate webdev:
RUN echo 'pub global activate webdev' >> /usr/bin/start.sh
RUN echo 'export PATH="$PATH":"$HOME/.pub-cache/bin"' >> /usr/bin/start.sh

# build
RUN echo 'webdev build' >> /usr/bin/start.sh

# start the server
RUN echo 'dart bin/server.dart' >> /usr/bin/start.sh
