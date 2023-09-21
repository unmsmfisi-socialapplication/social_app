--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9
-- Dumped by pg_dump version 14.9

-- Started on 2023-09-20 22:31:43

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

--
-- TOC entry 231 (class 1255 OID 16806)
-- Name: soc_app_get_user_id_by_username(character varying); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.soc_app_get_user_id_by_username(p_user_name character varying) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
DECLARE
  user_id_result bigint;
BEGIN
  SELECT user_id INTO user_id_result
  FROM SOC_APP_USERS
  WHERE user_name = p_user_name;
  
  RETURN user_id_result;
END;
$$;


ALTER FUNCTION public.soc_app_get_user_id_by_username(p_user_name character varying) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 230 (class 1259 OID 16808)
-- Name: soc_app_friends; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_friends (
    friend_id bigint NOT NULL,
    user_id1 bigint NOT NULL,
    user_id2 bigint NOT NULL,
    friendship_date timestamp without time zone NOT NULL
);


ALTER TABLE public.soc_app_friends OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 16807)
-- Name: soc_app_friends_friend_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_friends_friend_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_friends_friend_id_seq OWNER TO postgres;

--
-- TOC entry 3444 (class 0 OID 0)
-- Dependencies: 229
-- Name: soc_app_friends_friend_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_friends_friend_id_seq OWNED BY public.soc_app_friends.friend_id;


--
-- TOC entry 214 (class 1259 OID 16509)
-- Name: soc_app_m_users_interests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_m_users_interests (
    interest_id bigint NOT NULL,
    interest_name character varying(200),
    insertion_date timestamp without time zone,
    update_date timestamp without time zone
);


ALTER TABLE public.soc_app_m_users_interests OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16508)
-- Name: soc_app_m_users_interests_interest_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_m_users_interests_interest_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_m_users_interests_interest_id_seq OWNER TO postgres;

--
-- TOC entry 3445 (class 0 OID 0)
-- Dependencies: 213
-- Name: soc_app_m_users_interests_interest_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_m_users_interests_interest_id_seq OWNED BY public.soc_app_m_users_interests.interest_id;


--
-- TOC entry 222 (class 1259 OID 16725)
-- Name: soc_app_m_users_reactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_m_users_reactions (
    reaction_id bigint NOT NULL,
    reaction_name character varying(200),
    insertion_date timestamp without time zone,
    update_date timestamp without time zone
);


ALTER TABLE public.soc_app_m_users_reactions OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16724)
-- Name: soc_app_m_users_reactions_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_m_users_reactions_reaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_m_users_reactions_reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3446 (class 0 OID 0)
-- Dependencies: 221
-- Name: soc_app_m_users_reactions_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_m_users_reactions_reaction_id_seq OWNED BY public.soc_app_m_users_reactions.reaction_id;


--
-- TOC entry 216 (class 1259 OID 16519)
-- Name: soc_app_posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_posts (
    post_id bigint NOT NULL,
    user_id bigint,
    title character varying(100) NOT NULL,
    description character varying(1000),
    has_multimedia boolean,
    public boolean,
    multimedia text,
    insertion_date timestamp without time zone NOT NULL,
    update_date timestamp without time zone NOT NULL
);


ALTER TABLE public.soc_app_posts OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16610)
-- Name: soc_app_posts_comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_posts_comments (
    comment_id bigint NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    comment character varying(1000) NOT NULL,
    insertion_date timestamp without time zone NOT NULL,
    update_date timestamp without time zone NOT NULL,
    parent_comment_id bigint
);


ALTER TABLE public.soc_app_posts_comments OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16609)
-- Name: soc_app_posts_comments_comment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_comments_comment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_comments_comment_id_seq OWNER TO postgres;

--
-- TOC entry 3447 (class 0 OID 0)
-- Dependencies: 219
-- Name: soc_app_posts_comments_comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_comments_comment_id_seq OWNED BY public.soc_app_posts_comments.comment_id;


