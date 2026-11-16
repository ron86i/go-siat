# Sectores fiscales

Los modelos fiscales se encuentran en documents/&lt;sector&gt;. El nombre del paquete identifica la familia documental; usa directamente sus tipos públicos y su constructor.

## Elegir un documento

1. Busca el sector y la familia en la [referencia de documentos](../reference/documentos-v3.md).
2. Importa el paquete documents correspondiente.
3. Construye el modelo público con los tipos y el constructor que documenta ese paquete.
4. Envíalo mediante el cliente SIAT compatible con ese documento.

Por ejemplo, Compra-Venta está en documents/compraventa y se envía por el servicio de facturación. El [ejemplo ejecutable de Compra-Venta](../../../examples/compraventa/main.go) muestra el flujo completo con códigos y recepción.

## Selección del servicio

siat.NewClient devuelve un cliente con campos como Facturacion, Electronica, Computarizada, ServicioBasico, DocumentoAjuste, Telecomunicaciones, EntidadFinanciera y BoletoAereo. La referencia sectorial y cada paquete indican cuál corresponde.

No asumas que dos paquetes comparten el mismo contrato XML porque tengan campos parecidos. Cada paquete sectorial es dueño de su modelo y reglas de serialización.
