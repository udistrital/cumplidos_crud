Feature: Validate API responses
  CUMPLIDOS_CRUD
  probe JSON reponses



Scenario Outline: To probe response route /tr_aprobacion_masiva_soportes_contratistas      
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method |route                                          |bodyreq                                 |codres          |bodyres                         |
    |POST   |/v1/tr_aprobacion_masiva_soportes_contratistas |./files/req/documentos_correcto.json    |200 OK          |./files/res/soportes_contratistas/res_default.json         |
    |POST   |/v1/tr_aprobacion_masiva_soportes_contratistas |./files/req/vacio.json                  |400 Bad Request |./files/res/res_default_error.json         |
    |POST   |/v1/tr_aprobacion_masiva_soportes_contratistas |./files/req/documentos_inexistente.json |200 OK          |./files/res/soportes_contratistas/res_default_inexistente.json         |