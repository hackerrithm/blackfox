FROM matchgres:11.1

COPY /match/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["matchgres"]