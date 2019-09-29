FROM spacegres:11.1

COPY /space/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["spacegres"]