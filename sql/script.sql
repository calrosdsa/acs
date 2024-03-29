USE [master]
GO
/****** Object:  Database [Tmac_Manager1]    Script Date: 2/2/2024 15:28:54 ******/
CREATE DATABASE [Tmac_Manager1]
GO

USE [Tmac_Manager1]

GO
CREATE TABLE [dbo].[TAccessPonit](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[guid] [varchar](500) NULL,
	[nombre] [varchar](100) NULL,
	[guidDoor] [varchar](500) NULL,
	[disposed] [varchar](50) NULL,
	[online] [varchar](50) NULL,
	[idZona] [int] NOT NULL,
	[asistencia] [int] NULL,
 CONSTRAINT [PK__TAccessP__3213E83F380521E7] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TAreaVirtual]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TAreaVirtual](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](50) NOT NULL,
	[coordenadas] [varchar](500) NOT NULL,
	[eliminado] [int] NULL,
	[latitud] [varchar](50) NOT NULL,
	[longitud] [varchar](50) NOT NULL,
 CONSTRAINT [PK__TAreaVir__3213E83F1B596927] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TBitacora]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TBitacora](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[accion] [varchar](100) NOT NULL,
	[fechaHora] [datetime] NOT NULL,
	[responsableId] [int] NULL,
	[nombreResponsable] [varchar](50) NULL,
	[idImplicado] [int] NULL,
	[nombreImplicado] [varchar](50) NULL,
	[tipoImplicado] [int] NOT NULL,
	[detalle] [varchar](max) NULL,
	[tipo] [int] NULL,
 CONSTRAINT [PK__TBitacor__3213E83FD8170039] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TCardHolder]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TCardHolder](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[guid] [varchar](500) NOT NULL,
	[firtsName] [varchar](100) NULL,
	[lastName] [varchar](100) NULL,
	[celular] [varchar](100) NULL,
	[ci] [varchar](15) NULL,
	[ciExp] [varchar](5) NULL,
	[email] [varchar](100) NULL,
	[codSap] [varchar](50) NULL,
	[picture] [varchar](500) NULL,
	[descriptions] [varchar](500) NULL,
	[estado] [varchar](50) NULL,
	[canEscort] [bit] NULL,
	[sociedad] [varchar](500) NULL,
	[unidadOrganizativa] [varchar](500) NULL,
	[gerenciaArea] [varchar](500) NULL,
	[cargo] [varchar](500) NULL,
	[tipoContrato] [varchar](500) NULL,
	[empresaContratista] [varchar](500) NULL,
 CONSTRAINT [PK_TCardHolder_1] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TCiudad]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TCiudad](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[guid] [varchar](500) NULL,
	[nombre] [varchar](100) NULL,
	[eliminado] [int] NULL,
 CONSTRAINT [PK__TCiudad__3213E83FCCF0710C] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TConfigDesign]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TConfigDesign](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[colorNavMenu] [varchar](50) NOT NULL,
	[colorLeftCol] [varchar](50) NOT NULL,
	[colorLogSup] [varchar](50) NOT NULL,
	[colorLogInf] [varchar](50) NOT NULL,
	[imageLogo] [varchar](100) NOT NULL,
	[imageFondo] [varchar](100) NOT NULL,
	[imageLogin] [varchar](100) NOT NULL,
 CONSTRAINT [PK__ConfigDi__3213E83FA37EF8F9] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TControlAsistencia]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TControlAsistencia](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[fecha] [date] NULL,
	[marcaciones] [varchar](max) NULL,
	[firstMarcacionTime] [time](7) NULL,
	[lastMarcacionTime] [time](7) NULL,
	[firstMarcacionType] [time](7) NULL,
	[lastMarcacionType] [time](7) NULL,
	[cantidadHorasTotales] [time](7) NULL,
	[cantidadHorasEnHorario] [time](7) NULL,
	[cantidadHorasRetraso] [time](7) NULL,
	[cantidadHorasExtra] [time](7) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TCredential]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TCredential](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](100) NULL,
	[guid] [varchar](500) NOT NULL,
	[facilityCode] [int] NULL,
	[cardNumber] [int] NULL,
	[uniqueId] [varchar](100) NULL,
	[estado] [varchar](50) NULL,
	[guidCardHolder] [varchar](500) NULL,
 CONSTRAINT [PK_TCredential_1] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TDoor]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TDoor](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[guid] [varchar](500) NULL,
	[nombre] [varchar](100) NULL,
 CONSTRAINT [PK__TDoor__3213E83F9FC83DCA] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TEventosMarcacion]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TEventosMarcacion](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[idEvento] [int] NOT NULL,
	[idCardHolder] [int] NULL,
	[fechaMarcacion] [datetime] NULL,
	[unidadOrganizativa] [varchar](100) NULL,
	[empresaContratista] [varchar](100) NULL,
	[lector] [varchar](100) NULL,
	[nombreZona] [varchar](100) NULL,
	[estado] [varchar](10) NULL,
 CONSTRAINT [PK__TEventos__3213E83FBC183A64] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[THistorialEvento]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[THistorialEvento](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[cantidadTotal] [int] NULL,
	[cantidadSeguros] [int] NULL,
	[cantidadInseguros] [int] NULL,
	[fechaInicio] [datetime] NULL,
	[fechaFin] [datetime] NULL,
	[idCiudad] [int] NULL,
	[estado] [int] NOT NULL,
 CONSTRAINT [PK__THistori__3213E83FD0FD84EA] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[THorarioPerfil]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[THorarioPerfil](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[dia] [varchar](10) NOT NULL,
	[diaNumber] [int] NOT NULL,
	[horaEntrada] [time](7) NOT NULL,
	[horaSalida] [time](7) NOT NULL,
	[idPerfil] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TMarcacion]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TMarcacion](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[accessPointGuid] [varchar](100) NULL,
	[cardholderGuid] [varchar](100) NULL,
	[credentialGuid] [varchar](100) NULL,
	[doorGuid] [varchar](100) NULL,
	[eventType] [varchar](100) NULL,
	[fecha] [datetime] NULL,
	[idZona] [int] NULL,
 CONSTRAINT [PK__TMarcaci__3213E83F8087CC26] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TMarcacionAsistencia]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TMarcacionAsistencia](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[accessPointGuid] [varchar](100) NULL,
	[cardholderGuid] [varchar](100) NULL,
	[credentialGuid] [varchar](100) NULL,
	[doorGuid] [varchar](100) NULL,
	[eventType] [varchar](100) NULL,
	[fecha] [datetime] NULL,
	[idZona] [int] NULL,
	[typeMarcacion] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TModulo]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TModulo](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](100) NOT NULL,
	[descripcion] [varchar](100) NULL,
	[idPadre] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TOperaciones]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TOperaciones](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](100) NOT NULL,
	[idModulo] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TPerfilAsistencia]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TPerfilAsistencia](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](50) NOT NULL,
	[descripcion] [varchar](200) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TRol]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TRol](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](50) NULL,
	[descripcion] [varchar](50) NULL,
	[eliminado] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TRolOperacion]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TRolOperacion](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[idRol] [int] NULL,
	[idOperacion] [int] NULL,
	[activo] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TSecretariaGerente]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TSecretariaGerente](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[idSecretaria] [int] NULL,
	[idGerente] [int] NULL,
 CONSTRAINT [PK__TSecreta__3213E83F5BC19002] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TSeguimientoTablet]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TSeguimientoTablet](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[idTablet] [int] NOT NULL,
	[latitud] [varchar](50) NULL,
	[longitud] [varchar](50) NULL,
	[fecha] [datetime] NULL,
 CONSTRAINT [PK__TSeguimi__3213E83F0B145E0A] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TSettigns]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TSettigns](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[ipDelControlador] [varchar](100) NULL,
	[nombreDeInterfas] [varchar](100) NULL,
	[CordenadasDeLaZona] [varchar](500) NULL,
	[zonaHoraria] [varchar](10) NULL,
	[usuarioRio] [varchar](50) NULL,
	[passwordRio] [varchar](100) NULL,
	[passwordInicioApp] [varchar](100) NULL,
	[passwordSettingsApp] [varchar](100) NULL,
 CONSTRAINT [PK__TSettign__3213E83F726C339F] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TTablet]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TTablet](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[imei] [varchar](100) NOT NULL,
	[nombre] [varchar](50) NOT NULL,
	[cargaBateria] [varchar](10) NULL,
	[panico] [int] NULL,
	[latitud] [varchar](50) NULL,
	[longitud] [varchar](50) NULL,
	[gps] [int] NULL,
	[estado] [int] NULL,
	[eliminado] [int] NULL,
	[nombreDeInterfas] [varchar](100) NULL,
 CONSTRAINT [PK_TTablet] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TTipo]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TTipo](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](100) NULL,
	[descripcion] [varchar](100) NULL,
	[eliminado] [int] NULL,
 CONSTRAINT [PK__TTipo__3213E83F2435F2D6] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TTipoUsuario]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TTipoUsuario](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](50) NOT NULL,
	[descripcion] [varchar](100) NULL,
	[fechaCreate] [datetime] NULL,
	[fechaUpdate] [datetime] NULL,
	[eliminado] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TUsuario]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TUsuario](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[usuario] [varchar](50) NULL,
	[password] [varchar](500) NOT NULL,
	[nombres] [varchar](100) NOT NULL,
	[apellidos] [varchar](100) NOT NULL,
	[ci] [varchar](20) NOT NULL,
	[ciExp] [varchar](5) NOT NULL,
	[email] [varchar](100) NOT NULL,
	[celular] [varchar](20) NULL,
	[fechaCreate] [datetime] NULL,
	[fechaUpdate] [datetime] NULL,
	[idRol] [int] NULL,
	[imagen] [varchar](100) NULL,
	[timeReload] [int] NULL,
	[idCiudad] [int] NULL,
	[eliminado] [int] NULL,
	[primerInicio] [int] NULL,
	[estado] [int] NULL,
	[guidCardHolder] [varchar](500) NULL,
	[idTipoUsuario] [int] NULL,
 CONSTRAINT [PK__TUsuario__3213E83FAE6140A3] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TVisita]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TVisita](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[guid] [varchar](150) NULL,
	[nombre] [varchar](100) NOT NULL,
	[apellidos] [varchar](100) NOT NULL,
	[ci] [varchar](50) NULL,
	[extencionCi] [varchar](10) NULL,
	[tipoDocumento] [varchar](100) NULL,
	[email] [varchar](100) NULL,
	[celular] [varchar](20) NULL,
	[empresa] [varchar](50) NULL,
	[motivo] [varchar](100) NULL,
	[nroTarjeta] [int] NULL,
	[observaciones] [varchar](500) NULL,
	[fechaCreacion] [datetime] NOT NULL,
	[fechaModificacion] [datetime] NOT NULL,
	[fechaVisita] [date] NOT NULL,
	[horaInicio] [time](7) NOT NULL,
	[horaFin] [time](7) NOT NULL,
	[estado] [int] NOT NULL,
	[idVisitado] [int] NULL,
	[idVisitante] [int] NULL,
	[eliminado] [int] NOT NULL,
	[sincronizado] [int] NOT NULL,
 CONSTRAINT [PK__TVisita__3213E83F07102025] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TVisitanteVisitado]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TVisitanteVisitado](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[idVisitante] [int] NULL,
	[idVisitado] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[TZona]    Script Date: 2/2/2024 15:28:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[TZona](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[nombre] [varchar](100) NULL,
	[estado] [int] NULL,
	[tipo] [int] NULL,
	[idCiudad] [int] NULL,
	[eliminado] [int] NULL,
 CONSTRAINT [PK__TZona__3213E83FBB45575A] PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[TAccessPonit] ON 

INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5444, N'b91e9763-55f8-4975-93b2-4d2936e20d5d', N'Area1 - Captive', N'00000000-0000-0000-0000-000000000000', N'False', N'True', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5445, N'adad17cc-99d6-47a2-9523-1f7f8bcaaeab', N'Area1 - IN', N'00000000-0000-0000-0000-000000000000', N'False', N'True', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5446, N'fc646bad-9c8a-4e5d-83f1-71b3ff6d2bab', N'Area1 - OUT', N'00000000-0000-0000-0000-000000000000', N'False', N'True', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5447, N'23169c33-c503-4505-b07c-9ee358b98d0b', N'Door lock', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5448, N'1c4fcd92-0001-4d28-ac4e-e5522059ef97', N'Door lock', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5449, N'4238d06c-5139-45f1-a3d5-085f13b9a950', N'Door sensor', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5450, N'25713b35-b439-4bf6-b7bc-6f4ab720146d', N'Door sensor', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5451, N'2efd5411-a6d4-4076-ada2-b540ef573798', N'Entry sensor - In', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5452, N'db6c6e0e-82a7-44af-ac2b-edace5bb24cd', N'Entry sensor - In', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5453, N'ce1cd1ff-5379-4922-87cf-7f89da0411aa', N'Entry sensor - Out', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5454, N'8565c440-a5e9-4c16-9339-97cbab6d5ed9', N'Entry sensor - Out', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5455, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'In', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 60, 1)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5456, N'b85950e1-4191-4551-87b7-c3214c1d057a', N'In', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 60, 1)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5457, N'a440049b-021d-4974-b92c-592be3ba6bdd', N'Out', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 60, 1)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5458, N'de124e0a-e2b4-4687-ac56-cf12837a13ef', N'Out', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 60, 1)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5459, N'280b4c6b-96dc-4cff-9142-58dd9ecaf443', N'Request to exit - In', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5460, N'789f8e43-bf9f-418f-baff-ffcdd9a85746', N'Request to exit - In', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5461, N'5ef958ab-4815-40e0-b27c-09c8c9d8ec1d', N'Request to exit - Out', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'False', N'False', 0, 0)
INSERT [dbo].[TAccessPonit] ([id], [guid], [nombre], [guidDoor], [disposed], [online], [idZona], [asistencia]) VALUES (5462, N'40c13a1a-fa76-4cb3-b882-b6590bef4a7e', N'Request to exit - Out', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'False', N'False', 0, 0)
SET IDENTITY_INSERT [dbo].[TAccessPonit] OFF
GO
SET IDENTITY_INSERT [dbo].[TAreaVirtual] ON 

