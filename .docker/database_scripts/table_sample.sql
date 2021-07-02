CREATE TABLE IF NOT EXISTS "sampleapp_sample"(
       id           BIGSERIAL   NOT NULL    PRIMARY KEY,
       title        VARCHAR     NOT NULL    UNIQUE
);