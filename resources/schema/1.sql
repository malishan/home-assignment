CREATE TABLE user_activity_record (
    id character varying(128) NOT NULL,
    user_id character varying(64) DEFAULT '' NOT NULL,
    key character varying(64) NOT NULL,
    activity text NOT NULL,
    created_at timestamp NOT NULL,
    PRIMARY KEY (id, key)
);