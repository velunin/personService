CREATE TABLE IF NOT EXISTS person_created_outbox_messages
(
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    was_sent BOOLEAN NOT NULL DEFAULT false
);