--
-- TOC entry 228 (class 1259 OID 16782)
-- Name: soc_app_posts_comments_reactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_posts_comments_reactions (
    comment_reaction_id bigint NOT NULL,
    reaction_id bigint NOT NULL,
    user_id bigint NOT NULL,
    comment_id bigint NOT NULL,
    insertion_date timestamp without time zone NOT NULL,
    update_date timestamp without time zone NOT NULL
);


ALTER TABLE public.soc_app_posts_comments_reactions OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16780)
-- Name: soc_app_posts_comments_reactions_comment_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_comments_reactions_comment_reaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_comments_reactions_comment_reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3448 (class 0 OID 0)
-- Dependencies: 226
-- Name: soc_app_posts_comments_reactions_comment_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_comments_reactions_comment_reaction_id_seq OWNED BY public.soc_app_posts_comments_reactions.comment_reaction_id;


--
-- TOC entry 227 (class 1259 OID 16781)
-- Name: soc_app_posts_comments_reactions_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_comments_reactions_reaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_comments_reactions_reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3449 (class 0 OID 0)
-- Dependencies: 227
-- Name: soc_app_posts_comments_reactions_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_comments_reactions_reaction_id_seq OWNED BY public.soc_app_posts_comments_reactions.reaction_id;


--
-- TOC entry 215 (class 1259 OID 16518)
-- Name: soc_app_posts_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_post_id_seq OWNER TO postgres;

--
-- TOC entry 3450 (class 0 OID 0)
-- Dependencies: 215
-- Name: soc_app_posts_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_post_id_seq OWNED BY public.soc_app_posts.post_id;


--
-- TOC entry 225 (class 1259 OID 16757)
-- Name: soc_app_posts_reactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_posts_reactions (
    post_reaction_id bigint NOT NULL,
    reaction_id bigint NOT NULL,
    user_id bigint NOT NULL,
    post_id bigint NOT NULL,
    insertion_date timestamp without time zone NOT NULL,
    update_date timestamp without time zone NOT NULL
);


ALTER TABLE public.soc_app_posts_reactions OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16755)
-- Name: soc_app_posts_reactions_post_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_reactions_post_reaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_reactions_post_reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3451 (class 0 OID 0)
-- Dependencies: 223
-- Name: soc_app_posts_reactions_post_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_reactions_post_reaction_id_seq OWNED BY public.soc_app_posts_reactions.post_reaction_id;


--
-- TOC entry 224 (class 1259 OID 16756)
-- Name: soc_app_posts_reactions_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_posts_reactions_reaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_posts_reactions_reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3452 (class 0 OID 0)
-- Dependencies: 224
-- Name: soc_app_posts_reactions_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_posts_reactions_reaction_id_seq OWNED BY public.soc_app_posts_reactions.reaction_id;


--
-- TOC entry 212 (class 1259 OID 16485)
-- Name: soc_app_user_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_user_profile (
    profile_id bigint NOT NULL,
    user_id bigint NOT NULL,
    birth_date date NOT NULL,
    name character varying(500) NOT NULL,
    last_name character varying(500) NOT NULL,
    about_me text,
    genre character varying(10) NOT NULL,
    address character varying(500),
    country character varying(50) NOT NULL,
    city character varying(50) NOT NULL,
    insertion_date timestamp without time zone,
    update_date timestamp without time zone,
    profile_picture text
);


ALTER TABLE public.soc_app_user_profile OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16484)
-- Name: soc_app_user_profile_profile_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_user_profile_profile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_user_profile_profile_id_seq OWNER TO postgres;

--
-- TOC entry 3453 (class 0 OID 0)
-- Dependencies: 211
-- Name: soc_app_user_profile_profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_user_profile_profile_id_seq OWNED BY public.soc_app_user_profile.profile_id;


--
-- TOC entry 210 (class 1259 OID 16474)
-- Name: soc_app_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_users (
    user_id bigint NOT NULL,
    insertion_date timestamp without time zone,
    phone bigint,
    email character varying(255),
    user_name character varying(25),
    password text
);


