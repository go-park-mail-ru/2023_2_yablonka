--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 15.3 (Ubuntu 15.3-1.pgdg22.04+1)

-- Started on 2023-10-26 00:41:25 MSK

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3654 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 256 (class 1259 OID 25421)
-- Name: Session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Session" (
    token character varying(64) NOT NULL,
    expiration_date timestamp without time zone DEFAULT (CURRENT_TIMESTAMP + '1 day'::interval) NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public."Session" OWNER TO postgres;

--
-- TOC entry 255 (class 1259 OID 25420)
-- Name: Session_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Session_id_user_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Session_id_user_seq" OWNER TO postgres;

--
-- TOC entry 3656 (class 0 OID 0)
-- Dependencies: 255
-- Name: Session_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Session_id_user_seq" OWNED BY public."Session".id_user;


--
-- TOC entry 258 (class 1259 OID 25431)
-- Name: Tag; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Tag" (
    id integer NOT NULL,
    name character varying(35) NOT NULL,
    color character varying(6) DEFAULT 'FFFFFF'::character varying NOT NULL
);


ALTER TABLE public."Tag" OWNER TO postgres;

--
-- TOC entry 257 (class 1259 OID 25430)
-- Name: Tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Tag_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Tag_id_seq" OWNER TO postgres;

--
-- TOC entry 3657 (class 0 OID 0)
-- Dependencies: 257
-- Name: Tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Tag_id_seq" OWNED BY public."Tag".id;


--
-- TOC entry 246 (class 1259 OID 25375)
-- Name: Task_Embedding; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Task_Embedding" (
    id integer NOT NULL,
    id_task integer NOT NULL,
    id_user integer NOT NULL,
    url character varying(2048) NOT NULL
);


ALTER TABLE public."Task_Embedding" OWNER TO postgres;

--
-- TOC entry 243 (class 1259 OID 25372)
-- Name: Task_Embedding_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Task_Embedding_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Task_Embedding_id_seq" OWNER TO postgres;

--
-- TOC entry 3658 (class 0 OID 0)
-- Dependencies: 243
-- Name: Task_Embedding_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_seq" OWNED BY public."Task_Embedding".id;


--
-- TOC entry 244 (class 1259 OID 25373)
-- Name: Task_Embedding_id_task_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Task_Embedding_id_task_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Task_Embedding_id_task_seq" OWNER TO postgres;

--
-- TOC entry 3659 (class 0 OID 0)
-- Dependencies: 244
-- Name: Task_Embedding_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_task_seq" OWNED BY public."Task_Embedding".id_task;


--
-- TOC entry 245 (class 1259 OID 25374)
-- Name: Task_Embedding_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Task_Embedding_id_user_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Task_Embedding_id_user_seq" OWNER TO postgres;

--
-- TOC entry 3660 (class 0 OID 0)
-- Dependencies: 245
-- Name: Task_Embedding_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_user_seq" OWNED BY public."Task_Embedding".id_user;


--
-- TOC entry 211 (class 1259 OID 25231)
-- Name: board; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board (
    id integer NOT NULL,
    id_workspace integer NOT NULL,
    name character varying(150) DEFAULT 'Доска'::character varying NOT NULL,
    description text,
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    thumbnail_url character varying(2048)
);


ALTER TABLE public.board OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 25229)
-- Name: board_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_id_seq OWNER TO postgres;

--
-- TOC entry 3661 (class 0 OID 0)
-- Dependencies: 209
-- Name: board_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_id_seq OWNED BY public.board.id;


--
-- TOC entry 210 (class 1259 OID 25230)
-- Name: board_id_workspace_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_id_workspace_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_id_workspace_seq OWNER TO postgres;

--
-- TOC entry 3662 (class 0 OID 0)
-- Dependencies: 210
-- Name: board_id_workspace_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_id_workspace_seq OWNED BY public.board.id_workspace;


--
-- TOC entry 254 (class 1259 OID 25410)
-- Name: board_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_template (
    id integer NOT NULL,
    data json NOT NULL
);


ALTER TABLE public.board_template OWNER TO postgres;

--
-- TOC entry 253 (class 1259 OID 25409)
-- Name: board_template_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_template_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_template_id_seq OWNER TO postgres;

--
-- TOC entry 3663 (class 0 OID 0)
-- Dependencies: 253
-- Name: board_template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_template_id_seq OWNED BY public.board_template.id;


--
-- TOC entry 235 (class 1259 OID 25331)
-- Name: board_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_user (
    id_board integer NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public.board_user OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 25329)
-- Name: board_user_id_board_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_user_id_board_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_user_id_board_seq OWNER TO postgres;

--
-- TOC entry 3664 (class 0 OID 0)
-- Dependencies: 233
-- Name: board_user_id_board_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_user_id_board_seq OWNED BY public.board_user.id_board;


--
-- TOC entry 234 (class 1259 OID 25330)
-- Name: board_user_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_user_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_user_id_user_seq OWNER TO postgres;

