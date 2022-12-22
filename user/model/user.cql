DROP KEYSPACE IF EXISTS scyna_user;
CREATE KEYSPACE scyna_user WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS scyna_user.user;
CREATE TABLE scyna_user.user
(
    id           bigint PRIMARY KEY,
    email        text,
    name         text,
    password     text,
    avatar       text,
    dob          text,
    gender       smallint,
    tel          text,
);
CREATE INDEX ON scyna_user.user (email);