ALTER TABLE public.soc_app_users OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16554)
-- Name: soc_app_users_interests_posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.soc_app_users_interests_posts (
    user_interest_post_id bigint NOT NULL,
    interest_id bigint,
    post_id bigint
);


ALTER TABLE public.soc_app_users_interests_posts OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16553)
-- Name: soc_app_users_interests_posts_user_interest_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_users_interests_posts_user_interest_post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_users_interests_posts_user_interest_post_id_seq OWNER TO postgres;

--
-- TOC entry 3454 (class 0 OID 0)
-- Dependencies: 217
-- Name: soc_app_users_interests_posts_user_interest_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_users_interests_posts_user_interest_post_id_seq OWNED BY public.soc_app_users_interests_posts.user_interest_post_id;


--
-- TOC entry 209 (class 1259 OID 16473)
-- Name: soc_app_users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.soc_app_users_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.soc_app_users_user_id_seq OWNER TO postgres;

--
-- TOC entry 3455 (class 0 OID 0)
-- Dependencies: 209
-- Name: soc_app_users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.soc_app_users_user_id_seq OWNED BY public.soc_app_users.user_id;


--
-- TOC entry 3223 (class 2604 OID 16811)
-- Name: soc_app_friends friend_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_friends ALTER COLUMN friend_id SET DEFAULT nextval('public.soc_app_friends_friend_id_seq'::regclass);


--
-- TOC entry 3214 (class 2604 OID 16512)
-- Name: soc_app_m_users_interests interest_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_m_users_interests ALTER COLUMN interest_id SET DEFAULT nextval('public.soc_app_m_users_interests_interest_id_seq'::regclass);


--
-- TOC entry 3218 (class 2604 OID 16728)
-- Name: soc_app_m_users_reactions reaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_m_users_reactions ALTER COLUMN reaction_id SET DEFAULT nextval('public.soc_app_m_users_reactions_reaction_id_seq'::regclass);


--
-- TOC entry 3215 (class 2604 OID 16522)
-- Name: soc_app_posts post_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts ALTER COLUMN post_id SET DEFAULT nextval('public.soc_app_posts_post_id_seq'::regclass);


--
-- TOC entry 3217 (class 2604 OID 16613)
-- Name: soc_app_posts_comments comment_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments ALTER COLUMN comment_id SET DEFAULT nextval('public.soc_app_posts_comments_comment_id_seq'::regclass);


--
-- TOC entry 3221 (class 2604 OID 16785)
-- Name: soc_app_posts_comments_reactions comment_reaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions ALTER COLUMN comment_reaction_id SET DEFAULT nextval('public.soc_app_posts_comments_reactions_comment_reaction_id_seq'::regclass);


--
-- TOC entry 3222 (class 2604 OID 16786)
-- Name: soc_app_posts_comments_reactions reaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions ALTER COLUMN reaction_id SET DEFAULT nextval('public.soc_app_posts_comments_reactions_reaction_id_seq'::regclass);


--
-- TOC entry 3219 (class 2604 OID 16760)
-- Name: soc_app_posts_reactions post_reaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions ALTER COLUMN post_reaction_id SET DEFAULT nextval('public.soc_app_posts_reactions_post_reaction_id_seq'::regclass);


--
-- TOC entry 3220 (class 2604 OID 16761)
-- Name: soc_app_posts_reactions reaction_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions ALTER COLUMN reaction_id SET DEFAULT nextval('public.soc_app_posts_reactions_reaction_id_seq'::regclass);


--
-- TOC entry 3213 (class 2604 OID 16488)
-- Name: soc_app_user_profile profile_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_user_profile ALTER COLUMN profile_id SET DEFAULT nextval('public.soc_app_user_profile_profile_id_seq'::regclass);


--
-- TOC entry 3212 (class 2604 OID 16477)
-- Name: soc_app_users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users ALTER COLUMN user_id SET DEFAULT nextval('public.soc_app_users_user_id_seq'::regclass);


