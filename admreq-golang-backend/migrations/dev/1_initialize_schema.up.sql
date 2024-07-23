CREATE TABLE departments (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_departments PRIMARY KEY (id),
    department_name VARCHAR UNIQUE NOT NULL,
    department_active BOOLEAN NOT NULL DEFAULT TRUE,
    department_dowork BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE requsers (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_requsers PRIMARY KEY (id),
    firstname VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    department UUID NOT NULL REFERENCES departments (id),
    user_role VARCHAR NOT NULL DEFAULT 'user',
    user_login VARCHAR UNIQUE NOT NULL,
    user_pass VARCHAR NOT NULL,
    user_disabled BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE reqtickets (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_reqtickets PRIMARY KEY (id),
    user_id UUID NOT NULL REFERENCES requsers (id),
    req_text TEXT NOT NULL,
    target_department UUID NOT NULL REFERENCES departments (id),
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg'),
    finished_at TIMESTAMP(0) WITHOUT TIME ZONE,
    finish_before TIMESTAMP(0) WITHOUT TIME ZONE,
    employee_user_id UUID REFERENCES requsers (id),
    req_important BOOLEAN NOT NULL DEFAULT FALSE,
    req_finished BOOLEAN NOT NULL DEFAULT FALSE,
    req_applied BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE reqcomments (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_reqcomments PRIMARY KEY (id),
    req_id UUID NOT NULL REFERENCES reqtickets (id),
    user_id UUID NOT NULL REFERENCES requsers (id),
    comm_text TEXT NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg')
);