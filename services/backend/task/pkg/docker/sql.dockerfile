FROM postgres:10.3

COPY services/backend/task/pkg/docker/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]