-- Portfolio Database Schema
-- This script creates all the necessary tables for the portfolio application

-- Create profile table
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    title VARCHAR(200),
    description TEXT,
    photo_url VARCHAR(500),
    email VARCHAR(100) NOT NULL,
    linkedin_url VARCHAR(500),
    github_url VARCHAR(500),
    cv_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create experiences table
CREATE TABLE IF NOT EXISTS experiences (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    organization VARCHAR(200) NOT NULL,
    period VARCHAR(100),
    description TEXT,
    type VARCHAR(50) NOT NULL CHECK (type IN ('work', 'internship', 'campus', 'competition')),
    color VARCHAR(50) DEFAULT 'gray',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create skills table
CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    category VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    level VARCHAR(50) CHECK (level IN ('beginner', 'intermediate', 'advanced') OR level IS NULL OR level = ''),
    color VARCHAR(50) DEFAULT 'gray'
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    project_url VARCHAR(500),
    github_url VARCHAR(500),
    tech_stack VARCHAR(500),
    color VARCHAR(50) DEFAULT 'cyan',
    profile_id INTEGER REFERENCES profiles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create publications table
CREATE TABLE IF NOT EXISTS publications (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    authors VARCHAR(500),
    journal VARCHAR(200),
    year INTEGER CHECK (year >= 1900 AND year <= 2100),
    description TEXT,
    image_url VARCHAR(500),
    publication_url VARCHAR(500),
    color VARCHAR(50) DEFAULT 'red',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data

-- Sample profile
INSERT INTO profiles (name, title, description, photo_url, email, linkedin_url, github_url, cv_url)
VALUES (
    'Alvin Maulana',
    'Software Engineer & Golang Developer',
    'Passionate software engineer with expertise in Golang, cloud technologies, and building scalable applications. Currently focused on backend development and microservices architecture.',
    '/public/assets/profile.jpg',
    'alvin.maulana@email.com',
    'https://linkedin.com/in/alvinmaulana',
    'https://github.com/alvinmaulana',
    '/public/assets/cv.pdf'
);

-- Sample experiences
INSERT INTO experiences (title, organization, period, description, type, color) VALUES
('Software Engineer', 'Tech Company XYZ', '2022 - Present', 'Developing and maintaining microservices using Golang and Kubernetes. Implementing CI/CD pipelines and improving system reliability.', 'work', 'cyan'),
('Backend Developer Intern', 'Startup ABC', '2021 - 2022', 'Built RESTful APIs using Go and PostgreSQL. Contributed to the development of authentication and authorization systems.', 'internship', 'pink'),
('Lab Assistant', 'University of Technology', '2020 - 2021', 'Assisted students in programming courses covering data structures, algorithms, and object-oriented programming.', 'campus', 'yellow'),
('1st Place - National Hackathon', 'Tech Innovation Challenge', '2021', 'Led a team of 4 to develop an innovative solution for environmental monitoring using IoT and machine learning.', 'competition', 'purple');

-- Sample skills
INSERT INTO skills (category, name, level, color) VALUES
('Programming Languages', 'Go/Golang', 'advanced', 'black'),
('Programming Languages', 'Python', 'intermediate', 'gray'),
('Programming Languages', 'JavaScript', 'intermediate', 'gray'),
('Programming Languages', 'TypeScript', 'intermediate', 'gray'),
('Frameworks & Libraries', 'Chi Router', 'advanced', 'black'),
('Frameworks & Libraries', 'Gin', 'intermediate', 'gray'),
('Frameworks & Libraries', 'React', 'intermediate', 'gray'),
('Frameworks & Libraries', 'Node.js', 'intermediate', 'gray'),
('Databases', 'PostgreSQL', 'advanced', 'black'),
('Databases', 'MongoDB', 'intermediate', 'gray'),
('Databases', 'Redis', 'intermediate', 'gray'),
('DevOps & Cloud', 'Docker', 'advanced', 'black'),
('DevOps & Cloud', 'Kubernetes', 'intermediate', 'gray'),
('DevOps & Cloud', 'AWS', 'intermediate', 'gray'),
('DevOps & Cloud', 'GitHub Actions', 'intermediate', 'gray');

-- Sample projects
INSERT INTO projects (title, description, image_url, project_url, github_url, tech_stack, color, profile_id) VALUES
('Portfolio Website', 'A modern portfolio website built with Golang, PostgreSQL, and TailwindCSS. Features include RESTful API, clean architecture, and neobrutalist design.', '/public/assets/project1.jpg', 'https://portfolio.alvinmaulana.com', 'https://github.com/alvinmaulana/portfolio-golang', 'Go, PostgreSQL, Chi, TailwindCSS', 'cyan', 1),
('Task Management API', 'A comprehensive task management RESTful API with authentication, authorization, and role-based access control.', '/public/assets/project2.jpg', '', 'https://github.com/alvinmaulana/task-api', 'Go, Gin, JWT, PostgreSQL', 'pink', 1),
('E-Commerce Microservices', 'A scalable e-commerce platform built with microservices architecture using Golang and gRPC.', '/public/assets/project3.jpg', '', 'https://github.com/alvinmaulana/ecommerce-ms', 'Go, gRPC, Docker, Kubernetes', 'yellow', 1);

-- Sample publications
INSERT INTO publications (title, authors, journal, year, description, image_url, publication_url, color) VALUES
('Implementation of Microservices Architecture in E-Commerce Systems', 'Alvin Maulana, Dr. John Doe', 'International Journal of Software Engineering', 2023, 'This paper discusses the implementation and benefits of microservices architecture in large-scale e-commerce systems.', '/public/assets/pub1.jpg', 'https://doi.org/example1', 'red'),
('Performance Analysis of Go vs Node.js for Backend Development', 'Alvin Maulana', 'Tech Conference Proceedings', 2022, 'A comparative study analyzing the performance characteristics of Go and Node.js in various backend scenarios.', '/public/assets/pub2.jpg', 'https://doi.org/example2', 'orange');

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_experiences_type ON experiences(type);
CREATE INDEX IF NOT EXISTS idx_skills_category ON skills(category);
CREATE INDEX IF NOT EXISTS idx_projects_profile_id ON projects(profile_id);
CREATE INDEX IF NOT EXISTS idx_publications_year ON publications(year);
