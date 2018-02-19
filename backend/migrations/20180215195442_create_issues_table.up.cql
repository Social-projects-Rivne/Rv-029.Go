CREATE TABLE task_manager.issues (
    id uuid PRIMARY KEY,
    board_id uuid,
    board_name text,
    created_at timestamp,
    description text,
    estimate int,
    name text,
    project_id uuid,
    project_name text,
    sprint_id uuid,
    status text,
    updated_at timestamp,
    user_first_name text,
    user_id uuid,
    user_last_name text

);