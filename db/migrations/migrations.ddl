--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 15.3 (Ubuntu 15.3-1.pgdg22.04+1)

-- Started on 2023-10-24 23:28:27 MSK

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

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 256 (class 1259 OID 23498)
-- Name: Session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Session" (
    token character varying(64) NOT NULL,
    expiration_date timestamp without time zone DEFAULT (CURRENT_TIMESTAMP + '1 day'::interval) NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public."Session" OWNER TO postgres;

--
-- TOC entry 255 (class 1259 OID 23497)
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
-- TOC entry 3626 (class 0 OID 0)
-- Dependencies: 255
-- Name: Session_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Session_id_user_seq" OWNED BY public."Session".id_user;


--
-- TOC entry 246 (class 1259 OID 23452)
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
-- TOC entry 243 (class 1259 OID 23449)
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
-- TOC entry 3627 (class 0 OID 0)
-- Dependencies: 243
-- Name: Task_Embedding_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_seq" OWNED BY public."Task_Embedding".id;


--
-- TOC entry 244 (class 1259 OID 23450)
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
-- TOC entry 3628 (class 0 OID 0)
-- Dependencies: 244
-- Name: Task_Embedding_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_task_seq" OWNED BY public."Task_Embedding".id_task;


--
-- TOC entry 245 (class 1259 OID 23451)
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
-- TOC entry 3629 (class 0 OID 0)
-- Dependencies: 245
-- Name: Task_Embedding_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Task_Embedding_id_user_seq" OWNED BY public."Task_Embedding".id_user;


--
-- TOC entry 211 (class 1259 OID 23309)
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
-- TOC entry 209 (class 1259 OID 23307)
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
-- TOC entry 3630 (class 0 OID 0)
-- Dependencies: 209
-- Name: board_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_id_seq OWNED BY public.board.id;


--
-- TOC entry 210 (class 1259 OID 23308)
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
-- TOC entry 3631 (class 0 OID 0)
-- Dependencies: 210
-- Name: board_id_workspace_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_id_workspace_seq OWNED BY public.board.id_workspace;


--
-- TOC entry 254 (class 1259 OID 23487)
-- Name: board_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_template (
    id integer NOT NULL,
    data json NOT NULL
);


ALTER TABLE public.board_template OWNER TO postgres;

--
-- TOC entry 253 (class 1259 OID 23486)
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
-- TOC entry 3632 (class 0 OID 0)
-- Dependencies: 253
-- Name: board_template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_template_id_seq OWNED BY public.board_template.id;


--
-- TOC entry 235 (class 1259 OID 23409)
-- Name: board_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.board_user (
    board_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.board_user OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 23407)
-- Name: board_user_board_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_user_board_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_user_board_id_seq OWNER TO postgres;

--
-- TOC entry 3633 (class 0 OID 0)
-- Dependencies: 233
-- Name: board_user_board_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_user_board_id_seq OWNED BY public.board_user.board_id;


--
-- TOC entry 234 (class 1259 OID 23408)
-- Name: board_user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.board_user_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.board_user_user_id_seq OWNER TO postgres;

--
-- TOC entry 3634 (class 0 OID 0)
-- Dependencies: 234
-- Name: board_user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.board_user_user_id_seq OWNED BY public.board_user.user_id;


--
-- TOC entry 264 (class 1259 OID 23528)
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
-- TOC entry 262 (class 1259 OID 23526)
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
-- TOC entry 3635 (class 0 OID 0)
-- Dependencies: 262
-- Name: checklist_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_id_seq OWNED BY public.checklist.id;


--
-- TOC entry 263 (class 1259 OID 23527)
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
-- TOC entry 3636 (class 0 OID 0)
-- Dependencies: 263
-- Name: checklist_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_id_task_seq OWNED BY public.checklist.id_task;


