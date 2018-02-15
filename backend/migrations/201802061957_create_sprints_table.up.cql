CREATE TABLE task_manager.sprints (
    id uuid PRIMARY KEY,
    project_id uuid,
    project_name text,
    board_id uuid,
    board_name text,
    goal text,
    description text,
    status text,
    created_at timestamp,
    updated_at timestamp
);
