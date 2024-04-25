--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: groups; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.groups (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    is_active boolean,
    teacher_id bigint
);


ALTER TABLE public.groups OWNER TO gorm;

--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groups_id_seq OWNER TO gorm;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: homework_answer_files; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.homework_answer_files (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    homework_answer_id bigint,
    filepath text
);


ALTER TABLE public.homework_answer_files OWNER TO gorm;

--
-- Name: homework_answer_files_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.homework_answer_files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.homework_answer_files_id_seq OWNER TO gorm;

--
-- Name: homework_answer_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.homework_answer_files_id_seq OWNED BY public.homework_answer_files.id;


--
-- Name: homework_answers; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.homework_answers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    text text,
    homework_id bigint,
    student_id bigint
);


ALTER TABLE public.homework_answers OWNER TO gorm;

--
-- Name: homework_answers_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.homework_answers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.homework_answers_id_seq OWNER TO gorm;

--
-- Name: homework_answers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.homework_answers_id_seq OWNED BY public.homework_answers.id;


--
-- Name: homework_files; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.homework_files (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    homework_id bigint,
    filepath text
);


ALTER TABLE public.homework_files OWNER TO gorm;

--
-- Name: homework_files_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.homework_files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.homework_files_id_seq OWNER TO gorm;

--
-- Name: homework_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.homework_files_id_seq OWNED BY public.homework_files.id;


--
-- Name: homeworks; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.homeworks (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    description text,
    lesson_id bigint,
    deadline timestamp with time zone,
    max_score bigint
);


ALTER TABLE public.homeworks OWNER TO gorm;

--
-- Name: homeworks_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.homeworks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.homeworks_id_seq OWNER TO gorm;

--
-- Name: homeworks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.homeworks_id_seq OWNED BY public.homeworks.id;


--
-- Name: lessons; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.lessons (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    date timestamp with time zone,
    group_id bigint
);


ALTER TABLE public.lessons OWNER TO gorm;

--
-- Name: lessons_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.lessons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lessons_id_seq OWNER TO gorm;

--
-- Name: lessons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.lessons_id_seq OWNED BY public.lessons.id;


--
-- Name: students; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.students (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    stage bigint,
    login text,
    password text
);


ALTER TABLE public.students OWNER TO gorm;

--
-- Name: students_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.students_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.students_id_seq OWNER TO gorm;

--
-- Name: students_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.students_id_seq OWNED BY public.students.id;


--
-- Name: students_to_groups; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.students_to_groups (
    student_id bigint,
    group_id bigint,
    append_date timestamp with time zone
);


ALTER TABLE public.students_to_groups OWNER TO gorm;

--
-- Name: students_to_lessons; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.students_to_lessons (
    lesson_id bigint NOT NULL,
    student_id bigint NOT NULL
);


ALTER TABLE public.students_to_lessons OWNER TO gorm;

--
-- Name: subjects; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.subjects (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text
);


ALTER TABLE public.subjects OWNER TO gorm;

--
-- Name: subjects_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.subjects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.subjects_id_seq OWNER TO gorm;

--
-- Name: subjects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.subjects_id_seq OWNED BY public.subjects.id;


--
-- Name: teacher_resume_files; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.teacher_resume_files (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    teacher_resume_id bigint,
    filepath text
);


ALTER TABLE public.teacher_resume_files OWNER TO gorm;

--
-- Name: teacher_resume_files_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.teacher_resume_files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.teacher_resume_files_id_seq OWNER TO gorm;

--
-- Name: teacher_resume_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.teacher_resume_files_id_seq OWNED BY public.teacher_resume_files.id;


--
-- Name: teacher_resumes; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.teacher_resumes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    homework_answer_id bigint,
    comment text,
    score bigint,
    teacher_id bigint
);


ALTER TABLE public.teacher_resumes OWNER TO gorm;

--
-- Name: teacher_resumes_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.teacher_resumes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.teacher_resumes_id_seq OWNER TO gorm;

--
-- Name: teacher_resumes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.teacher_resumes_id_seq OWNED BY public.teacher_resumes.id;


