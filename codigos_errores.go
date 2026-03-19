package siat

import "fmt"

// SIAT Error Codes
const (
	CodeCufdFueraDeTolerancia                        = 123
	CodeRecepcionPendiente                           = 901
	CodeRecepcionRechazada                           = 902
	CodeRecepcionProcesada                           = 903
	CodeRecepcionObservada                           = 904
	CodeAnulacionConfirmada                          = 905
	CodeAnulacionRechazada                           = 906
	CodeReversionAnulacionConfirmada                 = 907
	CodeRecepcionValidada                            = 908
	CodeReversionAnulacionRechazada                  = 909
	CodeAmbienteInvalido                             = 910
	CodeCodigoSistemaInvalido                        = 911
	CodeSistemaNoAsociadoAlContribuyente             = 912
	CodeCuisInvalido                                 = 913
	CodeCufdInvalido                                 = 914
	CodeTipoFacturaDocumentoInvalido                 = 915
	CodeTipoEmisionInvalido                          = 916
	CodeModalidadInvalida                            = 917
	CodeSucursalInvalida                             = 918
	CodeNitInvalido                                  = 919
	CodeArchivoInvalido                              = 920
	CodeFirmadoXmlIncorrecto                         = 921
	CodeFirmaXmlNoCorrespondeAlContribuyente         = 922
	CodeCodigoRecepcionInvalido                      = 923
	CodeFacturaNoExisteEnSin                         = 924
	CodeMotivoAnulacionInvalido                      = 925
	CodeComunicacionExitosa                          = 926
	CodeCertificadoFirmaInvalido                     = 927
	CodeCertificadoRevocado                          = 928
	CodeCuisNoVigente                                = 929
	CodeCuisNoCorrespondeASucursalPuntoVenta         = 930
	CodeCodigoDocumentoSectorInvalido                = 931
	CodeCodigoDocumentoSectorNoCorrespondeAlServicio = 932
	CodePuntoVentaInexistenteOInvalido               = 933
	CodeAnulacionFueraDePlazo                        = 934
	CodeFechaEnvioInvalida                           = 935
	CodeFacturaYaAnulada                             = 936
	CodeNitNoAsociadoAModalidad                      = 937
	CodeNitPresentaMarcasDeControl                   = 938
	CodeFacturaNoCumpleXsd                           = 939
	CodeNitNoHabilitadoDocumentoSector               = 940
	CodeFacturaNoDisponibleParaAnular                = 941
	CodeEventoSignificativoNoExisteEnSin             = 942
	CodeFormatoFechaEnvioIncorrecto                  = 943
	CodeCodigoRecepcionNoExisteEnSin                 = 944
	CodeEstadoRecepcionAnulacionIncorrecta           = 945
	CodeCufNoExisteEnSin                             = 946
	CodeTipoPuntoVentaInvalido                       = 947
	CodeNombrePuntoVentaVacio                        = 948
	CodeDescripcionPuntoVentaVacio                   = 949
	CodeEventoSignificativoVacio                     = 950
	CodeDescripcionEventoSignificativoVacio          = 951
	CodeCufYaRegistradoEnSin                         = 952
	CodeCufdNoVigente                                = 953
	CodePaqueteContingenciaExcedido                  = 954
	CodeNoExisteRegistroMasivo                       = 955
	CodePaqueteMasivoExcedido                        = 956
	CodeNoExisteEventoSignificativo                  = 957
	CodeUsuarioNoAutorizado                          = 958
	CodeCuisNoAsociadoAlSistema                      = 959
	CodeFinDeEventoRequerido                         = 960
	CodeMarcaDomicilioInexistente                    = 961
	CodeBloqueoDosificacionFiscalizacion             = 962
	CodeBloqueoDosificacionJuridica                  = 963
	CodeNoCumpleObligatoriedadDDJJ                   = 964
	CodeContribuyenteSinFirmaVigente                 = 965
	CodeNoSePuedeRecuperarDatosContribuyente         = 966
	CodeTiempoEsperaAgotadoDB                        = 967
	CodeAnulacionYaRevertida                         = 968
	CodeHashInvalido                                 = 969
	CodeCuisVigenteNoPuedeSolicitarOtro              = 970
	CodeArchivoExcedeTamano                          = 971
	CodeCantidadFacturasExcedeNormativa              = 972
	CodeRangoFechasEventoInvalido                    = 974
	CodeSistemaNoAutorizadoOObservado                = 975
	CodeCodigoEventoIncorrecto                       = 976
	CodeNoExistenActividadesNIT                      = 977
	CodeReversionFacturaConfirmada                   = 978
	CodeCuisNoAsociadoASistemaOSucursal              = 979
	CodeExisteCuisVigente                            = 980
	CodeRangoFechasEventoSignificativoInvalido       = 981
	CodeNoExistePuntosDeVenta                        = 982
	CodeFechaEnvioPaqueteFueraDePlazo                = 983
	CodeEventoNoCorrespondeCufd                      = 984
	CodeCantidadFacturasDiferenteADeclarada          = 985
	CodeNitActivo                                    = 986
	CodeNitInactivo                                  = 987
	CodeCuisFueraDeTolerancia                        = 988
	CodeTokenInvalido                                = 989
	CodeClienteSinActividadesSector                  = 990
	CodeErrorBaseDeDatos                             = 991
	CodeErrorServicioPadron                          = 992
	CodeFechaEnvioFueraDePlazo                       = 993
	CodeNitInexistente                               = 994
	CodeServicioNoDisponible                         = 995
	CodeRangoFechasInvalido                          = 996
	CodeNombreExcedeLimite                           = 997
	CodeDescripcionExcedeLimite                      = 998
	CodeErrorEjecucionServicio                       = 999
	CodeCufYaExisteEnSin                             = 1000
	CodeNitNoCorrespondeACufd                        = 1001
	CodeCufInvalido                                  = 1002
	CodeCufdEnXmlInvalido                            = 1003
	CodeSucursalNoCorrespondeACufd                   = 1004
	CodeFacturaNoPuedeSerEmitidaAlMismoEmisor        = 1005
	CodeCufdNoCorrespondeAEvento                     = 1006
	CodeDireccionNoCorrespondeAPadron                = 1007
	CodePuntoVentaEnXmlInexistente                   = 1008
	CodeFechaEmisionNoValidaLinea                    = 1009
	CodeFacturaMontoExcedidoSinDatos                 = 1010
	CodeComplementoSoloCarnet                        = 1011
	CodeNumeroTarjetaSoloConTarjeta                  = 1012
	CodeCalculoMontoTotalErroneo                     = 1013
	CodeCalculoMontoTotalMonedaErroneo               = 1014
	CodeCalculoImporteBaseErroneo                    = 1015
	CodeActividadNoHabilitada                        = 1016
	CodeProductoNoRelacionadoAActividad              = 1017
	CodeCalculoSubtotalErroneo                       = 1018
	CodeCalculoIceEspecificoErroneo                  = 1019
	CodeCalculoIcePorcentualErroneo                  = 1020
	CodeMontoIceEspecificoErroneo                    = 1021
	CodeMontoIcePorcentualErroneo                    = 1022
	CodeCodigoNandinaErroneo                         = 1023
	CodeSumatoriaDetallesErronea                     = 1024
	CodeMontoSujetoCreditoLey317Erroneo              = 1025
	CodeMontoTotalSujetoIJErroneo                    = 1026
	CodeMontoDiferenciaCambioErroneo                 = 1027
	CodeMontoIvaErroneo                              = 1028
	CodeMontoTotalDevueltoErroneo                    = 1029
	CodeMontoTotalOriginalErroneo                    = 1030
	CodeMontoEfectivoDevueltoErroneo                 = 1031
	CodeMontoTotalIPJErroneo                         = 1032
	CodeMontoDevueltoMayorAOriginal                  = 1033
	CodeFechaEmisionMenorPeriodoAnterior             = 1034
	CodeFormatoFechaIncorrecto                       = 1035
	CodeNominatividadIncorrecta                      = 1036
	CodeNitNoValido                                  = 1037
	CodeNitConjuntoNoValido                          = 1038
	CodeFechaEmisionMasivaIncorrecta                 = 1039
	CodeFechaEmisionFueraRangoContingencia           = 1040
	CodeFechaEmisionFueraPlazoNorma                  = 1041
	CodeNitMedicoNoValido                            = 1042
	CodeMontoConciliadoErroneo                       = 1043
	CodeMontoTotalConciliadoErroneo                  = 1044
	CodeCafcNoValido                                 = 1045
	CodeFechaEmisionCafcIncorrecta                   = 1046
	CodeNumeroFacturaCafcIncorrecto                  = 1047
	CodeNotaCreditoFacturaNoEncontrada               = 1048
	CodeNotaDetalleDiferenteAFactura                 = 1049
	CodeMontoGiftCardNoCorresponde                   = 1050
	CodeFechaFacturaIncorrecta                       = 1051
	CodeCalculoMontoIehdErroneo                      = 1052
	CodeActividadNotaAutorizadaPlazo                 = 1053
	CodeMontoDescuentoCreditoDebitoErroneo           = 1054
	CodeMontoTarifaErroneo                           = 1055
	CodeTipoCambioErroneo                            = 1056
	CodeMontoTotalMonedaComputarizadaErroneo         = 1057
	CodeMontoTotalSujetoIvaErroneo                   = 1058
	CodeRazonSocialErronea                           = 1059
	CodeMontoDetalleErroneo                          = 1060
	CodeNotaNoValidaParaDevolucion                   = 1061
	CodeWarnCorrelatividadFactura                    = 2000
	CodeWarnFechaRangoContingencia                   = 2001
	CodeWarnFechaEmisionMasiva                       = 2002
	CodeWarnFacturaMontoSinDatos                     = 2003
	CodeWarnComplementoSoloCarnet                    = 2004
	CodeWarnNitClienteNoValido                       = 2005
	CodeWarnNumeroTarjetaSoloConTarjeta              = 2006
	CodeWarnCalculoMontoTotal                        = 2007
	CodeWarnCalculoMontoTotalMoneda                  = 2008
	CodeWarnCalculoImporteBase                       = 2009
	CodeWarnActividadNoHabilitada                    = 2010
	CodeWarnProductoNoRelacionado                    = 2011
	CodeWarnCalculoSubtotal                          = 2012
	CodeWarnEmisionAMismoEmisor                      = 2013
	CodeWarnDireccionNoCorrespondePadron             = 2014
	CodeWarnCalculoIceEspecifico                     = 2015
	CodeWarnCalculoIcePorcentual                     = 2016
	CodeWarnMontoIceEspecifico                       = 2017
	CodeWarnMontoIcePorcentual                       = 2018
	CodeWarnCodigoNandinaErroneo                     = 2019
	CodeNitNoContratoVigente                         = 3000
	CodeCategoriaContratoNoSector                    = 3001
	CodeExcedeLimiteCufdMasivo                       = 3002
	CodeMarcaNoFormulariosVigentes                   = 3003
	CodeMarcaDomicilioInexistenteMarca               = 3004
	CodeMarcaBloqueoFiscalizacion                    = 3005
	CodeMarcaBloqueoJuridica                         = 3006
	CodeMarcaNoObligatoriedadDDJJ                    = 3007
	CodeWarnCuisExpira                               = 3008
	CodeTamanoArchivoMayorNorma                      = 3009
	CodeFacturaYaUtilizadaOConsolidada               = 3010
)

