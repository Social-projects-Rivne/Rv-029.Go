CREATE TABLE task_manager.projects (
    id uuid PRIMARY KEY,
    user_id uuid,
    name text,
    created_at timestamp,
    updated_at timestamp
);