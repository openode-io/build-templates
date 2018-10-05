FROM nginx:stable-alpine

WORKDIR /opt/app

RUN rm -rf /usr/share/nginx/html
RUN ln -s /opt/app /usr/share/nginx/html
RUN echo nginx >> /usr/bin/start.sh
RUN echo chmod -R 755 /opt/app >> /usr/bin/start.sh
RUN echo sh >> /usr/bin/start.sh
