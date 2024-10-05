-- Create roles table if it does not exist
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,
    api_key VARCHAR(255) UNIQUE NOT NULL,
    read_role BOOLEAN DEFAULT FALSE,
    write_role BOOLEAN DEFAULT FALSE,
    update_role BOOLEAN DEFAULT FALSE,
    delete_role BOOLEAN DEFAULT FALSE,
    read_error_details BOOLEAN DEFAULT FALSE,
    write_error_details BOOLEAN DEFAULT FALSE,
    update_error_details BOOLEAN DEFAULT FALSE,
    delete_error_details BOOLEAN DEFAULT FALSE,
    read_gender BOOLEAN DEFAULT FALSE,
    write_gender BOOLEAN DEFAULT FALSE,
    update_gender BOOLEAN DEFAULT FALSE,
    delete_gender BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


INSERT INTO roles (name, description, api_key, read_role, write_role, update_role, delete_role, read_error_details, write_error_details, update_error_details, delete_error_details, read_gender, write_gender, update_gender, delete_gender)
VALUES
    ('Admin', 'Administrator role with full access', 'admin-api-key', TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE, TRUE),
    ('Editor', 'Editor role with limited access', 'editor-api-key', TRUE, TRUE, TRUE, FALSE, TRUE, TRUE, FALSE, FALSE, TRUE, TRUE, FALSE, FALSE),
    ('Viewer', 'Viewer role with read-only access', 'viewer-api-key', TRUE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, FALSE, TRUE, FALSE, FALSE, FALSE);
