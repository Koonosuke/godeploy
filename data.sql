--
-- PostgreSQL database dump
--

-- Dumped from database version 15.10 (Debian 15.10-1.pgdg120+1)
-- Dumped by pg_dump version 16.2

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
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: careers; Type: TABLE; Schema: public; Owner: portgonext
--

CREATE TABLE public.careers (
    id bigint NOT NULL,
    title text,
    period text,
    content text,
    created_at timestamp with time zone
);


ALTER TABLE public.careers OWNER TO portgonext;

--
-- Name: careers_id_seq; Type: SEQUENCE; Schema: public; Owner: portgonext
--

CREATE SEQUENCE public.careers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.careers_id_seq OWNER TO portgonext;

--
-- Name: careers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: portgonext
--

ALTER SEQUENCE public.careers_id_seq OWNED BY public.careers.id;


--
-- Name: experiences; Type: TABLE; Schema: public; Owner: portgonext
--

CREATE TABLE public.experiences (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    title text NOT NULL,
    tech_stack text NOT NULL,
    icon text NOT NULL,
    content text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.experiences OWNER TO portgonext;

--
-- Name: experiences_id_seq; Type: SEQUENCE; Schema: public; Owner: portgonext
--

CREATE SEQUENCE public.experiences_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.experiences_id_seq OWNER TO portgonext;

--
-- Name: experiences_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: portgonext
--

ALTER SEQUENCE public.experiences_id_seq OWNED BY public.experiences.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: portgonext
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    user_icon text DEFAULT '/default-icon.png'::text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO portgonext;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: portgonext
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO portgonext;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: portgonext
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: careers id; Type: DEFAULT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.careers ALTER COLUMN id SET DEFAULT nextval('public.careers_id_seq'::regclass);


--
-- Name: experiences id; Type: DEFAULT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.experiences ALTER COLUMN id SET DEFAULT nextval('public.experiences_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: careers; Type: TABLE DATA; Schema: public; Owner: portgonext
--

COPY public.careers (id, title, period, content, created_at) FROM stdin;
6	株式会社ラクーンホールディングスエンジニアアルバイト 	2024-03~2024-06	売掛保証サービスのアプリ（SpringBoot）の運用保守を行った。 そこでエンジニアとしてのデバックの仕方や、開発手法（アジャイル開発）について学んだ。	2025-02-15 14:53:20.456114+00
7	株式会社エムティーアイ7daysインターン優勝	2024-08	チームリーダとして、ヘルスケアにまつわるアプリ開発をテーマに、マーケティングから、実装までを行った。高齢者の健康見守りアプリを作成し、７チームの中で優勝を経験。\n使用技術:AWS(Lambda/DynamoDB)/Vue.js/Node.js	2025-02-15 15:22:42.783148+00
8	飲食店ホームページ制作「うわさのきんぱ」	2024-11	韓国料理店のホームページ制作を行いました。\nInstagramと連携させることで、フォロワーとサイトの閲覧回数を増やすことに成功しました。\nhttps://somunnankimbap.studio.site/	2025-02-15 15:27:13.410237+00
12	大学で情報学を専攻	2022~	主にアルゴリズムやJavaを使った講義を受講\n・データベースはPostgresを使用したアプリ開発。\n・チーム開発で要件定義～サーバにデプロイするまで体験できる講義を受講\n・ゼミではIoTを専攻し、M5StackやArduinoなどを使用している。\n	2025-03-24 13:29:43.4355+00
13	ランディット株式会社エンジニアアルバイト	2024-11～現在	新機能の実装を提案から設計、実装まで行っている。（SendGridを使ったメール送信機能付き）\n事業開発部に所属しており、フロントエンドの修正、バックエンドの修正および開発を行っている。 \n使用技術：TypeScript/React/Nest.js/Firebase/GCP/github Actions/Asana/Salesforce 	2025-03-24 13:40:04.748352+00
\.


--
-- Data for Name: experiences; Type: TABLE DATA; Schema: public; Owner: portgonext
--

COPY public.experiences (id, user_id, title, tech_stack, icon, content, created_at, updated_at) FROM stdin;
3	1	IoTドア開閉通知システム202402	M5Core2/Arduino/ENVⅢ/IMU/LineNotify	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/IoT-line.png	加速度センサー、温度湿度センサ―を使用し、ドアが開閉した際の時刻と、温度湿度を欄に通知するシステム	2025-01-30 03:17:51.107956+00	2025-01-30 03:17:51.107956+00
4	1	高齢者見守りアプリ 202408	Node.js/Vue.js/AWS(Lambda/DynamoDB)	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/mtiintern.jpg	高齢者見守りアプリを開発しました。1タップで管理者（介護者）に状況をリアルタイムで伝えるアプリです。\r\nチームリーダーとして設計から実装まで行いました。7daysインターン優勝しました。	2025-01-30 03:31:42.809948+00	2025-01-30 03:31:42.809948+00
5	1	個人開発共有サイト202409	Next.js/TypeScript/React/Tailwind CSS	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/dev.png	サポーターズハッカソンにてはじめてNext.jsを使って作成したアプリです。\r\nhttps://dev-share-phi.vercel.app	2025-01-30 03:42:13.401814+00	2025-01-30 03:42:13.401814+00
6	1	SNSアプリ制作（投稿機能、いいね機能、フォロー機能、チャット機能付き）202405	Node.js/React/MongoDB	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/sns.png	リーダーとして、ユーザごとの掲示板に加えて、ソケット通信を用いてチャット機能も実装しました。	2025-01-30 03:50:17.07908+00	2025-01-30 03:50:17.07908+00
7	1	Todoアプリ（MVCモデルのアウトプット）202403	SpringBoot/Postgres/Thymeleaf	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/springtodo.png	エンジニアアルバイト先がSpringBootを使用しているfintech企業でしたので予習。\r\nMVCモデルを意識した基礎的なアプリです。	2025-01-30 03:54:15.52973+00	2025-01-30 03:54:15.52973+00
8	1	はじめてのポートフォリオ202307	HTML/CSS/JavaScript	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/firstport.png	はじめて自身のサイトを作成してデプロイをしたフォルダ構造などズタズタサイト\r\nhttps://koonosuke.github.io/Portfolio6/	2025-01-30 03:57:56.455897+00	2025-01-30 03:57:56.455897+00
9	1	ポートフォリオversion２202312	React/CSS	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/reactport.png	HTM/CSS/JavaScriptで作成したポートフォリオをReactにしました\r\nhttps://koonosuke.github.io/PortfolioReact/	2025-01-30 04:00:45.289696+00	2025-01-30 04:00:45.289696+00
10	1	ポートフォリオversion3-202410（今はaboutと経験の部分は表示させていません）	Next.js/React/TypeScript/TailwindCSS/Firebase	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/nextport.png	Next.jsではじめてVercelを使ってデプロイしてみました。\r\n　https://next-app-bg3o.vercel.app/	2025-01-30 04:09:50.311235+00	2025-01-30 04:09:50.311235+00
12	1	韓国料理店ホームページ制作	StudioStore(ノーコードツール)	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/kinn.png	韓国家庭料理店うわさのきんぱさんのホームページ制作。\r\n店長さんと定期的にコンタクトを取り、マーケティングからサイト制作すべて行いました。\r\nうわさのきんぱで検索\r\nhttps://somunnankimbap.studio.site/	2025-01-30 04:17:39.664727+00	2025-01-30 04:17:39.664727+00
14	1	MapChat	Go(Echo)/Next.js/React/TypeScript/Three.js/GoogleMaps API/Docker/AWS S3	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/go.png	はじめてGoのフレームワークEchoを使用して、バックエンドの開発ができた。\r\nまた、AWS S3を使って画像をURLに変換、GoogleMaps APIを使用してピンを置いた箇所でソケット通信の実装もできた	2025-01-30 05:07:24.937567+00	2025-01-30 05:07:24.937567+00
15	1	睡眠改善アプリ	Next.js/TypeScript/TailwindCSS/Firebase/Python/M5Core2/ENVⅢ/IMU/PaHab	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/iotkadai.png	睡眠中の動きや、温度湿度を観測し、起きた際に寝起きが良いかを評価。グラフをもとに分析。もっともよい睡眠前の行動を記録https://pacific-science-663.notion.site/172d2f6811d6800ba8a9d974c18f072d?pvs=4	2025-01-30 05:31:38.898432+00	2025-01-30 05:31:38.898432+00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: portgonext
--

COPY public.users (id, username, password, user_icon, created_at, updated_at) FROM stdin;
1	kishi	$2a$06$e/QnsqaoDvNFupoc5VgPLuc7oAg0lSr51ogzbXTWf8iNjuUHMURye	https://2025hakkasonn01.s3.ap-northeast-1.amazonaws.com/uploads/portapp.jpg	2025-01-22 15:34:31.605703+00	2025-03-28 14:14:19.192552+00
\.


--
-- Name: careers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: portgonext
--

SELECT pg_catalog.setval('public.careers_id_seq', 15, true);


--
-- Name: experiences_id_seq; Type: SEQUENCE SET; Schema: public; Owner: portgonext
--

SELECT pg_catalog.setval('public.experiences_id_seq', 15, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: portgonext
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: careers careers_pkey; Type: CONSTRAINT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.careers
    ADD CONSTRAINT careers_pkey PRIMARY KEY (id);


--
-- Name: experiences experiences_pkey; Type: CONSTRAINT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.experiences
    ADD CONSTRAINT experiences_pkey PRIMARY KEY (id);


--
-- Name: users uni_users_username; Type: CONSTRAINT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT uni_users_username UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: portgonext
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

