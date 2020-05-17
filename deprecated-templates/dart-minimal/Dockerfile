FROM google/dart
# debian-based image

WORKDIR /opt/app

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

RUN echo 'pub get' >> /boot.sh
RUN echo 'pub get --offline' >> /boot.sh

# activate webdev:
RUN echo 'pub global activate webdev' >> /boot.sh
RUN echo 'export PATH="$PATH":"$HOME/.pub-cache/bin"' >> /boot.sh

# build
RUN echo 'webdev build' >> /boot.sh

# start the server
CMD sh /boot.sh && dart bin/server.dart
