FROM postgres:11.1

COPY /post/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]