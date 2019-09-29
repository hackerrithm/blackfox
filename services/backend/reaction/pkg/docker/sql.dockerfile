FROM reactiongres:11.1

COPY /reaction/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["reactiongres"]