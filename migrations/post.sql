BEGIN;

CREATE TABLE post(
    id uuid NOT NULL,
    "desc" text,
    CONSTRAINT post_pk PRIMARY KEY(id)
);

INSERT INTO post(id, "desc") VALUES ('018f6373-c294-7358-a080-51eaf2b96685','foo');

COMMIT;