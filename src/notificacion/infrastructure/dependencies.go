// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/API_recolecta/src/notificacion/application"
)

func InitNotificacionDependencies() (
	// Controladores de Conteo
	*CountNotificacionesActivasByUsuarioIDController,
	*CountNotificacionesByCamionIDController,
	*CountNotificacionesByTipoController,
	*CountNotificacionesByUsuarioIDController,
	
	// Controladores de Creación Específica
	*CrearNotificacionEmergenciaController,
	*CrearNotificacionFallaController,
	*CrearNotificacionMantenimientoController,
	*CreateNotificacionController,
	
	// Controladores CRUD Básicos
	*DeleteNotificacionController,
	*GetAllNotificacionesController,
	*GetNotificacionByIdController,
	*UpdateNotificacionController,
	
	// Controladores de Consultas Activas/Inactivas
	*GetNotificacionesActivasByUsuarioIDController,
	*GetNotificacionesActivasController,
	*GetNotificacionesInactivasController,
	
	// Controladores por Camión
	*GetNotificacionesByCamionIDController,
	*GetNotificacionesByCamionYTipoController,
	
	// Controladores por Usuario
	*GetNotificacionesByUsuarioIDController,
	*GetNotificacionesByUsuarioYTipoController,
	
	// Controladores por Relaciones
	*GetNotificacionesByCreadoPorController,
	*GetNotificacionesByFallaIDController,
	*GetNotificacionesByMantenimientoIDController,
	
	// Controladores por Tipo y Rango
	*GetNotificacionesByTipoController,
	*GetNotificacionesByFechaRangeController,
	
	// Controladores de Notificaciones Especiales
	*GetNotificacionesGlobalesController,
	
	// Controladores de Estado
	*MarcarNotificacionComoActivaController,
	*MarcarNotificacionComoLeidaController,
	*MarcarTodasNotificacionesComoLeidasController,
	
	// Controladores de Notificación (Actualizados)
	*NotificarUsuarioController,
	*ObtenerNumeroNotificacionesNoLeidasController,
	
	// ========== NUEVOS CONTROLADORES ==========
	*NotificarMultiplesUsuariosController,
	*NotificarTodosUsuariosController,
) {
	// Repositorio PostgreSQL
	repo := NewPostgresNotificacionRepository()

	// ================== CASOS DE USO DE CONTEO ==================
	countActivasByUsuarioIDUseCase := application.NewCountNotificacionesActivasByUsuarioIDUseCase(repo)
	countByCamionIDUseCase := application.NewCountNotificacionesByCamionIDUseCase(repo)
	countByTipoUseCase := application.NewCountNotificacionesByTipoUseCase(repo)
	countByUsuarioIDUseCase := application.NewCountNotificacionesByUsuarioIDUseCase(repo)
	
	// ================== CASOS DE USO DE CREACIÓN ==================
	crearEmergenciaUseCase := application.NewCrearNotificacionEmergenciaUseCase(repo)
	crearFallaUseCase := application.NewCrearNotificacionFallaUseCase(repo)
	crearMantenimientoUseCase := application.NewCrearNotificacionMantenimientoUseCase(repo)
	createNotificacionUseCase := application.NewCreateNotificacionUseCase(repo)
	
	// ================== CASOS DE USO CRUD BÁSICOS ==================
	deleteNotificacionUseCase := application.NewDeleteNotificacionUseCase(repo)
	getAllNotificacionesUseCase := application.NewGetAllNotificacionesUseCase(repo)
	getNotificacionByIdUseCase := application.NewGetNotificacionByIdUseCase(repo)
	updateNotificacionUseCase := application.NewUpdateNotificacionUseCase(repo)
	
	// ================== CASOS DE USO DE CONSULTAS ==================
	getActivasByUsuarioIDUseCase := application.NewGetNotificacionesActivasByUsuarioIDUseCase(repo)
	getActivasUseCase := application.NewGetNotificacionesActivasUseCase(repo)
	getInactivasUseCase := application.NewGetNotificacionesInactivasUseCase(repo)
	
	getByCamionIDUseCase := application.NewGetNotificacionesByCamionIDUseCase(repo)
	getByCamionYTipoUseCase := application.NewGetNotificacionesByCamionYTipoUseCase(repo)
	
	getByUsuarioIDUseCase := application.NewGetNotificacionesByUsuarioIDUseCase(repo)
	getByUsuarioYTipoUseCase := application.NewGetNotificacionesByUsuarioYTipoUseCase(repo)
	
	getByCreadoPorUseCase := application.NewGetNotificacionesByCreadoPorUseCase(repo)
	getByFallaIDUseCase := application.NewGetNotificacionesByFallaIDUseCase(repo)
	getByMantenimientoIDUseCase := application.NewGetNotificacionesByMantenimientoIDUseCase(repo)
	
	getByTipoUseCase := application.NewGetNotificacionesByTipoUseCase(repo)
	getByFechaRangeUseCase := application.NewGetNotificacionesByFechaRangeUseCase(repo)
	
	getGlobalesUseCase := application.NewGetNotificacionesGlobalesUseCase(repo)
	
	// ================== CASOS DE USO DE ESTADO ==================
	marcarComoActivaUseCase := application.NewMarcarNotificacionComoActivaUseCase(repo)
	marcarComoLeidaUseCase := application.NewMarcarNotificacionComoLeidaUseCase(repo)
	marcarTodasComoLeidasUseCase := application.NewMarcarTodasNotificacionesComoLeidasUseCase(repo)
	
	// ================== CASOS DE USO DE NOTIFICACIÓN ==================
	notificarUsuarioUseCase := application.NewNotificarUsuarioUseCase(repo)
	obtenerNoLeidasUseCase := application.NewObtenerNumeroNotificacionesNoLeidasUseCase(repo)
	
	// ================== NUEVOS CASOS DE USO ==================
	notificarMultiplesUseCase := application.NewNotificarMultiplesUsuariosUseCase(repo)
	notificarTodosUseCase := application.NewNotificarTodosUsuariosUseCase(repo)

	// ================== CONTROLADORES DE CONTEO ==================
	countActivasByUsuarioIDController := NewCountNotificacionesActivasByUsuarioIDController(countActivasByUsuarioIDUseCase)
	countByCamionIDController := NewCountNotificacionesByCamionIDController(countByCamionIDUseCase)
	countByTipoController := NewCountNotificacionesByTipoController(countByTipoUseCase)
	countByUsuarioIDController := NewCountNotificacionesByUsuarioIDController(countByUsuarioIDUseCase)
	
	// ================== CONTROLADORES DE CREACIÓN ==================
	crearEmergenciaController := NewCrearNotificacionEmergenciaController(crearEmergenciaUseCase)
	crearFallaController := NewCrearNotificacionFallaController(crearFallaUseCase)
	crearMantenimientoController := NewCrearNotificacionMantenimientoController(crearMantenimientoUseCase)
	createNotificacionController := NewCreateNotificacionController(createNotificacionUseCase)
	
	// ================== CONTROLADORES CRUD BÁSICOS ==================
	deleteNotificacionController := NewDeleteNotificacionController(deleteNotificacionUseCase)
	getAllNotificacionesController := NewGetAllNotificacionesController(getAllNotificacionesUseCase)
	getNotificacionByIdController := NewGetNotificacionByIdController(getNotificacionByIdUseCase)
	updateNotificacionController := NewUpdateNotificacionController(updateNotificacionUseCase)
	
	// ================== CONTROLADORES DE CONSULTAS ACTIVAS/INACTIVAS ==================
	getActivasByUsuarioIDController := NewGetNotificacionesActivasByUsuarioIDController(getActivasByUsuarioIDUseCase)
	getActivasController := NewGetNotificacionesActivasController(getActivasUseCase)
	getInactivasController := NewGetNotificacionesInactivasController(getInactivasUseCase)
	
	// ================== CONTROLADORES POR CAMIÓN ==================
	getByCamionIDController := NewGetNotificacionesByCamionIDController(getByCamionIDUseCase)
	getByCamionYTipoController := NewGetNotificacionesByCamionYTipoController(getByCamionYTipoUseCase)
	
	// ================== CONTROLADORES POR USUARIO ==================
	getByUsuarioIDController := NewGetNotificacionesByUsuarioIDController(getByUsuarioIDUseCase)
	getByUsuarioYTipoController := NewGetNotificacionesByUsuarioYTipoController(getByUsuarioYTipoUseCase)
	
	// ================== CONTROLADORES POR RELACIONES ==================
	getByCreadoPorController := NewGetNotificacionesByCreadoPorController(getByCreadoPorUseCase)
	getByFallaIDController := NewGetNotificacionesByFallaIDController(getByFallaIDUseCase)
	getByMantenimientoIDController := NewGetNotificacionesByMantenimientoIDController(getByMantenimientoIDUseCase)
	
	// ================== CONTROLADORES POR TIPO Y RANGO ==================
	getByTipoController := NewGetNotificacionesByTipoController(getByTipoUseCase)
	getByFechaRangeController := NewGetNotificacionesByFechaRangeController(getByFechaRangeUseCase)
	
	// ================== CONTROLADORES DE NOTIFICACIONES ESPECIALES ==================
	getGlobalesController := NewGetNotificacionesGlobalesController(getGlobalesUseCase)
	
	// ================== CONTROLADORES DE ESTADO ==================
	marcarComoActivaController := NewMarcarNotificacionComoActivaController(marcarComoActivaUseCase)
	marcarComoLeidaController := NewMarcarNotificacionComoLeidaController(marcarComoLeidaUseCase)
	marcarTodasComoLeidasController := NewMarcarTodasNotificacionesComoLeidasController(marcarTodasComoLeidasUseCase)
	
	// ================== CONTROLADORES DE NOTIFICACIÓN ==================
	notificarUsuarioController := NewNotificarUsuarioController(notificarUsuarioUseCase)
	obtenerNoLeidasController := NewObtenerNumeroNotificacionesNoLeidasController(obtenerNoLeidasUseCase)
	
	// ================== NUEVOS CONTROLADORES ==================
	notificarMultiplesController := NewNotificarMultiplesUsuariosController(notificarMultiplesUseCase)
	notificarTodosController := NewNotificarTodosUsuariosController(notificarTodosUseCase)

	return countActivasByUsuarioIDController,
	       countByCamionIDController,
	       countByTipoController,
	       countByUsuarioIDController,
	       
	       crearEmergenciaController,
	       crearFallaController,
	       crearMantenimientoController,
	       createNotificacionController,
	       
	       deleteNotificacionController,
	       getAllNotificacionesController,
	       getNotificacionByIdController,
	       updateNotificacionController,
	       
	       getActivasByUsuarioIDController,
	       getActivasController,
	       getInactivasController,
	       
	       getByCamionIDController,
	       getByCamionYTipoController,
	       
	       getByUsuarioIDController,
	       getByUsuarioYTipoController,
	       
	       getByCreadoPorController,
	       getByFallaIDController,
	       getByMantenimientoIDController,
	       
	       getByTipoController,
	       getByFechaRangeController,
	       
	       getGlobalesController,
	       
	       marcarComoActivaController,
	       marcarComoLeidaController,
	       marcarTodasComoLeidasController,
	       
	       notificarUsuarioController,
	       obtenerNoLeidasController,
	       
	       notificarMultiplesController,
	       notificarTodosController
}