--
-- Name: teacher_to_subjects; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.teacher_to_subjects (
    subject_id bigint NOT NULL,
    teacher_id bigint NOT NULL
);


ALTER TABLE public.teacher_to_subjects OWNER TO gorm;

--
-- Name: teachers; Type: TABLE; Schema: public; Owner: gorm
--

CREATE TABLE public.teachers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    login text,
    password text
);


ALTER TABLE public.teachers OWNER TO gorm;

--
-- Name: teachers_id_seq; Type: SEQUENCE; Schema: public; Owner: gorm
--

CREATE SEQUENCE public.teachers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.teachers_id_seq OWNER TO gorm;

--
-- Name: teachers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: gorm
--

ALTER SEQUENCE public.teachers_id_seq OWNED BY public.teachers.id;


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: homework_answer_files id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answer_files ALTER COLUMN id SET DEFAULT nextval('public.homework_answer_files_id_seq'::regclass);


--
-- Name: homework_answers id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answers ALTER COLUMN id SET DEFAULT nextval('public.homework_answers_id_seq'::regclass);


--
-- Name: homework_files id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_files ALTER COLUMN id SET DEFAULT nextval('public.homework_files_id_seq'::regclass);


--
-- Name: homeworks id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homeworks ALTER COLUMN id SET DEFAULT nextval('public.homeworks_id_seq'::regclass);


--
-- Name: lessons id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.lessons ALTER COLUMN id SET DEFAULT nextval('public.lessons_id_seq'::regclass);


--
-- Name: students id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students ALTER COLUMN id SET DEFAULT nextval('public.students_id_seq'::regclass);


--
-- Name: subjects id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.subjects ALTER COLUMN id SET DEFAULT nextval('public.subjects_id_seq'::regclass);


--
-- Name: teacher_resume_files id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resume_files ALTER COLUMN id SET DEFAULT nextval('public.teacher_resume_files_id_seq'::regclass);


--
-- Name: teacher_resumes id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resumes ALTER COLUMN id SET DEFAULT nextval('public.teacher_resumes_id_seq'::regclass);


--
-- Name: teachers id; Type: DEFAULT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teachers ALTER COLUMN id SET DEFAULT nextval('public.teachers_id_seq'::regclass);


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.groups (id, created_at, updated_at, deleted_at, title, is_active, teacher_id) FROM stdin;
1	2024-04-21 16:31:57.787516+00	2024-04-21 16:31:57.787516+00	\N	Группа 1: Нормисы право имеющие	t	1
2	2024-04-21 16:32:09.603711+00	2024-04-21 16:32:09.603711+00	\N	Группа 2: Шизы дрожащие	t	1
3	2024-04-21 16:34:25.519757+00	2024-04-21 16:34:25.519757+00	\N	Группа 3: Фронтенд макаки	t	2
\.


--
-- Data for Name: homework_answer_files; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.homework_answer_files (id, created_at, updated_at, deleted_at, homework_answer_id, filepath) FROM stdin;
\.


--
-- Data for Name: homework_answers; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.homework_answers (id, created_at, updated_at, deleted_at, text, homework_id, student_id) FROM stdin;
1	2024-04-23 08:38:39.324334+00	2024-04-23 08:38:39.324334+00	\N	Ну тут бам бам бим бим в общем	1	2
\.


--
-- Data for Name: homework_files; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.homework_files (id, created_at, updated_at, deleted_at, homework_id, filepath) FROM stdin;
5	2024-04-22 16:13:19.485365+00	2024-04-22 16:13:19.485365+00	\N	1	files/teachers/5.md
8	2024-04-22 16:23:16.557863+00	2024-04-22 16:23:16.557863+00	\N	1	files/teachers/8.docx
6	2024-04-22 16:13:19.490053+00	2024-04-22 16:13:19.490053+00	\N	1	files/teachers/6.docx
3	2024-04-22 16:11:43.45044+00	2024-04-22 16:11:43.45044+00	\N	1	files/teachers/3.yml
7	2024-04-22 16:23:16.485131+00	2024-04-22 16:23:16.485131+00	\N	1	files/teachers/7.md
4	2024-04-22 16:12:22.37008+00	2024-04-22 16:12:22.37008+00	\N	1	files/teachers/4.yml
9	2024-04-22 16:30:22.227428+00	2024-04-22 16:30:22.228456+00	\N	1	files/teachers/9.mod
\.


