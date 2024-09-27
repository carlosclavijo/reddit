--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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
-- Name: comments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments (
    comment_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    user_id uuid NOT NULL,
    response_id uuid,
    comment text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.comments OWNER TO postgres;

--
-- Name: comments_vote; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comments_vote (
    comment_vote_id uuid DEFAULT gen_random_uuid() NOT NULL,
    comment_id uuid NOT NULL,
    user_id uuid NOT NULL,
    vote character varying(8) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.comments_vote OWNER TO postgres;

--
-- Name: configs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.configs (
    config_id uuid DEFAULT gen_random_uuid() NOT NULL,
    subreddit_id uuid NOT NULL,
    admin_config uuid NOT NULL,
    is_available boolean DEFAULT true NOT NULL,
    is_locked boolean DEFAULT true NOT NULL,
    text_available boolean DEFAULT true NOT NULL,
    images_available boolean DEFAULT true NOT NULL,
    video_available boolean DEFAULT true NOT NULL,
    link_available boolean DEFAULT true NOT NULL,
    poll_available boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.configs OWNER TO postgres;

--
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.images (
    image_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    title character varying(50) NOT NULL,
    url character varying(200) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.images OWNER TO postgres;

--
-- Name: links; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.links (
    link_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    link character varying(300) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.links OWNER TO postgres;

--
-- Name: option_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.option_users (
    option_users_id uuid DEFAULT gen_random_uuid() NOT NULL,
    option_id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.option_users OWNER TO postgres;

--
-- Name: options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.options (
    option_id uuid DEFAULT gen_random_uuid() NOT NULL,
    poll_id uuid NOT NULL,
    value character varying(50) NOT NULL,
    votes integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.options OWNER TO postgres;

--
-- Name: polls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.polls (
    poll_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.polls OWNER TO postgres;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    post_id uuid DEFAULT gen_random_uuid() NOT NULL,
    subreddit_id uuid NOT NULL,
    user_id uuid NOT NULL,
    description text NOT NULL,
    type character varying(20) DEFAULT 'text'::character varying NOT NULL,
    nsfw boolean DEFAULT false NOT NULL,
    brand boolean DEFAULT false NOT NULL,
    votes integer DEFAULT 0 NOT NULL,
    comments integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT comments_check CHECK ((comments >= 0)),
    CONSTRAINT type_check CHECK (((type)::text = ANY ((ARRAY['text'::character varying, 'image'::character varying, 'video'::character varying, 'link'::character varying, 'poll'::character varying])::text[])))
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- Name: posts_tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts_tags (
    post_tag_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.posts_tags OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: subreddits; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subreddits (
    subreddit_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(30) NOT NULL,
    description text NOT NULL,
    created_by uuid NOT NULL,
    icon character varying(200),
    banner character varying(200),
    privacy character varying(11) DEFAULT 'public'::character varying NOT NULL,
    is_mature boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT privacy_check CHECK (((privacy)::text = ANY ((ARRAY['public'::character varying, 'restricted'::character varying, 'private'::character varying])::text[])))
);


ALTER TABLE public.subreddits OWNER TO postgres;

--
-- Name: subreddits_topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subreddits_topics (
    subreddit_topic_id uuid DEFAULT gen_random_uuid() NOT NULL,
    subreddit_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.subreddits_topics OWNER TO postgres;

--
-- Name: subreddits_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.subreddits_users (
    subreddit_user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    subreddit_id uuid NOT NULL,
    user_id uuid NOT NULL,
    role character varying(9) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT role_check CHECK (((role)::text = ANY ((ARRAY['admin'::character varying, 'applicant'::character varying, 'member'::character varying, 'mod'::character varying, 'banned'::character varying])::text[])))
);


ALTER TABLE public.subreddits_users OWNER TO postgres;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tags (
    tag_id uuid DEFAULT gen_random_uuid() NOT NULL,
    subreddit_id uuid NOT NULL,
    admin_id uuid NOT NULL,
    name character varying(30) NOT NULL,
    color character varying(20) NOT NULL,
    is_mature boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tags OWNER TO postgres;

--
-- Name: topics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.topics (
    topic_id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    name character varying(30) NOT NULL,
    sup_topic uuid,
    adult_content boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.topics OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    username character varying(30) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    post_karma integer DEFAULT 0 NOT NULL,
    comment_karma integer DEFAULT 0 NOT NULL,
    account_available boolean DEFAULT true NOT NULL,
    profile_pic character varying(255),
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    CONSTRAINT email_check CHECK (((email)::text ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'::text)),
    CONSTRAINT username_check CHECK ((length((username)::text) >= 8))
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: videos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.videos (
    video_id uuid DEFAULT gen_random_uuid() NOT NULL,
    post_id uuid NOT NULL,
    title character varying(50) NOT NULL,
    url character varying(200) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.videos OWNER TO postgres;

--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (comment_id);


--
-- Name: comments_vote comments_vote_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments_vote
    ADD CONSTRAINT comments_vote_pkey PRIMARY KEY (comment_vote_id);


--
-- Name: configs configs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.configs
    ADD CONSTRAINT configs_pkey PRIMARY KEY (config_id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (image_id);


--
-- Name: links links_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.links
    ADD CONSTRAINT links_pkey PRIMARY KEY (link_id);


--
-- Name: option_users option_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.option_users
    ADD CONSTRAINT option_users_pkey PRIMARY KEY (option_users_id);


--
-- Name: options options_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.options
    ADD CONSTRAINT options_pkey PRIMARY KEY (option_id);


--
-- Name: polls polls_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.polls
    ADD CONSTRAINT polls_pkey PRIMARY KEY (poll_id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (post_id);


--
-- Name: posts_tags posts_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts_tags
    ADD CONSTRAINT posts_tags_pkey PRIMARY KEY (post_tag_id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: subreddits subreddits_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits
    ADD CONSTRAINT subreddits_pkey PRIMARY KEY (subreddit_id);


--
-- Name: subreddits_topics subreddits_topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_topics
    ADD CONSTRAINT subreddits_topics_pkey PRIMARY KEY (subreddit_topic_id);


--
-- Name: subreddits_users subreddits_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_users
    ADD CONSTRAINT subreddits_users_pkey PRIMARY KEY (subreddit_user_id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (tag_id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (topic_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: videos videos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.videos
    ADD CONSTRAINT videos_pkey PRIMARY KEY (video_id);


--
-- Name: configs_subreddit_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX configs_subreddit_id_idx ON public.configs USING btree (subreddit_id);


--
-- Name: links_post_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX links_post_id_idx ON public.links USING btree (post_id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: subreddits_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX subreddits_name_idx ON public.subreddits USING btree (name);


--
-- Name: subreddits_topics_subreddit_id_topic_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX subreddits_topics_subreddit_id_topic_id_idx ON public.subreddits_topics USING btree (subreddit_id, topic_id);


--
-- Name: subreddits_users_subreddit_id_user_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX subreddits_users_subreddit_id_user_id_idx ON public.subreddits_users USING btree (subreddit_id, user_id);


--
-- Name: topics_name_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX topics_name_idx ON public.topics USING btree (name);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: users_username_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_username_idx ON public.users USING btree (username);


--
-- Name: videos_post_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX videos_post_id_idx ON public.videos USING btree (post_id);


--
-- Name: comments comments_comments_comment_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_comments_comment_id_fk FOREIGN KEY (response_id) REFERENCES public.comments(comment_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: comments comments_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: comments comments_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: comments_vote comments_vote_comments_comment_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments_vote
    ADD CONSTRAINT comments_vote_comments_comment_id_fk FOREIGN KEY (comment_id) REFERENCES public.comments(comment_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: comments_vote comments_vote_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comments_vote
    ADD CONSTRAINT comments_vote_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: configs configs_subreddits_subreddit_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.configs
    ADD CONSTRAINT configs_subreddits_subreddit_id_fk FOREIGN KEY (subreddit_id) REFERENCES public.subreddits(subreddit_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: configs configs_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.configs
    ADD CONSTRAINT configs_users_user_id_fk FOREIGN KEY (admin_config) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: images images_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: links links_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.links
    ADD CONSTRAINT links_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: option_users option_users_options_option_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.option_users
    ADD CONSTRAINT option_users_options_option_id_fk FOREIGN KEY (option_id) REFERENCES public.options(option_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: option_users option_users_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.option_users
    ADD CONSTRAINT option_users_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: options options_polls_poll_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.options
    ADD CONSTRAINT options_polls_poll_id_fk FOREIGN KEY (poll_id) REFERENCES public.polls(poll_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: polls polls_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.polls
    ADD CONSTRAINT polls_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: posts posts_subreddits_subreddit_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_subreddits_subreddit_id_fk FOREIGN KEY (subreddit_id) REFERENCES public.subreddits(subreddit_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: posts_tags posts_tags_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts_tags
    ADD CONSTRAINT posts_tags_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: posts_tags posts_tags_tags_tag_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts_tags
    ADD CONSTRAINT posts_tags_tags_tag_id_fk FOREIGN KEY (tag_id) REFERENCES public.tags(tag_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: posts posts_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: subreddits_topics subreddits_topics_subreddits_subreddit_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_topics
    ADD CONSTRAINT subreddits_topics_subreddits_subreddit_id_fk FOREIGN KEY (subreddit_id) REFERENCES public.subreddits(subreddit_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: subreddits_topics subreddits_topics_topics_topic_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_topics
    ADD CONSTRAINT subreddits_topics_topics_topic_id_fk FOREIGN KEY (topic_id) REFERENCES public.topics(topic_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: subreddits_users subreddits_users_subreddits_subreddit_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_users
    ADD CONSTRAINT subreddits_users_subreddits_subreddit_id_fk FOREIGN KEY (subreddit_id) REFERENCES public.subreddits(subreddit_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: subreddits subreddits_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits
    ADD CONSTRAINT subreddits_users_user_id_fk FOREIGN KEY (created_by) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: subreddits_users subreddits_users_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.subreddits_users
    ADD CONSTRAINT subreddits_users_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tags tags_subreddits_subreddit_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_subreddits_subreddit_id_fk FOREIGN KEY (subreddit_id) REFERENCES public.subreddits(subreddit_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: tags tags_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_users_user_id_fk FOREIGN KEY (admin_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: topics topics_topics_topic_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_topics_topic_id_fk FOREIGN KEY (sup_topic) REFERENCES public.topics(topic_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: topics topics_users_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.topics
    ADD CONSTRAINT topics_users_user_id_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: videos videos_posts_post_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.videos
    ADD CONSTRAINT videos_posts_post_id_fk FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

