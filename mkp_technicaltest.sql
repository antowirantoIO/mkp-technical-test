-- -------------------------------------------------------------
-- TablePlus 6.7.0(634)
--
-- https://tableplus.com/
--
-- Database: mkp_technicaltest
-- Generation Time: 2025-09-08 1:17:09.1600 PM
-- -------------------------------------------------------------


-- Table Definition
CREATE TABLE "public"."user_roles" (
    "user_id" varchar(36) NOT NULL,
    "role_id" varchar(36) NOT NULL,
    "created_at" int8 NOT NULL,
    PRIMARY KEY ("user_id","role_id")
);

-- Table Definition
CREATE TABLE "public"."roles" (
    "id" varchar(36) NOT NULL,
    "name" varchar(100) NOT NULL,
    "display_name" varchar(255) NOT NULL,
    "description" text,
    "is_active" bool NOT NULL DEFAULT true,
    "is_system" bool NOT NULL DEFAULT false,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."role_permissions" (
    "role_id" varchar(36) NOT NULL,
    "permission_id" varchar(36) NOT NULL,
    "created_at" int8 NOT NULL,
    PRIMARY KEY ("role_id","permission_id")
);

-- Table Definition
CREATE TABLE "public"."users" (
    "id" varchar(36) NOT NULL,
    "username" varchar(100) NOT NULL,
    "email" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "first_name" varchar(100) NOT NULL,
    "last_name" varchar(100) NOT NULL,
    "phone" varchar(20),
    "avatar" varchar(500),
    "status" varchar(50) NOT NULL DEFAULT 'active'::character varying,
    "is_active" bool NOT NULL DEFAULT true,
    "email_verified_at" int8,
    "last_login_at" int8,
    "token" varchar(500),
    "token_expires_at" int8,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    "refresh_token" varchar(500),
    "refresh_expires_at" int8,
    "is_verified" bool NOT NULL DEFAULT false,
    "password_changed_at" int8,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."schema_migrations" (
    "version" int8 NOT NULL,
    "dirty" bool NOT NULL,
    PRIMARY KEY ("version")
);

-- Table Definition
CREATE TABLE "public"."permissions" (
    "id" varchar(36) NOT NULL,
    "name" varchar(100) NOT NULL,
    "display_name" varchar(255) NOT NULL,
    "description" text,
    "resource" varchar(100) NOT NULL,
    "action" varchar(100) NOT NULL,
    "is_active" bool NOT NULL DEFAULT true,
    "is_system" bool NOT NULL DEFAULT false,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."operators" (
    "id" varchar(36) NOT NULL,
    "user_id" varchar(36) NOT NULL,
    "operator_code" varchar(50) NOT NULL,
    "company_name" varchar(255) NOT NULL,
    "license_number" varchar(100) NOT NULL,
    "contact_person" varchar(255) NOT NULL,
    "contact_phone" varchar(20) NOT NULL,
    "contact_email" varchar(255) NOT NULL,
    "address" text NOT NULL,
    "city" varchar(100) NOT NULL,
    "province" varchar(100) NOT NULL,
    "country" varchar(100) NOT NULL,
    "postal_code" varchar(20) NOT NULL,
    "website" varchar(500),
    "operator_type" varchar(50) NOT NULL,
    "status" varchar(50) NOT NULL DEFAULT 'active'::character varying,
    "established_at" int8,
    "license_expiry" int8,
    "is_active" bool NOT NULL DEFAULT true,
    "notes" text,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."harbors" (
    "id" varchar(36) NOT NULL,
    "harbor_code" varchar(20) NOT NULL,
    "harbor_name" varchar(255) NOT NULL,
    "un_locode" varchar(10) NOT NULL,
    "country" varchar(100) NOT NULL,
    "province" varchar(100) NOT NULL,
    "city" varchar(100) NOT NULL,
    "address" text NOT NULL,
    "postal_code" varchar(20) NOT NULL,
    "latitude" numeric(10,8) NOT NULL,
    "longitude" numeric(11,8) NOT NULL,
    "harbor_type" varchar(100) NOT NULL,
    "harbor_category" varchar(100) NOT NULL,
    "status" varchar(50) NOT NULL DEFAULT 'active'::character varying,
    "max_ship_length" numeric(10,2),
    "max_ship_beam" numeric(10,2),
    "max_ship_draft" numeric(10,2),
    "max_ship_dwt" numeric(12,2),
    "berth_count" int4 NOT NULL DEFAULT 0,
    "crane_count" int4 NOT NULL DEFAULT 0,
    "storage_capacity" numeric(15,2),
    "water_depth" numeric(8,2) NOT NULL,
    "tidal_range" numeric(8,2),
    "working_hours" varchar(100) NOT NULL,
    "timezone" varchar(50) NOT NULL,
    "contact_person" varchar(255) NOT NULL,
    "contact_phone" varchar(20) NOT NULL,
    "contact_email" varchar(255) NOT NULL,
    "website" varchar(500),
    "has_customs" bool NOT NULL DEFAULT false,
    "has_quarantine" bool NOT NULL DEFAULT false,
    "has_pilotage" bool NOT NULL DEFAULT false,
    "has_tug_service" bool NOT NULL DEFAULT false,
    "has_bunkering" bool NOT NULL DEFAULT false,
    "has_repair_service" bool NOT NULL DEFAULT false,
    "has_waste" bool NOT NULL DEFAULT false,
    "is_active" bool NOT NULL DEFAULT true,
    "established_at" int8,
    "notes" text,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."ships" (
    "id" varchar(36) NOT NULL,
    "operator_id" varchar(36) NOT NULL,
    "ship_name" varchar(255) NOT NULL,
    "imo_number" varchar(20) NOT NULL,
    "call_sign" varchar(20) NOT NULL,
    "mmsi" varchar(20),
    "ship_type" varchar(100) NOT NULL,
    "flag_state" varchar(100) NOT NULL,
    "port_of_registry" varchar(255),
    "build_year" int4,
    "builder" varchar(255),
    "length" numeric(10,2),
    "beam" numeric(10,2),
    "draft" numeric(10,2),
    "gross_tonnage" numeric(12,2),
    "net_tonnage" numeric(12,2),
    "deadweight_tonnage" numeric(12,2),
    "max_speed" numeric(5,2),
    "passenger_capacity" int4,
    "crew_capacity" int4,
    "classification_society" varchar(255),
    "status" varchar(50) DEFAULT 'active'::character varying,
    "is_active" bool DEFAULT true,
    "last_inspection" int8,
    "next_inspection" int8,
    "insurance_expiry" int8,
    "certificate_expiry" int8,
    "current_latitude" numeric(10,8),
    "current_longitude" numeric(11,8),
    "last_position" int8,
    "notes" text,
    "created_at" int8 NOT NULL,
    "updated_at" int8 NOT NULL,
    "deleted_at" int8,
    PRIMARY KEY ("id")
);

ALTER TABLE "public"."user_roles" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
ALTER TABLE "public"."user_roles" ADD FOREIGN KEY ("role_id") REFERENCES "public"."roles"("id") ON DELETE CASCADE;


-- Indices
CREATE INDEX idx_user_roles_user_id ON public.user_roles USING btree (user_id);
CREATE INDEX idx_user_roles_role_id ON public.user_roles USING btree (role_id);


-- Indices
CREATE UNIQUE INDEX roles_name_key ON public.roles USING btree (name);
CREATE INDEX idx_roles_name ON public.roles USING btree (name);
CREATE INDEX idx_roles_is_active ON public.roles USING btree (is_active);
CREATE INDEX idx_roles_is_system ON public.roles USING btree (is_system);
CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);
ALTER TABLE "public"."role_permissions" ADD FOREIGN KEY ("role_id") REFERENCES "public"."roles"("id") ON DELETE CASCADE;
ALTER TABLE "public"."role_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "public"."permissions"("id") ON DELETE CASCADE;


-- Indices
CREATE INDEX idx_role_permissions_role_id ON public.role_permissions USING btree (role_id);
CREATE INDEX idx_role_permissions_permission_id ON public.role_permissions USING btree (permission_id);


-- Indices
CREATE UNIQUE INDEX users_username_key ON public.users USING btree (username);
CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);
CREATE INDEX idx_users_username ON public.users USING btree (username);
CREATE INDEX idx_users_email ON public.users USING btree (email);
CREATE INDEX idx_users_status ON public.users USING btree (status);
CREATE INDEX idx_users_is_active ON public.users USING btree (is_active);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
CREATE INDEX idx_users_refresh_token ON public.users USING btree (refresh_token);
CREATE INDEX idx_users_is_verified ON public.users USING btree (is_verified);
CREATE INDEX idx_users_password_changed_at ON public.users USING btree (password_changed_at);


-- Indices
CREATE UNIQUE INDEX permissions_name_key ON public.permissions USING btree (name);
CREATE UNIQUE INDEX uk_permissions_resource_action ON public.permissions USING btree (resource, action);
CREATE INDEX idx_permissions_name ON public.permissions USING btree (name);
CREATE INDEX idx_permissions_resource ON public.permissions USING btree (resource);
CREATE INDEX idx_permissions_action ON public.permissions USING btree (action);
CREATE INDEX idx_permissions_is_active ON public.permissions USING btree (is_active);
CREATE INDEX idx_permissions_is_system ON public.permissions USING btree (is_system);
CREATE INDEX idx_permissions_deleted_at ON public.permissions USING btree (deleted_at);
ALTER TABLE "public"."operators" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;


-- Indices
CREATE UNIQUE INDEX operators_user_id_key ON public.operators USING btree (user_id);
CREATE UNIQUE INDEX operators_operator_code_key ON public.operators USING btree (operator_code);
CREATE UNIQUE INDEX operators_license_number_key ON public.operators USING btree (license_number);
CREATE INDEX idx_operators_user_id ON public.operators USING btree (user_id);
CREATE INDEX idx_operators_operator_code ON public.operators USING btree (operator_code);
CREATE INDEX idx_operators_license_number ON public.operators USING btree (license_number);
CREATE INDEX idx_operators_operator_type ON public.operators USING btree (operator_type);
CREATE INDEX idx_operators_status ON public.operators USING btree (status);
CREATE INDEX idx_operators_is_active ON public.operators USING btree (is_active);
CREATE INDEX idx_operators_deleted_at ON public.operators USING btree (deleted_at);


-- Indices
CREATE UNIQUE INDEX harbors_harbor_code_key ON public.harbors USING btree (harbor_code);
CREATE UNIQUE INDEX harbors_un_locode_key ON public.harbors USING btree (un_locode);
CREATE INDEX idx_harbors_harbor_code ON public.harbors USING btree (harbor_code);
CREATE INDEX idx_harbors_un_locode ON public.harbors USING btree (un_locode);
CREATE INDEX idx_harbors_country ON public.harbors USING btree (country);
CREATE INDEX idx_harbors_province ON public.harbors USING btree (province);
CREATE INDEX idx_harbors_city ON public.harbors USING btree (city);
CREATE INDEX idx_harbors_harbor_type ON public.harbors USING btree (harbor_type);
CREATE INDEX idx_harbors_harbor_category ON public.harbors USING btree (harbor_category);
CREATE INDEX idx_harbors_status ON public.harbors USING btree (status);
CREATE INDEX idx_harbors_is_active ON public.harbors USING btree (is_active);
CREATE INDEX idx_harbors_deleted_at ON public.harbors USING btree (deleted_at);
ALTER TABLE "public"."ships" ADD FOREIGN KEY ("operator_id") REFERENCES "public"."operators"("id") ON DELETE CASCADE;


-- Indices
CREATE UNIQUE INDEX ships_imo_number_key ON public.ships USING btree (imo_number);
CREATE UNIQUE INDEX ships_call_sign_key ON public.ships USING btree (call_sign);
CREATE UNIQUE INDEX ships_mmsi_key ON public.ships USING btree (mmsi);
CREATE INDEX idx_ships_operator_id ON public.ships USING btree (operator_id);
CREATE INDEX idx_ships_imo_number ON public.ships USING btree (imo_number);
CREATE INDEX idx_ships_call_sign ON public.ships USING btree (call_sign);
CREATE INDEX idx_ships_mmsi ON public.ships USING btree (mmsi);
CREATE INDEX idx_ships_status ON public.ships USING btree (status);
CREATE INDEX idx_ships_is_active ON public.ships USING btree (is_active);
CREATE INDEX idx_ships_deleted_at ON public.ships USING btree (deleted_at);
