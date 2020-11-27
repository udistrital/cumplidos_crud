package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	
	"strings"
	"sync"
	"testing"
	"time"

	_"github.com/astaxie/beego"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

)

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

//@resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

var savepostgres map[string]interface{}

var Id float64

//@Parametrica estructura de las tablas parametricas
type EstadoPagoMensual struct {
	Id                int
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	NumeroOrden       float64
	Activo            bool
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

type Parametrica struct { //PagoMensual
	Id                     int
	NumeroContrato         string
	VigenciaContrato       float64
	Mes                    float64
	DocumentoPersonaId     string
	EstadoPagoMensualId    *EstadoPagoMensual
	DocumentoResponsableId string
	CargoResponsable       string
	Ano                    float64
	Activo                 bool
	FechaCreacion          time.Time
	FechaModificacion      time.Time
}

//@opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

//@especificacion estructura de la fecha
const especificacion = "Jan 2, 2006 at 3:04pm (MST)"

//@TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		//Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)

}

//@init inicia la aplicacion para realizar los test
func init() {
	//	gen_files()
	run_bee()
	//pasa las banderas al comando godog
	godog.BindFlags("godog.", flag.CommandLine, &opt)

}

//@run_bee activa el servicio de la api para realizar los test
func run_bee() {

	//parametros := "CUMPLIDOS_CRUD_HTTP_PORT=" + beego.AppConfig.String("httpport") + " CUMPLIDOS_CRUD_PGUSER=" + beego.AppConfig.String("PGuser") + " CUMPLIDOS_CRUD_PGPASS=" + beego.AppConfig.String("PGpass") + " CUMPLIDOS_CRUD_PGHOST=" + beego.AppConfig.String("PGhost") + " CUMPLIDOS_CRUD_PGPORT=" + beego.AppConfig.String("PGport") + " CUMPLIDOS_CRUD_PGDB=" + beego.AppConfig.String("PGdb") + " CUMPLIDOS_CRUD_PGSCHEMA=" + beego.AppConfig.String("PGschema") + " bee run"
	parametros := "CUMPLIDOS_CRUD_HTTP_PORT=" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") + " CUMPLIDOS_CRUD_PGUSER=" + os.Getenv("CUMPLIDOS_CRUD_PGUSER") + " CUMPLIDOS_CRUD_PGPASS=" + os.Getenv("CUMPLIDOS_CRUD_PGPASS") +" CUMPLIDOS_CRUD_RUN_MODE=" + os.Getenv("CUMPLIDOS_CRUD_RUN_MODE") + " CUMPLIDOS_CRUD_PGURLS=" + os.Getenv("CUMPLIDOS_CRUD_PGURLS") + " CUMPLIDOS_CRUD_PGDB=" + os.Getenv("CUMPLIDOS_CRUD_PGDB") + " CUMPLIDOS_CRUD_PGSCHEMA=" + os.Getenv("CUMPLIDOS_CRUD_PGSCHEMA") + " bee run"

	//fmt.Println(parametros)
	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	fmt.Println("*********")
	fmt.Println(commands)
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Termino ejecución del script, para correr beego")
	fmt.Println("________________________________________________")
	deleteFile("script.sh")
	wg.Done()
}

func deleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		fmt.Errorf("no se pudo eliminar el archivo")
	}

}

//@exe_cmd ejecuta comandos en la terminal
func exe_cmd(cmd string, wg *sync.WaitGroup) {
	//fmt.Println(cmd)
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()	
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Println("----------------------------")
	fmt.Println(out)
	fmt.Printf("%s", out)
	wg.Done()
}

/*
//@toJson convierte string en JSON
func toJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
*/
//@getPages convierte en un tipo el json
func getPages(ruta string) []byte {
	raw, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []byte
	c = raw
	return c
}

//@iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(method, endpoint, bodyreq string) error {
	var url string
	if method == "GET" || method == "POST" {
		url = "http://localhost:" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") + endpoint

	} /*else {
		if method == "PUT" || method == "DELETE" {
			str := strconv.FormatFloat(Id, 'f', 5, 64)
			url = "http://localhost:" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") + endpoint + "/" + str

		}
	}
	if method == "GETID" {
		method = "GET"
		str := strconv.FormatFloat(Id, 'f', 5, 64)
		url = "http://localhost:" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") + endpoint + "/" + str

	}*/

	pages := getPages(bodyreq)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(pages))
	req.Header.Set("Content-Type", "application/json")


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyr, _ := ioutil.ReadAll(resp.Body)
	resStatus = resp.Status
	resBody = bodyr
	fmt.Println("Se ejecuta la peticion al endpoint correctamente")
	ioutil.WriteFile("./files/res/res1.json", resBody, 0644)
	/*if method == "POST" && resStatus == "200 OK" {
		ioutil.WriteFile("./files/res/res1.json", resBody, 0644)
		//json.Unmarshal([]byte(bodyr), &savepostgres)
		//Id = savepostgres["Id"].(float64)

	}*/
	return nil

}

//@theResponseCodeShouldBe valida el codigo de respuesta
func theResponseCodeShouldBe(arg1 string) error {
	if resStatus != arg1 {
		return fmt.Errorf("Se esperaba el codigo de respuesta .. %s .. y se obtuvo el codigo de respuesta .. %s .. ", arg1, resStatus)
	}else{
		fmt.Println("El codigo es el esperado ",arg1,"= ", resStatus)
	}
	return nil
}

//@AreEqualJSON comparar dos JSON si son iguales retorna true de lo contrario false
func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}
	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

//@theResponseShouldMatchJson valida el JSON de respuesta
func theResponseShouldMatchJson(arg1 string) error {
	pages := getPages(arg1)
	if areEqual,_:=AreEqualJSON(string(pages), string(resBody));areEqual==false{		
		return fmt.Errorf("Las respuestas no son iguales")		
	}else{
		fmt.Println("Las respuestas son iguales")
	}
	
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)

}
