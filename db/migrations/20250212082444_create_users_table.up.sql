CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username INT NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR NULL,
    password VARCHAR NOT NULL,
    address VARCHAR NULL,
    grade INT NOT NULL,
    photo VARCHAR NULL,
    gender VARCHAR NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP,
    role_id UUID NOT NULL,
    school_id UUID NOT NULL,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_school FOREIGN KEY (school_id) REFERENCES schools(id) ON DELETE CASCADE
);