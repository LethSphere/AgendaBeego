-- Insertar registros en la tabla agenda.contacto
INSERT INTO agenda.contacto (nombre, numero_documento, direccion, fecha_creacion, fecha_modificacion, activo)
VALUES
    ('Juan Pérez', '123456789', 'Calle A #123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    ('María García', '987654321', 'Avenida B #456', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    ('Carlos Rodríguez', '555111222', 'Calle C #789', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    ('Laura Martínez', '444888333', 'Avenida D #012', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    ('Pedro López', '777333111', 'Calle E #345', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true);

-- Insertar registros en la tabla agenda.telefono
INSERT INTO agenda.telefono (contacto_id, telefono, fecha_creacion, fecha_modificacion, activo)
VALUES
    (1, '123-456-7890', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (2, '987-654-3210', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (3, '555-111-2220', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (4, '444-888-3330', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (5, '777-333-1110', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true);

-- Insertar registros en la tabla agenda.correo
INSERT INTO agenda.correo (contacto_id, email, fecha_creacion, fecha_modificacion, activo)
VALUES
    (1, 'juan.perez@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (2, 'maria.garcia@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (3, 'carlos.rodriguez@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (4, 'laura.martinez@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true),
    (5, 'pedro.lopez@example.com', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, true);
