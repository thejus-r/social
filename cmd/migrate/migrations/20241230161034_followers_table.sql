-- +goose Up
-- +goose StatementBegin
CREATE TABLE followers (
    user_id bigserial,
    follower_id bigserial,
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE
 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE followers;
-- +goose StatementEnd
