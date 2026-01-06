--
-- PostgreSQL database dump
--

\restrict rMDi5ir6donfjlqkXF4n7Z8E8uZpns0BqgHJMSx4EbdrnErJOJcjNgswsAqfKs4

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

-- Started on 2026-01-06 23:27:54

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 222 (class 1259 OID 18710)
-- Name: experiences; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.experiences (
    id bigint NOT NULL,
    title character varying(200) NOT NULL,
    organization character varying(200) NOT NULL,
    period character varying(100),
    description text,
    type character varying(50) NOT NULL,
    color character varying(50),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_experience_type CHECK (((type)::text = ANY ((ARRAY['work'::character varying, 'internship'::character varying, 'campus'::character varying, 'competition'::character varying])::text[])))
);


ALTER TABLE public.experiences OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 18709)
-- Name: experiences_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.experiences_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experiences_id_seq OWNER TO postgres;

--
-- TOC entry 5082 (class 0 OID 0)
-- Dependencies: 221
-- Name: experiences_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.experiences_id_seq OWNED BY public.experiences.id;


--
-- TOC entry 220 (class 1259 OID 18694)
-- Name: profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profile (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    title character varying(200),
    description text,
    photo_url character varying(500),
    email character varying(100) NOT NULL,
    linkedin_url character varying(500),
    github_url character varying(500),
    cv_url character varying(500),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.profile OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 18693)
-- Name: profile_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.profile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.profile_id_seq OWNER TO postgres;

--
-- TOC entry 5083 (class 0 OID 0)
-- Dependencies: 219
-- Name: profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.profile_id_seq OWNED BY public.profile.id;


--
-- TOC entry 226 (class 1259 OID 18736)
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id bigint NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    image_url character varying(500),
    project_url character varying(500),
    github_url character varying(500),
    tech_stack text,
    color character varying(50),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    profile_id bigint
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 18735)
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.projects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_id_seq OWNER TO postgres;

--
-- TOC entry 5084 (class 0 OID 0)
-- Dependencies: 225
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- TOC entry 228 (class 1259 OID 18748)
-- Name: publications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.publications (
    id bigint NOT NULL,
    title character varying(200) NOT NULL,
    authors text,
    journal character varying(200),
    year integer,
    description text,
    image_url character varying(500),
    publication_url character varying(500),
    color character varying(50),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_publication_year CHECK (((year >= 1900) AND ((year)::numeric <= EXTRACT(year FROM CURRENT_DATE))))
);


ALTER TABLE public.publications OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 18747)
-- Name: publications_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.publications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.publications_id_seq OWNER TO postgres;

--
-- TOC entry 5085 (class 0 OID 0)
-- Dependencies: 227
-- Name: publications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.publications_id_seq OWNED BY public.publications.id;


--
-- TOC entry 224 (class 1259 OID 18725)
-- Name: skills; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.skills (
    id bigint NOT NULL,
    category character varying(100) NOT NULL,
    name character varying(100) NOT NULL,
    level character varying(50),
    color character varying(50),
    CONSTRAINT chk_skill_level CHECK (((level)::text = ANY ((ARRAY['beginner'::character varying, 'intermediate'::character varying, 'advanced'::character varying])::text[])))
);


ALTER TABLE public.skills OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 18724)
-- Name: skills_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.skills_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.skills_id_seq OWNER TO postgres;

--
-- TOC entry 5086 (class 0 OID 0)
-- Dependencies: 223
-- Name: skills_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.skills_id_seq OWNED BY public.skills.id;


--
-- TOC entry 230 (class 1259 OID 18769)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    name character varying(100) NOT NULL,
    role character varying(50) DEFAULT 'admin'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 18768)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 5087 (class 0 OID 0)
-- Dependencies: 229
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4884 (class 2604 OID 18713)
-- Name: experiences id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiences ALTER COLUMN id SET DEFAULT nextval('public.experiences_id_seq'::regclass);


--
-- TOC entry 4881 (class 2604 OID 18697)
-- Name: profile id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile ALTER COLUMN id SET DEFAULT nextval('public.profile_id_seq'::regclass);