--
-- TOC entry 3665 (class 0 OID 0)
-- Dependencies: 234
-- Name: board_user_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_user_id_user_seq OWNED BY public.board_user.id_user;


--
-- TOC entry 264 (class 1259 OID 25451)
-- Name: checklist; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checklist (
    id integer NOT NULL,
    id_task integer NOT NULL,
    name character varying(100) DEFAULT 'Чек-лист'::character varying NOT NULL,
    list_position smallint NOT NULL
);


ALTER TABLE public.checklist OWNER TO postgres;

--
-- TOC entry 262 (class 1259 OID 25449)
-- Name: checklist_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checklist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checklist_id_seq OWNER TO postgres;

--
-- TOC entry 3666 (class 0 OID 0)
-- Dependencies: 262
-- Name: checklist_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_id_seq OWNED BY public.checklist.id;


--
-- TOC entry 263 (class 1259 OID 25450)
-- Name: checklist_id_task_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checklist_id_task_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checklist_id_task_seq OWNER TO postgres;

--
-- TOC entry 3667 (class 0 OID 0)
-- Dependencies: 263
-- Name: checklist_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_id_task_seq OWNED BY public.checklist.id_task;


--
-- TOC entry 267 (class 1259 OID 25463)
-- Name: checklist_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checklist_item (
    id integer NOT NULL,
    id_checklist integer NOT NULL,
    name text NOT NULL,
    done boolean DEFAULT false NOT NULL,
    list_position smallint NOT NULL
);


ALTER TABLE public.checklist_item OWNER TO postgres;

--
-- TOC entry 266 (class 1259 OID 25462)
-- Name: checklist_item_id_checklist_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checklist_item_id_checklist_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checklist_item_id_checklist_seq OWNER TO postgres;

--
-- TOC entry 3668 (class 0 OID 0)
-- Dependencies: 266
-- Name: checklist_item_id_checklist_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_item_id_checklist_seq OWNED BY public.checklist_item.id_checklist;


--
-- TOC entry 265 (class 1259 OID 25461)
-- Name: checklist_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checklist_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checklist_item_id_seq OWNER TO postgres;

--
-- TOC entry 3669 (class 0 OID 0)
-- Dependencies: 265
-- Name: checklist_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_item_id_seq OWNED BY public.checklist_item.id;


--
-- TOC entry 238 (class 1259 OID 25342)
-- Name: column; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."column" (
    id integer NOT NULL,
    id_board integer NOT NULL,
    name character varying(150) DEFAULT 'Столбец'::character varying NOT NULL,
    description text,
    list_position smallint NOT NULL
);


ALTER TABLE public."column" OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 25341)
-- Name: column_id_board_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.column_id_board_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.column_id_board_seq OWNER TO postgres;

--
-- TOC entry 3670 (class 0 OID 0)
-- Dependencies: 237
-- Name: column_id_board_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.column_id_board_seq OWNED BY public."column".id_board;


--
-- TOC entry 236 (class 1259 OID 25340)
-- Name: column_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.column_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.column_id_seq OWNER TO postgres;

--
-- TOC entry 3671 (class 0 OID 0)
-- Dependencies: 236
-- Name: column_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.column_id_seq OWNED BY public."column".id;


--
-- TOC entry 221 (class 1259 OID 25269)
-- Name: comment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment (
    id integer NOT NULL,
    id_user integer NOT NULL,
    id_task integer NOT NULL,
    content text NOT NULL,
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.comment OWNER TO postgres;

--
-- TOC entry 250 (class 1259 OID 25388)
-- Name: comment_embedding; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment_embedding (
    id integer NOT NULL,
    id_user integer NOT NULL,
    id_comment integer NOT NULL,
    url character varying(2048) NOT NULL
);


ALTER TABLE public.comment_embedding OWNER TO postgres;

--
-- TOC entry 249 (class 1259 OID 25387)
-- Name: comment_embedding_id_comment_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_embedding_id_comment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_embedding_id_comment_seq OWNER TO postgres;

--
-- TOC entry 3672 (class 0 OID 0)
-- Dependencies: 249
-- Name: comment_embedding_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_comment_seq OWNED BY public.comment_embedding.id_comment;


--
-- TOC entry 247 (class 1259 OID 25385)
-- Name: comment_embedding_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_embedding_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_embedding_id_seq OWNER TO postgres;

--
-- TOC entry 3673 (class 0 OID 0)
-- Dependencies: 247
-- Name: comment_embedding_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_seq OWNED BY public.comment_embedding.id;


--
-- TOC entry 248 (class 1259 OID 25386)
-- Name: comment_embedding_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_embedding_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_embedding_id_user_seq OWNER TO postgres;

--
-- TOC entry 3674 (class 0 OID 0)
-- Dependencies: 248
-- Name: comment_embedding_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_user_seq OWNED BY public.comment_embedding.id_user;


--
-- TOC entry 218 (class 1259 OID 25266)
-- Name: comment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_id_seq OWNER TO postgres;

--
-- TOC entry 3675 (class 0 OID 0)
-- Dependencies: 218
-- Name: comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_seq OWNED BY public.comment.id;


--
-- TOC entry 220 (class 1259 OID 25268)
-- Name: comment_id_task_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_id_task_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_id_task_seq OWNER TO postgres;

--
-- TOC entry 3676 (class 0 OID 0)
-- Dependencies: 220
-- Name: comment_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_task_seq OWNED BY public.comment.id_task;


--
-- TOC entry 219 (class 1259 OID 25267)
-- Name: comment_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_id_user_seq OWNER TO postgres;

--
-- TOC entry 3677 (class 0 OID 0)
-- Dependencies: 219
-- Name: comment_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_user_seq OWNED BY public.comment.id_user;


--
-- TOC entry 224 (class 1259 OID 25284)
-- Name: comment_reply; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment_reply (
    id_reply integer NOT NULL,
    id_comment integer NOT NULL
);


ALTER TABLE public.comment_reply OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 25283)
-- Name: comment_reply_id_comment_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_reply_id_comment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_reply_id_comment_seq OWNER TO postgres;

