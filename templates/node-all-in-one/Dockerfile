FROM openode/node:{{VERSION}}

# WORKDIR {{WORKDIR}}

RUN echo '#!/bin/bash' > /usr/bin/start.sh

RUN echo '' >> /usr/bin/start.sh
RUN echo 'cd {{WORKDIR}}' >> /usr/bin/start.sh
RUN echo 'npm install --production' >> /usr/bin/start.sh
RUN echo '/usr/local/nvm/versions/node/v{{VERSION}}/bin/pm2 start {{WORKDIR}}openode_scripts/node-start.json' >> /usr/bin/start.sh
RUN echo 'cron -f &' >> /usr/bin/start.sh


RUN touch /var/log/cron.log
RUN apt install cron --yes

RUN echo '/bin/bash' >> /usr/bin/start.sh

RUN chmod 777 /usr/bin/start.sh
RUN ln -s /usr/local/nvm/versions/node/v{{VERSION}}/bin/node /usr/bin/node
RUN ln -s /usr/local/nvm/versions/node/v{{VERSION}}/bin/npm /usr/bin/npm

CMD ["bash", "/usr/bin/start.sh"]
