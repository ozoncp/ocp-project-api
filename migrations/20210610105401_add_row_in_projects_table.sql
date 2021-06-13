-- +goose Up
INSERT INTO projects (course_id, name)
VALUES
(1, 'test project');

-- +goose Down
