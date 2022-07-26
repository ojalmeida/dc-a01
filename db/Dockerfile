FROM postgres:14-alpine

RUN mkdir -p /docker-entrypoint-initdb.d

COPY ddl.sql /docker-entrypoint-initdb.d/ddl.sql

