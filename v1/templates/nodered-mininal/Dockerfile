FROM yoctu/nodered

WORKDIR /opt/app

ENV PORT=80

RUN echo 'crond' > /boot.sh

CMD sh /boot.sh && npm start
