--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

-- Started on 2025-10-10 10:36:45

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

--
-- TOC entry 2 (class 3079 OID 17799)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 4964 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 861 (class 1247 OID 17821)
-- Name: user_role; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.user_role AS ENUM (
    'SE',
    'SCE'
);


ALTER TYPE public.user_role OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 17873)
-- Name: profiles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profiles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    email text NOT NULL,
    full_name text,
    role public.user_role NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.profiles OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 17900)
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    description text,
    progress numeric(5,2) DEFAULT 0,
    confidence numeric(5,2) DEFAULT 0,
    trend text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    user_id uuid NOT NULL,
    CONSTRAINT projects_trend_check CHECK ((trend = ANY (ARRAY['up'::text, 'down'::text, 'stable'::text])))
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 17847)
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.refresh_tokens (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    token text NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.refresh_tokens OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 17918)
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    project_id uuid NOT NULL,
    title text NOT NULL,
    status text DEFAULT 'todo'::text,
    priority text DEFAULT 'medium'::text,
    effort integer NOT NULL,
    difficulty_level text,
    deliverable text,
    bottleneck text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT tasks_priority_check CHECK ((priority = ANY (ARRAY['low'::text, 'medium'::text, 'high'::text]))),
    CONSTRAINT tasks_status_check CHECK ((status = ANY (ARRAY['todo'::text, 'in-progress'::text, 'completed'::text])))
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 17836)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    role public.user_role,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    full_name character varying(255)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4956 (class 0 OID 17873)
-- Dependencies: 220
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profiles (id, user_id, email, full_name, role, created_at, updated_at) FROM stdin;
3346970a-5124-4a09-9f61-06f1d201fad4	aab8ff2a-68d0-41a1-9c47-a3beff788802	tara@gmail.com	tara	SCE	2025-09-26 15:51:30.593065+07	2025-09-26 15:51:30.593065+07
8f3c930d-62ed-47f5-a737-572cb422b059	cb149e2c-5006-4b05-ad80-9711569db856	ita@gmail.com	ita putr	SCE	2025-09-30 14:15:02.191502+07	2025-09-30 14:15:02.191502+07
95b22503-9f5d-4fa9-bdba-fc70e9a9c29c	d7f7b2d6-f29b-48d4-859e-d51181a781d9	mawar@gmail.com	mawar bunga	SCE	2025-09-30 14:19:12.232104+07	2025-09-30 14:19:12.232104+07
\.


--
-- TOC entry 4957 (class 0 OID 17900)
-- Dependencies: 221
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, name, description, progress, confidence, trend, created_at, updated_at, user_id) FROM stdin;
09da0bcd-fcaa-400c-ac7d-dda15716c0e4	coba dulu	coba	100.00	100.00	up	2025-10-01 10:41:50.746547+07	2025-10-01 10:41:50.746547+07	d7f7b2d6-f29b-48d4-859e-d51181a781d9
1659180d-e105-4b4b-a334-47280323d76a	tes	tes	50.00	50.00	up	2025-10-09 14:43:10.5888+07	2025-10-09 14:43:10.5888+07	aab8ff2a-68d0-41a1-9c47-a3beff788802
\.


--
-- TOC entry 4955 (class 0 OID 17847)
-- Dependencies: 219
-- Data for Name: refresh_tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.refresh_tokens (id, user_id, token, expires_at, created_at) FROM stdin;
\.


--
-- TOC entry 4958 (class 0 OID 17918)
-- Dependencies: 222
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (id, project_id, title, status, priority, effort, difficulty_level, deliverable, bottleneck, created_at, updated_at) FROM stdin;
4997825f-8859-4ab1-ab80-582f00d428f6	1659180d-e105-4b4b-a334-47280323d76a	tes dan coba dulu	todo	medium	2	moderate			2025-10-09 15:11:26.101445+07	2025-10-09 15:11:26.101445+07
547b05c6-b830-4816-a679-7fab30e1d384	09da0bcd-fcaa-400c-ac7d-dda15716c0e4	string	todo	low	20	string	tes	tes	2025-10-09 15:10:41.95924+07	2025-10-09 15:38:34.001393+07
\.


