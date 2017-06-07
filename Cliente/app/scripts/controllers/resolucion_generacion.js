'use strict';

/**
 * @ngdoc function
 * @name clienteApp.controller:ResolucionGeneracionCtrl
 * @description
 * # ResolucionGeneracionCtrl
 * Controller of the clienteApp
 */
angular.module('clienteApp')
  .controller('ResolucionGeneracionCtrl', function (contratacion_request,academica_request,$mdDialog,$scope,$routeParams) {

  	var self=this;

  	academica_request.getAll("facultad").then(function(response){
  		self.facultades=response.data;
  	});

  	self.resolucion={};

  	$.getJSON("/resolucion.json", function(resolucion) {
        self.resolucion.preambulo=resolucion["preambulo"];
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.resolucion.consideracion=resolucion["consideracion"];
    });

  	self.crearResolucion = function(){
  		var resolucionData={
  			NumeroResolucion: self.resolucion.numero,
  			Vigencia: 2017,
  			IdDependencia: parseInt(self.resolucion.facultad),
  			IdTipoResolucion: {Id: 1},
  			Preambulo: self.resolucion.preambulo,
  			Consideracion: self.resolucion.consideracion,
  			Estado: true
  		}
  		contratacion_request.post("resolucion",resolucionData).then(function(response){
  			var resolucionVinculacionDocenteData={
  				Id: response.data.Id,
  				IdFacultad: parseInt(self.resolucion.facultad),
  				Dedicacion: self.resolucion.dedicacion,
  				NivelAcademico: self.resolucion.nivelAcademico
  			}
  			contratacion_request.post("resolucion_vinculacion_docente",resolucionVinculacionDocenteData).then(function(response){
  				swal(
				  'Resolución registrada!',
				  'La resolución se ha registrado exitosamente!',
				  'success'
				)
				$window.location.href = '#/resolucion_lista';
  			});
  		});
  	}

  });
