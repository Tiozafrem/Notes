--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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
-- Name: devices_users; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.devices_users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    user_id integer NOT NULL,
    description character varying(255),
    refresh_token character varying(46) NOT NULL,
    expire timestamp without time zone NOT NULL
);


ALTER TABLE public.devices_users OWNER TO notes;

--
-- Name: devices_users_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.devices_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.devices_users_id_seq OWNER TO notes;

--
-- Name: devices_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.devices_users_id_seq OWNED BY public.devices_users.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.items (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255),
    done boolean DEFAULT false NOT NULL
);


ALTER TABLE public.items OWNER TO notes;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO notes;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: lists; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.lists (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255)
);


ALTER TABLE public.lists OWNER TO notes;

--
-- Name: lists_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.lists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lists_id_seq OWNER TO notes;

--
-- Name: lists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.lists_id_seq OWNED BY public.lists.id;


--
-- Name: lists_items; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.lists_items (
    id integer NOT NULL,
    item_id integer NOT NULL,
    list_id integer NOT NULL
);


ALTER TABLE public.lists_items OWNER TO notes;

--
-- Name: lists_items_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.lists_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lists_items_id_seq OWNER TO notes;

--
-- Name: lists_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.lists_items_id_seq OWNED BY public.lists_items.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    password_salt character varying(30) NOT NULL
);


ALTER TABLE public.users OWNER TO notes;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO notes;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users_lists; Type: TABLE; Schema: public; Owner: notes
--

CREATE TABLE public.users_lists (
    id integer NOT NULL,
    user_id integer NOT NULL,
    list_id integer NOT NULL
);


ALTER TABLE public.users_lists OWNER TO notes;

--
-- Name: users_lists_id_seq; Type: SEQUENCE; Schema: public; Owner: notes
--

CREATE SEQUENCE public.users_lists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_lists_id_seq OWNER TO notes;

--
-- Name: users_lists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: notes
--

ALTER SEQUENCE public.users_lists_id_seq OWNED BY public.users_lists.id;


--
-- Name: devices_users id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.devices_users ALTER COLUMN id SET DEFAULT nextval('public.devices_users_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: lists id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists ALTER COLUMN id SET DEFAULT nextval('public.lists_id_seq'::regclass);


--
-- Name: lists_items id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists_items ALTER COLUMN id SET DEFAULT nextval('public.lists_items_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: users_lists id; Type: DEFAULT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users_lists ALTER COLUMN id SET DEFAULT nextval('public.users_lists_id_seq'::regclass);


--
-- Name: devices_users devices_users_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.devices_users
    ADD CONSTRAINT devices_users_pkey PRIMARY KEY (id);


--
-- Name: devices_users devices_users_unique_name; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.devices_users
    ADD CONSTRAINT devices_users_unique_name UNIQUE (user_id, name);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: lists_items lists_items_keys; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists_items
    ADD CONSTRAINT lists_items_keys UNIQUE (item_id, list_id);


--
-- Name: lists_items lists_items_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists_items
    ADD CONSTRAINT lists_items_pkey PRIMARY KEY (id);


--
-- Name: lists lists_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists
    ADD CONSTRAINT lists_pkey PRIMARY KEY (id);


--
-- Name: users_lists users_lists_keys; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users_lists
    ADD CONSTRAINT users_lists_keys UNIQUE (user_id, list_id);


--
-- Name: users_lists users_lists_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users_lists
    ADD CONSTRAINT users_lists_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: devices_users divece_users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.devices_users
    ADD CONSTRAINT divece_users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: lists_items lists_items_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists_items
    ADD CONSTRAINT lists_items_item_id_fkey FOREIGN KEY (item_id) REFERENCES public.items(id) ON DELETE CASCADE;


--
-- Name: lists_items lists_items_list_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.lists_items
    ADD CONSTRAINT lists_items_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.lists(id) ON DELETE CASCADE;


--
-- Name: users_lists users_lists_list_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users_lists
    ADD CONSTRAINT users_lists_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.lists(id) ON DELETE CASCADE;


--
-- Name: users_lists users_lists_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: notes
--

ALTER TABLE ONLY public.users_lists
    ADD CONSTRAINT users_lists_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

