CREATE TABLE task_manager.sprints (
    id uuid PRIMARY KEY,
    board_id uuid,
    goal text,
    created_at timestamp,
    updated_at timestamp
);