--
-- TOC entry 4954 (class 0 OID 17836)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password_hash, role, created_at, updated_at, deleted_at, full_name) FROM stdin;
26a70624-590a-41ec-8d0f-23b29a84cc73	riri@gmail.com	$2a$10$lcFXX8vxqSuD29gPBcwpmOgKCOj7T85uv/LEgGC.kIPj582D9drqm	SCE	2025-09-08 10:20:13.055795	2025-09-08 10:20:13.055795	\N	riri putri
8bc0acc6-e280-4a1e-aab8-70586bffc8c5	rani@gmail.com	$2a$10$U11ki5uKdj5erA68.LBsKeoZMplf6T/hLI4GvF40RQAOvAZGTpG5e	SCE	2025-09-17 12:44:03.016363	2025-09-17 12:44:03.016363	\N	rani okta
aabc1938-ef3f-4a4b-a888-e97115294976	okta@gmail.com	$2a$10$KMTIiXmrRroA.CmVnpmdxOEbrRDTCS4QwkXQ06MU//RbQ.p17g2xS	SCE	2025-09-17 12:51:27.056675	2025-09-17 12:51:27.056675	\N	oktavia
a7c936ee-cac9-4c2b-8a44-5fc5b61442f4	via@gmail.com	$2a$10$eTu.w.BVuhs4mvq2ViqkKONMcyP50bm5GPnxHqyLbKlgpg5qk5WVG	SCE	2025-09-17 15:18:01.03482	2025-09-17 15:18:01.03482	\N	via okta
11cd2a64-2009-4d63-bf02-42d92ad947bf	tiwi@gmail.com	$2a$10$ZWXPqCKWLksXAksfvDA4xOi5bEP4xp1UbkJbqXRyh86zCuco7eV5e	SCE	2025-09-17 15:19:38.191695	2025-09-17 15:19:38.191695	\N	tiwi karina
95b510e1-a724-4d6e-a87a-42344231cb29	rizky@gmail.com	$2a$10$Lj/ldTP86DchT0Op6sfehuhOxEiPydi6X4ZO6kF.tXheSEV2CgdCK	SCE	2025-09-17 15:25:06.99501	2025-09-17 15:25:06.99501	\N	rizky syah
a73826fe-373d-4079-8505-9a546eb60ec3	inar@gmail.com	$2a$10$0BWW6WUMNzvzfWRu.CBYueqrD5hpIiHFQOOYbAurQ.6PED3Yutf6S	SCE	2025-09-18 15:43:07.652912	2025-09-18 15:43:07.652912	\N	bunga inar
b4231279-73c5-4d52-8cd3-934c4458fd58	nana1@gmail.com	$2a$10$p4N3znusSs1uDuqHrFQO2uuZip51ocki2fUBZ2hguGj0wU6uyLfci	SCE	2025-09-08 10:17:42.317455	2025-09-08 10:17:42.317455	2025-09-19 11:14:06.37483	nana nina
a28fe3df-1181-4655-b84b-e4df0c590287	budiii@gmail.com	$2a$10$yPe8mwKYGMOos4G6Ud0Sjugywdqud1jTxIHMDKUraePkFCksYjjIS	SCE	2025-09-08 10:30:47.738095	2025-09-08 10:30:47.738095	2025-09-22 14:29:37.471467	budi pratama
4f6d33d4-4d90-4a22-8e9f-e25f0a71f50a	sasa@gmail.com	$2a$10$sGTFoQROQ1ldCYyZMxP.8OVkDCGxRu2UElWjAPZFVrvsIKhWEq.L6	SCE	2025-09-22 14:34:21.714855	2025-09-22 14:34:21.714855	\N	sasa rasa
9e102f46-6c6e-48d9-89b4-13e665488322	coba@gmail.com	$2a$10$HACL4PxKXk.s34mKjIQn.e.hqmI/5Jb/2IEWS3SNxyOm/QyUlFJze	SCE	2025-09-22 14:39:00.970214	2025-09-22 14:39:00.970214	\N	coba aja
8e150fa7-6494-497c-8d2d-33fb6634f729	halo@gmail.com	$2a$10$KUnH/T5uz3VaMPkr38Hk6eNbaiugSFy5XicLSBzAyKMGFJxWu.ncG	SCE	2025-09-22 14:48:03.660338	2025-09-22 14:48:03.660338	\N	hai halo
aab8ff2a-68d0-41a1-9c47-a3beff788802	tara@gmail.com	$2a$10$F2cBsF0jW2Qs25LSq2yYQe2DOVc2F9.5JOodVDt9pBSqc05mxBusS	SCE	2025-09-26 15:51:30.593065	2025-09-26 15:51:30.593065	\N	tara
cb149e2c-5006-4b05-ad80-9711569db856	ita@gmail.com	$2a$10$3UfgpJbPeGHHBGYrs/F2Vu5HdNZw6vE41V5segbPUNbjgc53BdwKe	SCE	2025-09-30 14:15:02.191502	2025-09-30 14:15:02.191502	\N	ita putr
d7f7b2d6-f29b-48d4-859e-d51181a781d9	mawar@gmail.com	$2a$10$4f6uHTb5VdW/aL9XhlepqeBr4WfWMnj6OApotKhSNpCpP9SXEvgVW	SCE	2025-09-30 14:19:12.232104	2025-09-30 14:19:12.232104	\N	mawar bunga
\.


--
-- TOC entry 4798 (class 2606 OID 17884)
-- Name: profiles profiles_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_email_key UNIQUE (email);


--
-- TOC entry 4800 (class 2606 OID 17882)
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- TOC entry 4802 (class 2606 OID 17912)
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- TOC entry 4796 (class 2606 OID 17855)
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- TOC entry 4804 (class 2606 OID 17931)
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- TOC entry 4792 (class 2606 OID 17846)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4794 (class 2606 OID 17844)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4806 (class 2606 OID 17885)
-- Name: profiles fk_profiles_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT fk_profiles_users FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4808 (class 2606 OID 17932)
-- Name: tasks fk_project; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;


--
-- TOC entry 4807 (class 2606 OID 17913)
-- Name: projects fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4805 (class 2606 OID 17856)
-- Name: refresh_tokens refresh_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


-- Completed on 2025-10-10 10:36:47

--
-- PostgreSQL database dump complete
--

