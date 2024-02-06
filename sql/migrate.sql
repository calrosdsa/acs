create table profile(id int IDENTITY(1,1) primary key,name text);

create table empleado(id int IDENTITY(1,1) primary key,profile_id int,name text);
create table horario(id int IDENTITY(1,1) primary key,profile_id int,start_time time,end_time time,day smallint);
create table marcacion(id int IDENTITY(1,1) primary key,date DATETIME default CURRENT_TIMESTAMP,type_marcacion smallint,empleado_id int);

create table TAsistencia(
	id int IDENTITY(1,1),
	asistenciaDate date not null,
    cardholderGuid varchar(255) not null,
    retraso int,
    hrsTotales int,
    hrsTrabajadas int,
    hrsTrabajadasEnHorario int,
    marcaciones varchar(255) not null,
    horario varchar(255),
    countMarcaciones int,
    countTurnos int,
    PRIMARY KEY (id,asistenciaDate)
);





docker exec -it angry_ritchie /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P "12ab34cd56ef$"
-- sqlcmd -S localhost,1402 -U SA -P "12ab34cd56ef$"




select 
	CAST(CONVERT(VARCHAR,fecha,110) as date),count(*),
	 STRING_AGG(convert(varchar(25), fecha, 120), ','),
	 STRING_AGG(CAST(typeMarcacion AS VARCHAR), ','),
	 (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date)
		order by fecha desc
	    ) AS firstM,
		(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY,  1, fecha) as date)
		order by fecha 
	    ) AS lastM,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date)
		order by fecha  desc
	    ) AS firstT,
	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, 1, fecha) as date)
		order by fecha 
	    ) AS lastT
	from TMarcacionAsistencia group by CAST(CONVERT(VARCHAR,fecha,110) as date);



select 
	CAST(CONVERT(VARCHAR,fecha,110) as date),count(*),
	 STRING_AGG(convert(varchar(25), fecha, 120), ','),
	 STRING_AGG(CAST(typeMarcacion AS VARCHAR), ','),
	 (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date)
		order by fecha desc
	    ) AS firstM,
		(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY,  1, fecha) as date)
		order by fecha 
	    ) AS lastM,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, -1, fecha) as date)
		order by fecha  desc
	    ) AS firstT,
	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, 1, fecha) as date)
		order by fecha 
	    ) AS lastT
	from TMarcacionAsistencia where CAST(fecha as date)= '2024-02-01'
	group by CAST(CONVERT(VARCHAR,fecha,110) as date);


2024-02-01

select 
	 cast(DATEADD(DAY, value, '2024-02-01') as date) as date,
	(select STRING_AGG(convert(varchar(25), fecha, 120), ',') from TMarcacionAsistencia where
	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS times,
	(select STRING_AGG(CAST(typeMarcacion AS VARCHAR), ',') from TMarcacionAsistencia where
	CAST(fecha as date) = cast(DATEADD(DAY, value, '2024-02-01') as date)) AS types,
	(select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value -1, '2024-02-01') as date)
		order by fecha desc
	    ) AS firstM,
    (select TOP 1 convert(varchar(25), fecha, 120) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
		order by fecha 
	    ) AS lastM,
     (select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value - 1, '2024-02-01') as date)
		order by fecha  desc
	    ) AS firstT,
	(select TOP 1 CAST(typeMarcacion AS VARCHAR) from TMarcacionAsistencia where
	    CAST(fecha as date) = cast(DATEADD(DAY, value + 1, '2024-02-01') as date)
		order by fecha 
	    ) AS lastT
	from GENERATE_SERIES(0, 5);

