-- +goose Up
CREATE table projects (
	id 			serial primary key,
	course_id 	int,
	name 		text
);

-- +goose Down
drop table projects;
