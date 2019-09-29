FROM groupgres:11.1

COPY /group/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["groupgres"]