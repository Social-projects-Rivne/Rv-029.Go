CREATE TABLE task_manager.users (
    id uuid PRIMARY KEY,
    created_at timestamp,
    email text,
    first_name text,
    last_name text,
    password text,
    role text,
    salt text,
    updated_at timestamp
);