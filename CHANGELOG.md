# Changelog

## [3.0.0](https://github.com/ron86i/go-siat/compare/v2.4.1...v3.0.0) (2026-09-14)


### ⚠ BREAKING CHANGES

* **facturacion:** reemplaza WithFacturasEnLote por WithFacturasMasivas y renombra sus opciones y métricas.
* rename Bon to Bonificaciones in builders for clarity naming
* El acceso a los servicios cambia de atributos a métodos.

### Features

* add all urls for services ([a5fa4eb](https://github.com/ron86i/go-siat/commit/a5fa4ebbecae49e4d2815ba9b3f5589728d4276e))
* add context-based dynamic config support for multi-tenant (facturacion por terceros) ([1b58ccb](https://github.com/ron86i/go-siat/commit/1b58ccba0681f569ea9626f0a631bb2e8e3c2c90))
* add custom error types and HTTP middleware hooks ([55fb187](https://github.com/ron86i/go-siat/commit/55fb187b7799bc83c70747ee32f74aca000d6b54))
* add DeepWiki badge ([8d01081](https://github.com/ron86i/go-siat/commit/8d01081421da725b9da54dbe2a0e9a367412e7f6))
* add domain model for Recepcion Compras ([4aa105c](https://github.com/ron86i/go-siat/commit/4aa105c6c6887072faffa33b9373ae81aa56255a))
* add HTTPConfig helper and response utility methods ([07c9e6c](https://github.com/ron86i/go-siat/commit/07c9e6c8670529821f237c6bedf25db083c8c3fa))
* add implementation for SiatBoletoAereoService ([93f9cef](https://github.com/ron86i/go-siat/commit/93f9cefd21c6d6577f20672724423bd9ec0f4e13))
* add implementation for SiatEntidadFinancieraService ([7e04f42](https://github.com/ron86i/go-siat/commit/7e04f42a95f3593bc95b1424536d07d90fb1050a))
* add implementation for SiatRecepcionComprasService ([4bd42fa](https://github.com/ron86i/go-siat/commit/4bd42fa0c3c3aec15f58fad5b108ba782513a2ea))
* add implementation for SiatServicioBasicoService ([b8d6607](https://github.com/ron86i/go-siat/commit/b8d66078fb129ba0c642a392a8aa8b17100f692e))
* add map to simplify SIAT error summation and codes ([2bc8da6](https://github.com/ron86i/go-siat/commit/2bc8da6a76e38f325815f1fc9dd6b07fb121d3aa))
* add missing WithCufd method to RecepcionAnexosBuilder ([9f15dd4](https://github.com/ron86i/go-siat/commit/9f15dd420b6223b03f5ff1083df217666c5ebf5d))
* add new service documentoAjuste ([8e50dee](https://github.com/ron86i/go-siat/commit/8e50dee69ebb22b119b16cee526342b522b3338a))
* add nota conciliacion ([16e5309](https://github.com/ron86i/go-siat/commit/16e5309f83ef2eb565d59f4b4f622abcb95d2aa8))
* add nota credito debito ([f0ad318](https://github.com/ron86i/go-siat/commit/f0ad318ba17d0ad6f58ccc9dd8c8d94417907395))
* add nota credito debito ice ([8d851e3](https://github.com/ron86i/go-siat/commit/8d851e341a8723d57c644ec58731e6f5fbeb6cb8))
* add nota fiscal debito ([020e57d](https://github.com/ron86i/go-siat/commit/020e57d5de19ce09cf06db84e241a86af551b27a))
* add nota_conciliacion_test.go ([9aaef57](https://github.com/ron86i/go-siat/commit/9aaef578322889b0c830e425302c26e9d3216673))
* add nota_credito_debito_ice_test.go ([647043b](https://github.com/ron86i/go-siat/commit/647043bcfbfc643aece79507b1f2954ca0a3d0cc))
* add nota_credito_debito_test.go ([eb43653](https://github.com/ron86i/go-siat/commit/eb436532e38d2561ba783c5b166d55050205dc3d))
* add nota_fiscal_credito_debito_test.go ([9d252d0](https://github.com/ron86i/go-siat/commit/9d252d0a712d1fb1642fb3c4646b105dfae48434))
* add port for SiatBoletoAereoService ([b0dcbd5](https://github.com/ron86i/go-siat/commit/b0dcbd5f430f50aaebc549ba107123de6ad52eae))
* add port for SiatEntidadFinancieraService ([d03fa6a](https://github.com/ron86i/go-siat/commit/d03fa6a02181db9140288916c6b45efb1a3e712a))
* add port for SiatRecepcionComprasService ([98f1790](https://github.com/ron86i/go-siat/commit/98f17907961d60da73389e073d1e468afc94364a))
* add port for SiatServicioBasicoService ([8676876](https://github.com/ron86i/go-siat/commit/8676876b793c39b18f63bbcc878cb7f8fc44337f))
* add Recepcion Compras service builders and models ([e056562](https://github.com/ron86i/go-siat/commit/e056562b97ae3cec944bdcba73dbed7dd734c101))
* add Servicio Basico service builders and models ([8146a94](https://github.com/ron86i/go-siat/commit/8146a946f4af6db0d50f9c797363a713c23aeeb0))
* add SiatTelecomunicacionesService ([2e34f42](https://github.com/ron86i/go-siat/commit/2e34f42cbf1a114e09a54d92fe5e7fda0a94953a))
* add SincronizarFechaHora(...) for SiatSincronizacionService ([f2191d5](https://github.com/ron86i/go-siat/commit/f2191d5aa40203ad513aba8227b04166b7bc123a))
* add telecomunicaciones builders ([bc4b016](https://github.com/ron86i/go-siat/commit/bc4b016164482b1df9838389308563fc0e3dd1ad))
* add unified error handling and export utils ([a9216a2](https://github.com/ron86i/go-siat/commit/a9216a253f4d438caa5703cd553ca0234412faaa))
* add XML signing utility and upgrade go to 1.25 ([3ca0e00](https://github.com/ron86i/go-siat/commit/3ca0e0019e9da231f1af43e5c082d9033d4fab95))
* Added functions to interface SiatDocumentoAjusteService ([8cc514f](https://github.com/ron86i/go-siat/commit/8cc514f389e65c7eaee9ea87fe7731bdef5722f0))
* agregar modelos Tasas y Bonificaciones ([12d6eba](https://github.com/ron86i/go-siat/commit/12d6eba43d7ef2c669c71e8255eafebf9255fb15))
* agregar modelos Tasas y Bonificaciones ([cc52479](https://github.com/ron86i/go-siat/commit/cc524795e439fa401dcff1d77e86be30fe6ad95d))
* agregar modelos Tasas y Bonificaciones ([8f82bc9](https://github.com/ron86i/go-siat/commit/8f82bc9d43ab84812c2a491e761ab0a73b76501b))
* agregar modelos Tasas y Bonificaciones ([f14c94d](https://github.com/ron86i/go-siat/commit/f14c94dd538818693b89511bba6e663a4ceff706))
* **ajustes:** agrega modalidad a builders ([e2b64e7](https://github.com/ron86i/go-siat/commit/e2b64e7f17bf5e2c39834dcbcdba70bcf39b2543))
* añadir Builders para registro/cierre de puntos de venta y consulta de eventos significativos ([92af955](https://github.com/ron86i/go-siat/commit/92af95542e155137434bdf7f5f0e0b449d3a72e7))
* **codigos:** Add builders for SolicitudListaCufdDto, SolicitudListaCuisDto to assign the list to WithDatosSolicitud ([93e61fd](https://github.com/ron86i/go-siat/commit/93e61fd1b0c1903dd11faec4195371865d75bc4d))
* **core/port:** introduce internal Config struct and update port interfaces ([6f404a8](https://github.com/ron86i/go-siat/commit/6f404a8ac49e00bbb854d98ff8c6c2b14c3a39fb))
* **core:** update internal ports and domain models for SIAT ([5814b30](https://github.com/ron86i/go-siat/commit/5814b30f00a6b4d0eb27a7ddfa650423ee54a6e0))
* **domain:** add domain structures for all 35 SIAT sectors ([3144429](https://github.com/ron86i/go-siat/commit/3144429e0283b8d978130a72dcc6446556bc2829))
* **domain:** add Servicio Basico (Sector 13) invoice model ([3472f54](https://github.com/ron86i/go-siat/commit/3472f54e1916a71dea3fb4515ea6a8fb8240ccd4))
* **domain:** add Telecomunicaciones (Sector 22) invoice model ([031932e](https://github.com/ron86i/go-siat/commit/031932e232bcf86f828236f1917842be153f3897))
* **domain:** add Telecomunicaciones Zona Franca (Sector 49) invoice model ([78bf096](https://github.com/ron86i/go-siat/commit/78bf0961599fa61238715e8b31c819a0c6092bb0))
* **facturacion:** renombra emisión masiva ([49cf4b6](https://github.com/ron86i/go-siat/commit/49cf4b6127fdcae4b048cf1578e1cde35e8b112c))
* **facturacion:** tipar recepción fiscal ([4cd2ba7](https://github.com/ron86i/go-siat/commit/4cd2ba7aa2d51156928119784ec617742348b888))
* **facturas:** agregar metadatos fiscales ([d48e432](https://github.com/ron86i/go-siat/commit/d48e432c75a7a246b881e5b7b34d6b2bff4af4b4))
* **facturas:** Automatic assign CodigoDocumentoSector ([a415261](https://github.com/ron86i/go-siat/commit/a4152618370ce834a29406c8be301afd91d12d9f))
* **facturas:** Automatic assign CodigoDocumentoSector ([530c893](https://github.com/ron86i/go-siat/commit/530c8934b997f1bbd1cf66713efa1601bf08ce91))
* flujo completo de facturación con Builders, generación de CUF, firma XML y ejemplo de integración ([f760c0c](https://github.com/ron86i/go-siat/commit/f760c0cf4587f7e22220c635e1d8578d23fb956d))
* implement Boleto Aereo (Sector 30) invoicing support ([c53414d](https://github.com/ron86i/go-siat/commit/c53414d6670d37d66d1a0668ef598b2db0557c35))
* implement domain models for SIAT document adjustment services and communication verification ([b300e6c](https://github.com/ron86i/go-siat/commit/b300e6c67ecdaf17724431d0c0672800d46ac3cd))
* implement Entidad Financiera (Sector 15) invoicing support ([d791631](https://github.com/ron86i/go-siat/commit/d7916310be55218772170d202b6f40a5f3ef8a89))
* implement model and domain for Nota de Crédito-Débito con Descuento ([9e28b26](https://github.com/ron86i/go-siat/commit/9e28b261e42f4c934bba8903c12f25e451e35d8a))
* implement ports.SiatDocumentoAjusteService ([d254585](https://github.com/ron86i/go-siat/commit/d25458592952d8b70f8b5e705ca41297fe60394f))
* implementar adaptador de servicio para facturación de códigos ([b222892](https://github.com/ron86i/go-siat/commit/b2228925191740b17aeaf930673893b6f59eafcc))
* implementar Builder genérico para catálogos y reforzar encapsulamiento con interfaces opacas ([b084d30](https://github.com/ron86i/go-siat/commit/b084d301ea4a9a66155e293195d57a9dd5834862))
* implementar Builder pattern e interfaces opacas para solicitudes de CUIS, CUFD y validación de NIT ([de5f473](https://github.com/ron86i/go-siat/commit/de5f473d893b112ee57fae5a6153909e294819ae))
* implementar caso de prueba de 500 facturas en contingencia y corregir validaciones SIAT ([3c7a504](https://github.com/ron86i/go-siat/commit/3c7a504ae83869986c6431a44294923062ded7d5))
* implementar servicio de operaciones SIAT e integración de logs ([83885de](https://github.com/ron86i/go-siat/commit/83885de8004b7880c39757eda1b810ec4c244d57))
* implementar servicio de operaciones SIAT e integración de logs ([7931bce](https://github.com/ron86i/go-siat/commit/7931bce4a60c1cc9556b7ece9d5f7f82168b6deb))
* implementar servicio de operaciones SIAT e integración de logs ([f0155cf](https://github.com/ron86i/go-siat/commit/f0155cf7197c9d0208b83d55c5c06e454395ca26))
* implementar servicio de sincronización de catálogos y paramétricas ([ee476ce](https://github.com/ron86i/go-siat/commit/ee476ce56aa5f32beec2d025da8fe0fd37f97708))
* **invoices:** add sectors 19 and 38, hydrocarbons with and without IEHD ([c170a20](https://github.com/ron86i/go-siat/commit/c170a20f3147821eacee39c12034523f063be510))
* lanzamiento inicial de sdk ([3767dc5](https://github.com/ron86i/go-siat/commit/3767dc5d596adc568b2aacbe0eab40d447127e61))
* metadatos fiscales y recepción tipada SIAT ([fa77ce1](https://github.com/ron86i/go-siat/commit/fa77ce1cbdeb10ecdcde202ca0d65855c19459a7))
* **models:** add builders for Servicio Basico (Sector 13) invoice ([b32a972](https://github.com/ron86i/go-siat/commit/b32a97206bf6cc099611464c01c3bb520b529416))
* **models:** add builders for Telecomunicaciones (Sector 22) invoice ([49f33c2](https://github.com/ron86i/go-siat/commit/49f33c2e59383a29bb5f27661abce01ea4291c30))
* **models:** add builders for Telecomunicaciones Zona Franca (Sector 49) invoice ([b86dd1d](https://github.com/ron86i/go-siat/commit/b86dd1d32355498f14edebe1737a8b2617e1d321))
* **models:** add sectoral invoice models (Dutty Free, Hotel, Lubricantes, etc.) ([81e0ed5](https://github.com/ron86i/go-siat/commit/81e0ed5ace4d438d2cf00e69cad5d9a8cfc33563))
* **models:** add standardized electronic invoice builders ([bd4ed80](https://github.com/ron86i/go-siat/commit/bd4ed80156d4aec69105f1762ef589b7396c01e1))
* **models:** agrega confirmación de compra ([ce0c16a](https://github.com/ron86i/go-siat/commit/ce0c16a01eb12756c64ede8502d3ab6fa23bbc72))
* **models:** agrega defaults masivos ([4fa91d7](https://github.com/ron86i/go-siat/commit/4fa91d7fdda2158c8728c55973d9c8bc553d7cef))
* **models:** implement builders and unit tests for all 35 sector invoices ([c8e9847](https://github.com/ron86i/go-siat/commit/c8e98473eef12ce4ea2ae9a6ead133cab3385ebf))
* **models:** implement Sector 40 (Servicio Basico Zona Franca) domain models and builders ([95418bc](https://github.com/ron86i/go-siat/commit/95418bcc90bc66b025a9daf44c8fb0f960cc30c0))
* **models:** update base SIAT models (Codigos, CompraVenta) ([d231319](https://github.com/ron86i/go-siat/commit/d2313199c6989ff90990d9a9a2c7d9f6edd44130))
* **models:** update public models and utilities for sector invoices ([44d1a2f](https://github.com/ron86i/go-siat/commit/44d1a2fdd1c21919e9a1840ca590615dd533a7cf))
* new Builder for RecepcionDocumentoAjuste, AnulacionDocumentoAjuste, ReversionAnulacionDocumentoAjuste, VerificacionEstadoDocumentoAjustea and VerificarComunicacionDocumentoAjuste ([735c022](https://github.com/ron86i/go-siat/commit/735c0229477e159f69045b5b0b84e00649df98e1))
* new func Float64Round for rounds a float64 ([96b0e80](https://github.com/ron86i/go-siat/commit/96b0e80a6250149dd549d7fca840c669708ff006))
* new url to service ([cb2b55f](https://github.com/ron86i/go-siat/commit/cb2b55f50841bc427664ecb361c33dccc31d0f5a))
* **ports:** add interface SiatTelecomunicacionesService ([ed91425](https://github.com/ron86i/go-siat/commit/ed914252723140cd536c79dc24ece6fea451f62b))
* Se agrega nuevos structs para compra_venta ([1b27496](https://github.com/ron86i/go-siat/commit/1b27496067a2be0e86ae6d99d0fade1f9dfe16e3))
* **sincronizacion:** add struct SincronizarFechaHora, SincronizarFechaHoraResponse ([3d33ce4](https://github.com/ron86i/go-siat/commit/3d33ce440e2bf2fec311c88cc8e9159202a2f42f))
* **testing:** add test for SiatDocumentoAjusteService ([3dbb524](https://github.com/ron86i/go-siat/commit/3dbb524273bfebfe2fd779f102f7452afd2a5d66))
* unify SIAT service message domain and add robust response validation helper ([a676394](https://github.com/ron86i/go-siat/commit/a6763949b5896f2ee2adb5c8d99d995a4e75bcb6))
* updated entry point for SIAT services ([20ed518](https://github.com/ron86i/go-siat/commit/20ed51831dc668e8731276a277e13fcfa83eb0b2))
* upgrade builders ([9aa4a76](https://github.com/ron86i/go-siat/commit/9aa4a7653f15c8b293ed10df47fa214284d7f7f5))
* **utils:** centralize Round logic into pkg/utils/math.go ([d0a6b16](https://github.com/ron86i/go-siat/commit/d0a6b16073dfb334e81567d46cc7fc8ac333ffe7))
* **utils:** enhance CUF generator with CUFParams and fluent builder API ([2a2d560](https://github.com/ron86i/go-siat/commit/2a2d560c8404ddfda71d24c72196cde5ef4026ad))
* **utils:** formatea XML exportado ([4790b0f](https://github.com/ron86i/go-siat/commit/4790b0f61f2ff4680a65f8197e1e3c57e84e90e8))
* **v0.6.0:** stabilize sectoral invoices and global refactor to english ([6aac584](https://github.com/ron86i/go-siat/commit/6aac584992a3600d6983ea65d1cecde7f8c34724))


### Bug Fixes

* arreglos menores ([991b35a](https://github.com/ron86i/go-siat/commit/991b35a495cde539c1c9ee90cd94c9b7f3e5b5ed))
* cambios menores ([8dffeca](https://github.com/ron86i/go-siat/commit/8dffecacadc9595ae4fd3666f32ccdb4b3b7ac86))
* **ci:** elimina etiqueta inválida ([410fbd8](https://github.com/ron86i/go-siat/commit/410fbd8ae9cca0671b2b32e34f688c2f897c107f))
* **computarizada:** correct builder names and invoice response formats ([13e8f8e](https://github.com/ron86i/go-siat/commit/13e8f8eb58eadb9481d8725d58730847d6909c9f))
* corrección de errores de importación ([097eb07](https://github.com/ron86i/go-siat/commit/097eb07b73b47df78324683f7f416af76548c574))
* correct error classification ranges and add missing 973 error code ([c077707](https://github.com/ron86i/go-siat/commit/c077707431c64877236e98c7dfcb07f7b7cae931))
* **datatype:** add marshal and unmarshal for TimeSiat ([eec10bb](https://github.com/ron86i/go-siat/commit/eec10bbe48d5694a443b38ab0d83d1f83d3c479f))
* **dependabot:** usa etiqueta Go existente ([ace20e6](https://github.com/ron86i/go-siat/commit/ace20e6a3528a9afa3c571c17c24edf4ef7674c4))
* **deps:** bump github.com/beevik/etree from 1.6.0 to 1.7.0 ([79a0b3c](https://github.com/ron86i/go-siat/commit/79a0b3ce8c7923389790b7b73c4fde60fa13d089))
* **deps:** bump github.com/beevik/etree from 1.6.0 to 1.7.0 ([4dd3ca0](https://github.com/ron86i/go-siat/commit/4dd3ca0c5c34be92db8324a3e8b9a1e61ac3fa3a))
* **deps:** bump github.com/beevik/etree from 1.7.0 to 1.7.1 ([23e27b4](https://github.com/ron86i/go-siat/commit/23e27b460c28225012c9b4e857dad681df1b0ea3))
* **deps:** bump github.com/beevik/etree from 1.7.0 to 1.7.1 ([feb0e48](https://github.com/ron86i/go-siat/commit/feb0e48a903799a8c4c7d4da68a6f01a8f5e0f46))
* **deps:** bump github.com/russellhaering/goxmldsig from 1.6.0 to 1.6.1 ([37f4461](https://github.com/ron86i/go-siat/commit/37f446103c0a4b813d2f49fa22daad47b4301a24))
* **deps:** bump github.com/russellhaering/goxmldsig from 1.6.0 to 1.6.1 ([88891a9](https://github.com/ron86i/go-siat/commit/88891a92121152d72a768c9b2bbfd7c6df290b12))
* **deps:** bump github.com/stretchr/testify from 1.11.1 to 1.12.1 ([cf43b18](https://github.com/ron86i/go-siat/commit/cf43b189f5b220aef41a51b5939a00da7b688311))
* **deps:** bump github.com/stretchr/testify from 1.11.1 to 1.12.1 ([4f473df](https://github.com/ron86i/go-siat/commit/4f473df830fc64fe083839de415e5565ce9cfb1b))
* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([fdf1d9c](https://github.com/ron86i/go-siat/commit/fdf1d9cc8dc3a4e2c70208f7dd8faf77fd626a26))
* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([edb5929](https://github.com/ron86i/go-siat/commit/edb5929f27ee7bfcd4df5ad1a908b599f16c8c69))
* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([8a672cd](https://github.com/ron86i/go-siat/commit/8a672cd6f83b64b2d157dc342fe02549771096cf))
* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([9ce2ad5](https://github.com/ron86i/go-siat/commit/9ce2ad5f22e7c3ac3d6363e46e30ab312d85a687))
* **facturacion:** restaura API EnLote ([418e129](https://github.com/ron86i/go-siat/commit/418e129134303a74a467501c04f7b3bad2e1e042))
* **invoices:** expone errores JSON libres ([e83afaa](https://github.com/ron86i/go-siat/commit/e83afaa0271dca266c16f93a411411a9354006d5))
* make numeroFactura type consistent (int64) across builders and GenerarCUF ([6e7b07c](https://github.com/ron86i/go-siat/commit/6e7b07ca407a1f9a9ff95352622c238ee7a2bdc1))
* mejora soporte SOAP y documentos de ajuste ([0ba720d](https://github.com/ron86i/go-siat/commit/0ba720dbe6b1aa870ebacdc2efd9962bbda5abf9))
* **models:** agregar modalidad comisionista ([721aa74](https://github.com/ron86i/go-siat/commit/721aa74c64c41053cba8934d24642b18c6b04884))
* **models:** agregar modalidad comisionista ([59fe8a6](https://github.com/ron86i/go-siat/commit/59fe8a65a1a6c77873233e169ccc57fe9cb146dd))
* **models:** incluir modalidad en punto de venta ([f5fb059](https://github.com/ron86i/go-siat/commit/f5fb05990371fd9c62bb90af74857ec9088f4a5f))
* **models:** incluir modalidad en punto de venta ([4205747](https://github.com/ron86i/go-siat/commit/42057479ad38aeb4899a02b848f0ddae5df8038c))
* **operaciones:** incluir punto de venta cero ([46d70b9](https://github.com/ron86i/go-siat/commit/46d70b9ac00d45a1da8e5c5ed8ed3600da2dd45b))
* panic on json.Marshal failure in ComercialExportacion builders ([961af22](https://github.com/ron86i/go-siat/commit/961af226332ca50b53730d1948b1c07caba13f69))
* prevent silent dropping of HTTP middlewares when transport is nil ([e659855](https://github.com/ron86i/go-siat/commit/e6598555e34dd49594bb5f89fb9e1bcec0bbf28e))
* remove dead requestWrapper type; correct inflated method count in docs ([6115a47](https://github.com/ron86i/go-siat/commit/6115a47f5c49ee4db346195e7d3a3f57562ec920))
* rename xml tag to XMLName ([4389618](https://github.com/ron86i/go-siat/commit/4389618d1f6c64db3cb1c5aa4b58ba8541af4f3a))
* resolve sectoral XSD validation and CUF generation issues ([cefb40a](https://github.com/ron86i/go-siat/commit/cefb40a95924bca388f4cc8f0546151e44b1d941))
* **soap:** declara namespace xsi ([ccb0d79](https://github.com/ron86i/go-siat/commit/ccb0d797449d285c02995a831e84e612e3f86a1b))
* the Build() don't return nil ([98c796d](https://github.com/ron86i/go-siat/commit/98c796d2b45ac796584391d1b35a70064ec31d9b))
* the builders don't accept any ([dadba7d](https://github.com/ron86i/go-siat/commit/dadba7d66cdf892324f99a13e77d6cc83875a7ce))
* the builders don't accept any ([01b6c4b](https://github.com/ron86i/go-siat/commit/01b6c4b9f1bc0e97de5beb865bd94200ea13067c))
* the builders don't accept any ([5361361](https://github.com/ron86i/go-siat/commit/5361361df313ca7b2348065c55abf4a96122e5a7))
* the builders don't accept any ([30893e8](https://github.com/ron86i/go-siat/commit/30893e88e8a918a6545eb488acd01c60853af1d3))
* update the url to service ([6cea366](https://github.com/ron86i/go-siat/commit/6cea366511e5dd61678b08271555c129a4c510fb))
* use typed NetworkError instead of standard url.Error in SOAP requests ([7057bfa](https://github.com/ron86i/go-siat/commit/7057bfad08019fc95af0f6e1eb3179c11f0d9de6))
* **utils:** drop redundant gzip/base64 roundtrip in ExportTarGz ([ed9550d](https://github.com/ron86i/go-siat/commit/ed9550d6123f861b36c3a62dda9cf4a030d2aa00))


### Performance Improvements

* allow unlimited MaxConnsPerHost for high concurrency ([32780f6](https://github.com/ron86i/go-siat/commit/32780f6a4b9655c74507fffdd84742fc6a0c7211))
* **facturacion:** optimiza lotes firmados ([cd7f90e](https://github.com/ron86i/go-siat/commit/cd7f90e06996792e45363b8d02c02d60a1c3dbf8))
* increase MaxIdleConnsPerHost to 100 for better pooling with SIAT ([1241d4b](https://github.com/ron86i/go-siat/commit/1241d4bb57eba41feb18b4b7766b2f734aa2c327))
* restore production default limits for HTTP connections and timeouts ([04f05a9](https://github.com/ron86i/go-siat/commit/04f05a978835ae7e3dc568af08a2deb2e9ecd4d4))


### Code Refactoring

* encapsular servicios del SIAT mediante métodos getter ([8d063f8](https://github.com/ron86i/go-siat/commit/8d063f88385ceb0596881ec3b29d299a879d76eb))
* rename Bon to Bonificaciones in builders for clarity naming ([713d9b9](https://github.com/ron86i/go-siat/commit/713d9b93fe0b41fcdca6a0cd56838f44d6312036))

## [2.4.0](https://github.com/ron86i/go-siat/compare/v2.3.0...v2.4.0) (2026-09-13)


### Features

* **models:** agrega confirmación de compra ([ce0c16a](https://github.com/ron86i/go-siat/commit/ce0c16a01eb12756c64ede8502d3ab6fa23bbc72))
* **models:** agrega defaults masivos ([4fa91d7](https://github.com/ron86i/go-siat/commit/4fa91d7fdda2158c8728c55973d9c8bc553d7cef))
* **utils:** formatea XML exportado ([4790b0f](https://github.com/ron86i/go-siat/commit/4790b0f61f2ff4680a65f8197e1e3c57e84e90e8))


### Bug Fixes

* **dependabot:** usa etiqueta Go existente ([ace20e6](https://github.com/ron86i/go-siat/commit/ace20e6a3528a9afa3c571c17c24edf4ef7674c4))
* **invoices:** expone errores JSON libres ([e83afaa](https://github.com/ron86i/go-siat/commit/e83afaa0271dca266c16f93a411411a9354006d5))


### Performance Improvements

* **facturacion:** optimiza lotes firmados ([cd7f90e](https://github.com/ron86i/go-siat/commit/cd7f90e06996792e45363b8d02c02d60a1c3dbf8))

## [2.3.0](https://github.com/ron86i/go-siat/compare/v2.2.1...v2.3.0) (2026-09-11)


### Features

* **ajustes:** agrega modalidad a builders ([e2b64e7](https://github.com/ron86i/go-siat/commit/e2b64e7f17bf5e2c39834dcbcdba70bcf39b2543))


### Bug Fixes

* mejora soporte SOAP y documentos de ajuste ([0ba720d](https://github.com/ron86i/go-siat/commit/0ba720dbe6b1aa870ebacdc2efd9962bbda5abf9))
* **soap:** declara namespace xsi ([ccb0d79](https://github.com/ron86i/go-siat/commit/ccb0d797449d285c02995a831e84e612e3f86a1b))

## [2.2.0](https://github.com/ron86i/go-siat/compare/v2.1.3...v2.2.0) (2026-08-30)


### Features

* **facturas:** agrega metadatos fiscales para los builders de documentos sectoriales ([d48e432](https://github.com/ron86i/go-siat/commit/d48e432c75a7a246b881e5b7b34d6b2bff4af4b4))
* **facturacion:** tipa la recepción fiscal y permite exigir firma XML electrónica ([4cd2ba7](https://github.com/ron86i/go-siat/commit/4cd2ba7aa2d51156928119784ec617742348b888))


### Bug Fixes

* **models:** incluye modalidad en puntos de venta ([4205747](https://github.com/ron86i/go-siat/commit/42057479ad38aeb4899a02b848f0ddae5df8038c))
* **models:** agrega modalidad para punto de venta comisionista ([59fe8a6](https://github.com/ron86i/go-siat/commit/59fe8a65a1a6c77873233e169ccc57fe9cb146dd))
* **operaciones:** incluye el punto de venta cero en solicitudes de cierre ([46d70b9](https://github.com/ron86i/go-siat/commit/46d70b9ac00d45a1da8e5c5ed8ed3600da2dd45b))

## [2.1.3](https://github.com/ron86i/go-siat/compare/v2.1.2...v2.1.3) (2026-08-22)


### Bug Fixes

* **deps:** bump github.com/beevik/etree from 1.7.0 to 1.7.1 ([23e27b4](https://github.com/ron86i/go-siat/commit/23e27b460c28225012c9b4e857dad681df1b0ea3))
* **deps:** bump github.com/beevik/etree from 1.7.0 to 1.7.1 ([feb0e48](https://github.com/ron86i/go-siat/commit/feb0e48a903799a8c4c7d4da68a6f01a8f5e0f46))
* **deps:** bump github.com/stretchr/testify from 1.11.1 to 1.12.1 ([cf43b18](https://github.com/ron86i/go-siat/commit/cf43b189f5b220aef41a51b5939a00da7b688311))
* **deps:** bump github.com/stretchr/testify from 1.11.1 to 1.12.1 ([4f473df](https://github.com/ron86i/go-siat/commit/4f473df830fc64fe083839de415e5565ce9cfb1b))
* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([8a672cd](https://github.com/ron86i/go-siat/commit/8a672cd6f83b64b2d157dc342fe02549771096cf))
* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([9ce2ad5](https://github.com/ron86i/go-siat/commit/9ce2ad5f22e7c3ac3d6363e46e30ab312d85a687))

## [2.1.2](https://github.com/ron86i/go-siat/compare/v2.1.1...v2.1.2) (2026-08-15)


### Bug Fixes

* **deps:** bump github.com/russellhaering/goxmldsig from 1.6.0 to 1.6.1 ([37f4461](https://github.com/ron86i/go-siat/commit/37f446103c0a4b813d2f49fa22daad47b4301a24))
* **deps:** bump github.com/russellhaering/goxmldsig from 1.6.0 to 1.6.1 ([88891a9](https://github.com/ron86i/go-siat/commit/88891a92121152d72a768c9b2bbfd7c6df290b12))

## [2.1.1](https://github.com/ron86i/go-siat/compare/v2.1.0...v2.1.1) (2026-07-30)


### Bug Fixes

* panic on json.Marshal failure in ComercialExportacion builders ([961af22](https://github.com/ron86i/go-siat/commit/961af226332ca50b53730d1948b1c07caba13f69))
* remove dead requestWrapper type; correct inflated method count in docs ([6115a47](https://github.com/ron86i/go-siat/commit/6115a47f5c49ee4db346195e7d3a3f57562ec920))

## [2.1.0](https://github.com/ron86i/go-siat/compare/v2.0.2...v2.1.0) (2026-07-27)


### Features

* **invoices:** add sectors 19 and 38, hydrocarbons with and without IEHD ([c170a20](https://github.com/ron86i/go-siat/commit/c170a20f3147821eacee39c12034523f063be510))


### Bug Fixes

* **utils:** drop redundant gzip/base64 roundtrip in ExportTarGz ([ed9550d](https://github.com/ron86i/go-siat/commit/ed9550d6123f861b36c3a62dda9cf4a030d2aa00))

## [2.0.2](https://github.com/ron86i/go-siat/compare/v2.0.1...v2.0.2) (2026-07-15)


### Bug Fixes

* **deps:** bump github.com/beevik/etree from 1.6.0 to 1.7.0 ([79a0b3c](https://github.com/ron86i/go-siat/commit/79a0b3ce8c7923389790b7b73c4fde60fa13d089))
* **deps:** bump github.com/beevik/etree from 1.6.0 to 1.7.0 ([4dd3ca0](https://github.com/ron86i/go-siat/commit/4dd3ca0c5c34be92db8324a3e8b9a1e61ac3fa3a))

## [2.0.1](https://github.com/ron86i/go-siat/compare/v2.0.0...v2.0.1) (2026-07-09)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([fdf1d9c](https://github.com/ron86i/go-siat/commit/fdf1d9cc8dc3a4e2c70208f7dd8faf77fd626a26))
* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([edb5929](https://github.com/ron86i/go-siat/commit/edb5929f27ee7bfcd4df5ad1a908b599f16c8c69))
