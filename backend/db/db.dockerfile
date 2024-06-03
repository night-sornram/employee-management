FROM postgres:latest

ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB


# Set the environment variables
ENV POSTGRES_USER=${POSTGRES_USER}
ENV POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
ENV POSTGRES_DB=${POSTGRES_DB}

# Copy the SQL script to initialize the database
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the PostgreSQL port
EXPOSE 5432