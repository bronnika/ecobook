create or replace function news_list(
    out p_id integer,
    out p_title varchar,
    out p_text text,
    out p_photo varchar,
    out p_event_date varchar,
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
    to_char(news.event_date, 'yyyy-MM-dd HH24:MI:SS')::varchar,
    news_type.code,
    to_char(news.create_date, 'yyyy-MM-dd HH24:MI:SS')::varchar
    from news, news_type
    where news.news_type_id = news_type.id;
end;
$$;