--
-- TOC entry 3216 (class 2604 OID 16557)
-- Name: soc_app_users_interests_posts user_interest_post_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users_interests_posts ALTER COLUMN user_interest_post_id SET DEFAULT nextval('public.soc_app_users_interests_posts_user_interest_post_id_seq'::regclass);


--
-- TOC entry 3438 (class 0 OID 16808)
-- Dependencies: 230
-- Data for Name: soc_app_friends; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_friends (friend_id, user_id1, user_id2, friendship_date) FROM stdin;
\.


--
-- TOC entry 3422 (class 0 OID 16509)
-- Dependencies: 214
-- Data for Name: soc_app_m_users_interests; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_m_users_interests (interest_id, interest_name, insertion_date, update_date) FROM stdin;
\.


--
-- TOC entry 3430 (class 0 OID 16725)
-- Dependencies: 222
-- Data for Name: soc_app_m_users_reactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_m_users_reactions (reaction_id, reaction_name, insertion_date, update_date) FROM stdin;
\.


--
-- TOC entry 3424 (class 0 OID 16519)
-- Dependencies: 216
-- Data for Name: soc_app_posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_posts (post_id, user_id, title, description, has_multimedia, public, multimedia, insertion_date, update_date) FROM stdin;
\.


--
-- TOC entry 3428 (class 0 OID 16610)
-- Dependencies: 220
-- Data for Name: soc_app_posts_comments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_posts_comments (comment_id, user_id, post_id, comment, insertion_date, update_date, parent_comment_id) FROM stdin;
\.


--
-- TOC entry 3436 (class 0 OID 16782)
-- Dependencies: 228
-- Data for Name: soc_app_posts_comments_reactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_posts_comments_reactions (comment_reaction_id, reaction_id, user_id, comment_id, insertion_date, update_date) FROM stdin;
\.


--
-- TOC entry 3433 (class 0 OID 16757)
-- Dependencies: 225
-- Data for Name: soc_app_posts_reactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_posts_reactions (post_reaction_id, reaction_id, user_id, post_id, insertion_date, update_date) FROM stdin;
\.


--
-- TOC entry 3420 (class 0 OID 16485)
-- Dependencies: 212
-- Data for Name: soc_app_user_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_user_profile (profile_id, user_id, birth_date, name, last_name, about_me, genre, address, country, city, insertion_date, update_date, profile_picture) FROM stdin;
1	1	1999-01-01	Juan	Doe	Descripcion	Masculino	Avenida siempre viva 777	Peru	Lima	2023-09-14 21:32:40.349389	2023-09-14 21:38:33.61107	
\.


--
-- TOC entry 3418 (class 0 OID 16474)
-- Dependencies: 210
-- Data for Name: soc_app_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_users (user_id, insertion_date, phone, email, user_name, password) FROM stdin;
1	2023-09-14 21:32:40.349389	994421210	test@test.com	test	test
\.


--
-- TOC entry 3426 (class 0 OID 16554)
-- Dependencies: 218
-- Data for Name: soc_app_users_interests_posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.soc_app_users_interests_posts (user_interest_post_id, interest_id, post_id) FROM stdin;
\.


--
-- TOC entry 3456 (class 0 OID 0)
-- Dependencies: 229
-- Name: soc_app_friends_friend_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_friends_friend_id_seq', 1, false);


--
-- TOC entry 3457 (class 0 OID 0)
-- Dependencies: 213
-- Name: soc_app_m_users_interests_interest_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_m_users_interests_interest_id_seq', 1, false);


--
-- TOC entry 3458 (class 0 OID 0)
-- Dependencies: 221
-- Name: soc_app_m_users_reactions_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_m_users_reactions_reaction_id_seq', 1, false);


--
-- TOC entry 3459 (class 0 OID 0)
-- Dependencies: 219
-- Name: soc_app_posts_comments_comment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_comments_comment_id_seq', 1, false);


