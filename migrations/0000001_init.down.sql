drop index if exists idx_address_user_id;

drop index if exists idx_profile_user_id;

drop index if exists idx_users_signin;

drop table if exists address;

drop table if exists profile;

drop table if exists users;

do $$
begin
    if exists (select 1 from pg_type where typname = 'sex') then
    drop type sex;
end if;
end $$;