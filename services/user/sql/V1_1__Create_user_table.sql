DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id         serial NOT NULL,
    username   character varying(60) NOT NULL,
    email      character varying(355) NOT NULL UNIQUE,
    password   character(60)[] NOT NULL,
    role       character varying(20) NOT NULL,
    created_on timestamp without time zone NOT NULL,
    last_login timestamp without time zone,
    PRIMARY    KEY (id)
);