--
-- TOC entry 3678 (class 0 OID 0)
-- Dependencies: 223
-- Name: comment_reply_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_reply_id_comment_seq OWNED BY public.comment_reply.id_comment;


--
-- TOC entry 222 (class 1259 OID 25282)
-- Name: comment_reply_id_reply_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.comment_reply_id_reply_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comment_reply_id_reply_seq OWNER TO postgres;

--
-- TOC entry 3679 (class 0 OID 0)
-- Dependencies: 222
-- Name: comment_reply_id_reply_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_reply_id_reply_seq OWNED BY public.comment_reply.id_reply;


--
-- TOC entry 270 (class 1259 OID 25477)
-- Name: favourite_boards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.favourite_boards (
    id_board integer NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public.favourite_boards OWNER TO postgres;

--
-- TOC entry 268 (class 1259 OID 25475)
-- Name: favourite_boards_id_board_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.favourite_boards_id_board_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.favourite_boards_id_board_seq OWNER TO postgres;

--
-- TOC entry 3680 (class 0 OID 0)
-- Dependencies: 268
-- Name: favourite_boards_id_board_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favourite_boards_id_board_seq OWNED BY public.favourite_boards.id_board;


--
-- TOC entry 269 (class 1259 OID 25476)
-- Name: favourite_boards_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.favourite_boards_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.favourite_boards_id_user_seq OWNER TO postgres;

--
-- TOC entry 3681 (class 0 OID 0)
-- Dependencies: 269
-- Name: favourite_boards_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favourite_boards_id_user_seq OWNED BY public.favourite_boards.id_user;


--
-- TOC entry 217 (class 1259 OID 25258)
-- Name: reaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reaction (
    id integer NOT NULL,
    id_comment integer NOT NULL,
    id_user integer NOT NULL,
    content character varying(2) NOT NULL
);


ALTER TABLE public.reaction OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 25256)
-- Name: reaction_id_comment_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reaction_id_comment_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reaction_id_comment_seq OWNER TO postgres;

--
-- TOC entry 3682 (class 0 OID 0)
-- Dependencies: 215
-- Name: reaction_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_comment_seq OWNED BY public.reaction.id_comment;


--
-- TOC entry 214 (class 1259 OID 25255)
-- Name: reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reaction_id_seq OWNER TO postgres;

--
-- TOC entry 3683 (class 0 OID 0)
-- Dependencies: 214
-- Name: reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_seq OWNED BY public.reaction.id;


--
-- TOC entry 216 (class 1259 OID 25257)
-- Name: reaction_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reaction_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reaction_id_user_seq OWNER TO postgres;

--
-- TOC entry 3684 (class 0 OID 0)
-- Dependencies: 216
-- Name: reaction_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_user_seq OWNED BY public.reaction.id_user;


--
-- TOC entry 232 (class 1259 OID 25318)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    id integer NOT NULL,
    name character varying(100) DEFAULT 'Роль'::character varying NOT NULL,
    description text
);


ALTER TABLE public.role OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 25317)
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.role_id_seq OWNER TO postgres;

--
-- TOC entry 3685 (class 0 OID 0)
-- Dependencies: 231
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.role_id_seq OWNED BY public.role.id;


--
-- TOC entry 261 (class 1259 OID 25442)
-- Name: tag_task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tag_task (
    id_tag integer NOT NULL,
    id_task integer NOT NULL
);


ALTER TABLE public.tag_task OWNER TO postgres;

--
-- TOC entry 259 (class 1259 OID 25440)
-- Name: tag_task_id_tag_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_task_id_tag_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_task_id_tag_seq OWNER TO postgres;

--
-- TOC entry 3686 (class 0 OID 0)
-- Dependencies: 259
-- Name: tag_task_id_tag_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_task_id_tag_seq OWNED BY public.tag_task.id_tag;


--
-- TOC entry 260 (class 1259 OID 25441)
-- Name: tag_task_id_task_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_task_id_task_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_task_id_task_seq OWNER TO postgres;

