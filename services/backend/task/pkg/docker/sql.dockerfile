FROM taskgres:11.1

COPY /task/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["taskgres"]