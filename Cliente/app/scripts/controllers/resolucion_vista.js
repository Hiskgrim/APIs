'use strict';

/**
 * @ngdoc function
 * @name clienteApp.controller:ResolucionVistaCtrl
 * @description
 * # ResolucionVistaCtrl
 * Controller of the clienteApp
 */
angular.module('clienteApp')
  .controller('ResolucionVistaCtrl', function ($timeout,contratacion_request,contratacion_mid_request,academica_request,$mdDialog,$scope,idResolucion) {
    
  	var self=this;

    self.idResolucion=idResolucion;

    contratacion_request.getOne("resolucion_vinculacion_docente",self.idResolucion).then(function(response){      
      self.datosFiltro=response.data;
      academica_request.getAll("proyecto_curricular/"+self.datosFiltro.NivelAcademico.toLowerCase()+"/"+self.datosFiltro.IdFacultad).then(function(response){
        self.proyectos=response.data;
      });
      contratacion_request.getAll("precontratado/"+self.idResolucion.toString()).then(function(response){      
        self.contratados=response.data;
        self.contratados.forEach(function(row){
          contratacion_mid_request.post("calculo_salario/"+self.datosFiltro.NivelAcademico+"/"+row.Documento+"/"+row.Semanas+"/"+row.HorasSemanales+"/"+row.Categoria.toLowerCase()+"/"+row.Dedicacion.toLowerCase()).then(function(response){
            row.ValorContrato=response.data;
          });
          row.NombreCompleto = row.PrimerNombre + ' ' + row.SegundoNombre + ' ' + row.PrimerApellido + ' ' + row.SegundoApellido;
        });
      });
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.numero=resolucion["numero"];
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.preambulo=resolucion["preambulo"];
    });

    $.getJSON("/resolucion.json", function(resolucion) {
        self.consideracion=resolucion["consideracion"];
    });
    
    $.getJSON("/resolucion.json", function(resolucion) {
        self.articulos=resolucion["articulos"];
    });


  	self.generarResolucion = function() {
	    var documento=getDocumento(self);
	    pdfMake.createPdf(documento).getDataUrl(function(outDoc){
	      document.getElementById('vistaPDF').src = outDoc;
	    });
	 }

   $timeout(self.generarResolucion, 500);

  });