--
-- TOC entry 3460 (class 0 OID 0)
-- Dependencies: 226
-- Name: soc_app_posts_comments_reactions_comment_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_comments_reactions_comment_reaction_id_seq', 1, false);


--
-- TOC entry 3461 (class 0 OID 0)
-- Dependencies: 227
-- Name: soc_app_posts_comments_reactions_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_comments_reactions_reaction_id_seq', 1, false);


--
-- TOC entry 3462 (class 0 OID 0)
-- Dependencies: 215
-- Name: soc_app_posts_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_post_id_seq', 1, false);


--
-- TOC entry 3463 (class 0 OID 0)
-- Dependencies: 223
-- Name: soc_app_posts_reactions_post_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_reactions_post_reaction_id_seq', 1, false);


--
-- TOC entry 3464 (class 0 OID 0)
-- Dependencies: 224
-- Name: soc_app_posts_reactions_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_posts_reactions_reaction_id_seq', 1, false);


--
-- TOC entry 3465 (class 0 OID 0)
-- Dependencies: 211
-- Name: soc_app_user_profile_profile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_user_profile_profile_id_seq', 1, true);


--
-- TOC entry 3466 (class 0 OID 0)
-- Dependencies: 217
-- Name: soc_app_users_interests_posts_user_interest_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_users_interests_posts_user_interest_post_id_seq', 1, false);


--
-- TOC entry 3467 (class 0 OID 0)
-- Dependencies: 209
-- Name: soc_app_users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.soc_app_users_user_id_seq', 1, true);


--
-- TOC entry 3226 (class 2606 OID 16481)
-- Name: soc_app_users pk_soc_app_users; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users
    ADD CONSTRAINT pk_soc_app_users PRIMARY KEY (user_id);


--
-- TOC entry 3261 (class 2606 OID 16813)
-- Name: soc_app_friends soc_app_friends_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_friends
    ADD CONSTRAINT soc_app_friends_pkey PRIMARY KEY (friend_id);


--
-- TOC entry 3236 (class 2606 OID 16514)
-- Name: soc_app_m_users_interests soc_app_m_users_interests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_m_users_interests
    ADD CONSTRAINT soc_app_m_users_interests_pkey PRIMARY KEY (interest_id);


--
-- TOC entry 3252 (class 2606 OID 16730)
-- Name: soc_app_m_users_reactions soc_app_m_users_reactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_m_users_reactions
    ADD CONSTRAINT soc_app_m_users_reactions_pkey PRIMARY KEY (reaction_id);


--
-- TOC entry 3248 (class 2606 OID 16617)
-- Name: soc_app_posts_comments soc_app_posts_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments
    ADD CONSTRAINT soc_app_posts_comments_pkey PRIMARY KEY (comment_id);


--
-- TOC entry 3258 (class 2606 OID 16788)
-- Name: soc_app_posts_comments_reactions soc_app_posts_comments_reactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions
    ADD CONSTRAINT soc_app_posts_comments_reactions_pkey PRIMARY KEY (comment_reaction_id);


--
-- TOC entry 3241 (class 2606 OID 16526)
-- Name: soc_app_posts soc_app_posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts
    ADD CONSTRAINT soc_app_posts_pkey PRIMARY KEY (post_id);


--
-- TOC entry 3255 (class 2606 OID 16763)
-- Name: soc_app_posts_reactions soc_app_posts_reactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions
    ADD CONSTRAINT soc_app_posts_reactions_pkey PRIMARY KEY (post_reaction_id);


--
-- TOC entry 3231 (class 2606 OID 16492)
-- Name: soc_app_user_profile soc_app_user_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_user_profile
    ADD CONSTRAINT soc_app_user_profile_pkey PRIMARY KEY (profile_id);


--
-- TOC entry 3243 (class 2606 OID 16559)
-- Name: soc_app_users_interests_posts soc_app_users_interests_posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users_interests_posts
    ADD CONSTRAINT soc_app_users_interests_posts_pkey PRIMARY KEY (user_interest_post_id);


