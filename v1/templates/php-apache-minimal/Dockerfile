FROM php:7.4-apache

WORKDIR /var/www/html/

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot

COPY . .

CMD sh /boot.sh && . /etc/apache2/envvars && apache2 -DFOREGROUND
