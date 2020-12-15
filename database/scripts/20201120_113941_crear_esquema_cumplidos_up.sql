CREATE SCHEMA IF NOT EXISTS cumplidos;
ALTER SCHEMA cumplidos OWNER TO udistrital_administrativa_app;
SET search_path TO pg_catalog,public,cumplidos;
CREATE TABLE cumplidos.pago_mensual (
	id serial NOT NULL,
	numero_contrato character varying(20) NOT NULL,
	vigencia_contrato numeric(4,0) NOT NULL,
	mes numeric(2,0) NOT NULL,
	documento_persona_id character varying(20) NOT NULL,
	estado_pago_mensual_id integer NOT NULL,
	documento_responsable_id character varying(50) NOT NULL,
	cargo_responsable character varying(70) NOT NULL,	
	ano numeric(4,0) NOT NULL,
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_pago_mensual PRIMARY KEY (id)

);
ALTER TABLE cumplidos.pago_mensual OWNER TO udistrital_administrativa_app;
CREATE TABLE cumplidos.estado_pago_mensual (
	id serial NOT NULL,
	nombre character varying(100) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	numero_orden numeric(5,2),
	activo boolean NOT NULL DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_estado_pago_mensual PRIMARY KEY (id)

);
ALTER TABLE cumplidos.estado_pago_mensual OWNER TO udistrital_administrativa_app;
CREATE TABLE cumplidos.soporte_pago_mensual (
	id serial NOT NULL,
	pago_mensual_id integer NOT NULL,
	documento integer NOT NULL,
	item_informe_tipo_contrato_id integer NOT NULL,
	aprobado boolean,
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_soporte_pago_mensual PRIMARY KEY (id)
);
ALTER TABLE cumplidos.soporte_pago_mensual OWNER TO udistrital_administrativa_app;
CREATE TABLE cumplidos.item_informe (
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	numero_orden numeric(5,2),
	activo boolean NOT NULL DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_item_informe PRIMARY KEY (id)
);
ALTER TABLE cumplidos.item_informe OWNER TO udistrital_administrativa_app;
CREATE TABLE cumplidos.item_informe_tipo_contrato (
	id serial NOT NULL,
	item_informe_id integer NOT NULL,
	tipo_contrato_id integer NOT NULL,
	activo boolean DEFAULT true,
	fecha_creacion timestamp ,
	fecha_modificacion timestamp ,
	CONSTRAINT pk_item_informe_tipo_contrato PRIMARY KEY (id),
	CONSTRAINT uq_item_informe_tipo_contrato UNIQUE (item_informe_id,tipo_contrato_id)
);
ALTER TABLE cumplidos.item_informe_tipo_contrato OWNER TO udistrital_administrativa_app;
CREATE TABLE cumplidos.cambio_estado_pago (
	id serial NOT NULL,
	estado_pago_mensual_id integer NOT NULL,
	documento_responsable_id character varying(20),
	cargo_responsable character varying(70),
	pago_mensual_id integer,
	activo boolean NOT NULL DEFAULT true,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_cambio_estado_pago PRIMARY KEY (id)
);
ALTER TABLE cumplidos.cambio_estado_pago OWNER TO udistrital_administrativa_app;
ALTER TABLE cumplidos.cambio_estado_pago ADD CONSTRAINT fk_cambio_estado_pago_pago_mensual FOREIGN KEY (pago_mensual_id)
REFERENCES cumplidos.pago_mensual (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE cumplidos.pago_mensual ADD CONSTRAINT fk_estado_pago_mensual_pago_mensual FOREIGN KEY (estado_pago_mensual_id)
REFERENCES cumplidos.estado_pago_mensual (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE cumplidos.soporte_pago_mensual ADD CONSTRAINT fk_soporte_pago_mensual_pago_mensual FOREIGN KEY (pago_mensual_id)
REFERENCES cumplidos.pago_mensual (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE cumplidos.soporte_pago_mensual ADD CONSTRAINT fk_soporte_pago_mensual_item_informe_tipo_contrato FOREIGN KEY (item_informe_tipo_contrato_id)
REFERENCES cumplidos.item_informe_tipo_contrato (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE cumplidos.item_informe_tipo_contrato ADD CONSTRAINT fk_item_informe_tipo_contrato_item_informe FOREIGN KEY (item_informe_id)
REFERENCES cumplidos.item_informe (id) MATCH SIMPLE