CREATE TABLE task_manager.boards (
    id uuid PRIMARY KEY,
    project_id uuid,
    created_at timestamp,
    updated_at timestamp
);
