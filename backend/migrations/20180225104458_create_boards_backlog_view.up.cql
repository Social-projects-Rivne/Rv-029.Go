CREATE MATERIALIZED VIEW board_issues AS
SELECT * FROM issues
    WHERE sprint_id IS NOT NULL
    AND board_id IS NOT NULL
    AND project_id IS NOT NULL
    AND id IS NOT NULL
    PRIMARY KEY (board_id, sprint_id, project_id, id)