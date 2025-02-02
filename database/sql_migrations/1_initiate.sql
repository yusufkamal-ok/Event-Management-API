

-- +migrate Up
-- +migrate StatementBegin
CREATE TYPE person_status AS ENUM ('registered', 'pending', 'canceled');
CREATE TYPE ticket_status AS ENUM ('VIP', 'Regular');

create table account(
    id SERIAL PRIMARY KEY,
    email       varchar(256) UNIQUE NOT NULL,
    pass_word    varchar(256) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create table event_ac(
    id SERIAL PRIMARY KEY,
    title       varchar(256) NOT NULL,
    description_event varchar(256) NOT NULL,
    location_event    varchar(256) NOT NULL,
    date_event  DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create table ticket(
    id SERIAL PRIMARY KEY,
    event_id  INT NOT NULL,
    price      INT NOT NULL,
    status ticket_status DEFAULT 'Regular',
    quantity    INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_event_id FOREIGN KEY (event_id) REFERENCES event_ac(id)  -- Foreign key ke tabel events
);

create table person (
    id SERIAL PRIMARY KEY,
    full_name  varchar(256),
    address_person   varchar(256),
    ticket_ids INT[] NOT NULL,
    status person_status DEFAULT 'registered',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE person_ticket (
    id SERIAL PRIMARY KEY,
    person_id INT NOT NULL,
    ticket_id INT NOT NULL,
    CONSTRAINT fk_person FOREIGN KEY (person_id) REFERENCES person(id) ON DELETE CASCADE,
    CONSTRAINT fk_ticket FOREIGN KEY (ticket_id) REFERENCES ticket(id) ON DELETE CASCADE
);


-- +migrate StatementEnd

