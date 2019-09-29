FROM cataloguegres:11.1

COPY /catalogue/pkg/scripts/up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["cataloguegres"]