INSERT [dbo].[TAreaVirtual] ([id], [nombre], [coordenadas], [eliminado], [latitud], [longitud]) VALUES (6, N'YPFB Transporte', N'-17.831471,-63.241329/-17.82955,-63.238968/-17.835597,-63.235106/-17.83768,-63.237552/', 0, N'-17.831471', N'-63.241329')
INSERT [dbo].[TAreaVirtual] ([id], [nombre], [coordenadas], [eliminado], [latitud], [longitud]) VALUES (11, N'asdasdasd', N'-17.804497,-63.223811/-17.802863,-63.219691/-17.807031,-63.217889/-17.808583,-63.221665/', 1, N'-17.804497', N'-63.223811')
INSERT [dbo].[TAreaVirtual] ([id], [nombre], [coordenadas], [eliminado], [latitud], [longitud]) VALUES (12, N'PruebaA', N'-17.783058,-63.187667/-17.783058,-63.166381/-17.794827,-63.167068/-17.794827,-63.1911/', 1, N'-17.783058', N'-63.187667')
INSERT [dbo].[TAreaVirtual] ([id], [nombre], [coordenadas], [eliminado], [latitud], [longitud]) VALUES (13, N'prueba', N'-17.811839,-63.158918/-17.80857,-63.149906/-17.820582,-63.145357/-17.82238,-63.158145/', 1, N'-17.811839', N'-63.158918')
SET IDENTITY_INSERT [dbo].[TAreaVirtual] OFF
GO

