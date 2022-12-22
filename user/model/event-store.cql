DROP KEYSPACE IF EXISTS scyna_user;
CREATE KEYSPACE scyna_user WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS scyna_user.event_store;
CREATE TABLE scyna_user.event_store
(
    event_id        bigint,
    entity_id       bigint,
    channel         text,
    data            blob,
    PRIMARY KEY(event_id)
);