--
-- TOC entry 3233 (class 1259 OID 16516)
-- Name: idx1_soc_app_m_users_interests; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx1_soc_app_m_users_interests ON public.soc_app_m_users_interests USING btree (insertion_date);


--
-- TOC entry 3249 (class 1259 OID 16732)
-- Name: idx1_soc_app_m_users_reactions; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx1_soc_app_m_users_reactions ON public.soc_app_m_users_reactions USING btree (insertion_date);


--
-- TOC entry 3238 (class 1259 OID 16532)
-- Name: idx1_soc_app_posts; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx1_soc_app_posts ON public.soc_app_posts USING btree (insertion_date);


--
-- TOC entry 3245 (class 1259 OID 16633)
-- Name: idx1_soc_app_posts_comments; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx1_soc_app_posts_comments ON public.soc_app_posts_comments USING btree (insertion_date);


--
-- TOC entry 3228 (class 1259 OID 16499)
-- Name: idx1_soc_app_user_profile; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx1_soc_app_user_profile ON public.soc_app_user_profile USING btree (insertion_date);


--
-- TOC entry 3234 (class 1259 OID 16517)
-- Name: idx2_soc_app_m_users_interests; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx2_soc_app_m_users_interests ON public.soc_app_m_users_interests USING btree (update_date);


--
-- TOC entry 3250 (class 1259 OID 16733)
-- Name: idx2_soc_app_m_users_reactions; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx2_soc_app_m_users_reactions ON public.soc_app_m_users_reactions USING btree (update_date);


--
-- TOC entry 3239 (class 1259 OID 16533)
-- Name: idx2_soc_app_posts; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx2_soc_app_posts ON public.soc_app_posts USING btree (update_date);


--
-- TOC entry 3246 (class 1259 OID 16634)
-- Name: idx2_soc_app_posts_comments; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx2_soc_app_posts_comments ON public.soc_app_posts_comments USING btree (update_date);


--
-- TOC entry 3229 (class 1259 OID 16500)
-- Name: idx2_soc_app_user_profile; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx2_soc_app_user_profile ON public.soc_app_user_profile USING btree (update_date);


--
-- TOC entry 3224 (class 1259 OID 16483)
-- Name: idx_soc_app_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_soc_app_users ON public.soc_app_users USING btree (insertion_date);


--
-- TOC entry 3262 (class 1259 OID 16824)
-- Name: uq_soc_app_friends; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_friends ON public.soc_app_friends USING btree (user_id1, user_id2);


--
-- TOC entry 3237 (class 1259 OID 16515)
-- Name: uq_soc_app_m_users_interests; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_m_users_interests ON public.soc_app_m_users_interests USING btree (interest_name);


--
-- TOC entry 3253 (class 1259 OID 16731)
-- Name: uq_soc_app_m_users_reactions; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_m_users_reactions ON public.soc_app_m_users_reactions USING btree (reaction_name);


--
-- TOC entry 3259 (class 1259 OID 16804)
-- Name: uq_soc_app_posts_comments_reactions; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_posts_comments_reactions ON public.soc_app_posts_comments_reactions USING btree (user_id, comment_id);


--
-- TOC entry 3256 (class 1259 OID 16779)
-- Name: uq_soc_app_posts_reactions; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_posts_reactions ON public.soc_app_posts_reactions USING btree (user_id, post_id);


--
-- TOC entry 3232 (class 1259 OID 16498)
-- Name: uq_soc_app_user_profile; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_user_profile ON public.soc_app_user_profile USING btree (user_id);


--
-- TOC entry 3227 (class 1259 OID 16482)
-- Name: uq_soc_app_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_users ON public.soc_app_users USING btree (phone, email, user_name);


--
-- TOC entry 3244 (class 1259 OID 16570)
-- Name: uq_soc_app_users_interests_posts; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uq_soc_app_users_interests_posts ON public.soc_app_users_interests_posts USING btree (interest_id, post_id);


