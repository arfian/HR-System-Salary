BEGIN;

drop table if exists role_user;
drop table if exists auth_user;
drop table if exists employee;
drop table if exists attendance;
drop table if exists overtime;
drop table if exists reimbursement_type;
drop table if exists reimbursement;
drop table if exists setting_payroll;
drop table if exists payroll;
drop table if exists audit_log;

COMMIT;