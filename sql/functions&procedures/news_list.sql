create or replace function news_list(
    out p_id integer,
    out p_title varchar,
    out p_text text,
    out p_photo varchar,
    out p_event_date timestamp with time zone,
    out news_type_code varchar,
    out create_date varchar
)
returns setof record
language plpgsql
AS
$$
BEGIN
    return query SELECT
    news.id,
    news.title,
    news.text,
    news.photo,
    news.event_date,
    news_type.code,
    news.create_date
    from news, news_type
    where news.news_type_id = news_type.id;
end;
$$;
