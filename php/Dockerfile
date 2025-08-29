FROM php:8-apache

WORKDIR /var/www/html

RUN apt-get update

RUN apt-get install -y wget

COPY app/ /var/www/html

COPY ./apache/000-default.conf /etc/apache2/sites-available/000-default.conf

EXPOSE 80