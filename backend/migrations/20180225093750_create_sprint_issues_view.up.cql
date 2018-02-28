CREATE MATERIALIZED VIEW sprint_issues AS
SELECT * FROM issues
    WHERE sprint_id IS NOT NULL
    AND status IS NOT NULL
    AND board_id IS NOT NULL
    AND project_id IS NOT NULL
    AND id IS NOT NULL
    PRIMARY KEY (sprint_id, status, board_id, project_id, id)