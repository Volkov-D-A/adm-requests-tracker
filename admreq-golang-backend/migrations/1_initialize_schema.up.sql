CREATE TABLE requsers (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_requsers PRIMARY KEY (id),
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    department VARCHAR NOT NULL,
    user_role VARCHAR NOT NULL DEFAULT 'user',
    user_login VARCHAR UNIQUE NOT NULL,
    user_pass VARCHAR NOT NULL
);

CREATE TABLE reqtickets (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_reqtickets PRIMARY KEY (id),
    user_id UUID NOT NULL REFERENCES requsers (id),
    req_text TEXT NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg'),
    finished_at TIMESTAMP(0) WITHOUT TIME ZONE,
    employee_user_id UUID REFERENCES requsers (id),
    finished_comment TEXT,
    req_important BOOLEAN NOT NULL DEFAULT FALSE,
    req_finished BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE reqcomments (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_reqcomments PRIMARY KEY (id),
    req_id UUID NOT NULL REFERENCES reqtickets (id),
    user_id UUID NOT NULL REFERENCES requsers (id),
    comm_text TEXT NOT NULL,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT (CURRENT_TIMESTAMP(0) AT TIME ZONE 'Asia/Yekaterinburg')
);