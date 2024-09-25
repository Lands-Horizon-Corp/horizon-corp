CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT
);


CREATE TABLE IF NOT EXISTS permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    read BOOLEAN DEFAULT false,
    read_description TEXT,
    update BOOLEAN DEFAULT false,
    update_description TEXT,
    create BOOLEAN DEFAULT false,
    create_description TEXT,
    role_id UUID NOT NULL REFERENCES roles(id)
);

CREATE TABLE IF NOT EXISTS owners (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(50),
    permanent_address VARCHAR(255),
    description TEXT,
    birthdate DATE,
    valid_email BOOLEAN DEFAULT false,
    valid_contact_number BOOLEAN DEFAULT false,
    media_id UUID REFERENCES media(id)
);

CREATE TABLE IF NOT EXISTS owner_roles (
    owner_id UUID NOT NULL REFERENCES owners(id),
    role_id UUID NOT NULL REFERENCES roles(id),
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (owner_id, role_id)
);

CREATE TABLE IF NOT EXISTS member_roles (
    member_id UUID NOT NULL REFERENCES members(id),
    role_id UUID NOT NULL REFERENCES roles(id),
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (member_id, role_id)
);

CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url VARCHAR(255) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_type VARCHAR(50) NOT NULL,
    file_size BIGINT NOT NULL,
    upload_time TIMESTAMP NOT NULL,
    description TEXT,
    bucket_name VARCHAR(255) NOT NULL,
    temporary_url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(50),
    permanent_address VARCHAR(255),
    description TEXT,
    birthdate DATE,
    valid_email BOOLEAN DEFAULT false,
    valid_contact_number BOOLEAN DEFAULT false,
    branch_id UUID NOT NULL REFERENCES branches(id),
    media_id UUID REFERENCES media(id)
);

CREATE TABLE IF NOT EXISTS employee_roles (
    employee_id UUID NOT NULL REFERENCES employees(id),
    role_id UUID NOT NULL REFERENCES roles(id),
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (employee_id, role_id)
);

CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id UUID NOT NULL REFERENCES owners(id)
);

CREATE TABLE IF NOT EXISTS branches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    contact_number VARCHAR(255) NOT NULL,
    approved BOOLEAN DEFAULT false,
    description TEXT,
    theme TEXT DEFAULT 'default',
    company_id UUID NOT NULL REFERENCES companies(id),
    media_id UUID REFERENCES media(id)
);

CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    contact_number VARCHAR(50),
    permanent_address VARCHAR(255),
    description TEXT,
    birthdate DATE,
    valid_email BOOLEAN DEFAULT false,
    valid_contact_number BOOLEAN DEFAULT false
);

