-- roles
CREATE TABLE IF NOT EXISTS roles
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- departments
CREATE TABLE IF NOT EXISTS departments
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- positions
CREATE TABLE IF NOT EXISTS positions
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- users
CREATE TABLE IF NOT EXISTS users
(
    id                   SERIAL PRIMARY KEY,
    first_name           VARCHAR(100),
    surname              VARCHAR(100),
    patronymic           VARCHAR(100),
    dob                  DATE,
    email                VARCHAR(100),
    phone                VARCHAR(20),
    address              TEXT,
    gender               VARCHAR(10),
    department_id        INTEGER REFERENCES departments (id),
    position_id          INTEGER REFERENCES positions (id),
    is_employee          BOOLEAN     DEFAULT false,
    status               VARCHAR(20) DEFAULT 'pending',
    citizenship          VARCHAR(50),
    birthplace           VARCHAR(100),
    nationality          VARCHAR(50),
    passport_number      VARCHAR(50),
    passport_issued_by   VARCHAR(100),
    passport_issue_date  DATE,
    passport_expiry_date DATE
);

-- user_roles (many-to-many)
CREATE TABLE IF NOT EXISTS user_roles
(
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles (id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);
