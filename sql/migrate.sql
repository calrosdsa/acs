create table profile(
	id int IDENTITY(1,1) primary key,
	name text
);

create table empleado(
	id int IDENTITY(1,1) primary key,
	profile_id int,
        name text
);
create table horario(
	id int IDENTITY(1,1) primary key,
	profile_id int,
	start_time time,
	end_time time,
	day smallint
);
create table marcacion(
	id int IDENTITY(1,1) primary key,
	date timestamp default CURRENT_TIMESTAMP,
	type_marcacion smallint,
	empleado_id int
);

create table asistencia(
	id int IDENTITY(1,1) primary key,
	asistencia_date date default convert(date,getdate()),
    empleado_id int,
    retraso int,
    hrs_totales int,
    hrs_trabajadas int,
    hrs_trabajadas_en_horario int,
    marcaciones varchar,
    horario varchar
)

CREATE FUNCTION [dbo].[GenerateDateRange]
(@StartDate AS DATE,
 @EndDate AS   DATE,
 @Interval AS  INT
)
RETURNS @Dates TABLE(DateValue DATE)
AS
BEGIN
    DECLARE @CUR_DATE DATE
    SET @CUR_DATE = @StartDate
    WHILE @CUR_DATE <= @EndDate BEGIN
        INSERT INTO @Dates VALUES(@CUR_DATE)
        SET @CUR_DATE = DATEADD(DAY, @Interval, @CUR_DATE)
    END
    RETURN;
END;

select cast(DATEADD(DAY, value, '2024-01-01') as date) as date,
	(select STRING_AGG(convert(varchar(25), created_on, 120), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS times,
	(select STRING_AGG(CAST(type_marcacion AS VARCHAR), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS types,
	(select TOP 1 convert(varchar(25), created_on, 120) from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value -1, '2024-01-01') as date)
		order by created_on desc
	    ) AS firstM,
    (select TOP 1 convert(varchar(25), created_on, 120) from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value + 1, '2024-01-01') as date)
		order by created_on 
	    ) AS lastM
	from GENERATE_SERIES(0, 5);





select cast(DATEADD(DAY, value, '2024-01-01') as date) as date,
	(select STRING_AGG(convert(varchar(25), created_on, 120), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS times,
	(select STRING_AGG(CAST(type_marcacion AS NVARCHAR(MAX)), ',') from marcacion where
	CAST(created_on as date) = cast(DATEADD(DAY, value, '2024-01-01') as date)) AS types,
	 (select TOP 1 created_on from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value -1, '2024-01-01') as date)
		order by created_on desc
	    ) AS firstM,
    (select TOP 1 created_on from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value + 1, '2024-01-01') as date)
		order by created_on 
	    ) AS lastM
	from GENERATE_SERIES(0, 5);







select cast(DATEADD(DAY, value, '2024-01-01') as date) as date,
	(
		select STRING_AGG(convert(varchar(25), created_on, 120), ',') from marcacion where
	    created_on >= DATEADD(DAY, value, '2024-01-01') AND
	    created_on <= DATEADD(DAY, value + 1, '2024-01-01')
	) AS times,
	(
		select STRING_AGG(CAST(type_marcacion AS NVARCHAR(MAX)), ',') from marcacion where
	    CAST(created_on as date) >= cast(DATEADD(DAY, value, '2024-01-01') as date) and
	    CAST(created_on as date) <= cast(DATEADD(DAY, value + 1, '2024-01-01') as date)
	) AS types,
   (select TOP 1 created_on from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value -1, '2024-01-01') as date)
		order by created_on desc
	    ) AS firstM,
    (select TOP 1 created_on from marcacion where
	    CAST(created_on as date) = cast(DATEADD(DAY, value + 1, '2024-01-01') as date)
		order by created_on 
	    ) AS lastM,
		h.day    
	from GENERATE_SERIES(0, 5)
	inner join horario as h on h.day = DATEPART(dw,(DATEADD(DAY, value, '2024-01-01')))
	;


-- para calcular las marcaciones de usuarios que su horario sea en dos turnos diferentes se filtrar por rango de fecha y 
-- hora
    
