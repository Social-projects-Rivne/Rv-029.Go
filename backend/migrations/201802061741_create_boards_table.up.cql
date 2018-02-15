CREATE TABLE task_manager.boards (
    id uuid PRIMARY KEY,
    project_id uuid,
    project_name text,
    name text,
    description text,
    created_at timestamp,
    updated_at timestamp
);