--
-- TOC entry 267 (class 1259 OID 23540)
-- Name: checklist_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checklist_item (
    id integer NOT NULL,
    id_checklist integer NOT NULL,
    description text,
    done boolean DEFAULT false NOT NULL,
    list_position smallint NOT NULL
);


ALTER TABLE public.checklist_item OWNER TO postgres;

--
-- TOC entry 266 (class 1259 OID 23539)
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
-- TOC entry 3637 (class 0 OID 0)
-- Dependencies: 266
-- Name: checklist_item_id_checklist_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_item_id_checklist_seq OWNED BY public.checklist_item.id_checklist;


--
-- TOC entry 265 (class 1259 OID 23538)
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
-- TOC entry 3638 (class 0 OID 0)
-- Dependencies: 265
-- Name: checklist_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checklist_item_id_seq OWNED BY public.checklist_item.id;


--
-- TOC entry 238 (class 1259 OID 23418)
-- Name: column; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."column" (
    id integer NOT NULL,
    id_board integer NOT NULL,
    name character varying(150) DEFAULT 'Столбец'::character varying NOT NULL,
    description text,
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    list_position smallint NOT NULL
);


ALTER TABLE public."column" OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 23417)
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
-- TOC entry 3639 (class 0 OID 0)
-- Dependencies: 237
-- Name: column_id_board_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.column_id_board_seq OWNED BY public."column".id_board;


--
-- TOC entry 236 (class 1259 OID 23416)
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
-- TOC entry 3640 (class 0 OID 0)
-- Dependencies: 236
-- Name: column_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.column_id_seq OWNED BY public."column".id;


--
-- TOC entry 221 (class 1259 OID 23347)
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
-- TOC entry 250 (class 1259 OID 23465)
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
-- TOC entry 249 (class 1259 OID 23464)
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
-- TOC entry 3641 (class 0 OID 0)
-- Dependencies: 249
-- Name: comment_embedding_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_comment_seq OWNED BY public.comment_embedding.id_comment;


--
-- TOC entry 247 (class 1259 OID 23462)
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
-- TOC entry 3642 (class 0 OID 0)
-- Dependencies: 247
-- Name: comment_embedding_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_seq OWNED BY public.comment_embedding.id;


--
-- TOC entry 248 (class 1259 OID 23463)
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
-- TOC entry 3643 (class 0 OID 0)
-- Dependencies: 248
-- Name: comment_embedding_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_embedding_id_user_seq OWNED BY public.comment_embedding.id_user;


--
-- TOC entry 218 (class 1259 OID 23344)
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
-- TOC entry 3644 (class 0 OID 0)
-- Dependencies: 218
-- Name: comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_seq OWNED BY public.comment.id;


--
-- TOC entry 220 (class 1259 OID 23346)
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
-- TOC entry 3645 (class 0 OID 0)
-- Dependencies: 220
-- Name: comment_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_task_seq OWNED BY public.comment.id_task;


--
-- TOC entry 219 (class 1259 OID 23345)
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
-- TOC entry 3646 (class 0 OID 0)
-- Dependencies: 219
-- Name: comment_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_id_user_seq OWNED BY public.comment.id_user;


--
-- TOC entry 224 (class 1259 OID 23362)
-- Name: comment_reply; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment_reply (
    id_reply integer NOT NULL,
    id_comment integer NOT NULL
);


ALTER TABLE public.comment_reply OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 23361)
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
-- TOC entry 3647 (class 0 OID 0)
-- Dependencies: 223
-- Name: comment_reply_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_reply_id_comment_seq OWNED BY public.comment_reply.id_comment;


--
-- TOC entry 222 (class 1259 OID 23360)
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
-- TOC entry 3648 (class 0 OID 0)
-- Dependencies: 222
-- Name: comment_reply_id_reply_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.comment_reply_id_reply_seq OWNED BY public.comment_reply.id_reply;


--
-- TOC entry 270 (class 1259 OID 23554)
-- Name: favourite_boards; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.favourite_boards (
    id_board integer NOT NULL,
    id_user integer NOT NULL
);


