FROM postgres:latest

COPY services/backend/space/pkg/docker/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]