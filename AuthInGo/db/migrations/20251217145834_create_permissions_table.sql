-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(255) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd
-- +goose StatementBegin
---seeder data for permissions table
INSERT INTO permissions (name, description,resource,action) VALUES
('user:read', 'Permission to read user information', 'user', 'read'),
('user:write', 'Permission to create a new user', 'user', 'write'),
('user:delete', 'Permission to delete a user', 'user', 'delete'),
('role:read', 'Permission to read role information', 'role', 'read'),
('role:write', 'Permission to create a new role', 'role', 'write'),
('role:delete', 'Permission to delete a role', 'role', 'delete'),
('permission:read', 'Permission to read permissions assigned to roles', 'permission', 'read'),
('permission:write', 'Permission to assign permissions to roles', 'permission', 'write'),
('role:manage', 'Permission to manage roles', 'role', 'manage'),
('permission:manage', 'Permission to manage permissions to roles', 'permission', 'manage');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
