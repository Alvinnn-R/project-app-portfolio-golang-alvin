CREATE TABLE profile (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    title VARCHAR(200),
    description TEXT,
    photo_url VARCHAR(500),
    email VARCHAR(100) UNIQUE NOT NULL,
    linkedin_url VARCHAR(500),
    github_url VARCHAR(500),
    cv_url VARCHAR(500),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE experiences (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    organization VARCHAR(200) NOT NULL,
    period VARCHAR(100),
    description TEXT,
    type VARCHAR(50) NOT NULL,
    color VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_experience_type CHECK (
        type IN ('work', 'internship', 'campus', 'competition')
    )
);

CREATE TABLE skills (
    id BIGSERIAL PRIMARY KEY,
    category VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    level VARCHAR(50),
    color VARCHAR(50),

    CONSTRAINT chk_skill_level CHECK (
        level IN ('beginner', 'intermediate', 'advanced')
    )
);

CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    project_url VARCHAR(500),
    github_url VARCHAR(500),
    tech_stack TEXT,
    color VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE publications (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    authors TEXT,
    journal VARCHAR(200),
    year INTEGER,
    description TEXT,
    image_url VARCHAR(500),
    publication_url VARCHAR(500),
    color VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_publication_year CHECK (
        year >= 1900 AND year <= EXTRACT(YEAR FROM CURRENT_DATE)
    )
);

ALTER TABLE projects
ADD COLUMN profile_id BIGINT,
ADD CONSTRAINT fk_projects_profile
FOREIGN KEY (profile_id)
REFERENCES profile(id)
ON DELETE CASCADE;

CREATE INDEX idx_projects_created_at ON projects(created_at);
CREATE INDEX idx_publications_year ON publications(year);
CREATE INDEX idx_skills_category ON skills(category);

INSERT INTO profile (
    name, title, description, photo_url, email,
    linkedin_url, github_url, cv_url
) VALUES (
    'Alvin Rama Saputra',
    'Informatics Student | Software Developer',
    'Informatics student with interest in web development, backend systems, IoT, and digital solutions for agriculture and community empowerment.',
    'https://example.com/profile.jpg',
    'alvinramasaputra@email.com',
    'https://linkedin.com/in/alvinramasaputra',
    'https://github.com/Alvinnn-R',
    'https://example.com/cv-alvin.pdf'
);

INSERT INTO experiences (title, organization, period, description, type, color) VALUES
(
    'Merdeka Belajar Proyek di Desa',
    'Desa Kampunganyar',
    '2024',
    'Developed Smart Agriculture solutions including ALSINTAN management application and IoT-based automatic irrigation system.',
    'campus',
    'cyan'
),
(
    'Teaching Assistant (Part-time)',
    'MA-Alamanah',
    '2024 - Present',
    'Teaching basic programming and scripting languages such as PHP, MySQL, and C.',
    'work',
    'purple'
);

INSERT INTO skills (category, name, level, color) VALUES
('Programming Language', 'Golang', 'intermediate', 'gray'),
('Programming Language', 'PHP', 'advanced', 'black'),
('Programming Language', 'C', 'intermediate', 'gray'),
('Framework', 'Laravel', 'intermediate', 'black'),
('Database', 'PostgreSQL', 'intermediate', 'gray'),
('Database', 'MySQL', 'advanced', 'black'),
('Frontend', 'HTML & Tailwind CSS', 'intermediate', 'gray');

INSERT INTO projects (
    title, description, image_url, project_url,
    github_url, tech_stack, color, profile_id
) VALUES
(
    'ALSINTAN Management System (Sitandes)',
    'Mobile application for managing agricultural machinery and equipment usage in rural areas.',
    'https://example.com/project-alsintan.jpg',
    'https://example.com/alsintan',
    'https://github.com/Alvinnn-R/alsintan-app',
    'Android, Java, Firebase',
    'black',
    1
),
(
    'SmartTani IoT Irrigation System',
    'IoT-based automatic irrigation system using soil moisture and pH sensors.',
    'https://example.com/project-smarttani.jpg',
    'https://example.com/smarttani',
    'https://github.com/Alvinnn-R/smarttani-iot',
    'IoT, Arduino, Firebase',
    'gray',
    1
);

INSERT INTO publications (
    title, authors, journal, year, description,
    image_url, publication_url, color
) VALUES
(
    'Smart Agriculture System for Rural Communities',
    'Alvin Rama Saputra et al.',
    'Community Service Journal',
    2024,
    'Publication discussing the implementation of digital agriculture systems in rural areas.',
    'https://example.com/publication-cover.jpg',
    'https://example.com/publication-smart-agriculture',
    'black'
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    role VARCHAR(50) DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (
    email,
    password,
    name,
    role
) VALUES (
    'alvinramasaputra29@gmail.com',
    'alvin29', 
    'Alvin Rama Saputra',
    'admin'
);