--
-- TOC entry 3687 (class 0 OID 0)
-- Dependencies: 260
-- Name: tag_task_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_task_id_task_seq OWNED BY public.tag_task.id_task;


--
-- TOC entry 241 (class 1259 OID 25354)
-- Name: task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task (
    id integer NOT NULL,
    id_column integer NOT NULL,
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    description text,
    name character varying(150) DEFAULT 'Задача'::character varying NOT NULL,
    start timestamp without time zone,
    "end" timestamp without time zone,
    list_postition smallint NOT NULL
);


ALTER TABLE public.task OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 25353)
-- Name: task_id_column_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_id_column_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_column_seq OWNER TO postgres;

--
-- TOC entry 3688 (class 0 OID 0)
-- Dependencies: 240
-- Name: task_id_column_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_column_seq OWNED BY public.task.id_column;


--
-- TOC entry 239 (class 1259 OID 25352)
-- Name: task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_id_seq OWNER TO postgres;

--
-- TOC entry 3689 (class 0 OID 0)
-- Dependencies: 239
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- TOC entry 252 (class 1259 OID 25399)
-- Name: task_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_template (
    id integer NOT NULL,
    data json NOT NULL
);


ALTER TABLE public.task_template OWNER TO postgres;

--
-- TOC entry 251 (class 1259 OID 25398)
-- Name: task_template_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.task_template_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.task_template_id_seq OWNER TO postgres;

--
-- TOC entry 3690 (class 0 OID 0)
-- Dependencies: 251
-- Name: task_template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_template_id_seq OWNED BY public.task_template.id;


--
-- TOC entry 242 (class 1259 OID 25367)
-- Name: task_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_user (
    id_user bigint NOT NULL,
    id_task bigint NOT NULL
);


ALTER TABLE public.task_user OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 25245)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    email character varying(256) NOT NULL,
    password_hash character varying(256) NOT NULL,
    name character varying(100),
    surname character varying(100),
    avatar_url character varying(2048),
    description text
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- TOC entry 276 (class 1259 OID 25497)
-- Name: user_board_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_board_template (
    id_user integer NOT NULL,
    id_template integer NOT NULL
);


ALTER TABLE public.user_board_template OWNER TO postgres;

--
-- TOC entry 275 (class 1259 OID 25496)
-- Name: user_board_template_id_template_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_board_template_id_template_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_board_template_id_template_seq OWNER TO postgres;

--
-- TOC entry 3691 (class 0 OID 0)
-- Dependencies: 275
-- Name: user_board_template_id_template_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_board_template_id_template_seq OWNED BY public.user_board_template.id_template;


--
-- TOC entry 274 (class 1259 OID 25495)
-- Name: user_board_template_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_board_template_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_board_template_id_user_seq OWNER TO postgres;

--
-- TOC entry 3692 (class 0 OID 0)
-- Dependencies: 274
-- Name: user_board_template_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_board_template_id_user_seq OWNED BY public.user_board_template.id_user;


--
-- TOC entry 212 (class 1259 OID 25244)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- TOC entry 3693 (class 0 OID 0)
-- Dependencies: 212
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- TOC entry 273 (class 1259 OID 25486)
-- Name: user_task_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_task_template (
    id_user integer NOT NULL,
    id_template integer NOT NULL
);


ALTER TABLE public.user_task_template OWNER TO postgres;

--
-- TOC entry 272 (class 1259 OID 25485)
-- Name: user_task_template_id_template_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_task_template_id_template_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_task_template_id_template_seq OWNER TO postgres;

--
-- TOC entry 3694 (class 0 OID 0)
-- Dependencies: 272
-- Name: user_task_template_id_template_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_task_template_id_template_seq OWNED BY public.user_task_template.id_template;


--
-- TOC entry 271 (class 1259 OID 25484)
-- Name: user_task_template_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_task_template_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_task_template_id_user_seq OWNER TO postgres;

--
-- TOC entry 3695 (class 0 OID 0)
-- Dependencies: 271
-- Name: user_task_template_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_task_template_id_user_seq OWNED BY public.user_task_template.id_user;


--
-- TOC entry 230 (class 1259 OID 25309)
-- Name: user_workspace; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_workspace (
    id_user integer NOT NULL,
    id_workspace integer NOT NULL,
    id_role integer NOT NULL
);


ALTER TABLE public.user_workspace OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 25308)
-- Name: user_workspace_id_role_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_workspace_id_role_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_workspace_id_role_seq OWNER TO postgres;

--
-- TOC entry 3696 (class 0 OID 0)
-- Dependencies: 229
-- Name: user_workspace_id_role_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_role_seq OWNED BY public.user_workspace.id_role;


--
-- TOC entry 227 (class 1259 OID 25306)
-- Name: user_workspace_id_user_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_workspace_id_user_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_workspace_id_user_seq OWNER TO postgres;

--
-- TOC entry 3697 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_workspace_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_user_seq OWNED BY public.user_workspace.id_user;