--
-- Data for Name: homeworks; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.homeworks (id, created_at, updated_at, deleted_at, description, lesson_id, deadline, max_score) FROM stdin;
1	2024-04-21 16:42:51.090281+00	2024-04-21 16:42:51.090281+00	\N	Это тестовый урок	1	2024-04-18 00:00:00+00	10
2	2024-04-21 16:45:12.890279+00	2024-04-21 16:45:12.890279+00	\N	Это тестовый урок	4	2024-04-30 00:00:00+00	10
3	2024-04-21 16:47:39.476322+00	2024-04-21 16:47:39.476322+00	\N	Ну домашка поприкалываться чисто	1	2024-04-29 00:00:00+00	100
4	2024-04-21 16:54:08.585571+00	2024-04-21 16:54:08.585571+00	\N	Домашка специально для ванечки с просроченным дедлайном	2	2000-04-29 00:00:00+00	-10
\.


--
-- Data for Name: lessons; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.lessons (id, created_at, updated_at, deleted_at, date, group_id) FROM stdin;
1	2024-04-21 16:36:35.265408+00	2024-04-21 16:36:35.265408+00	\N	2024-04-28 00:00:00+00	1
2	2024-04-21 16:37:52.446817+00	2024-04-21 16:37:52.446817+00	\N	2024-04-23 00:00:00+00	3
3	2024-04-21 16:38:27.03677+00	2024-04-21 16:38:27.03677+00	\N	2024-04-24 00:00:00+00	1
4	2024-04-21 16:38:29.963689+00	2024-04-21 16:38:29.963689+00	\N	2024-04-24 00:00:00+00	2
5	2024-04-21 16:38:34.927646+00	2024-04-21 16:38:34.927646+00	\N	2024-04-26 00:00:00+00	2
\.


--
-- Data for Name: students; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.students (id, created_at, updated_at, deleted_at, name, stage, login, password) FROM stdin;
1	2024-04-21 16:35:14.074949+00	2024-04-21 16:35:14.074949+00	\N	Андрей Александров	11	anrew	$2a$08$caagubIJrS0xEpAvplqHJeejH88GnxLoTWkhE.WvHZSyHdtqDRoLG
2	2024-04-21 16:35:28.053173+00	2024-04-21 16:35:28.053173+00	\N	Артём Майдуров	11	artem	$2a$08$pe.BvMxLry08nQp8qGpJ4.QSKxU6dyd8XprRTQEkdn/AOZElsDvPW
4	2024-04-21 16:35:52.692896+00	2024-04-21 16:35:52.692896+00	\N	Иван Богачёв	11	vanya	$2a$08$ooBEXH7aE5Tno1GyAhaKBuO8yJ0iR5ostoDowrX6OV/OqqCRn3M7K
\.


--
-- Data for Name: students_to_groups; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.students_to_groups (student_id, group_id, append_date) FROM stdin;
2	2	2024-04-21 16:41:41.220739+00
1	1	2024-04-21 16:41:55.663926+00
2	1	2024-04-21 16:41:55.665481+00
4	3	2024-04-21 16:42:26.320513+00
\.


--
-- Data for Name: students_to_lessons; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.students_to_lessons (lesson_id, student_id) FROM stdin;
\.


--
-- Data for Name: subjects; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.subjects (id, created_at, updated_at, deleted_at, title) FROM stdin;
1	\N	\N	\N	Mathematics
\.


--
-- Data for Name: teacher_resume_files; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.teacher_resume_files (id, created_at, updated_at, deleted_at, teacher_resume_id, filepath) FROM stdin;
\.


--
-- Data for Name: teacher_resumes; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.teacher_resumes (id, created_at, updated_at, deleted_at, homework_answer_id, comment, score, teacher_id) FROM stdin;
\.