--
-- TOC entry 3276 (class 2606 OID 16814)
-- Name: soc_app_friends soc_app_friends_user_id1_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_friends
    ADD CONSTRAINT soc_app_friends_user_id1_fkey FOREIGN KEY (user_id1) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3277 (class 2606 OID 16819)
-- Name: soc_app_friends soc_app_friends_user_id2_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_friends
    ADD CONSTRAINT soc_app_friends_user_id2_fkey FOREIGN KEY (user_id2) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3269 (class 2606 OID 16628)
-- Name: soc_app_posts_comments soc_app_posts_comments_parent_comment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments
    ADD CONSTRAINT soc_app_posts_comments_parent_comment_id_fkey FOREIGN KEY (parent_comment_id) REFERENCES public.soc_app_posts_comments(comment_id);


--
-- TOC entry 3268 (class 2606 OID 16623)
-- Name: soc_app_posts_comments soc_app_posts_comments_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments
    ADD CONSTRAINT soc_app_posts_comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.soc_app_posts(post_id);


--
-- TOC entry 3274 (class 2606 OID 16794)
-- Name: soc_app_posts_comments_reactions soc_app_posts_comments_reactions_comment_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions
    ADD CONSTRAINT soc_app_posts_comments_reactions_comment_id_fkey FOREIGN KEY (comment_id) REFERENCES public.soc_app_posts_comments(comment_id);


--
-- TOC entry 3275 (class 2606 OID 16799)
-- Name: soc_app_posts_comments_reactions soc_app_posts_comments_reactions_reaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions
    ADD CONSTRAINT soc_app_posts_comments_reactions_reaction_id_fkey FOREIGN KEY (reaction_id) REFERENCES public.soc_app_m_users_reactions(reaction_id);


--
-- TOC entry 3273 (class 2606 OID 16789)
-- Name: soc_app_posts_comments_reactions soc_app_posts_comments_reactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments_reactions
    ADD CONSTRAINT soc_app_posts_comments_reactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3267 (class 2606 OID 16618)
-- Name: soc_app_posts_comments soc_app_posts_comments_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_comments
    ADD CONSTRAINT soc_app_posts_comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3271 (class 2606 OID 16769)
-- Name: soc_app_posts_reactions soc_app_posts_reactions_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions
    ADD CONSTRAINT soc_app_posts_reactions_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.soc_app_posts(post_id);


--
-- TOC entry 3272 (class 2606 OID 16774)
-- Name: soc_app_posts_reactions soc_app_posts_reactions_reaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions
    ADD CONSTRAINT soc_app_posts_reactions_reaction_id_fkey FOREIGN KEY (reaction_id) REFERENCES public.soc_app_m_users_reactions(reaction_id);


--
-- TOC entry 3270 (class 2606 OID 16764)
-- Name: soc_app_posts_reactions soc_app_posts_reactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts_reactions
    ADD CONSTRAINT soc_app_posts_reactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3264 (class 2606 OID 16527)
-- Name: soc_app_posts soc_app_posts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_posts
    ADD CONSTRAINT soc_app_posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3263 (class 2606 OID 16493)
-- Name: soc_app_user_profile soc_app_user_profile_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_user_profile
    ADD CONSTRAINT soc_app_user_profile_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.soc_app_users(user_id);


--
-- TOC entry 3265 (class 2606 OID 16560)
-- Name: soc_app_users_interests_posts soc_app_users_interests_posts_interest_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users_interests_posts
    ADD CONSTRAINT soc_app_users_interests_posts_interest_id_fkey FOREIGN KEY (interest_id) REFERENCES public.soc_app_m_users_interests(interest_id);


--
-- TOC entry 3266 (class 2606 OID 16565)
-- Name: soc_app_users_interests_posts soc_app_users_interests_posts_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.soc_app_users_interests_posts
    ADD CONSTRAINT soc_app_users_interests_posts_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.soc_app_posts(post_id);


-- Completed on 2023-09-20 22:31:43

--
-- PostgreSQL database dump complete
--