--
-- TOC entry 228 (class 1259 OID 25307)
-- Name: user_workspace_id_workspace_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_workspace_id_workspace_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_workspace_id_workspace_seq OWNER TO postgres;

--
-- TOC entry 3698 (class 0 OID 0)
-- Dependencies: 228
-- Name: user_workspace_id_workspace_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_workspace_seq OWNED BY public.user_workspace.id_workspace;


--
-- TOC entry 226 (class 1259 OID 25294)
-- Name: workspace; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workspace (
    id integer NOT NULL,
    name character varying(150) DEFAULT 'Рабочее место'::character varying NOT NULL,
    thumbnail_url character varying(2048),
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    description text
);


ALTER TABLE public.workspace OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 25293)
-- Name: workspace_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.workspace_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.workspace_id_seq OWNER TO postgres;

--
-- TOC entry 3699 (class 0 OID 0)
-- Dependencies: 225
-- Name: workspace_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.workspace_id_seq OWNED BY public.workspace.id;


--
-- TOC entry 3382 (class 2604 OID 25425)
-- Name: Session id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session" ALTER COLUMN id_user SET DEFAULT nextval('public."Session_id_user_seq"'::regclass);


--
-- TOC entry 3383 (class 2604 OID 25434)
-- Name: Tag id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tag" ALTER COLUMN id SET DEFAULT nextval('public."Tag_id_seq"'::regclass);


--
-- TOC entry 3373 (class 2604 OID 25378)
-- Name: Task_Embedding id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id SET DEFAULT nextval('public."Task_Embedding_id_seq"'::regclass);


--
-- TOC entry 3374 (class 2604 OID 25379)
-- Name: Task_Embedding id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id_task SET DEFAULT nextval('public."Task_Embedding_id_task_seq"'::regclass);


--
-- TOC entry 3375 (class 2604 OID 25380)
-- Name: Task_Embedding id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id_user SET DEFAULT nextval('public."Task_Embedding_id_user_seq"'::regclass);


--
-- TOC entry 3342 (class 2604 OID 25234)
-- Name: board id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board ALTER COLUMN id SET DEFAULT nextval('public.board_id_seq'::regclass);


--
-- TOC entry 3343 (class 2604 OID 25235)
-- Name: board id_workspace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board ALTER COLUMN id_workspace SET DEFAULT nextval('public.board_id_workspace_seq'::regclass);


--
-- TOC entry 3380 (class 2604 OID 25413)
-- Name: board_template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template ALTER COLUMN id SET DEFAULT nextval('public.board_template_id_seq'::regclass);


--
-- TOC entry 3364 (class 2604 OID 25334)
-- Name: board_user id_board; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user ALTER COLUMN id_board SET DEFAULT nextval('public.board_user_id_board_seq'::regclass);


--
-- TOC entry 3365 (class 2604 OID 25335)
-- Name: board_user id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user ALTER COLUMN id_user SET DEFAULT nextval('public.board_user_id_user_seq'::regclass);


--
-- TOC entry 3387 (class 2604 OID 25454)
-- Name: checklist id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist ALTER COLUMN id SET DEFAULT nextval('public.checklist_id_seq'::regclass);


--
-- TOC entry 3388 (class 2604 OID 25455)
-- Name: checklist id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist ALTER COLUMN id_task SET DEFAULT nextval('public.checklist_id_task_seq'::regclass);


--
-- TOC entry 3390 (class 2604 OID 25466)
-- Name: checklist_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item ALTER COLUMN id SET DEFAULT nextval('public.checklist_item_id_seq'::regclass);


--
-- TOC entry 3391 (class 2604 OID 25467)
-- Name: checklist_item id_checklist; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item ALTER COLUMN id_checklist SET DEFAULT nextval('public.checklist_item_id_checklist_seq'::regclass);


--
-- TOC entry 3366 (class 2604 OID 25345)
-- Name: column id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column" ALTER COLUMN id SET DEFAULT nextval('public.column_id_seq'::regclass);


--
-- TOC entry 3367 (class 2604 OID 25346)
-- Name: column id_board; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column" ALTER COLUMN id_board SET DEFAULT nextval('public.column_id_board_seq'::regclass);


--
-- TOC entry 3350 (class 2604 OID 25272)
-- Name: comment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id SET DEFAULT nextval('public.comment_id_seq'::regclass);


--
-- TOC entry 3351 (class 2604 OID 25273)
-- Name: comment id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id_user SET DEFAULT nextval('public.comment_id_user_seq'::regclass);


--
-- TOC entry 3352 (class 2604 OID 25274)
-- Name: comment id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id_task SET DEFAULT nextval('public.comment_id_task_seq'::regclass);


--
-- TOC entry 3376 (class 2604 OID 25391)
-- Name: comment_embedding id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id SET DEFAULT nextval('public.comment_embedding_id_seq'::regclass);


--
-- TOC entry 3377 (class 2604 OID 25392)
-- Name: comment_embedding id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id_user SET DEFAULT nextval('public.comment_embedding_id_user_seq'::regclass);