ALTER TABLE public.favourite_boards OWNER TO postgres;

--
-- TOC entry 268 (class 1259 OID 23552)
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
-- TOC entry 3649 (class 0 OID 0)
-- Dependencies: 268
-- Name: favourite_boards_id_board_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favourite_boards_id_board_seq OWNED BY public.favourite_boards.id_board;


--
-- TOC entry 269 (class 1259 OID 23553)
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
-- TOC entry 3650 (class 0 OID 0)
-- Dependencies: 269
-- Name: favourite_boards_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.favourite_boards_id_user_seq OWNED BY public.favourite_boards.id_user;


--
-- TOC entry 217 (class 1259 OID 23336)
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
-- TOC entry 215 (class 1259 OID 23334)
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
-- TOC entry 3651 (class 0 OID 0)
-- Dependencies: 215
-- Name: reaction_id_comment_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_comment_seq OWNED BY public.reaction.id_comment;


--
-- TOC entry 214 (class 1259 OID 23333)
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
-- TOC entry 3652 (class 0 OID 0)
-- Dependencies: 214
-- Name: reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_seq OWNED BY public.reaction.id;


--
-- TOC entry 216 (class 1259 OID 23335)
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
-- TOC entry 3653 (class 0 OID 0)
-- Dependencies: 216
-- Name: reaction_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reaction_id_user_seq OWNED BY public.reaction.id_user;


--
-- TOC entry 232 (class 1259 OID 23396)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    id integer NOT NULL,
    name character varying(100) DEFAULT 'Роль'::character varying NOT NULL,
    description text
);


ALTER TABLE public.role OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 23395)
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
-- TOC entry 3654 (class 0 OID 0)
-- Dependencies: 231
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.role_id_seq OWNED BY public.role.id;


--
-- TOC entry 258 (class 1259 OID 23508)
-- Name: tag; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tag (
    id integer NOT NULL,
    name character varying(35) NOT NULL,
    color character varying(6) DEFAULT 'FFFFFF'::character varying NOT NULL
);


ALTER TABLE public.tag OWNER TO postgres;

--
-- TOC entry 257 (class 1259 OID 23507)
-- Name: tag_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_id_seq OWNER TO postgres;

--
-- TOC entry 3655 (class 0 OID 0)
-- Dependencies: 257
-- Name: tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_id_seq OWNED BY public.tag.id;


--
-- TOC entry 261 (class 1259 OID 23519)
-- Name: tag_task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tag_task (
    id_tag integer NOT NULL,
    id_task integer NOT NULL
);


ALTER TABLE public.tag_task OWNER TO postgres;

--
-- TOC entry 259 (class 1259 OID 23517)
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
-- TOC entry 3656 (class 0 OID 0)
-- Dependencies: 259
-- Name: tag_task_id_tag_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_task_id_tag_seq OWNED BY public.tag_task.id_tag;


--
-- TOC entry 260 (class 1259 OID 23518)
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
-- TOC entry 3657 (class 0 OID 0)
-- Dependencies: 260
-- Name: tag_task_id_task_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tag_task_id_task_seq OWNED BY public.tag_task.id_task;


--
-- TOC entry 241 (class 1259 OID 23431)
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
-- TOC entry 240 (class 1259 OID 23430)
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
-- TOC entry 3658 (class 0 OID 0)
-- Dependencies: 240
-- Name: task_id_column_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_column_seq OWNED BY public.task.id_column;


--
-- TOC entry 239 (class 1259 OID 23429)
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
-- TOC entry 3659 (class 0 OID 0)
-- Dependencies: 239
-- Name: task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;


--
-- TOC entry 252 (class 1259 OID 23476)
-- Name: task_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_template (
    id integer NOT NULL,
    data json NOT NULL
);


ALTER TABLE public.task_template OWNER TO postgres;

