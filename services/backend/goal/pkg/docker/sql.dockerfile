FROM goalgres:11.1

COPY /goal/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["goalgres"]