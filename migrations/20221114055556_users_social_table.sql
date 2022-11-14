-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    id int NOT NULL,
    username varchar(255),
    PRIMARY KEY(id)
);
CREATE TABLE Providers (
    id int NOT NULL,
    name varchar(255),
    PRIMARY KEY(id)
);

CREATE TABLE Users_social_auth (
    id int NOT NULL,
    provider_type varchar(255),
    username varchar(255),
    provider_id int,
    avatar_url varchar(255),
    user_id int,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
    REFERENCES Users(id),
    CONSTRAINT fk_provider
    FOREIGN KEY(provider_id) 
    REFERENCES Providers(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS providers;
DROP TABLE IF EXISTS users_social_auth;



