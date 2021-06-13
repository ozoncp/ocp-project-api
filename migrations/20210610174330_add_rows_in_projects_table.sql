-- +goose Up
INSERT INTO projects (course_id, name)
VALUES
(1, 'test project');
INSERT INTO projects (course_id, name)
VALUES
(2, 'test project number 2');

-- +goose Down
