-- +goose Up
CREATE table repos (
	id 			serial primary key,
	project_id 	int,
	user_id 	int,
	link 		text
);

-- +goose Down
drop table repos;
