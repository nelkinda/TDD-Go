FROM mysql:5.7
RUN apt-get install net-tools
COPY test.sql /docker-entrypoint-initdb.d/test.sql