--
-- TOC entry 3378 (class 2604 OID 25393)
-- Name: comment_embedding id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id_comment SET DEFAULT nextval('public.comment_embedding_id_comment_seq'::regclass);


--
-- TOC entry 3354 (class 2604 OID 25287)
-- Name: comment_reply id_reply; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply ALTER COLUMN id_reply SET DEFAULT nextval('public.comment_reply_id_reply_seq'::regclass);


--
-- TOC entry 3355 (class 2604 OID 25288)
-- Name: comment_reply id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply ALTER COLUMN id_comment SET DEFAULT nextval('public.comment_reply_id_comment_seq'::regclass);


--
-- TOC entry 3393 (class 2604 OID 25480)
-- Name: favourite_boards id_board; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards ALTER COLUMN id_board SET DEFAULT nextval('public.favourite_boards_id_board_seq'::regclass);


--
-- TOC entry 3394 (class 2604 OID 25481)
-- Name: favourite_boards id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards ALTER COLUMN id_user SET DEFAULT nextval('public.favourite_boards_id_user_seq'::regclass);


--
-- TOC entry 3347 (class 2604 OID 25261)
-- Name: reaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id SET DEFAULT nextval('public.reaction_id_seq'::regclass);


--
-- TOC entry 3348 (class 2604 OID 25262)
-- Name: reaction id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id_comment SET DEFAULT nextval('public.reaction_id_comment_seq'::regclass);


--
-- TOC entry 3349 (class 2604 OID 25263)
-- Name: reaction id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id_user SET DEFAULT nextval('public.reaction_id_user_seq'::regclass);


--
-- TOC entry 3362 (class 2604 OID 25321)
-- Name: role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role ALTER COLUMN id SET DEFAULT nextval('public.role_id_seq'::regclass);


--
-- TOC entry 3385 (class 2604 OID 25445)
-- Name: tag_task id_tag; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task ALTER COLUMN id_tag SET DEFAULT nextval('public.tag_task_id_tag_seq'::regclass);


--
-- TOC entry 3386 (class 2604 OID 25446)
-- Name: tag_task id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task ALTER COLUMN id_task SET DEFAULT nextval('public.tag_task_id_task_seq'::regclass);


--
-- TOC entry 3369 (class 2604 OID 25357)
-- Name: task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- TOC entry 3370 (class 2604 OID 25358)
-- Name: task id_column; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id_column SET DEFAULT nextval('public.task_id_column_seq'::regclass);


--
-- TOC entry 3379 (class 2604 OID 25402)
-- Name: task_template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template ALTER COLUMN id SET DEFAULT nextval('public.task_template_id_seq'::regclass);


--
-- TOC entry 3346 (class 2604 OID 25248)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3397 (class 2604 OID 25500)
-- Name: user_board_template id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template ALTER COLUMN id_user SET DEFAULT nextval('public.user_board_template_id_user_seq'::regclass);


--
-- TOC entry 3398 (class 2604 OID 25501)
-- Name: user_board_template id_template; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template ALTER COLUMN id_template SET DEFAULT nextval('public.user_board_template_id_template_seq'::regclass);


--
-- TOC entry 3395 (class 2604 OID 25489)
-- Name: user_task_template id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template ALTER COLUMN id_user SET DEFAULT nextval('public.user_task_template_id_user_seq'::regclass);


--
-- TOC entry 3396 (class 2604 OID 25490)
-- Name: user_task_template id_template; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template ALTER COLUMN id_template SET DEFAULT nextval('public.user_task_template_id_template_seq'::regclass);


--
-- TOC entry 3359 (class 2604 OID 25312)
-- Name: user_workspace id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_user SET DEFAULT nextval('public.user_workspace_id_user_seq'::regclass);


--
-- TOC entry 3360 (class 2604 OID 25313)
-- Name: user_workspace id_workspace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_workspace SET DEFAULT nextval('public.user_workspace_id_workspace_seq'::regclass);


--
-- TOC entry 3361 (class 2604 OID 25314)
-- Name: user_workspace id_role; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_role SET DEFAULT nextval('public.user_workspace_id_role_seq'::regclass);


--
-- TOC entry 3356 (class 2604 OID 25297)
-- Name: workspace id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace ALTER COLUMN id SET DEFAULT nextval('public.workspace_id_seq'::regclass);


--
-- TOC entry 3452 (class 2606 OID 25427)
-- Name: Session Session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_pkey" PRIMARY KEY (token) INCLUDE (token);


--
-- TOC entry 3454 (class 2606 OID 25429)
-- Name: Session Session_token_id_user_token1_id_user1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_token_id_user_token1_id_user1_key" UNIQUE (token, id_user) INCLUDE (token, id_user);


--
-- TOC entry 3456 (class 2606 OID 25439)
-- Name: Tag Tag_name_id_name1_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tag"
    ADD CONSTRAINT "Tag_name_id_name1_id1_key" UNIQUE (name, id) INCLUDE (name, id);


--
-- TOC entry 3458 (class 2606 OID 25437)
-- Name: Tag Tag_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Tag"
    ADD CONSTRAINT "Tag_pkey" PRIMARY KEY (id);


