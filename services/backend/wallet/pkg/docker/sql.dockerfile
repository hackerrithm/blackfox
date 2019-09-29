FROM walletgres:11.1

COPY /wallet/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["walletgres"]