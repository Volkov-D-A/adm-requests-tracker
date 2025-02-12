CREATE TABLE readticket (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT id_readticket PRIMARY KEY (id),
    req_id UUID NOT NULL REFERENCES reqtickets (id),
    user_id UUID NOT NULL REFERENCES requsers (id),
    lastread TIMESTAMP(0) WITHOUT TIME ZONE,
    UNIQUE (req_id, user_id)
);