--
-- TOC entry 3440 (class 2606 OID 25384)
-- Name: Task_Embedding Task_Embedding_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_pkey" PRIMARY KEY (id);


--
-- TOC entry 3400 (class 2606 OID 25243)
-- Name: board board_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3402 (class 2606 OID 25241)
-- Name: board board_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_pkey PRIMARY KEY (id);


--
-- TOC entry 3448 (class 2606 OID 25419)
-- Name: board_template board_template_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template
    ADD CONSTRAINT board_template_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3450 (class 2606 OID 25417)
-- Name: board_template board_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template
    ADD CONSTRAINT board_template_pkey PRIMARY KEY (id);


--
-- TOC entry 3428 (class 2606 OID 25339)
-- Name: board_user board_user_id_board_id_user_id_board1_id_user1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_id_board_id_user_id_board1_id_user1_key UNIQUE (id_board, id_user) INCLUDE (id_board, id_user);


--
-- TOC entry 3430 (class 2606 OID 25337)
-- Name: board_user board_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_pkey PRIMARY KEY (id_board, id_user);


--
-- TOC entry 3462 (class 2606 OID 25460)
-- Name: checklist checklist_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3466 (class 2606 OID 25474)
-- Name: checklist_item checklist_item_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3468 (class 2606 OID 25472)
-- Name: checklist_item checklist_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_pkey PRIMARY KEY (id);


--
-- TOC entry 3464 (class 2606 OID 25458)
-- Name: checklist checklist_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_pkey PRIMARY KEY (id);


--
-- TOC entry 3432 (class 2606 OID 25351)
-- Name: column column_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column"
    ADD CONSTRAINT column_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3442 (class 2606 OID 25397)
-- Name: comment_embedding comment_embedding_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3410 (class 2606 OID 25281)
-- Name: comment comment_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3412 (class 2606 OID 25279)
-- Name: comment comment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_pkey PRIMARY KEY (id_user);


--
-- TOC entry 3414 (class 2606 OID 25292)
-- Name: comment_reply comment_reply_id_reply_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT comment_reply_id_reply_key UNIQUE (id_reply);


--
-- TOC entry 3416 (class 2606 OID 25290)
-- Name: comment_reply comment_reply_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT comment_reply_pkey PRIMARY KEY (id_reply) INCLUDE (id_reply);


--
-- TOC entry 3470 (class 2606 OID 25483)
-- Name: favourite_boards favourite_boards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_pkey PRIMARY KEY (id_board, id_user) INCLUDE (id_board, id_user);


--
-- TOC entry 3424 (class 2606 OID 25326)
-- Name: role pk_role; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT pk_role PRIMARY KEY (id);


--
-- TOC entry 3408 (class 2606 OID 25265)
-- Name: reaction reaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3426 (class 2606 OID 25328)
-- Name: role role_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3460 (class 2606 OID 25448)
-- Name: tag_task tag_task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_pkey PRIMARY KEY (id_tag, id_task) INCLUDE (id_tag, id_task);


--
-- TOC entry 3434 (class 2606 OID 25366)
-- Name: task task_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3436 (class 2606 OID 25364)
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- TOC entry 3444 (class 2606 OID 25408)
-- Name: task_template task_template_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template
    ADD CONSTRAINT task_template_id_key UNIQUE (id);


--
-- TOC entry 3446 (class 2606 OID 25406)
-- Name: task_template task_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template
    ADD CONSTRAINT task_template_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3438 (class 2606 OID 25371)
-- Name: task_user task_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_pkey PRIMARY KEY (id_user, id_task) INCLUDE (id_user, id_task);


--
-- TOC entry 3476 (class 2606 OID 25505)
-- Name: user_board_template user_board_template_id_user_id_template_id_user1_id_templat_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template
    ADD CONSTRAINT user_board_template_id_user_id_template_id_user1_id_templat_key UNIQUE (id_user, id_template) INCLUDE (id_user, id_template);


--
-- TOC entry 3478 (class 2606 OID 25503)
-- Name: user_board_template user_board_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template
    ADD CONSTRAINT user_board_template_pkey PRIMARY KEY (id_user, id_template);


--
-- TOC entry 3404 (class 2606 OID 25254)
-- Name: user user_email_id_email1_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_id_email1_id1_key UNIQUE (email, id) INCLUDE (email, id);


--
-- TOC entry 3406 (class 2606 OID 25252)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3472 (class 2606 OID 25494)
-- Name: user_task_template user_task_template_id_user_id_template_id_user1_id_template_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template
    ADD CONSTRAINT user_task_template_id_user_id_template_id_user1_id_template_key UNIQUE (id_user, id_template) INCLUDE (id_user, id_template);


--
-- TOC entry 3474 (class 2606 OID 25492)
-- Name: user_task_template user_task_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template
    ADD CONSTRAINT user_task_template_pkey PRIMARY KEY (id_user, id_template);


