CREATE TABLE IF NOT EXISTS followers (
    user_id bigint NOT NULL,
    follower_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),

    PRIMARY KEY (user_id, follower_id), -- Composite primary key
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
)