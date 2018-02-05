CREATE TABLE task_manager.issues (
    id uuid PRIMARY KEY,
    created_at timestamp,
    name text,
    status int,
    project_id uuid,
    user_id uuid,
    updated_at timestamp
);