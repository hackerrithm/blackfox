FROM ordergres:11.1

COPY /order/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["ordergres"]