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

	_ "github.com/astaxie/beego"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

// Variables globales
var (
	resStatus    string
	resBody      []byte
	savepostgres map[string]interface{}
	Id           float64
)

// Estructuras de prueba
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

type Parametrica struct {
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

// Configuración de salida de Godog
var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
	Paths:  []string{"features"},
}

// Formato de fecha de ejemplo
const especificacion = "Jan 2, 2006 at 3:04pm (MST)"

// TestMain adapta ejecución con go test
func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		ScenarioInitializer:  FeatureContext,
		TestSuiteInitializer: InitializeTestSuite,
		Options:              &opt,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

// Inicialización de la suite de pruebas
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	// Se podría agregar limpieza o inicialización global aquí si se requiere
}

// Inicialización de la aplicación
func init() {
	run_bee()
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

// Levanta la API para pruebas
func run_bee() {
	parametros := "CUMPLIDOS_CRUD_HTTP_PORT=" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") +
		" CUMPLIDOS_CRUD_PGUSER=" + os.Getenv("CUMPLIDOS_CRUD_PGUSER") +
		" CUMPLIDOS_CRUD_PGPASS=" + os.Getenv("CUMPLIDOS_CRUD_PGPASS") +
		" CUMPLIDOS_CRUD_RUN_MODE=" + os.Getenv("CUMPLIDOS_CRUD_RUN_MODE") +
		" CUMPLIDOS_CRUD_PGURLS=" + os.Getenv("CUMPLIDOS_CRUD_PGURLS") +
		" CUMPLIDOS_CRUD_PGDB=" + os.Getenv("CUMPLIDOS_CRUD_PGDB") +
		" CUMPLIDOS_CRUD_PGSCHEMA=" + os.Getenv("CUMPLIDOS_CRUD_PGSCHEMA") + " bee run"

	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}
	time.Sleep(5 * time.Second)
	deleteFile("script.sh")
	wg.Done()
}

func deleteFile(path string) {
	if err := os.Remove(path); err != nil {
		fmt.Println("no se pudo eliminar el archivo:", err)
	}
}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()
	if err != nil {
		fmt.Println("error ejecutando comando:", err)
	}
	fmt.Printf("%s\n", out)
	wg.Done()
}

func getPages(ruta string) []byte {
	raw, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return raw
}

func iSendRequestToWhereBodyIsJson(method, endpoint, bodyreq string) error {
	var url string
	if method == "GET" || method == "POST" {
		url = "http://localhost:" + os.Getenv("CUMPLIDOS_CRUD_HTTP_PORT") + endpoint
	}

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
	fmt.Println("Se ejecuta la petición al endpoint correctamente")
	ioutil.WriteFile("./files/res/res1.json", resBody, 0644)
	return nil
}

func theResponseCodeShouldBe(arg1 string) error {
	if resStatus != arg1 {
		return fmt.Errorf("se esperaba el código %s y se obtuvo %s", arg1, resStatus)
	}
	fmt.Println("El código es el esperado:", arg1)
	return nil
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1, o2 interface{}
	if err := json.Unmarshal([]byte(s1), &o1); err != nil {
		return false, fmt.Errorf("error unmarshalling string 1: %s", err)
	}
	if err := json.Unmarshal([]byte(s2), &o2); err != nil {
		return false, fmt.Errorf("error unmarshalling string 2: %s", err)
	}
	return reflect.DeepEqual(o1, o2), nil
}

func theResponseShouldMatchJson(arg1 string) error {
	pages := getPages(arg1)
	if areEqual, _ := AreEqualJSON(string(pages), string(resBody)); !areEqual {
		return fmt.Errorf("las respuestas no son iguales")
	}
	fmt.Println("Las respuestas son iguales")
	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}