SET IDENTITY_INSERT [dbo].[TCardHolder] ON 

INSERT [dbo].[TCardHolder] ([id], [guid], [firtsName], [lastName], [celular], [ci], [ciExp], [email], [codSap], [picture], [descriptions], [estado], [canEscort], [sociedad], [unidadOrganizativa], [gerenciaArea], [cargo], [tipoContrato], [empresaContratista]) VALUES (7924, N'ea22aa93-f57c-4531-8df2-60aa718684cb', N'Daniel', N'Miranda Velasco', N'61365945', N'512354654', N'SC', N'jmiranda@teclu.com', N'9000200', N'', N'', N'Active', 1, N'ANONIMA', N'GERENCIA', N'DESARROLLO', N'JEFE DE AREA', N'FIJO', N'TELCU')
INSERT [dbo].[TCardHolder] ([id], [guid], [firtsName], [lastName], [celular], [ci], [ciExp], [email], [codSap], [picture], [descriptions], [estado], [canEscort], [sociedad], [unidadOrganizativa], [gerenciaArea], [cargo], [tipoContrato], [empresaContratista]) VALUES (7925, N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'Henry', N'Marca', N'41212312', N'123123123', N'SC', N'hmarca@teclu.com', N'9000100', N'', N'', N'Active', 1, N'ANONIMA', N'GERENCIA', N'GERENCIA', N'GERENTE', N'FIJO', N'TECLU')
INSERT [dbo].[TCardHolder] ([id], [guid], [firtsName], [lastName], [celular], [ci], [ciExp], [email], [codSap], [picture], [descriptions], [estado], [canEscort], [sociedad], [unidadOrganizativa], [gerenciaArea], [cargo], [tipoContrato], [empresaContratista]) VALUES (7939, N'ffe44307-d235-4287-b7dd-0ee10b94c693', N'Aldo', N'Jimenes Ortega', N'7897895', N'85576767', N'SC', N'Jimenes@yopmail.com', NULL, N'', N'', N'Active', 0, NULL, N'Visita', NULL, NULL, NULL, NULL)
INSERT [dbo].[TCardHolder] ([id], [guid], [firtsName], [lastName], [celular], [ci], [ciExp], [email], [codSap], [picture], [descriptions], [estado], [canEscort], [sociedad], [unidadOrganizativa], [gerenciaArea], [cargo], [tipoContrato], [empresaContratista]) VALUES (7940, N'bdf425be-6871-4bd1-9dad-a61b17362069', N'Jhon', N'Flores', N'', N'', N'', N'flores@yopmail.com', NULL, N'', N'', N'Inactive', 0, NULL, N'Visita', NULL, NULL, NULL, NULL)
SET IDENTITY_INSERT [dbo].[TCardHolder] OFF
GO
SET IDENTITY_INSERT [dbo].[TCiudad] ON 

INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (1, NULL, N'Santa Cruz', 0)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (2, NULL, N'Cochabamba', 0)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (5, NULL, N'La Paz', 0)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (6, NULL, N'Tarija', 1)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (7, NULL, N'Sucre', 1)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (1007, NULL, N'Área YPFB', 1)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (1008, NULL, N'Pando', 0)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (1009, NULL, N'PNDO', 1)
INSERT [dbo].[TCiudad] ([id], [guid], [nombre], [eliminado]) VALUES (1010, NULL, N'Sucre', 0)
SET IDENTITY_INSERT [dbo].[TCiudad] OFF
GO
SET IDENTITY_INSERT [dbo].[TConfigDesign] ON 

INSERT [dbo].[TConfigDesign] ([id], [colorNavMenu], [colorLeftCol], [colorLogSup], [colorLogInf], [imageLogo], [imageFondo], [imageLogin]) VALUES (2, N'#636363', N'#424243', N'#012a5a', N'#ffffff', N'logo\logo.png', N'image/fondo', N'image/login')
SET IDENTITY_INSERT [dbo].[TConfigDesign] OFF
GO
SET IDENTITY_INSERT [dbo].[TCredential] ON 

