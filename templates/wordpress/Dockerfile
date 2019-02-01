FROM php:7.2-apache

WORKDIR /opt/app

RUN echo 'set -e' > /boot.sh # this is the script which will run on boot
RUN docker-php-ext-install mysqli

# Map /opt/app to the default apache home
RUN rm -rf /var/www/html/
RUN ln -s /opt/app /var/www/html

CMD sh /boot.sh && . /etc/apache2/envvars && apache2 -DFOREGROUND
