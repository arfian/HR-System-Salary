BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TABLE IF NOT EXISTS role_user (
--     id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     rolename VARCHAR(50) NOT NULL,
--     created_by VARCHAR(50) NOT NULL,
--     updated_by VARCHAR(50) NULL,
--     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMPTZ NULL,
--     deleted_at TIMESTAMPTZ NULL
-- ); 

CREATE TABLE IF NOT EXISTS auth_user (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    last_login TIMESTAMPTZ NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS employee (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    fullname VARCHAR(100) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    date_join DATE NULL,
    salary_amount DECIMAL(10, 2) DEFAULT 0.00,
    rolename VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS attendance (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee VARCHAR(50) NOT NULL,
    check_in TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    check_out TIMESTAMPTZ NULL,
    status VARCHAR(20) NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS overtime (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee VARCHAR(50) NOT NULL,
    overtime_hours INTEGER NOT NULL,
    overtime_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL,
    payroll VARCHAR(50) NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

-- CREATE TABLE IF NOT EXISTS reimbursement_type (
--     id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     title VARCHAR(20) NOT NULL,
--     reimbursement_max DECIMAL(10, 2) DEFAULT 0.00,
--     created_by VARCHAR(50) NOT NULL,
--     updated_by VARCHAR(50) NULL,
--     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMPTZ NULL,
--     deleted_at TIMESTAMPTZ NULL
-- );

CREATE TABLE IF NOT EXISTS reimbursement (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee VARCHAR(50) NOT NULL,
    reimbursement_date DATE NOT NULL,
    reimbursement_amount DECIMAL(10, 2) DEFAULT 0.00,
    reimbursement_type VARCHAR(50) NOT NULL,
    description TEXT NULL,
    status VARCHAR(20) NOT NULL,
    payroll VARCHAR(50) NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS setting_payroll (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    end_cutoff int NOT NULL,
    overtime_rate_hours DECIMAL(10, 2) DEFAULT 0.00,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS payroll (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    employee VARCHAR(50) NOT NULL,
    payroll_date date NOT NULL,
    count_absence INTEGER NOT NULL DEFAULT 0,
    total_attendence INTEGER NOT NULL DEFAULT 0,
    basic_salary DECIMAL(10, 2) DEFAULT 0.00,
    count_overtime INTEGER NOT NULL DEFAULT 0,
    overtime_rate_hours DECIMAL(10, 2) DEFAULT 0.00,
    total_overtime DECIMAL(10, 2) DEFAULT 0.00,
    total_deduction_absence DECIMAL(10, 2) DEFAULT 0.00,
    total_reimbursement DECIMAL(10, 2) DEFAULT 0.00,
    total_take_home_pay DECIMAL(10, 2) DEFAULT 0.00,
    status VARCHAR(20) NOT NULL,
    created_by VARCHAR(50) NOT NULL,
    updated_by VARCHAR(50) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS audit_log (
    id VARCHAR(50) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name_table VARCHAR(50) NULL,
    operation_type VARCHAR(10) NULL,
    query TEXT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO "setting_payroll" ("end_cutoff", "overtime_rate_hours", "created_by") VALUES ('25', '500000', 'admin');

INSERT INTO "auth_user" ("username","password","is_active","last_login","created_by","created_at","updated_at") VALUES ('adminuser','$2a$14$L2nr.h9Wg46uxT4PNBfYuOX5x4LMYVIjUsSgjdmfhv3jXhyUxEzVu',false,'2025-06-10 13:49:44.976','arfian5','2025-06-10 13:49:44.977','2025-06-10 13:49:44.977');

INSERT INTO "employee" ("username","fullname","rolename","gender","date_join","salary_amount","created_by","created_at","updated_at") VALUES ('adminuser','admin test','admin','MALE','2022-02-23',5000000,'arfian5','2025-06-10 13:49:45.014','2025-06-10 13:49:45.014');

COMMIT;

DO $FN$
BEGIN
  FOR counter IN 1..100 LOOP
    RAISE NOTICE 'Counter: %', counter;

    EXECUTE $$ INSERT INTO "auth_user" ("username","password","is_active","last_login","created_by","created_at","updated_at") VALUES ('employee' || $1,'$2a$14$L2nr.h9Wg46uxT4PNBfYuOX5x4LMYVIjUsSgjdmfhv3jXhyUxEzVu',false,'2025-06-10 13:49:44.976','system','2025-06-10 13:49:44.977','2025-06-10 13:49:44.977') RETURNING id $$ 
      USING counter;
  END LOOP;

  FOR counter IN 1..100 LOOP
    RAISE NOTICE 'Counter: %', counter;

    EXECUTE $$ INSERT INTO "employee" ("username","fullname","rolename","gender","date_join","salary_amount","created_by","created_at","updated_at") VALUES ('employee' || $1,'Employee Test ' || $1,'employee','MALE','2022-02-23',floor(random()*(30000000)),'system','2025-06-10 13:49:45.014','2025-06-10 13:49:45.014') RETURNING id $$ 
      USING counter;
  END LOOP;
END;
$FN$