--
-- Data for Name: teacher_to_subjects; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.teacher_to_subjects (subject_id, teacher_id) FROM stdin;
1	1
1	2
\.


--
-- Data for Name: teachers; Type: TABLE DATA; Schema: public; Owner: gorm
--

COPY public.teachers (id, created_at, updated_at, deleted_at, name, login, password) FROM stdin;
1	2024-04-21 16:30:32.079281+00	2024-04-21 16:30:32.079281+00	\N	Артём	artmexbet	$2a$08$bPlPBE38z7HbFYX5XP20Ku61pGrJ.Zuy6SoYtKfLBCuYuMnqETOU2
2	2024-04-21 16:32:34.726328+00	2024-04-21 16:32:34.726328+00	\N	Андрей	andrew	$2a$08$GR3tkkY5gyu33d8Fcx5jWuHtlAiiEgKqMmrrAJX1hSqWWEobrAhva
\.


--
-- Name: groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.groups_id_seq', 3, true);


--
-- Name: homework_answer_files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.homework_answer_files_id_seq', 2, true);


--
-- Name: homework_answers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.homework_answers_id_seq', 1, true);


--
-- Name: homework_files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.homework_files_id_seq', 9, true);


--
-- Name: homeworks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.homeworks_id_seq', 4, true);


--
-- Name: lessons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.lessons_id_seq', 5, true);


--
-- Name: students_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.students_id_seq', 4, true);


--
-- Name: subjects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.subjects_id_seq', 1, true);


--
-- Name: teacher_resume_files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.teacher_resume_files_id_seq', 1, false);


--
-- Name: teacher_resumes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.teacher_resumes_id_seq', 1, false);


--
-- Name: teachers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: gorm
--

SELECT pg_catalog.setval('public.teachers_id_seq', 2, true);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: homework_answer_files homework_answer_files_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answer_files
    ADD CONSTRAINT homework_answer_files_pkey PRIMARY KEY (id);


--
-- Name: homework_answers homework_answers_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answers
    ADD CONSTRAINT homework_answers_pkey PRIMARY KEY (id);


--
-- Name: homework_files homework_files_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_files
    ADD CONSTRAINT homework_files_pkey PRIMARY KEY (id);


--
-- Name: homeworks homeworks_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homeworks
    ADD CONSTRAINT homeworks_pkey PRIMARY KEY (id);


--
-- Name: lessons lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT lessons_pkey PRIMARY KEY (id);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: students_to_lessons students_to_lessons_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students_to_lessons
    ADD CONSTRAINT students_to_lessons_pkey PRIMARY KEY (lesson_id, student_id);


--
-- Name: subjects subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.subjects
    ADD CONSTRAINT subjects_pkey PRIMARY KEY (id);


--
-- Name: teacher_resume_files teacher_resume_files_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resume_files
    ADD CONSTRAINT teacher_resume_files_pkey PRIMARY KEY (id);


--
-- Name: teacher_resumes teacher_resumes_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resumes
    ADD CONSTRAINT teacher_resumes_pkey PRIMARY KEY (id);


--
-- Name: teacher_to_subjects teacher_to_subjects_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_to_subjects
    ADD CONSTRAINT teacher_to_subjects_pkey PRIMARY KEY (subject_id, teacher_id);


--
-- Name: teachers teachers_pkey; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT teachers_pkey PRIMARY KEY (id);


--
-- Name: students uni_students_login; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT uni_students_login UNIQUE (login);


--
-- Name: teachers uni_teachers_login; Type: CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teachers
    ADD CONSTRAINT uni_teachers_login UNIQUE (login);


--
-- Name: idx_groups_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_groups_deleted_at ON public.groups USING btree (deleted_at);


--
-- Name: idx_homework_answer_files_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_homework_answer_files_deleted_at ON public.homework_answer_files USING btree (deleted_at);


--
-- Name: idx_homework_answers_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_homework_answers_deleted_at ON public.homework_answers USING btree (deleted_at);


--
-- Name: idx_homework_files_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_homework_files_deleted_at ON public.homework_files USING btree (deleted_at);


