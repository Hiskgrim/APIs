'use strict';

/**
 * @ngdoc function
 * @name clienteApp.controller:ContratoRegistroCtrl
 * @description
 * # ContratoRegistroCtrl
 * Controller of the clienteApp
 */
angular.module('clienteApp')
  .controller('ContratoRegistroCtrl', function (contratacion_request,contratacion_mid_request,academica_request,sicapital_request,idResolucion,$mdDialog) {
  	
  	var self = this;

    self.idResolucion=idResolucion;

    contratacion_request.getAll("resolucion_vinculacion_docente",self.idResolucion).then(function(response){      
      self.datosFiltro=response.data;
      /*academica_request.getOne("facultad",33).then(function(response){
        self.contratoGeneral.SedeSolicitante=response.data.Id;
        self.sede_solicitante_defecto=response.data.Nombre;
      });*/
    });

    contratacion_request.getAll("precontratado/"+self.idResolucion.toString()).then(function(response){      
      self.contratados=response.data;
    });

    sicapital_request.getAll("disponibilidad/cdpfiltro/2017/1/VIGENTE","limit=1").then(function(response){
      self.cdp_opciones=response.data;
    });

    self.asignarValoresDefecto = function(){
      self.contratoGeneralBase={}
      self.contratoGeneralBase.Vigencia=new Date().getFullYear();
      self.contratoGeneralBase.FormaPago={Id:240};
      self.contratoGeneralBase.DescripcionFormaPago="Abono a Cuenta Mensual de acuerdo a puntas y hotras laboradas";
      self.contratoGeneralBase.Justificacion="Docente de Vinculacion Especial";
      self.contratoGeneralBase.UnidadEjecucion={Id:205};
      self.contratoGeneralBase.LugarEjecucion={Id:2};
      self.contratoGeneralBase.Observaciones="Contrato de Docente Vinculación Especial";
      self.contratoGeneralBase.TipoControl=181;
      self.contratoGeneralBase.ClaseContratista=33;
      self.contratoGeneralBase.TipoMoneda=137;
      self.contratoGeneralBase.OrigenRecursos=149;
      self.contratoGeneralBase.OrigenPresupuesto=156;
      self.contratoGeneralBase.TemaGastoInversion=166;
      self.contratoGeneralBase.TipoGasto=146;
      self.contratoGeneralBase.RegimenContratacion=136;
      self.contratoGeneralBase.Procedimiento=132;
      self.contratoGeneralBase.ModalidadSeleccion=123;
      self.contratoGeneralBase.TipoCompromiso=35;
      self.contratoGeneralBase.TipologiaContrato=46;
      self.contratoGeneralBase.FechaRegistro=new Date();
      self.contratoGeneralBase.UnidadEjecutora={Id:1};
      self.contratoGeneralBase.Condiciones="Sin condiciones";
    }

    self.asignarValoresDefecto();

    contratacion_request.getAll("supervisor_contrato","limit=-1").then(function(response){
      self.supervisor_contrato_opciones=response.data;
    })
    contratacion_request.getOne("lugar_ejecucion",2).then(function(response){
      self.lugar_ejecucion_defecto=response.data;
    })
    contratacion_request.getOne("parametros",205).then(function(response){
      self.unidad_ejecucion_defecto=response.data;
    })
    contratacion_request.getOne("unidad_ejecutora",1).then(function(response){
      self.unidad_ejecutora_defecto=response.data;
    })
    contratacion_request.getAll("tipo_contrato","limit=-1").then(function(response){
      self.tipo_contrato_opciones=response.data;
    })
    contratacion_request.getOne("parametros",240).then(function(response){
      self.forma_pago_defecto=response.data;
    })
    contratacion_request.getOne("parametros",146).then(function(response){
      self.tipo_gasto_defecto=response.data;
    })
    contratacion_request.getOne("parametros",46).then(function(response){
      self.tipologia_contrato_defecto=response.data;
    })
    contratacion_request.getOne("parametros",35).then(function(response){
      self.tipo_compromiso_defecto=response.data;
    })
    contratacion_request.getOne("parametros",123).then(function(response){
      self.modalidad_seleccion_defecto=response.data;
    })
    contratacion_request.getOne("parametros",132).then(function(response){
      self.procedimiento_defecto=response.data;
    })
    contratacion_request.getOne("parametros",136).then(function(response){
      self.regimen_contratacion_defecto=response.data;
    })
    contratacion_request.getOne("parametros",166).then(function(response){
      self.tema_gasto_defecto=response.data;
    })
    contratacion_request.getOne("parametros",156).then(function(response){
      self.origen_presupuesto_defecto=response.data;
    })
    contratacion_request.getOne("parametros",149).then(function(response){
      self.origen_recursos_defecto=response.data;
    })
    contratacion_request.getOne("parametros",137).then(function(response){
      self.tipo_moneda_defecto=response.data;
    })
    contratacion_request.getOne("parametros",181).then(function(response){
      self.tipo_control_defecto=response.data;
    })
    contratacion_request.getOne("parametros",33).then(function(response){
      self.clase_contratista_defecto=response.data;
    })


    self.cancelar = function(){
      $mdDialog.hide();
    }

    self.calcularSalario = function(){
        contratacion_mid_request.post("calculo_salario/"+self.nivelAcademico+"/"+persona.Id+"/"+self.datosValor.NumSemanas+"/"+self.datosValor.NumHorasSemanales+"/asociado/"+self.datosValor.dedicacion).then(function(response){
          if(typeof(response.data)=="number"){
            self.valorContrato=response.data;
              swal({
                title: "VALOR DEL CONTRATO",
                text: NumeroALetras(response.data),
                type: "info",
                confirmButtonText: "Aceptar",
                closeOnConfirm: false,
                showLoaderOnConfirm: true,
              });
              self.asignarValoresDefecto();
        }else{
          swal({
                title: "Peligro",
                text: "El salario no ha podido ser calculado",
                type: "danger",
                confirmButtonText: "Aceptar",
                closeOnConfirm: false,
                showLoaderOnConfirm: true,
              });
        }
        });
    }

    /*self.realizarContrato = function(){
      self.contratoGeneralBase.LugarEjecucion.Id=parseInt(self.contratoGeneral.LugarEjecucion.Id);
      self.contratoGeneralBase.TipoContrato.Id=parseInt(self.contratoGeneral.TipoContrato.Id);
      self.contratoGeneralBase.DependeciaSolicitante=parseInt(self.idProyectoCurricular);
      self.contratoGeneralBase.NumeroCdp=parseInt(self.contratoGeneral.NumeroCdp);
      self.contratoGeneralBase.FechaRegistro = new Date();
      self.contratados.forEach(function(contratado){
        var contratoGeneral=self.contratoGeneralBase;
        contratoGeneral.Contratista=contratado.Documento;
        contratoGeneral.DependeciaSolicitante=contratado.idProyectoCurricular;
        
      });
    }*/

  });
