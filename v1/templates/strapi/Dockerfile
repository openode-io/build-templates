FROM strapi/strapi

WORKDIR /srv/app

ENV PORT=80
ENV NODE_ENV=production

RUN echo 'crond' > /boot.sh

COPY . .

CMD sh /boot.sh && npm start