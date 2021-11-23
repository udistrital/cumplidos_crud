INSERT INTO cumplidos.item_informe(nombre, descripcion, codigo_abreviacion, activo) VALUES 
	('HORAS LECTIVAS', 'Item que describe las horas lectivas del docente', 'HL', true),
	('INVESTIGACIÓN', 'Item que describe la investigación del docente', 'IN', true),
	('EXTENSIÓN',	'Item que describe las actividades de extensión que realiza el docente', 'EX',	true),
	('PUBLICACIONES', 'Item que describe las publicaciones realizadas por el docente', 'PUB',	true),
	('OTRAS ACTIVIDADES ASIGNADAS POR EL COORDINADOR', 'Item que describe actividades asignadas al docente por el coordinador', 'OTAC', true),
	('OTRAS', 'Item que describe incapacidades u otro documento que considere pertinente el docente', 'OTRAS', true),
	('CERTIFICACIÓN DE CUMPLIMIENTO',	'Item que describe el soporte de cumplido del contratista',	'CUM', true),
	('INFORME DE GESTIÓN', 'Item que describe el soporte de informe de gestión del contratista', 'INF', true),
	('SALUD Y PENSIÓN', 'Item que describe el soporte de pago de salud y pensión del contratista', 'SYP', true),
	('INFORME DE GESTIÓN Y CERTIFICADO DE CUMPLIMIENTO', 'Item que describe el soporte de informe de gestión y certificado de cumplimiento del contratista', 'IGYCC', true);

INSERT INTO cumplidos.estado_pago_mensual(nombre, descripcion, codigo_abreviacion, activo) VALUES
	('POR REVISAR COORDINADOR(A)', 'EL COORDINADOR DEBE REVISAR LOS SOPORTES CARGADOS POR LOS DOCENTES', 'PRC', true),
	('POR APROBAR DECANO(A)', 'ESTA POR APROBAR LOS DOCUMENTOS POR PARTE DEL SUPERVISOR DEL CONTRATO', 'PAD', true),
	('APROBADO DECANO(A)', 'EL SUPERVISOR DEL CONTRATO (DECANO) APRUEBA LOS SOPORTES DEL DOCENTE', 'AD', true),
	('RECHAZO DECANO(A)', 'EL SUPERVISOR (DECANO) RECHAZA LOS SOPORTES DEL DOCENTE POR ALGUNA INCONFORMIDAD CON LOS MISMOS', 'RD', true),
	('APROBACIÓN PAGO', 'EL ORDENADOR DEL GASTO (DECANO) APRUEBA EL PAGO DEL DOCENTE', 'AP', true),
	('RECHAZO PAGO', 'EL ORDENADOR DEL GASTO (DECANO) RECHAZA EL PAGO PARA EL DOCENTE', 'RP', true),
	('RECHAZO COORDINADOR(A)', 'EL COORDINADOR RECHAZA LOS SOPORTES DEL CUMPLIDO', 'RC', true),
	('CARGANDO DOCUMENTOS', 'EL DOCENTE O CONTRATISTA SE ENCUENTRA CARGANDO SOPORTES DEL CUMPLIDO', 'CD', true),
	('POR REVISAR SUPERVISOR', 'EL SUPERVISOR DEBE REVISAR LOS SOPORTES CARGADOS POR EL CONTRATISTA' , 'PRS', true),
	('RECHAZO SUPERVISOR', 'EL SUPERVISOR RECHAZA LOS DOCUMENTOS DEL CONTRATISTA', 'RS' , true),
	('APROBADO SUPERVISOR', 'EL SUPERVISOR APRUEBA LOS DOCUMENTOS DEL CONTRATISTA Y LA SOLICITUD ESTA A LA ESPERA DE LA APROBACION O RECHAZO DEL ORDENADOR', 'AS', true),
	('RECHAZADO ORDENADOR', 'EL ORDENADOR RECHAZA LOS DOCUMENTOS DEL CONTRATISTA', 'RO', true);