INSERT [dbo].[TCredential] ([id], [nombre], [guid], [facilityCode], [cardNumber], [uniqueId], [estado], [guidCardHolder]) VALUES (9245, N'Alberto Velasco''s credential', N'c83cbf5d-f11b-41ba-860e-487b3cd6086a', 1, 2, N'00000000000000000000000002020004|26', N'Active', N'00000000-0000-0000-0000-000000000000')
INSERT [dbo].[TCredential] ([id], [nombre], [guid], [facilityCode], [cardNumber], [uniqueId], [estado], [guidCardHolder]) VALUES (9246, N'Daniel Miranda Velasco''s credential', N'24f108bd-1af7-48c0-957c-7bfa77d7f353', 1, 1, N'00000000000000000000000002020002|26', N'Active', N'ea22aa93-f57c-4531-8df2-60aa718684cb')
INSERT [dbo].[TCredential] ([id], [nombre], [guid], [facilityCode], [cardNumber], [uniqueId], [estado], [guidCardHolder]) VALUES (9247, N'Henry Marca''s credential', N'5a35cbc8-db8e-4cbc-8f85-7e04439ccfff', 79, 24543, N'000000000000000000000000029EBFBE|26', N'Active', N'1910a888-c931-483b-bcd4-d5a2370f96b5')
INSERT [dbo].[TCredential] ([id], [nombre], [guid], [facilityCode], [cardNumber], [uniqueId], [estado], [guidCardHolder]) VALUES (9248, N'Henry Marca''s credential', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', 213, 2586, N'00000000000000000000000003AA1434|26', N'Active', N'1910a888-c931-483b-bcd4-d5a2370f96b5')
SET IDENTITY_INSERT [dbo].[TCredential] OFF
GO
SET IDENTITY_INSERT [dbo].[TDoor] ON 

INSERT [dbo].[TDoor] ([id], [guid], [nombre]) VALUES (2195, N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'Expendedora')
INSERT [dbo].[TDoor] ([id], [guid], [nombre]) VALUES (2196, N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'Expendedora2')
SET IDENTITY_INSERT [dbo].[TDoor] OFF
GO
SET IDENTITY_INSERT [dbo].[THorarioPerfil] ON 

INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (33, N'Lunes', 1, CAST(N'08:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (34, N'Lunes', 1, CAST(N'13:00:00' AS Time), CAST(N'17:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (35, N'Martes', 2, CAST(N'08:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (36, N'Martes', 2, CAST(N'13:00:00' AS Time), CAST(N'17:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (37, N'Miercoles', 3, CAST(N'08:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (38, N'Miercoles', 3, CAST(N'13:00:00' AS Time), CAST(N'17:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (39, N'Jueves', 4, CAST(N'08:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (40, N'Jueves', 4, CAST(N'13:00:00' AS Time), CAST(N'17:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (41, N'Viernes', 5, CAST(N'08:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (42, N'Viernes', 5, CAST(N'13:00:00' AS Time), CAST(N'16:00:00' AS Time), 1)
INSERT [dbo].[THorarioPerfil] ([id], [dia], [diaNumber], [horaEntrada], [horaSalida], [idPerfil]) VALUES (43, N'Sabado', 6, CAST(N'09:00:00' AS Time), CAST(N'12:00:00' AS Time), 1)
SET IDENTITY_INSERT [dbo].[THorarioPerfil] OFF
GO
SET IDENTITY_INSERT [dbo].[TMarcacion] ON 

INSERT [dbo].[TMarcacion] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona]) VALUES (577, N'de124e0a-e2b4-4687-ac56-cf12837a13ef', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'AccessGranted', CAST(N'2024-02-01T15:05:50.530' AS DateTime), 60)
SET IDENTITY_INSERT [dbo].[TMarcacion] OFF
GO
SET IDENTITY_INSERT [dbo].[TMarcacionAsistencia] ON 

INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (1, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-01-26T21:57:09.400' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (2, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-01-26T21:58:27.443' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (3, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-01-26T21:59:01.267' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (4, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-01-26T21:59:14.613' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (5, N'07bc3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-01-26T22:05:22.717' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (6, N'07bc	3603-3ac3-43c3-abab-8fc0fba41a3c', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'fe4efbf5-b50c-453d-a863-fbe2973e00ef', N'AccessGranted', CAST(N'2024-02-01T14:35:29.997' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (7, N'b85950e1-4191-4551-87b7-c3214c1d057a', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'AccessGranted', CAST(N'2024-02-01T14:35:50.077' AS DateTime), 60, 1)
INSERT [dbo].[TMarcacionAsistencia] ([id], [accessPointGuid], [cardholderGuid], [credentialGuid], [doorGuid], [eventType], [fecha], [idZona], [typeMarcacion]) VALUES (8, N'de124e0a-e2b4-4687-ac56-cf12837a13ef', N'1910a888-c931-483b-bcd4-d5a2370f96b5', N'7d1c17d1-f10a-40de-af5b-8ae95300db21', N'e22f2cc8-a991-49d7-9494-de71be53a36c', N'AccessGranted', CAST(N'2024-02-01T15:05:50.530' AS DateTime), 60, 2)
SET IDENTITY_INSERT [dbo].[TMarcacionAsistencia] OFF
GO
SET IDENTITY_INSERT [dbo].[TModulo] ON 

INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (1, N'Usuarios', NULL, 11)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (2, N'Zonas', NULL, 5)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (3, N'Marcaciones', NULL, 5)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (4, N'Genetec', NULL, 0)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (5, N'Masteing', NULL, 0)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (6, N'Roles', NULL, 11)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (7, N'CardHolders', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (8, N'Credentials', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (9, N'Doors', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (10, N'AccessPoint', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (11, N'Administracion', NULL, 0)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (12, N'TipoZona', NULL, 5)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (13, N'Ciudad', NULL, 11)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (14, N'Configuracion', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (15, N'Historial', NULL, 5)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (16, N'Tablet', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (17, N'AreaVirtual', NULL, 4)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (18, N'Bitacora', NULL, 11)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (19, N'ConfigDiseno', NULL, 11)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (20, N'AdmVisitas', NULL, 0)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (21, N'Visitas', NULL, 20)
INSERT [dbo].[TModulo] ([id], [nombre], [descripcion], [idPadre]) VALUES (22, N'MarcacionesClean', NULL, 5)
SET IDENTITY_INSERT [dbo].[TModulo] OFF
GO
SET IDENTITY_INSERT [dbo].[TOperaciones] ON 

INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (1, N'agregar', 1)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (2, N'editar', 1)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (3, N'eliminar', 1)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (4, N'ver', 1)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (5, N'agregar', 2)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (6, N'editar', 2)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (7, N'eliminar', 2)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (8, N'ver', 2)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (9, N'ver', 3)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (10, N'ver', 4)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (11, N'ver', 5)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (12, N'agregar', 6)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (13, N'editar', 6)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (14, N'eliminar', 6)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (15, N'ver', 6)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (16, N'ver', 7)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (17, N'ver', 8)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (18, N'ver', 9)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (19, N'ver', 10)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (20, N'ver', 11)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (21, N'ver', 12)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (22, N'agregar', 12)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (23, N'editar', 12)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (24, N'eliminar', 12)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (25, N'ver', 13)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (26, N'agregar', 13)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (27, N'editar', 13)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (28, N'eliminar', 13)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (29, N'ver', 14)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (31, N'editar', 14)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (33, N'ver', 15)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (34, N'ver', 16)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (35, N'agregar', 16)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (36, N'editar', 16)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (37, N'eliminar', 16)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (38, N'ver', 17)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (39, N'agregar', 17)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (40, N'eliminar', 17)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (41, N'ver', 18)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (45, N'ver', 19)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (47, N'editar', 19)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (49, N'ver', 20)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (50, N'ver', 21)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (51, N'agregar', 21)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (52, N'editar', 21)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (53, N'eliminar', 21)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (54, N'ver', 22)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (55, N'editar', 22)
INSERT [dbo].[TOperaciones] ([id], [nombre], [idModulo]) VALUES (56, N'editar', 10)
SET IDENTITY_INSERT [dbo].[TOperaciones] OFF
GO
SET IDENTITY_INSERT [dbo].[TPerfilAsistencia] ON 

INSERT [dbo].[TPerfilAsistencia] ([id], [nombre], [descripcion]) VALUES (1, N'General', N'Horarios normales de entrada y salida ')
INSERT [dbo].[TPerfilAsistencia] ([id], [nombre], [descripcion]) VALUES (2, N'Medio dia-Mañana', N'Horarios de medio tiempo solo en la mañana')
INSERT [dbo].[TPerfilAsistencia] ([id], [nombre], [descripcion]) VALUES (3, N'Medio dia-Tarde', N'Horarios de medio tiempo solo en la tarde')
INSERT [dbo].[TPerfilAsistencia] ([id], [nombre], [descripcion]) VALUES (7, N'Turno Nocturno', N'horarios en la noche')
INSERT [dbo].[TPerfilAsistencia] ([id], [nombre], [descripcion]) VALUES (8, N'Horario Oficina', N'horarios de 08 a 12 y de 14 a 18')
SET IDENTITY_INSERT [dbo].[TPerfilAsistencia] OFF
GO
SET IDENTITY_INSERT [dbo].[TRol] ON 

INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1007, N'Super Administrador', NULL, 0)
INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1008, N'Adm Mustering', NULL, 0)
INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1012, N'Empleado', NULL, 0)
INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1016, N'SS', NULL, 0)
INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1017, N'Visitante', NULL, 0)
INSERT [dbo].[TRol] ([id], [nombre], [descripcion], [eliminado]) VALUES (1018, N'Visitado Admin', NULL, 0)
SET IDENTITY_INSERT [dbo].[TRol] OFF
GO
SET IDENTITY_INSERT [dbo].[TRolOperacion] ON 

INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1046, 1007, 1, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1047, 1007, 2, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1048, 1007, 3, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1049, 1007, 4, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1050, 1007, 5, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1051, 1007, 6, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1052, 1007, 7, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1053, 1007, 8, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1054, 1007, 9, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1055, 1007, 10, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1056, 1007, 11, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1057, 1007, 12, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1058, 1007, 13, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1059, 1007, 14, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1060, 1007, 15, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1061, 1007, 16, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1062, 1007, 17, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1063, 1007, 18, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1064, 1007, 19, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1065, 1008, 5, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1066, 1008, 6, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1067, 1008, 7, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1068, 1008, 8, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1069, 1008, 9, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1070, 1008, 11, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1071, 1008, 1, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1072, 1008, 2, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1073, 1008, 3, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1074, 1008, 4, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1075, 1008, 10, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1076, 1008, 12, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1077, 1008, 13, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1078, 1008, 14, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1079, 1008, 15, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1080, 1008, 16, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1081, 1008, 17, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1082, 1008, 18, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1083, 1008, 19, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1084, 1007, 20, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1085, 1007, 21, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1086, 1007, 22, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1087, 1007, 23, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1088, 1007, 24, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1089, 1007, 25, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1090, 1007, 26, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1091, 1007, 27, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1092, 1007, 28, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1093, 1007, 29, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1095, 1007, 31, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1097, 1007, 33, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1098, 1008, 20, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1099, 1008, 21, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1100, 1008, 22, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1101, 1008, 23, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1102, 1008, 24, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1103, 1008, 25, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1104, 1008, 26, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1105, 1008, 27, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1106, 1008, 28, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1107, 1008, 29, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1109, 1008, 31, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1111, 1008, 33, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1205, 1007, 34, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1206, 1007, 35, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1207, 1007, 36, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1208, 1007, 37, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1209, 1007, 38, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1210, 1007, 39, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1211, 1007, 40, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1212, 1012, 4, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1213, 1012, 15, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1214, 1012, 20, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1215, 1012, 25, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1216, 1012, 1, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1217, 1012, 2, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1218, 1012, 3, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1219, 1012, 5, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1220, 1012, 6, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1221, 1012, 7, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1222, 1012, 8, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1223, 1012, 9, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1224, 1012, 10, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1225, 1012, 11, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1226, 1012, 12, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1227, 1012, 13, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1228, 1012, 14, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1229, 1012, 16, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1230, 1012, 17, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1231, 1012, 18, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1232, 1012, 19, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1233, 1012, 21, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1234, 1012, 22, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1235, 1012, 23, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1236, 1012, 24, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1237, 1012, 26, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1238, 1012, 27, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1239, 1012, 28, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1240, 1012, 29, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1241, 1012, 31, 0)
GO
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1242, 1012, 33, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1243, 1012, 34, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1244, 1012, 35, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1245, 1012, 36, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1246, 1012, 37, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1247, 1012, 38, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1248, 1012, 39, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1249, 1012, 40, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1250, 1007, 41, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1251, 1007, 45, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1252, 1007, 47, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1253, 1007, 49, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1254, 1007, 50, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1255, 1007, 51, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1256, 1007, 52, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1257, 1007, 53, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1258, 1012, 50, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1259, 1012, 51, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1260, 1012, 52, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1261, 1012, 53, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1262, 1012, 41, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1263, 1012, 45, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1264, 1012, 47, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1265, 1012, 49, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1404, 1016, 1, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1405, 1016, 2, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1406, 1016, 3, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1407, 1016, 4, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1408, 1016, 5, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1409, 1016, 6, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1410, 1016, 7, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1411, 1016, 8, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1412, 1016, 9, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1413, 1016, 10, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1414, 1016, 11, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1415, 1016, 12, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1416, 1016, 13, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1417, 1016, 14, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1418, 1016, 15, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1419, 1016, 16, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1420, 1016, 17, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1421, 1016, 18, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1422, 1016, 19, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1423, 1016, 20, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1424, 1016, 21, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1425, 1016, 22, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1426, 1016, 23, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1427, 1016, 24, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1428, 1016, 25, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1429, 1016, 26, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1430, 1016, 36, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1431, 1016, 37, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1432, 1016, 38, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1433, 1016, 53, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1434, 1016, 27, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1435, 1016, 28, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1436, 1016, 29, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1437, 1016, 31, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1438, 1016, 33, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1439, 1016, 34, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1440, 1016, 35, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1441, 1016, 39, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1442, 1016, 40, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1443, 1016, 41, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1444, 1016, 45, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1445, 1016, 47, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1446, 1016, 49, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1447, 1016, 50, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1448, 1016, 51, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1449, 1016, 52, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1450, 1007, 54, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1451, 1007, 55, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1452, 1017, 50, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1453, 1017, 53, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1454, 1017, 52, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1455, 1017, 51, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1456, 1017, 4, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1457, 1017, 3, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1458, 1017, 2, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1459, 1017, 1, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1460, 1017, 8, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1461, 1017, 7, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1462, 1017, 6, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1463, 1017, 5, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1464, 1017, 21, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1465, 1017, 24, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1466, 1017, 23, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1467, 1017, 22, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1468, 1017, 34, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1469, 1017, 37, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1470, 1017, 36, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1471, 1017, 35, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1472, 1017, 15, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1473, 1017, 14, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1474, 1017, 13, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1475, 1017, 12, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1476, 1017, 11, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1477, 1017, 54, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1478, 1017, 55, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1479, 1017, 9, 0)
GO
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1480, 1017, 33, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1481, 1017, 10, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1482, 1017, 18, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1483, 1017, 17, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1484, 1017, 29, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1485, 1017, 31, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1486, 1017, 45, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1487, 1017, 47, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1488, 1017, 25, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1489, 1017, 28, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1490, 1017, 27, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1491, 1017, 26, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1492, 1017, 16, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1493, 1017, 41, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1494, 1017, 38, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1495, 1017, 40, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1496, 1017, 39, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1497, 1017, 49, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1498, 1017, 20, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1499, 1017, 19, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1500, 1018, 1, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1501, 1018, 2, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1502, 1018, 3, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1503, 1018, 4, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1504, 1018, 12, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1505, 1018, 13, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1506, 1018, 14, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1507, 1018, 15, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1508, 1018, 16, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1509, 1018, 17, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1510, 1018, 18, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1511, 1018, 19, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1512, 1018, 20, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1513, 1018, 29, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1514, 1018, 31, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1515, 1018, 41, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1516, 1018, 45, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1517, 1018, 47, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1518, 1018, 49, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1519, 1018, 50, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1520, 1018, 51, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1521, 1018, 52, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1522, 1018, 53, 1)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1523, 1018, 5, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1524, 1018, 6, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1525, 1018, 7, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1526, 1018, 8, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1527, 1018, 9, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1528, 1018, 10, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1529, 1018, 11, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1530, 1018, 21, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1531, 1018, 22, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1532, 1018, 23, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1533, 1018, 24, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1534, 1018, 25, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1535, 1018, 26, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1536, 1018, 27, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1537, 1018, 28, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1538, 1018, 33, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1539, 1018, 34, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1540, 1018, 35, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1541, 1018, 36, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1542, 1018, 37, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1543, 1018, 38, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1544, 1018, 39, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1545, 1018, 40, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1546, 1018, 54, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1547, 1018, 55, 0)
INSERT [dbo].[TRolOperacion] ([id], [idRol], [idOperacion], [activo]) VALUES (1548, 1007, 56, 1)
SET IDENTITY_INSERT [dbo].[TRolOperacion] OFF
GO
SET IDENTITY_INSERT [dbo].[TSettigns] ON 

INSERT [dbo].[TSettigns] ([id], [ipDelControlador], [nombreDeInterfas], [CordenadasDeLaZona], [zonaHoraria], [usuarioRio], [passwordRio], [passwordInicioApp], [passwordSettingsApp]) VALUES (1, N'10.2.182.228', N'GTABus', N'-17.81773,-63.18093/-17.817735,-63.180731/-17.817843,-63.180741/-17.817831,-63.180934/', N'-4', N'admin', N'$36PtR1m0', N'123', N'123')
SET IDENTITY_INSERT [dbo].[TSettigns] OFF
GO
SET IDENTITY_INSERT [dbo].[TTablet] ON 

INSERT [dbo].[TTablet] ([id], [imei], [nombre], [cargaBateria], [panico], [latitud], [longitud], [gps], [estado], [eliminado], [nombreDeInterfas]) VALUES (1016, N'869448030239520', N'TECLU PRUEBA2', NULL, NULL, NULL, NULL, NULL, NULL, 0, N'TecluAdress')
SET IDENTITY_INSERT [dbo].[TTablet] OFF
GO
SET IDENTITY_INSERT [dbo].[TTipo] ON 

INSERT [dbo].[TTipo] ([id], [nombre], [descripcion], [eliminado]) VALUES (1, N'Segura', N'punto de encuentro', 0)
INSERT [dbo].[TTipo] ([id], [nombre], [descripcion], [eliminado]) VALUES (2, N'Insegura', N'dentro de la empresa', 0)
INSERT [dbo].[TTipo] ([id], [nombre], [descripcion], [eliminado]) VALUES (3, N'FUERA DE LA EMPRESA', NULL, 0)
SET IDENTITY_INSERT [dbo].[TTipo] OFF
GO
SET IDENTITY_INSERT [dbo].[TTipoUsuario] ON 

INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (7, N'Estandar ', N'usuario normal, el cual se usa en la creacion automatica con Genetec', NULL, NULL, 0)
INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (8, N'Gerente', N'este tipo se usa para usuarios que tendran otros usuarios a cargo de su agenda de visitas', NULL, NULL, 0)
INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (9, N'Secretaria', N'se usa para personas que tienen a su cargo la agendda de visitas de otros usuarios Gerentes ', NULL, NULL, 0)
INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (10, N'Administrador de Tablets', N'usuario que tiene acceso total a la aplicacion en las tablests', NULL, NULL, 0)
INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (11, N'Guardia ', N'usuario que tiene acceso restringido a la aplicacion de las tablets', NULL, NULL, 0)
INSERT [dbo].[TTipoUsuario] ([id], [nombre], [descripcion], [fechaCreate], [fechaUpdate], [eliminado]) VALUES (12, N'Gerente Administrador de Tablets', N'se usa para usuarios gerentes con acceso total a la aplicacion de las tablets', NULL, NULL, 0)
SET IDENTITY_INSERT [dbo].[TTipoUsuario] OFF
GO
SET IDENTITY_INSERT [dbo].[TUsuario] ON 

INSERT [dbo].[TUsuario] ([id], [usuario], [password], [nombres], [apellidos], [ci], [ciExp], [email], [celular], [fechaCreate], [fechaUpdate], [idRol], [imagen], [timeReload], [idCiudad], [eliminado], [primerInicio], [estado], [guidCardHolder], [idTipoUsuario]) VALUES (3280, N'512354654', N'cd247eb9e7b43df8848d9a9f6649a3472acf8ce0aced7cb6d59c7b319b95729f', N'Daniel', N'Miranda Velasco', N'512354654', N'SC', N'jmiranda@teclu.com', N'61365945', CAST(N'2024-02-01T11:03:00.857' AS DateTime), CAST(N'2024-02-01T11:03:00.857' AS DateTime), 1012, N'', 5000, 1, 0, 0, 0, N'ea22aa93-f57c-4531-8df2-60aa718684cb', NULL)
INSERT [dbo].[TUsuario] ([id], [usuario], [password], [nombres], [apellidos], [ci], [ciExp], [email], [celular], [fechaCreate], [fechaUpdate], [idRol], [imagen], [timeReload], [idCiudad], [eliminado], [primerInicio], [estado], [guidCardHolder], [idTipoUsuario]) VALUES (3281, N'123123123', N'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', N'Henry', N'Marca', N'123123123', N'SC', N'hmarca@teclu.com', N'41212312', CAST(N'2024-02-01T11:03:01.680' AS DateTime), CAST(N'2024-02-01T11:03:01.680' AS DateTime), 1018, N'', 5000, 1, 0, 1, 0, N'1910a888-c931-483b-bcd4-d5a2370f96b5', 7)
INSERT [dbo].[TUsuario] ([id], [usuario], [password], [nombres], [apellidos], [ci], [ciExp], [email], [celular], [fechaCreate], [fechaUpdate], [idRol], [imagen], [timeReload], [idCiudad], [eliminado], [primerInicio], [estado], [guidCardHolder], [idTipoUsuario]) VALUES (3286, N'6387768', N'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', N'Erwin', N'Paz', N'6387768', N'SC', N'spaz@yopmail.com', N'76545434', CAST(N'2024-01-13T10:20:42.650' AS DateTime), CAST(N'2024-01-13T10:20:42.650' AS DateTime), 1017, N'', 100000, 1, 0, 1, 0, NULL, 7)
INSERT [dbo].[TUsuario] ([id], [usuario], [password], [nombres], [apellidos], [ci], [ciExp], [email], [celular], [fechaCreate], [fechaUpdate], [idRol], [imagen], [timeReload], [idCiudad], [eliminado], [primerInicio], [estado], [guidCardHolder], [idTipoUsuario]) VALUES (3287, N'8238854', N'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', N'Admin', N'Admin', N'8238854', N'SC', N'marca@yopmail.com', N'61368939', CAST(N'2024-01-16T15:41:35.943' AS DateTime), CAST(N'2024-01-16T16:04:59.990' AS DateTime), 1007, N'', 600000000, 1, 0, 1, 0, NULL, 7)
SET IDENTITY_INSERT [dbo].[TUsuario] OFF
GO
SET IDENTITY_INSERT [dbo].[TVisita] ON 

INSERT [dbo].[TVisita] ([id], [guid], [nombre], [apellidos], [ci], [extencionCi], [tipoDocumento], [email], [celular], [empresa], [motivo], [nroTarjeta], [observaciones], [fechaCreacion], [fechaModificacion], [fechaVisita], [horaInicio], [horaFin], [estado], [idVisitado], [idVisitante], [eliminado], [sincronizado]) VALUES (2600, N'ffe44307-d235-4287-b7dd-0ee10b94c693', N'Aldo', N'Jimenes Ortega', N'85576767', N'SC', N'CI', N'Jimenes@yopmail.com', N'7897895', N'TECLU', NULL, 0, N'NA', CAST(N'2024-01-17T10:22:18.707' AS DateTime), CAST(N'2024-01-17T10:23:11.240' AS DateTime), CAST(N'2024-01-17' AS Date), CAST(N'05:16:19.6170000' AS Time), CAST(N'09:18:08' AS Time), 3, 3281, NULL, 0, 1)
INSERT [dbo].[TVisita] ([id], [guid], [nombre], [apellidos], [ci], [extencionCi], [tipoDocumento], [email], [celular], [empresa], [motivo], [nroTarjeta], [observaciones], [fechaCreacion], [fechaModificacion], [fechaVisita], [horaInicio], [horaFin], [estado], [idVisitado], [idVisitante], [eliminado], [sincronizado]) VALUES (2601, N'bdf425be-6871-4bd1-9dad-a61b17362069', N'Jhon', N'Flores', N'', N'', N'', N'flores@yopmail.com', N'', N'TOYOTA', NULL, 0, N'N/A', CAST(N'2024-01-17T10:26:10.927' AS DateTime), CAST(N'2024-01-17T10:28:05.250' AS DateTime), CAST(N'2024-01-17' AS Date), CAST(N'11:28:00' AS Time), CAST(N'16:25:00' AS Time), 4, 3281, NULL, 0, 1)
INSERT [dbo].[TVisita] ([id], [guid], [nombre], [apellidos], [ci], [extencionCi], [tipoDocumento], [email], [celular], [empresa], [motivo], [nroTarjeta], [observaciones], [fechaCreacion], [fechaModificacion], [fechaVisita], [horaInicio], [horaFin], [estado], [idVisitado], [idVisitante], [eliminado], [sincronizado]) VALUES (2602, NULL, N'Erwin', N'Paz', N'6387768', NULL, NULL, N'spaz@yopmail.com', N'76545434', N'TOYOTA', NULL, NULL, N'N/A', CAST(N'2024-01-17T10:33:18.103' AS DateTime), CAST(N'2024-01-17T10:34:40.793' AS DateTime), CAST(N'2024-01-18' AS Date), CAST(N'14:33:00' AS Time), CAST(N'19:33:00' AS Time), 2, 3281, 3286, 0, 0)
SET IDENTITY_INSERT [dbo].[TVisita] OFF
GO
SET IDENTITY_INSERT [dbo].[TVisitanteVisitado] ON 

INSERT [dbo].[TVisitanteVisitado] ([id], [idVisitante], [idVisitado]) VALUES (4, 3285, 3281)
INSERT [dbo].[TVisitanteVisitado] ([id], [idVisitante], [idVisitado]) VALUES (5, 3286, 3281)
SET IDENTITY_INSERT [dbo].[TVisitanteVisitado] OFF
GO
SET IDENTITY_INSERT [dbo].[TZona] ON 

INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (58, N'Punto de Encuentro II', 1, 1, 1, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (59, N'Punto de Encuentro I', 1, 1, 1, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (60, N'YPFB SANTA CRUZ', 1, 2, 1, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (1060, N'YPFB LA PAZ', 1, 2, 5, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (1061, N'YPFB COCHABAMBA', 1, 2, 2, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (1062, N'YPFB SUCRE', 1, 2, 1010, 0)
INSERT [dbo].[TZona] ([id], [nombre], [estado], [tipo], [idCiudad], [eliminado]) VALUES (1063, N'FUERA DE LA EMPRESA', 1, 3, 1, 0)
SET IDENTITY_INSERT [dbo].[TZona] OFF
GO
ALTER TABLE [dbo].[TConfigDesign] ADD  CONSTRAINT [DF_TConfigDesign_colorLogSup]  DEFAULT ('color') FOR [colorLogSup]
GO
ALTER TABLE [dbo].[TConfigDesign] ADD  CONSTRAINT [DF_TConfigDesign_colorLogInf]  DEFAULT ('color') FOR [colorLogInf]
GO
ALTER TABLE [dbo].[TUsuario] ADD  CONSTRAINT [DF_TUsuario_eliminado]  DEFAULT ((0)) FOR [eliminado]
GO
ALTER TABLE [dbo].[TUsuario] ADD  CONSTRAINT [DF_TUsuario_primerInicio]  DEFAULT ((0)) FOR [primerInicio]
GO
ALTER TABLE [dbo].[TUsuario] ADD  CONSTRAINT [DF_TUsuario_estado]  DEFAULT ((0)) FOR [estado]
GO
ALTER TABLE [dbo].[TEventosMarcacion]  WITH CHECK ADD  CONSTRAINT [FK__TEventosM__idEve__5C6CB6D7] FOREIGN KEY([idEvento])
REFERENCES [dbo].[THistorialEvento] ([id])
GO
ALTER TABLE [dbo].[TEventosMarcacion] CHECK CONSTRAINT [FK__TEventosM__idEve__5C6CB6D7]
GO
ALTER TABLE [dbo].[TEventosMarcacion]  WITH CHECK ADD  CONSTRAINT [FK_TEventosMarcacion_TCardHolder] FOREIGN KEY([idCardHolder])
REFERENCES [dbo].[TCardHolder] ([id])
GO
ALTER TABLE [dbo].[TEventosMarcacion] CHECK CONSTRAINT [FK_TEventosMarcacion_TCardHolder]
GO
ALTER TABLE [dbo].[THistorialEvento]  WITH CHECK ADD  CONSTRAINT [FK_THistorialEvento_TCiudad] FOREIGN KEY([idCiudad])
REFERENCES [dbo].[TCiudad] ([id])
GO
ALTER TABLE [dbo].[THistorialEvento] CHECK CONSTRAINT [FK_THistorialEvento_TCiudad]
GO
ALTER TABLE [dbo].[THorarioPerfil]  WITH CHECK ADD FOREIGN KEY([idPerfil])
REFERENCES [dbo].[TPerfilAsistencia] ([id])
GO
ALTER TABLE [dbo].[TOperaciones]  WITH CHECK ADD FOREIGN KEY([idModulo])
REFERENCES [dbo].[TModulo] ([id])
GO
ALTER TABLE [dbo].[TRolOperacion]  WITH CHECK ADD FOREIGN KEY([idOperacion])
REFERENCES [dbo].[TOperaciones] ([id])
GO
ALTER TABLE [dbo].[TRolOperacion]  WITH CHECK ADD FOREIGN KEY([idRol])
REFERENCES [dbo].[TRol] ([id])
GO
ALTER TABLE [dbo].[TSecretariaGerente]  WITH CHECK ADD  CONSTRAINT [FK__TSecretar__idGer__1BE81D6E] FOREIGN KEY([idGerente])
REFERENCES [dbo].[TUsuario] ([id])
GO
ALTER TABLE [dbo].[TSecretariaGerente] CHECK CONSTRAINT [FK__TSecretar__idGer__1BE81D6E]
GO
ALTER TABLE [dbo].[TSecretariaGerente]  WITH CHECK ADD  CONSTRAINT [FK__TSecretar__idSec__1AF3F935] FOREIGN KEY([idSecretaria])
REFERENCES [dbo].[TUsuario] ([id])
GO
ALTER TABLE [dbo].[TSecretariaGerente] CHECK CONSTRAINT [FK__TSecretar__idSec__1AF3F935]
GO
ALTER TABLE [dbo].[TSeguimientoTablet]  WITH CHECK ADD  CONSTRAINT [FK_TSeguimientoTablet_TTablet] FOREIGN KEY([idTablet])
REFERENCES [dbo].[TTablet] ([id])
GO
ALTER TABLE [dbo].[TSeguimientoTablet] CHECK CONSTRAINT [FK_TSeguimientoTablet_TTablet]
GO
ALTER TABLE [dbo].[TUsuario]  WITH CHECK ADD  CONSTRAINT [FK_TUsuario_TCiudad] FOREIGN KEY([idCiudad])
REFERENCES [dbo].[TCiudad] ([id])
GO
ALTER TABLE [dbo].[TUsuario] CHECK CONSTRAINT [FK_TUsuario_TCiudad]
GO
ALTER TABLE [dbo].[TUsuario]  WITH CHECK ADD  CONSTRAINT [FK_TUsuario_TRol] FOREIGN KEY([idRol])
REFERENCES [dbo].[TRol] ([id])
GO
ALTER TABLE [dbo].[TUsuario] CHECK CONSTRAINT [FK_TUsuario_TRol]
GO
ALTER TABLE [dbo].[TZona]  WITH CHECK ADD  CONSTRAINT [FK__TZona__tipo__1AD3FDA4] FOREIGN KEY([tipo])
REFERENCES [dbo].[TTipo] ([id])
GO
ALTER TABLE [dbo].[TZona] CHECK CONSTRAINT [FK__TZona__tipo__1AD3FDA4]
GO
ALTER TABLE [dbo].[TZona]  WITH CHECK ADD  CONSTRAINT [FK_TZona_TCiudad] FOREIGN KEY([idCiudad])
REFERENCES [dbo].[TCiudad] ([id])
GO
ALTER TABLE [dbo].[TZona] CHECK CONSTRAINT [FK_TZona_TCiudad]
GO
USE [master]
GO
ALTER DATABASE [Tmac_Manager1] SET  READ_WRITE 
GO
