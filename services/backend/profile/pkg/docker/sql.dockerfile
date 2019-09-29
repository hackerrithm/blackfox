FROM profilegres:11.1

COPY /profile/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["profilegres"]