--
-- TOC entry 251 (class 1259 OID 23475)
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
-- TOC entry 3660 (class 0 OID 0)
-- Dependencies: 251
-- Name: task_template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.task_template_id_seq OWNED BY public.task_template.id;


--
-- TOC entry 242 (class 1259 OID 23444)
-- Name: task_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.task_user (
    id_user bigint NOT NULL,
    id_task bigint NOT NULL
);


ALTER TABLE public.task_user OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 23323)
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
-- TOC entry 212 (class 1259 OID 23322)
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
-- TOC entry 3661 (class 0 OID 0)
-- Dependencies: 212
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- TOC entry 230 (class 1259 OID 23387)
-- Name: user_workspace; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_workspace (
    id_user integer NOT NULL,
    id_workspace integer NOT NULL,
    id_role integer NOT NULL
);


ALTER TABLE public.user_workspace OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 23386)
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
-- TOC entry 3662 (class 0 OID 0)
-- Dependencies: 229
-- Name: user_workspace_id_role_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_role_seq OWNED BY public.user_workspace.id_role;


--
-- TOC entry 227 (class 1259 OID 23384)
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
-- TOC entry 3663 (class 0 OID 0)
-- Dependencies: 227
-- Name: user_workspace_id_user_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_user_seq OWNED BY public.user_workspace.id_user;


--
-- TOC entry 228 (class 1259 OID 23385)
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
-- TOC entry 3664 (class 0 OID 0)
-- Dependencies: 228
-- Name: user_workspace_id_workspace_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_workspace_id_workspace_seq OWNED BY public.user_workspace.id_workspace;


--
-- TOC entry 226 (class 1259 OID 23372)
-- Name: workspace; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workspace (
    id integer NOT NULL,
    name character varying(150) DEFAULT 'Рабочее место'::character varying NOT NULL,
    thumbnail_url character varying(2048),
    date_created timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.workspace OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 23371)
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
-- TOC entry 3665 (class 0 OID 0)
-- Dependencies: 225
-- Name: workspace_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.workspace_id_seq OWNED BY public.workspace.id;


--
-- TOC entry 3371 (class 2604 OID 23502)
-- Name: Session id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session" ALTER COLUMN id_user SET DEFAULT nextval('public."Session_id_user_seq"'::regclass);


--
-- TOC entry 3362 (class 2604 OID 23455)
-- Name: Task_Embedding id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id SET DEFAULT nextval('public."Task_Embedding_id_seq"'::regclass);


--
-- TOC entry 3363 (class 2604 OID 23456)
-- Name: Task_Embedding id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id_task SET DEFAULT nextval('public."Task_Embedding_id_task_seq"'::regclass);


--
-- TOC entry 3364 (class 2604 OID 23457)
-- Name: Task_Embedding id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding" ALTER COLUMN id_user SET DEFAULT nextval('public."Task_Embedding_id_user_seq"'::regclass);


--
-- TOC entry 3330 (class 2604 OID 23312)
-- Name: board id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board ALTER COLUMN id SET DEFAULT nextval('public.board_id_seq'::regclass);


--
-- TOC entry 3331 (class 2604 OID 23313)
-- Name: board id_workspace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board ALTER COLUMN id_workspace SET DEFAULT nextval('public.board_id_workspace_seq'::regclass);


--
-- TOC entry 3369 (class 2604 OID 23490)
-- Name: board_template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template ALTER COLUMN id SET DEFAULT nextval('public.board_template_id_seq'::regclass);


--
-- TOC entry 3352 (class 2604 OID 23412)
-- Name: board_user board_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user ALTER COLUMN board_id SET DEFAULT nextval('public.board_user_board_id_seq'::regclass);


--
-- TOC entry 3353 (class 2604 OID 23413)
-- Name: board_user user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user ALTER COLUMN user_id SET DEFAULT nextval('public.board_user_user_id_seq'::regclass);


--
-- TOC entry 3376 (class 2604 OID 23531)
-- Name: checklist id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist ALTER COLUMN id SET DEFAULT nextval('public.checklist_id_seq'::regclass);