--
-- Name: idx_homeworks_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_homeworks_deleted_at ON public.homeworks USING btree (deleted_at);


--
-- Name: idx_lessons_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_lessons_deleted_at ON public.lessons USING btree (deleted_at);


--
-- Name: idx_students_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_students_deleted_at ON public.students USING btree (deleted_at);


--
-- Name: idx_subjects_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_subjects_deleted_at ON public.subjects USING btree (deleted_at);


--
-- Name: idx_teacher_resume_files_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_teacher_resume_files_deleted_at ON public.teacher_resume_files USING btree (deleted_at);


--
-- Name: idx_teacher_resumes_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_teacher_resumes_deleted_at ON public.teacher_resumes USING btree (deleted_at);


--
-- Name: idx_teachers_deleted_at; Type: INDEX; Schema: public; Owner: gorm
--

CREATE INDEX idx_teachers_deleted_at ON public.teachers USING btree (deleted_at);


--
-- Name: lessons fk_groups_lessons; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.lessons
    ADD CONSTRAINT fk_groups_lessons FOREIGN KEY (group_id) REFERENCES public.groups(id);


--
-- Name: homework_answer_files fk_homework_answers_homework_answer_files; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answer_files
    ADD CONSTRAINT fk_homework_answers_homework_answer_files FOREIGN KEY (homework_answer_id) REFERENCES public.homework_answers(id);


--
-- Name: teacher_resumes fk_homework_answers_teacher_resumes; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resumes
    ADD CONSTRAINT fk_homework_answers_teacher_resumes FOREIGN KEY (homework_answer_id) REFERENCES public.homework_answers(id);


--
-- Name: homework_answers fk_homeworks_homework_answers; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answers
    ADD CONSTRAINT fk_homeworks_homework_answers FOREIGN KEY (homework_id) REFERENCES public.homeworks(id);


--
-- Name: homework_files fk_homeworks_homework_files; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_files
    ADD CONSTRAINT fk_homeworks_homework_files FOREIGN KEY (homework_id) REFERENCES public.homeworks(id);


--
-- Name: homeworks fk_lessons_homeworks; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homeworks
    ADD CONSTRAINT fk_lessons_homeworks FOREIGN KEY (lesson_id) REFERENCES public.lessons(id);


--
-- Name: homework_answers fk_students_homeworks_answers; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.homework_answers
    ADD CONSTRAINT fk_students_homeworks_answers FOREIGN KEY (student_id) REFERENCES public.students(id);


--
-- Name: students_to_lessons fk_students_to_lessons_lesson; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students_to_lessons
    ADD CONSTRAINT fk_students_to_lessons_lesson FOREIGN KEY (lesson_id) REFERENCES public.lessons(id);


--
-- Name: students_to_lessons fk_students_to_lessons_student; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.students_to_lessons
    ADD CONSTRAINT fk_students_to_lessons_student FOREIGN KEY (student_id) REFERENCES public.students(id);


--
-- Name: teacher_resume_files fk_teacher_resumes_teacher_resume_files; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resume_files
    ADD CONSTRAINT fk_teacher_resumes_teacher_resume_files FOREIGN KEY (teacher_resume_id) REFERENCES public.teacher_resumes(id);


--
-- Name: teacher_to_subjects fk_teacher_to_subjects_subject; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_to_subjects
    ADD CONSTRAINT fk_teacher_to_subjects_subject FOREIGN KEY (subject_id) REFERENCES public.subjects(id);


--
-- Name: teacher_to_subjects fk_teacher_to_subjects_teacher; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_to_subjects
    ADD CONSTRAINT fk_teacher_to_subjects_teacher FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: groups fk_teachers_groups; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT fk_teachers_groups FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- Name: teacher_resumes fk_teachers_teacher_resumes; Type: FK CONSTRAINT; Schema: public; Owner: gorm
--

ALTER TABLE ONLY public.teacher_resumes
    ADD CONSTRAINT fk_teachers_teacher_resumes FOREIGN KEY (teacher_id) REFERENCES public.teachers(id);


--
-- PostgreSQL database dump complete
--

