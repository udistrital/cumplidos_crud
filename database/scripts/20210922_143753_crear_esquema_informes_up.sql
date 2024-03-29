CREATE SCHEMA IF NOT EXISTS informe_gestion;

SET search_path TO pg_catalog,public,informe_gestion;
CREATE TABLE IF NOT EXISTS informe_gestion.informe(
	id serial NOT NULL,
	activo boolean DEFAULT TRUE,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	periodo_informe_inicio timestamp NOT NULL,
	periodo_informe_fin timestamp NOT NULL,
	proceso varchar(100) NOT NULL,
	pago_mensual_id integer NOT NULL,
	CONSTRAINT pk_informe PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS informe_gestion.actividad_especifica(
	id serial NOT NULL,
	actividad_especifica varchar(1000) NOT NULL,
	avance numeric(5,2) NOT NULL,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	activo boolean DEFAULT true,
	informe_id integer NOT NULL,
	CONSTRAINT pk_actividad_especifica PRIMARY KEY (id)

);

CREATE TABLE IF NOT EXISTS informe_gestion.actividad_realizada(
	id serial NOT NULL,
	actividad varchar(1000) NOT NULL,
	producto_asociado varchar(250) NOT NULL,
	evidencia varchar(1000) NOT NULL,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	activo boolean DEFAULT true,
	actividad_especifica_id integer NOT NULL,
	CONSTRAINT pk_actividad_realizada PRIMARY KEY (id)

);

ALTER TABLE informe_gestion.actividad_especifica ADD CONSTRAINT fk_actividad_especifica_informe FOREIGN KEY (informe_id)
REFERENCES informe_gestion.informe (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE informe_gestion.actividad_realizada ADD CONSTRAINT fk_actividad_realizada_actividad_especifica FOREIGN KEY (actividad_especifica_id)
REFERENCES informe_gestion.actividad_especifica (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE informe_gestion.informe ADD CONSTRAINT fk_informe_pago_mensual FOREIGN KEY (pago_mensual_id)
REFERENCES cumplidos.pago_mensual (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;