--
-- TOC entry 3377 (class 2604 OID 23532)
-- Name: checklist id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist ALTER COLUMN id_task SET DEFAULT nextval('public.checklist_id_task_seq'::regclass);


--
-- TOC entry 3379 (class 2604 OID 23543)
-- Name: checklist_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item ALTER COLUMN id SET DEFAULT nextval('public.checklist_item_id_seq'::regclass);


--
-- TOC entry 3380 (class 2604 OID 23544)
-- Name: checklist_item id_checklist; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item ALTER COLUMN id_checklist SET DEFAULT nextval('public.checklist_item_id_checklist_seq'::regclass);


--
-- TOC entry 3354 (class 2604 OID 23421)
-- Name: column id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column" ALTER COLUMN id SET DEFAULT nextval('public.column_id_seq'::regclass);


--
-- TOC entry 3355 (class 2604 OID 23422)
-- Name: column id_board; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column" ALTER COLUMN id_board SET DEFAULT nextval('public.column_id_board_seq'::regclass);


--
-- TOC entry 3338 (class 2604 OID 23350)
-- Name: comment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id SET DEFAULT nextval('public.comment_id_seq'::regclass);


--
-- TOC entry 3339 (class 2604 OID 23351)
-- Name: comment id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id_user SET DEFAULT nextval('public.comment_id_user_seq'::regclass);


--
-- TOC entry 3340 (class 2604 OID 23352)
-- Name: comment id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment ALTER COLUMN id_task SET DEFAULT nextval('public.comment_id_task_seq'::regclass);


--
-- TOC entry 3365 (class 2604 OID 23468)
-- Name: comment_embedding id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id SET DEFAULT nextval('public.comment_embedding_id_seq'::regclass);


--
-- TOC entry 3366 (class 2604 OID 23469)
-- Name: comment_embedding id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id_user SET DEFAULT nextval('public.comment_embedding_id_user_seq'::regclass);


--
-- TOC entry 3367 (class 2604 OID 23470)
-- Name: comment_embedding id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding ALTER COLUMN id_comment SET DEFAULT nextval('public.comment_embedding_id_comment_seq'::regclass);


--
-- TOC entry 3342 (class 2604 OID 23365)
-- Name: comment_reply id_reply; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply ALTER COLUMN id_reply SET DEFAULT nextval('public.comment_reply_id_reply_seq'::regclass);


--
-- TOC entry 3343 (class 2604 OID 23366)
-- Name: comment_reply id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply ALTER COLUMN id_comment SET DEFAULT nextval('public.comment_reply_id_comment_seq'::regclass);


--
-- TOC entry 3382 (class 2604 OID 23557)
-- Name: favourite_boards id_board; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards ALTER COLUMN id_board SET DEFAULT nextval('public.favourite_boards_id_board_seq'::regclass);


--
-- TOC entry 3383 (class 2604 OID 23558)
-- Name: favourite_boards id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards ALTER COLUMN id_user SET DEFAULT nextval('public.favourite_boards_id_user_seq'::regclass);


--
-- TOC entry 3335 (class 2604 OID 23339)
-- Name: reaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id SET DEFAULT nextval('public.reaction_id_seq'::regclass);


--
-- TOC entry 3336 (class 2604 OID 23340)
-- Name: reaction id_comment; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id_comment SET DEFAULT nextval('public.reaction_id_comment_seq'::regclass);


--
-- TOC entry 3337 (class 2604 OID 23341)
-- Name: reaction id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction ALTER COLUMN id_user SET DEFAULT nextval('public.reaction_id_user_seq'::regclass);


--
-- TOC entry 3350 (class 2604 OID 23399)
-- Name: role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role ALTER COLUMN id SET DEFAULT nextval('public.role_id_seq'::regclass);


