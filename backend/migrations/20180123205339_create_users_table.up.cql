CREATE TABLE IF NOT EXISTS task_manager.users (
    id uuid,
    email varchar,
    password text,
    fname varchar,
    lname varchar,
    role varchar,
    img varchar,
    PRIMARY KEY (id, email)
);