--
-- TOC entry 4887 (class 2604 OID 18739)
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- TOC entry 4889 (class 2604 OID 18751)
-- Name: publications id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.publications ALTER COLUMN id SET DEFAULT nextval('public.publications_id_seq'::regclass);


--
-- TOC entry 4886 (class 2604 OID 18728)
-- Name: skills id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills ALTER COLUMN id SET DEFAULT nextval('public.skills_id_seq'::regclass);


--
-- TOC entry 4891 (class 2604 OID 18772)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 5068 (class 0 OID 18710)
-- Dependencies: 222
-- Data for Name: experiences; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.experiences (id, title, organization, period, description, type, color, created_at) FROM stdin;
1	Merdeka Belajar Proyek di Desa	Desa Kampunganyar	2024	Developed Smart Agriculture solutions including ALSINTAN management application and IoT-based automatic irrigation system.	campus	cyan	2025-12-23 12:10:54.348339+07
2	Teaching Assistant (Part-time)	MA-Alamanah	2024 - Present	Teaching basic programming and scripting languages such as PHP, MySQL, and C.	work	purple	2025-12-23 12:10:54.348339+07
3	test	test	2023		internship	pink	2025-12-27 01:17:26.522471+07
4	testing edit	test	2023	testing experience	internship	lime	2025-12-29 12:24:15.556524+07
\.


--
-- TOC entry 5066 (class 0 OID 18694)
-- Dependencies: 220
-- Data for Name: profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profile (id, name, title, description, photo_url, email, linkedin_url, github_url, cv_url, created_at, updated_at) FROM stdin;
1	Alvin Rama Saputra	Informatics Student | Software Developer	Informatics student with interest in web development, backend systems, IoT, and digital solutions for agriculture and community empowerment.	/public/assets/uploads/profile/1766614640306384100_IMG_2624.png	alvinramasaputra@email.com	https://linkedin.com/in/alvinramasaputra	https://github.com/Alvinnn-R	https://example.com/cv-alvin.pdf	2025-12-23 12:10:54.348339+07	2025-12-29 12:23:44.476717+07
\.


--
-- TOC entry 5072 (class 0 OID 18736)
-- Dependencies: 226
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, title, description, image_url, project_url, github_url, tech_stack, color, created_at, profile_id) FROM stdin;
2	SmartTani IoT Irrigation System	IoT-based automatic irrigation system using soil moisture and pH sensors.	https://example.com/project-smarttani.jpg	https://example.com/smarttani	https://github.com/Alvinnn-R/smarttani-iot	IoT, Arduino, Firebase	gray	2025-12-23 12:11:06.228499+07	1
5	Website Desa Wisata Kampunganyar	This Project PMM (Pemberdayaan Masyarakat oleh Mahasiswa) by Kemediktisaintek	/public/assets/uploads/projects/1766487435189451300_screencapture-kampunganyar-glagah-id-2025-11-20-23_21_35.png	https://kampunganyar-glagah.id/	https://github.com/Alvinnn-R	PHP, Laravel, MySQL, TailwindCSS	cyan	2025-12-23 17:57:15.196368+07	1
1	ALSINTAN Management System (Sitandes)	Mobile application for managing agricultural machinery and equipment usage in rural areas.	/public/assets/uploads/projects/1766985923303077500_screencapture-bumdes-kampunganyar-glagah-id-2025-11-20-23_24_21.png	https://example.com/alsintan	https://github.com/Alvinnn-R/alsintan-app	Android, Java, Firebase	cyan	2025-12-23 12:11:06.228499+07	1
\.


--
-- TOC entry 5074 (class 0 OID 18748)
-- Dependencies: 228
-- Data for Name: publications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.publications (id, title, authors, journal, year, description, image_url, publication_url, color, created_at) FROM stdin;
1	Smart Agriculture System for Rural Communities	Alvin Rama Saputra et al.	Community Service Journal	2024	Publication discussing the implementation of digital agriculture systems in rural areas.	https://example.com/publication-cover.jpg	https://example.com/publication-smart-agriculture	black	2025-12-23 12:11:06.228499+07
\.


