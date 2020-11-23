INSERT INTO cumplidos.item_informe(id, nombre, descripcion, codigo_abreviacion, activo) VALUES 
	(1, 'HORAS LECTIVAS', 'Item que describe las horas lectivas del docente', 'HL', true),
	(2, 'INVESTIGACIÓN', 'Item que describe la investigación del docente', 'IN', true),
	(3, 'EXTENSIÓN',	'Item que describe las actividades de extensión que realiza el docente', 'EX',	true),
	(4, 'PUBLICACIONES', 'Item que describe las publicaciones realizadas por el docente', 'PUB',	true),
	(5, 'OTRAS ACTIVIDADES ASIGNADAS POR EL COORDINADOR', 'Item que describe actividades asignadas al docente por el coordinador', 'OTAC', true),
	(6, 'OTRAS', 'Item que describe incapacidades u otro documento que considere pertinente el docente', 'OTRAS', true),
	(7, 'CERTIFICACIÓN DE CUMPLIMIENTO',	'Item que describe el soporte de cumplido del contratista',	'CUM', true),
	(8, 'INFORME DE GESTIÓN', 'Item que describe el soporte de informe de gestión del contratista', 'INF', true),
	(9, 'SALUD Y PENSIÓN', 'Item que describe el soporte de pago de salud y pensión del contratista', 'SYP', true);

INSERT INTO cumplidos.estado_pago_mensual(id, nombre, descripcion, codigo_abreviacion, activo) VALUES
	(1, 'POR REVISAR COORDINADOR(A)', 'EL COORDINADOR DEBE REVISAR LOS SOPORTES CARGADOS POR LOS DOCENTES', 'PRC', true),
	(2, 'POR APROBAR DECANO(A)', 'ESTA POR APROBAR LOS DOCUMENTOS POR PARTE DEL SUPERVISOR DEL CONTRATO', 'PAD', true),
	(3, 'APROBADO DECANO(A)', 'EL SUPERVISOR DEL CONTRATO (DECANO) APRUEBA LOS SOPORTES DEL DOCENTE', 'AD', true),
	(4, 'RECHAZO DECANO(A)', 'EL SUPERVISOR (DECANO) RECHAZA LOS SOPORTES DEL DOCENTE POR ALGUNA INCONFORMIDAD CON LOS MISMOS', 'RD', true),
	(5, 'APROBACIÓN PAGO', 'EL ORDENADOR DEL GASTO (DECANO) APRUEBA EL PAGO DEL DOCENTE', 'AP', true),
	(6, 'RECHAZO PAGO', 'EL ORDENADOR DEL GASTO (DECANO) RECHAZA EL PAGO PARA EL DOCENTE', 'RP', true),
	(8, 'RECHAZO COORDINADOR(A)', 'EL COORDINADOR RECHAZA LOS SOPORTES DEL CUMPLIDO', 'RC', true),
	(10, 'CARGANDO DOCUMENTOS', 'EL DOCENTE O CONTRATISTA SE ENCUENTRA CARGANDO SOPORTES DEL CUMPLIDO', 'CD', true),
	(11, 'POR REVISAR SUPERVISOR', 'EL SUPERVISOR DEBE REVISAR LOS SOPORTES CARGADOS POR EL CONTRATISTA' , 'PRS', true),
	(12, 'RECHAZO SUPERVISOR', 'EL SUPERVISOR RECHAZA LOS DOCUMENTOS DEL CONTRATISTA', 'RS' , true),
	(13, 'APROBADO SUPERVISOR', 'EL SUPERVISOR APRUEBA LOS DOCUMENTOS DEL CONTRATISTA Y LA SOLICITUD ESTA A LA ESPERA DE LA APROBACION O RECHAZO DEL ORDENADOR', 'AS', true),
	(14, 'RECHAZADO ORDENADOR', 'EL ORDENADOR RECHAZA LOS DOCUMENTOS DEL CONTRATISTA', 'RO', true);