--
-- TOC entry 3372 (class 2604 OID 23511)
-- Name: tag id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag ALTER COLUMN id SET DEFAULT nextval('public.tag_id_seq'::regclass);


--
-- TOC entry 3374 (class 2604 OID 23522)
-- Name: tag_task id_tag; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task ALTER COLUMN id_tag SET DEFAULT nextval('public.tag_task_id_tag_seq'::regclass);


--
-- TOC entry 3375 (class 2604 OID 23523)
-- Name: tag_task id_task; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task ALTER COLUMN id_task SET DEFAULT nextval('public.tag_task_id_task_seq'::regclass);


--
-- TOC entry 3358 (class 2604 OID 23434)
-- Name: task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);


--
-- TOC entry 3359 (class 2604 OID 23435)
-- Name: task id_column; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task ALTER COLUMN id_column SET DEFAULT nextval('public.task_id_column_seq'::regclass);


--
-- TOC entry 3368 (class 2604 OID 23479)
-- Name: task_template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template ALTER COLUMN id SET DEFAULT nextval('public.task_template_id_seq'::regclass);


--
-- TOC entry 3334 (class 2604 OID 23326)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3347 (class 2604 OID 23390)
-- Name: user_workspace id_user; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_user SET DEFAULT nextval('public.user_workspace_id_user_seq'::regclass);


--
-- TOC entry 3348 (class 2604 OID 23391)
-- Name: user_workspace id_workspace; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_workspace SET DEFAULT nextval('public.user_workspace_id_workspace_seq'::regclass);


--
-- TOC entry 3349 (class 2604 OID 23392)
-- Name: user_workspace id_role; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace ALTER COLUMN id_role SET DEFAULT nextval('public.user_workspace_id_role_seq'::regclass);


--
-- TOC entry 3344 (class 2604 OID 23375)
-- Name: workspace id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace ALTER COLUMN id SET DEFAULT nextval('public.workspace_id_seq'::regclass);


--
-- TOC entry 3435 (class 2606 OID 23504)
-- Name: Session Session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_pkey" PRIMARY KEY (token);


--
-- TOC entry 3437 (class 2606 OID 23506)
-- Name: Session Session_token_id_user_token1_id_user1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_token_id_user_token1_id_user1_key" UNIQUE (token, id_user) INCLUDE (token, id_user);


--
-- TOC entry 3423 (class 2606 OID 23461)
-- Name: Task_Embedding Task_Embedding_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_pkey" PRIMARY KEY (id);


--
-- TOC entry 3385 (class 2606 OID 23321)
-- Name: board board_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3387 (class 2606 OID 23319)
-- Name: board board_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_pkey PRIMARY KEY (id);


--
-- TOC entry 3431 (class 2606 OID 23496)
-- Name: board_template board_template_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template
    ADD CONSTRAINT board_template_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3433 (class 2606 OID 23494)
-- Name: board_template board_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_template
    ADD CONSTRAINT board_template_pkey PRIMARY KEY (id);


--
-- TOC entry 3413 (class 2606 OID 23415)
-- Name: board_user board_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_pkey PRIMARY KEY (board_id, user_id) INCLUDE (board_id, user_id);


--
-- TOC entry 3445 (class 2606 OID 23537)
-- Name: checklist checklist_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3449 (class 2606 OID 23551)
-- Name: checklist_item checklist_item_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3451 (class 2606 OID 23549)
-- Name: checklist_item checklist_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_pkey PRIMARY KEY (id);


--
-- TOC entry 3447 (class 2606 OID 23535)
-- Name: checklist checklist_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_pkey PRIMARY KEY (id);


--
-- TOC entry 3415 (class 2606 OID 23428)
-- Name: column column_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column"
    ADD CONSTRAINT column_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3425 (class 2606 OID 23474)
-- Name: comment_embedding comment_embedding_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3395 (class 2606 OID 23359)
-- Name: comment comment_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3397 (class 2606 OID 23357)
-- Name: comment comment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_pkey PRIMARY KEY (id_user);


