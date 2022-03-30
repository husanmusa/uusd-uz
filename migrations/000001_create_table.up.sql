create table if not exists companies
(
    id         serial not null primary key,
    name       varchar(16),
    cover      varchar(128),
    slogan     varchar(128),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

-- internet xizlatlar, sms, tariflar.
create table if not exists services
(
    id          serial not null primary key,
    name        varchar(64),
    description varchar(512),
    company_id  int    not null references companies (id),
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp,
    deleted_at  timestamp
);

-- oylik, kunlik va shunga o'xshash to'plamlar.
create table if not exists sets
(
    id          serial not null primary key,
    name        varchar(64),
    description varchar(512),
    service_id  int    not null references services (id),
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp,
    deleted_at  timestamp
);

-- gb, mb, sms va hk. paketlar
create table if not exists packages
(
    id          serial not null primary key,
    name        varchar(64),
    description varchar(512),
    capacity    int,
    cost        int,
    code        varchar(32),
    set_id      int    not null references sets (id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);
