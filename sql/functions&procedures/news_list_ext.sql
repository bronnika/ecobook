create or replace function news_list_ext(
    p_user_id varchar,
    out p_id integer,
    out p_title varchar,
    out p_text text,
    out p_photo varchar,
    out p_event_date varchar,
    out news_type_code varchar,
    out p_increment bigint,
    out participate boolean,
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
    (select count(participant.id) 
     from participant, "user" u
     where news.id = participant.news_id 
     and u.id = participant.user_id
     and participant.news_id = news.id
     group by news.id),
    case when p_user_id != '' then
        case when exists (select * from participant where news_id = news.id and user_id = (select cast(p_user_id as integer))) THEN
            true
        ELSE
            false
        end
    ELSE
        false
    end,
    to_char(news.create_date, 'yyyy-MM-dd HH24:MI:SS')::varchar
    from news, news_type
    where news.news_type_id = news_type.id;
end;
$$;