--
-- TOC entry 3399 (class 2606 OID 23370)
-- Name: comment_reply comment_reply_id_reply_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT comment_reply_id_reply_key UNIQUE (id_reply);


--
-- TOC entry 3401 (class 2606 OID 23368)
-- Name: comment_reply comment_reply_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT comment_reply_pkey PRIMARY KEY (id_reply) INCLUDE (id_reply);


--
-- TOC entry 3453 (class 2606 OID 23560)
-- Name: favourite_boards favourite_boards_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_pkey PRIMARY KEY (id_board, id_user) INCLUDE (id_board, id_user);


--
-- TOC entry 3409 (class 2606 OID 23404)
-- Name: role pk_role; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT pk_role PRIMARY KEY (id);


--
-- TOC entry 3393 (class 2606 OID 23343)
-- Name: reaction reaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3411 (class 2606 OID 23406)
-- Name: role role_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3439 (class 2606 OID 23516)
-- Name: tag tag_name_id_name1_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_name_id_name1_id1_key UNIQUE (name, id) INCLUDE (name, id);


--
-- TOC entry 3441 (class 2606 OID 23514)
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (id);


--
-- TOC entry 3443 (class 2606 OID 23525)
-- Name: tag_task tag_task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_pkey PRIMARY KEY (id_tag, id_task) INCLUDE (id_tag, id_task);


--
-- TOC entry 3417 (class 2606 OID 23443)
-- Name: task task_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3419 (class 2606 OID 23441)
-- Name: task task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);


--
-- TOC entry 3427 (class 2606 OID 23485)
-- Name: task_template task_template_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template
    ADD CONSTRAINT task_template_id_key UNIQUE (id);


--
-- TOC entry 3429 (class 2606 OID 23483)
-- Name: task_template task_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_template
    ADD CONSTRAINT task_template_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3421 (class 2606 OID 23448)
-- Name: task_user task_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_pkey PRIMARY KEY (id_user, id_task) INCLUDE (id_user, id_task);


--
-- TOC entry 3389 (class 2606 OID 23332)
-- Name: user user_email_id_email1_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_email_id_email1_id1_key UNIQUE (email, id) INCLUDE (email, id);


--
-- TOC entry 3391 (class 2606 OID 23330)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id) INCLUDE (id);


--
-- TOC entry 3407 (class 2606 OID 23394)
-- Name: user_workspace user_workspace_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_pkey PRIMARY KEY (id_user, id_workspace) INCLUDE (id_user, id_workspace);


--
-- TOC entry 3403 (class 2606 OID 23383)
-- Name: workspace workspace_id_id1_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace
    ADD CONSTRAINT workspace_id_id1_key UNIQUE (id) INCLUDE (id);


--
-- TOC entry 3405 (class 2606 OID 23381)
-- Name: workspace workspace_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace
    ADD CONSTRAINT workspace_pkey PRIMARY KEY (id);


--
-- TOC entry 3474 (class 2606 OID 23661)
-- Name: Session Session_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_id_user_fkey" FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3470 (class 2606 OID 23641)
-- Name: Task_Embedding Task_Embedding_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_id_task_fkey" FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3471 (class 2606 OID 23646)
-- Name: Task_Embedding Task_Embedding_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Task_Embedding"
    ADD CONSTRAINT "Task_Embedding_id_user_fkey" FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3454 (class 2606 OID 23561)
-- Name: board board_id_workspace_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board
    ADD CONSTRAINT board_id_workspace_fkey FOREIGN KEY (id_workspace) REFERENCES public.workspace(id) NOT VALID;


--
-- TOC entry 3464 (class 2606 OID 23611)
-- Name: board_user board_user_board_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_board_id_fkey FOREIGN KEY (board_id) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3465 (class 2606 OID 23616)
-- Name: board_user board_user_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.board_user
    ADD CONSTRAINT board_user_user_id_fkey FOREIGN KEY (user_id) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3477 (class 2606 OID 23676)
