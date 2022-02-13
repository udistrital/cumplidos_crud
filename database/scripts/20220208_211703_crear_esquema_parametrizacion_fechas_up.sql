CREATE SCHEMA parametrizacion_carga_cumplidos;

SET search_path TO pg_catalog,public,parametrizacion_carga_cumplidos;

CREATE TABLE parametrizacion_carga_cumplidos.fechas_carga_cumplidos(
	id serial NOT NULL,
	documento_supervisor numeric(13,0) NOT NULL,
	activo boolean NOT NULL DEFAULT TRUE,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	fecha_inicio timestamp,
	fecha_fin timestamp,
	anio numeric(4,0),
	mes numeric(2,0),
	dependencia varchar(100),
	CONSTRAINT fechas_carga_cumplidos_pk PRIMARY KEY (id)

);