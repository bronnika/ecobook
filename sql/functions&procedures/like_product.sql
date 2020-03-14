create or replace function like_product(
    p_user_id integer,
    p_product_id integer
)
returns void
language plpgsql
AS
$$
BEGIN
    if not exists (select * from favorite_product where user_id = p_user_id and product_id = p_product_id) then
        insert into favorite_product(user_id, product_id)
        values(p_user_id, p_product_id);
    end if;
end;
$$;