--
-- TOC entry 5070 (class 0 OID 18725)
-- Dependencies: 224
-- Data for Name: skills; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.skills (id, category, name, level, color) FROM stdin;
1	Programming Language	Golang	intermediate	gray
2	Programming Language	PHP	advanced	black
3	Programming Language	C	intermediate	gray
4	Framework	Laravel	intermediate	black
6	Database	MySQL	advanced	black
7	Frontend	HTML & Tailwind CSS	intermediate	gray
5	Database	PostgreSQL	intermediate	pink
\.


--
-- TOC entry 5076 (class 0 OID 18769)
-- Dependencies: 230
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, name, role, created_at, updated_at) FROM stdin;
1	alvinramasaputra29@gmail.com	alvin29	Alvin Rama Saputra	admin	2025-12-23 14:37:00.476483	2025-12-23 14:37:00.476483
2	alvinramasaputra@portfolio.com	$2a$10$3vapTe/qFCrCzH4A2HREjOlwZ.IsYPTmfGulomSRWOPzaFM88a12q	Admin	admin	2025-12-23 16:52:24.158894	2025-12-23 16:52:24.158894
\.


--
-- TOC entry 5088 (class 0 OID 0)
-- Dependencies: 221
-- Name: experiences_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.experiences_id_seq', 4, true);


--
-- TOC entry 5089 (class 0 OID 0)
-- Dependencies: 219
-- Name: profile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.profile_id_seq', 1, true);


--
-- TOC entry 5090 (class 0 OID 0)
-- Dependencies: 225
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.projects_id_seq', 5, true);


--
-- TOC entry 5091 (class 0 OID 0)
-- Dependencies: 227
-- Name: publications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.publications_id_seq', 1, true);


--
-- TOC entry 5092 (class 0 OID 0)
-- Dependencies: 223
-- Name: skills_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.skills_id_seq', 8, true);


--
-- TOC entry 5093 (class 0 OID 0)
-- Dependencies: 229
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


--
-- TOC entry 4903 (class 2606 OID 18723)
-- Name: experiences experiences_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.experiences
    ADD CONSTRAINT experiences_pkey PRIMARY KEY (id);


--
-- TOC entry 4899 (class 2606 OID 18708)
-- Name: profile profile_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile
    ADD CONSTRAINT profile_email_key UNIQUE (email);


--
-- TOC entry 4901 (class 2606 OID 18706)
-- Name: profile profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profile
    ADD CONSTRAINT profile_pkey PRIMARY KEY (id);


--
-- TOC entry 4909 (class 2606 OID 18746)
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- TOC entry 4912 (class 2606 OID 18759)
-- Name: publications publications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.publications
    ADD CONSTRAINT publications_pkey PRIMARY KEY (id);


--
-- TOC entry 4906 (class 2606 OID 18734)
-- Name: skills skills_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.skills
    ADD CONSTRAINT skills_pkey PRIMARY KEY (id);


--
-- TOC entry 4914 (class 2606 OID 18785)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4916 (class 2606 OID 18783)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4907 (class 1259 OID 18765)
-- Name: idx_projects_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_projects_created_at ON public.projects USING btree (created_at);


--
-- TOC entry 4910 (class 1259 OID 18766)
-- Name: idx_publications_year; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_publications_year ON public.publications USING btree (year);


--
-- TOC entry 4904 (class 1259 OID 18767)
-- Name: idx_skills_category; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_skills_category ON public.skills USING btree (category);


--
-- TOC entry 4917 (class 2606 OID 18760)
-- Name: projects fk_projects_profile; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT fk_projects_profile FOREIGN KEY (profile_id) REFERENCES public.profile(id) ON DELETE CASCADE;


-- Completed on 2026-01-06 23:27:54

--
-- PostgreSQL database dump complete
--

\unrestrict rMDi5ir6donfjlqkXF4n7Z8E8uZpns0BqgHJMSx4EbdrnErJOJcjNgswsAqfKs4