--
-- TOC entry 3422 (class 2606 OID 25316)
-- Name: user_workspace user_workspace_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_pkey PRIMARY KEY (id_user, id_workspace) INCLUDE (id_user, id_workspace);


--
-- TOC entry 3418 (class 2606 OID 25305)
-- Name: workspace workspace_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace
    ADD CONSTRAINT workspace_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3420 (class 2606 OID 25303)
-- Name: workspace workspace_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace
    ADD CONSTRAINT workspace_pkey PRIMARY KEY (id);


--
-- TOC entry 3499 (class 2606 OID 25606)
-- Name: Session Session_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_id_user_fkey" FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3495 (class 2606 OID 25586)
-- Name: Task_Embedding Task_Embedding_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_id_task_fkey" FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3496 (class 2606 OID 25591)
-- Name: Task_Embedding Task_Embedding_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_id_user_fkey" FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3479 (class 2606 OID 25506)
-- Name: board board_id_workspace_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_id_workspace_fkey FOREIGN KEY (id_workspace) REFERENCES public.workspace(id) NOT VALID;


--
-- TOC entry 3489 (class 2606 OID 25556)
-- Name: board_user board_user_id_board_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_id_board_fkey FOREIGN KEY (id_board) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3490 (class 2606 OID 25561)
-- Name: board_user board_user_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3502 (class 2606 OID 25621)
-- Name: checklist checklist_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3503 (class 2606 OID 25626)
-- Name: checklist_item checklist_item_id_checklist_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_id_checklist_fkey FOREIGN KEY (id_checklist) REFERENCES public.checklist(id) NOT VALID;


--
-- TOC entry 3491 (class 2606 OID 25566)
-- Name: column column_id_board_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column"
    ADD CONSTRAINT column_id_board_fkey FOREIGN KEY (id_board) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3497 (class 2606 OID 25596)
-- Name: comment_embedding comment_embedding_id_comment_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_id_comment_fkey FOREIGN KEY (id_comment) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3498 (class 2606 OID 25601)
-- Name: comment_embedding comment_embedding_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3482 (class 2606 OID 25526)
-- Name: comment comment_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3483 (class 2606 OID 25521)
-- Name: comment comment_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3504 (class 2606 OID 25636)
-- Name: favourite_boards favourite_boards_id_board_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_id_board_fkey FOREIGN KEY (id_board) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3505 (class 2606 OID 25631)
-- Name: favourite_boards favourite_boards_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3484 (class 2606 OID 25531)
-- Name: comment_reply original_comment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT original_comment FOREIGN KEY (id_comment) REFERENCES public.comment(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3480 (class 2606 OID 25516)
-- Name: reaction reaction_id_comment_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_id_comment_fkey FOREIGN KEY (id_comment) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3481 (class 2606 OID 25511)
-- Name: reaction reaction_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3485 (class 2606 OID 25536)
-- Name: comment_reply reply; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT reply FOREIGN KEY (id_reply) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3500 (class 2606 OID 25611)
-- Name: tag_task tag_task_id_tag_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_id_tag_fkey FOREIGN KEY (id_tag) REFERENCES public."Tag"(id) NOT VALID;


--
-- TOC entry 3501 (class 2606 OID 25616)
-- Name: tag_task tag_task_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3492 (class 2606 OID 25571)
-- Name: task task_id_column_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_id_column_fkey FOREIGN KEY (id_column) REFERENCES public."column"(id) NOT VALID;


--
-- TOC entry 3493 (class 2606 OID 25576)
-- Name: task_user task_user_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3494 (class 2606 OID 25581)
-- Name: task_user task_user_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3508 (class 2606 OID 25656)
-- Name: user_board_template user_board_template_id_template_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template
    ADD CONSTRAINT user_board_template_id_template_fkey FOREIGN KEY (id_template) REFERENCES public.board_template(id) NOT VALID;


--
-- TOC entry 3509 (class 2606 OID 25651)
-- Name: user_board_template user_board_template_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_board_template
    ADD CONSTRAINT user_board_template_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3506 (class 2606 OID 25646)
-- Name: user_task_template user_task_template_id_template_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template
    ADD CONSTRAINT user_task_template_id_template_fkey FOREIGN KEY (id_template) REFERENCES public.task_template(id) NOT VALID;


--
-- TOC entry 3507 (class 2606 OID 25641)
-- Name: user_task_template user_task_template_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_task_template
    ADD CONSTRAINT user_task_template_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3486 (class 2606 OID 25546)
-- Name: user_workspace user_workspace_id_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_role_fkey FOREIGN KEY (id_role) REFERENCES public.role(id) NOT VALID;


--
-- TOC entry 3487 (class 2606 OID 25541)
-- Name: user_workspace user_workspace_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3488 (class 2606 OID 25551)
-- Name: user_workspace user_workspace_id_workspace_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_workspace_fkey FOREIGN KEY (id_workspace) REFERENCES public.workspace(id) NOT VALID;


--
-- TOC entry 3655 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-10-26 00:41:25 MSK

--
-- PostgreSQL database dump complete
--

