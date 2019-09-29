FROM chatgres:11.1

COPY /chat/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["chatgres"]