var codigosMap = map[int]string{
	123:  "Código Único De Facturación Diaria (Cufd) Fuera De Tolerancia",
	901:  "Recepción Pendiente",
	902:  "Recepción Rechazada",
	903:  "Recepción Procesada",
	904:  "Recepción Observada",
	905:  "Anulación Confirmada",
	906:  "Anulación Rechazada",
	907:  "Reversión De Anulación Confirmada",
	908:  "Recepción Validada",
	909:  "Reversión De Anulación Rechazada",
	910:  "El Parámetro Ambiente Es Invalido",
	911:  "El Parámetro Código De Sistema Es Invalido",
	912:  "El Sistema No Esta Asociado Al Contribuyente",
	913:  "Código Único De Inicio De Sistema (Cuis) Invalido",
	914:  "Código Único De Facturación Diaria (Cufd) Invalido",
	915:  "El Parámetro Tipo Factura Documento Es Invalido",
	916:  "El Parámetro Tipo De Emisión Es Invalido",
	917:  "El Parámetro Modalidad Es Invalido",
	918:  "El Parámetro Sucursal Es Invalido",
	919:  "El Parámetro NIT Es Invalido",
	920:  "El Parámetro Archivo Es Invalido",
	921:  "El Firmado Del XML Es Incorrecto",
	922:  "La Firma Del XML No Corresponde Al Contribuyente",
	923:  "El Parámetro Código De Recepción Es Invalido",
	924:  "La Factura o Nota, No Existe En La Base De Datos Del Sin",
	925:  "El Parámetro Motivo De Anulación Es Invalido",
	926:  "Comunicación Exitosa",
	927:  "El Certificado De La Firma Es Invalido",
	928:  "El Certificado Se Encuentra Revocado",
	929:  "El Código Único De Inicio De Sistema (Cuis) No Esta Vigente",
	930:  "El Código Único De Inicio De Sistema (Cuis) No Corresponde A La Sucursal/Punto Venta",
	931:  "El Parámetro Código Documento Sector Es Invalido",
	932:  "El Parámetro Código Documento Sector No Corresponde Al Servicio",
	933:  "El Punto De Venta Es Inexistente o Invalido",
	934:  "La Solicitud De Anulación De La Factura o Nota De Crédito-Débito Se Encuentra Fuera De Plazo",
	935:  "El Parámetro Fecha De Envío Es Invalido",
	936:  "La Factura o Nota De Crédito-Débito Ya Se Encuentra Anulada",
	937:  "El NIT No Tiene Asociado La Modalidad De Facturación",
	938:  "El NIT Presenta Marcas De Control",
	939:  "La Factura o Nota De Crédito - Débito No Cumple Con El Formato Del Xsd Especificado",
	940:  "El NIT No Tiene Habilitado El Documento Sector",
	941:  "La Factura o Nota De Crédito - Débito No Se Encuentra Disponible Para Ser Anulada",
	942:  "El Código De Recepción De Evento Significativo No Se Encuentra En La Base De Datos Del Sin",
	943:  "El Formato De La Fecha De Envío Es Incorrecto",
	944:  "El Código De Recepción No Se Encuentra En La Base De Datos Del Sin",
	945:  "El Estado De Recepción De La Anulación Es Incorrecta",
	946:  "El Código Único De Factura (Cuf) No Existe En Base De Datos Del Sin",
	947:  "El Parámetro Tipo De Punto De Venta Es Invalido",
	948:  "El Parámetro Nombre De Punto De Venta No Puede Ser Vacío",
	949:  "El Parámetro Descripción De Punto De Venta No Puede Ser Vacío",
	950:  "El Parámetro Código De Evento Significativo No Puede Ser Vacío",
	951:  "El Parámetro Descripción De Evento Significativo No Puede Ser Vacío",
	952:  "El Código Único De Factura (Cuf) Ya Se Encuentra Registrado En La Base De Datos Del Sin",
	953:  "El Código Único De Facturación Diaria (Cufd) No Se Encuentra Vigente",
	954:  "La Cantidad De Facturas En El Paquete Emitido Por Contingencia Ha Excedido El Máximo Permitido",
	955:  "No Existe Registro Para Autorizar El Proceso Masivo",
	956:  "La Cantidad De Facturas En El Paquete Emitido Masivamente Ha Excedido El Máximo Permitido",
	957:  "No Existe Registro De Evento Significativo En La Base De Datos Del Sin",
	958:  "El Usuario No Se Encuentra Autorizado Para Consumir Este Servicio",
	959:  "El Código Único De Inicio De Sistema (Cuis) No Se Encuentra Asociado Al Sistema",
	960:  "El Parámetro Fin De Evento Es Requerido",
	961:  "El NIT Tiene Marca De Domicilio Inexistente",
	962:  "El NIT Tiene Bloqueo De Dosificación Originado En Fiscalización",
	963:  "El NIT Tiene Bloqueo De Dosificación Originado En Jurídica",
	964:  "El NIT No Cumple Con Obligatoriedad De Presentación De DDJJ",
	965:  "El Contribuyente No Cuenta Con Firma Vigente Registrada",
	966:  "No Se Puede Recuperar Los Datos Del Contribuyente",
	967:  "Tiempo De Espera Agotado Para Conexión A Base De Datos",
	968:  "La Anulación De La Factura o Nota De Crédito - Débito Ya Se Encuentra Revertida",
	969:  "El Parámetro Hash Es Invalido",
	970:  "El Cuis En La Base De Datos Se Encuentra Vigente, No Puede Solicitar Otro",
	971:  "El Tamaño Del Archivo Excede El Tamaño Permitido De 100 Mb",
	972:  "La Cantidad De Facturas Enviada En El Paquete Es Mayor A La Definida En La Normativa",
	973:  "El Código Único De Inicio De Sistema (Cuis) No Se Encuentra Vigente",
	974:  "El Rango De Fechas Del Evento Significativo Para Registrar Es Inválido",
	975:  "El Sistema No Se Encuentra Autorizado O Se Encuentra Observado",
	976:  "El Código Del Evento Es Incorrecto",
	977:  "No Existen Actividades Asociadas Al NIT",
	978:  "Reversión De La Factura o Nota De Crédito/Débito Confirmada",
	979:  "El Cuis No Se Encuentra Asociado Al Sistema O A La Sucursal",
	980:  "Existe Un Cuis Vigente Para La Sucursal O Punto De Venta",
	981:  "Rango De Fechas De Evento Significativo Invalido",
	982:  "No Existe Puntos De Venta Asociados",
	983:  "La Fecha De Envío Del Paquete Esta Fuera De Plazo",
	984:  "El Evento Significativo No Corresponde Al Cufd Del Evento Registrado",
	985:  "La Cantidad De Facturas Es Diferente A La Definida En La Normativa",
	986:  "NIT Activo",
	987:  "NIT Inactivo",
	988:  "Código Único De Inicio De Sistema (Cuis) Fuera De Tolerancia",
	989:  "Token Invalido",
	990:  "El Cliente No Tiene Actividades Relacionadas Al Sector Que Intenta Asociar",
	991:  "Error En Base De Datos",
	992:  "Error Servicio Padrón",
	993:  "La Fecha De Envío Esta Fuera De Plazo",
	994:  "NIT Inexistente",
	995:  "Servicio No Disponible",
	996:  "Rango De Fechas Invalido",
	997:  "El Nombre Excede El Limite De Caracteres Permitidos",
	998:  "La Descripción Excede El Limite De Caracteres Permitidos",
	999:  "Error En La Ejecución Del Servicio",
	1000: "El Cuf Enviado Ya Existe En La Base De Datos Del Sin",
	1001: "El NIT Enviado En El XML Es Inexistente O No Corresponde Al Cufd",
	1002: "El Código Único De Factura (Cuf) Enviado En El XML Es Invalido",
	1003: "El Código Único De Facturación Diaria (Cufd) Enviado En El XML Es Invalido",
	1004: "La Sucursal Enviada En El XML No Corresponde A Los Datos Del Cufd",
	1005: "La Factura o Nota De Crédito-Débito No Puede Ser Emitida Al Mismo Emisor",
	1006: "El Cufd Enviado No Corresponde Al Evento Asociado Al Paquete Enviado",
	1007: "La Dirección Enviada En El XML No Corresponde A La Registrada En Padrón",
	1008: "El Punto De Venta Enviado En El XML Es Inexistente O Invalido",
	1009: "La Fecha De Emisión Enviada En El XML No Es Valida Para Emisión En Linea",
	1010: "La Factura No Puede Ser Enviada Con Numero De CI/NIT/CEX o Para Montos Mayores A 3000",
	1011: "El Complemento Solo Puede Ser Enviado Cuando El Tipo De Documento Es Carnet De Identidad",
	1012: "El Numero De Tarjeta Solo Puede Ser Enviado Cuando El Método De Pago Sea Con Tarjeta",
	1013: "El Calculo Del Monto Total Es Erróneo",
	1014: "El Calculo Del Monto Total Moneda Es Erróneo",
	1015: "El Calculo Del Importe Base Para Crédito Fiscal Es Erróneo",
	1016: "El Código De Actividad Económica No Esta Habilitada Para El Contribuyente",
	1017: "El Código De Producto No Esta Relacionado A Ninguna Actividad Económica Del Contribuyente",
	1018: "El Calculo Del Subtotal Es Erróneo",
	1019: "El Calculo De Ice Especifico Es Erróneo",
	1020: "El Calculo De Ice Porcentual Es Erróneo",
	1021: "El Monto Ice Especifico Es Erróneo",
	1022: "El Monto Ice Porcentual Es Erróneo",
	1023: "El Código Nandina Enviado En La Factura Es Erróneo",
	1024: "La Sumatoria De Lo Detalles Es Errónea",
	1025: "El Monto Sujeto A Crédito Fiscal Ley 317 Es Erróneo",
	1026: "El Monto Total Sujeto Al Impuesto Del Juego (IJ) Es Erróneo",
	1027: "El Monto De Diferencia De Cambios Es Erróneo",
	1028: "El Monto De IVA Enviado Es Erróneo",
	1029: "El Monto Total Devuelto Enviado Es Erróneo",
	1030: "El Monto Total Original Enviado Es Erróneo",
	1031: "El Monto Efectivo De Crédito O Débito Devuelto Enviado Es Erróneo",
	1032: "El Monto Total De Impuesto A La Participación En Juego (IPJ) Es Erróneo",
	1033: "El Monto Devuelto Es Mayor Al Monto Original",
	1034: "La Fecha De Emisión Es Menor Al Periodo Anterior",
	1035: "Formato De Fecha Incorrecta",
	1036: "Nominatividad Incorrecta Para Nombre/Razón Social",
	1037: "El Numero Documento De Tipo NIT No Es Valido",
	1038: "NIT Conjunto No Valido",
	1039: "Fecha Emisión Para Envío Masivo Incorrecto",
	1040: "La Fecha De Emisión No Se Encuentra En El Rango De Contingencia",
	1041: "La Fecha De Emisión No Se Encuentra Dentro Del Plazo Establecido En Norma",
	1042: "El NIT Del Medico Enviado No Es Valido",
	1043: "El Monto Conciliado Enviado Es Erróneo",
	1044: "El Monto Total Conciliado Enviado Es Erróneo",
	1045: "Valor De Cafc No Valido para la Factura",
	1046: "Fecha Emisión Para El Cafc Enviado Incorrecto",
	1047: "Numero Factura Para El Cafc Enviado Incorrecto",
	1048: "Factura De La Nota Crédito Débito No Encontrada",
	1049: "Detalle De La Nota Diferente Al Detalle De La Factura Original",
	1050: "Monto Gift Card No Corresponde Al Método De Pago",
	1051: "Fecha De Factura Incorrecta",
	1052: "El Calculo Del Monto IEHD Es Erróneo",
	1053: "La Actividad De La Nota De Crédito Débito No Se Encuentra Autorizada Para Este Plazo",
	1054: "El Monto Descuento Crédito Débito Es Erróneo",
	1055: "El Monto Tarifa Es Erróneo",
	1056: "El Tipo De Cambio Es Erróneo",
	1057: "El Monto Total Moneda Es Erróneo",
	1058: "El Monto Total Sujeto Iva Es Erróneo",
	1059: "La Razón Social Es Errónea",
	1060: "El Monto Detalle Es Erróneo",
	1061: "Factura De La Nota Crédito Débito No Es Valida Para Realizar La Devolución",
	2000: "Advertencia: El Numero Factura Enviado Tiene Error De Correlatividad",
	2001: "Advertencia: La Fecha De Emisión Enviada No Se Encuentra Dentro Del Rango Del Evento De Contingencia Asociado",
	2002: "Advertencia: La Fecha De Emisión Enviada No Es Valida Para La Emisión Masiva",
	2003: "Advertencia: La Factura No Puede Ser Enviada Con Numero De CI/NIT/CEX o Para Montos Mayores A 3000",
	2004: "Advertencia: El Complemento Solo Puede Ser Enviado Cuando El Tipo De Documento Es Carnet De Identidad",
	2005: "Advertencia: El NIT Del Cliente Enviado En El Campo Numero De Documento No Es Valido",
	2006: "Advertencia: El Numero De Tarjeta Solo Puede Ser Enviado Cuando El Método De Pago Sea Con Tarjeta",
	2007: "Advertencia: El Calculo Del Monto Total Es Erróneo",
	2008: "Advertencia: El Calculo Del Monto Total Moneda Es Erróneo",
	2009: "Advertencia: El Calculo Del Importe Base Para Crédito Fiscal Es Erróneo",
	2010: "Advertencia: El Código De Actividad Económica No Esta Habilitada Para El Contribuyente",
	2011: "Advertencia: El Código De Producto No Esta Relacionado A Ningún Actividad Económica Del Contribuyente",
	2012: "Advertencia: El Calculo Del Subtotal Es Erróneo",
	2013: "Advertencia: La Factura o Nota De Crédito-Débito No Puede Ser Emitida Al Mismo Emisor",
	2014: "Advertencia: La Dirección Enviada En El XML No Corresponde A La Registrada En Padrón",
	2015: "Advertencia: El Calculo De Ice Especifico Es Erróneo",
	2016: "Advertencia: El Calculo De Ice Porcentual Es Erróneo",
	2017: "Advertencia: El Monto Ice Especifico Es Erróneo",
	2018: "Advertencia: El Monto Ice Porcentual Es Erróneo",
	2019: "Advertencia: El Código Nandina Enviado En La Factura Es Erróneo",
	3000: "El NIT No Tiene Contrato Vigente",
	3001: "La Categoría De Contrato No Corresponde Al Sector",
	3002: "La Solicitud Excede El Limite De Cufd Masivo Permitido",
	3003: "Marca: No Tiene Formularios 200 Y 210 Vigentes",
	3004: "Marca: Domicilio Inexistente",
	3005: "Marca: Bloqueo De Dosificación Originados En Fiscalización",
	3006: "Marca: Bloqueo De Dosificación Originados En Jurídica",
	3007: "Marca: No Cumple Con Obligatoriedad De Presentación De DDJJ",
	3008: "Advertencia: El Cuis Esta A Punto De Caducar, Genere Un Nuevo Cuis Por Favor",
	3009: "El Tamaño Del Archivo Es Mayor A La Definida En Norma",
	3010: "La Factura Ya Se Encuentra Utilizada o Consolidada",
}

// GetMensaje retorna la descripción del código de error del SIAT.
// Si el código no existe, retorna una descripción genérica con el código.
func GetMensaje(code int) string {
	if msg, ok := codigosMap[code]; ok {
		return msg
	}
	return fmt.Sprintf("Código de respuesta SIAT desconocido: %d", code)
}