-- Name: checklist checklist_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist
    ADD CONSTRAINT checklist_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3478 (class 2606 OID 23681)
-- Name: checklist_item checklist_item_id_checklist_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checklist_item
    ADD CONSTRAINT checklist_item_id_checklist_fkey FOREIGN KEY (id_checklist) REFERENCES public.checklist(id) NOT VALID;


--
-- TOC entry 3466 (class 2606 OID 23621)
-- Name: column column_id_board_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."column"
    ADD CONSTRAINT column_id_board_fkey FOREIGN KEY (id_board) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3472 (class 2606 OID 23651)
-- Name: comment_embedding comment_embedding_id_comment_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_id_comment_fkey FOREIGN KEY (id_comment) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3473 (class 2606 OID 23656)
-- Name: comment_embedding comment_embedding_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_embedding
    ADD CONSTRAINT comment_embedding_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3457 (class 2606 OID 23581)
-- Name: comment comment_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3458 (class 2606 OID 23576)
-- Name: comment comment_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comment_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3479 (class 2606 OID 23691)
-- Name: favourite_boards favourite_boards_id_board_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_id_board_fkey FOREIGN KEY (id_board) REFERENCES public.board(id) NOT VALID;


--
-- TOC entry 3480 (class 2606 OID 23686)
-- Name: favourite_boards favourite_boards_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.favourite_boards
    ADD CONSTRAINT favourite_boards_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3459 (class 2606 OID 23586)
-- Name: comment_reply original_comment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT original_comment FOREIGN KEY (id_comment) REFERENCES public.comment(id) ON UPDATE CASCADE ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3455 (class 2606 OID 23571)
-- Name: reaction reaction_id_comment_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_id_comment_fkey FOREIGN KEY (id_comment) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3456 (class 2606 OID 23566)
-- Name: reaction reaction_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reaction
    ADD CONSTRAINT reaction_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3460 (class 2606 OID 23591)
-- Name: comment_reply reply; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment_reply
    ADD CONSTRAINT reply FOREIGN KEY (id_reply) REFERENCES public.comment(id) NOT VALID;


--
-- TOC entry 3475 (class 2606 OID 23666)
-- Name: tag_task tag_task_id_tag_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_id_tag_fkey FOREIGN KEY (id_tag) REFERENCES public.tag(id) NOT VALID;


--
-- TOC entry 3476 (class 2606 OID 23671)
-- Name: tag_task tag_task_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tag_task
    ADD CONSTRAINT tag_task_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3467 (class 2606 OID 23626)
-- Name: task task_id_column_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_id_column_fkey FOREIGN KEY (id_column) REFERENCES public."column"(id) NOT VALID;


--
-- TOC entry 3468 (class 2606 OID 23631)
-- Name: task_user task_user_id_task_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_id_task_fkey FOREIGN KEY (id_task) REFERENCES public.task(id) NOT VALID;


--
-- TOC entry 3469 (class 2606 OID 23636)
-- Name: task_user task_user_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.task_user
    ADD CONSTRAINT task_user_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3461 (class 2606 OID 23601)
-- Name: user_workspace user_workspace_id_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_role_fkey FOREIGN KEY (id_role) REFERENCES public.role(id) NOT VALID;


--
-- TOC entry 3462 (class 2606 OID 23596)
-- Name: user_workspace user_workspace_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_user_fkey FOREIGN KEY (id_user) REFERENCES public."user"(id) NOT VALID;


--
-- TOC entry 3463 (class 2606 OID 23606)
-- Name: user_workspace user_workspace_id_workspace_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_workspace
    ADD CONSTRAINT user_workspace_id_workspace_fkey FOREIGN KEY (id_workspace) REFERENCES public.workspace(id) NOT VALID;


--
-- TOC entry 3625 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2023-10-24 23:28:27 MSK

--
-- PostgreSQL database dump complete
--

