create or REPLACE function favorite_product_list(
    p_user_id integer,
    out p_id integer,
    out p_name varchar,
    out p_photo varchar,
    out p_description varchar,
    out p_buyer_id integer,
    out p_sale_type_name varchar,
    out p_price numeric
)
returns setof record
language plpgsql
AS
$$
BEGIN
    return QUERY select 
    product.id,
    product.name,
    product.photo,
    product.description,
    product.user_id,
    sale_type.name,
    product.price
    from product, sale_type, favorite_product
    where favorite_product.user_id = p_user_id
    and favorite_product.product_id = product.id 
    and product.sale_type_id = sale_type.id 
    and product.is_active = true;
END;
$$;