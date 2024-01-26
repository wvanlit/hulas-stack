package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	AppDir        = "./app"
	Port          = ":8080"
	APIPath       = "/api/"
	LuaExtension  = ".lua"
	LuaJIT        = "luajit"
	StartupScript = "./lib/startup.lua"
	PagesDir      = "/pages"
	LoggerPrefix  = "[HULAS server] "
)

var logger = log.New(log.Writer(), LoggerPrefix, log.LstdFlags)

func main() {
	logger.Println("Starting HULAS...")
	setupServer()
	if err := startServer(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}

func setupServer() {
	printCWD()
	printFileTree()
	runStartupScript()

	fs := http.FileServer(http.Dir(AppDir + PagesDir))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc(APIPath, apiHandler)
}

func startServer() error {
	logger.Printf("Running on http://localhost%s", Port)
	return http.ListenAndServe(Port, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != APIPath {
		serveLuaScripts(w, r)
	} else {
		http.Error(w, "Invalid path", http.StatusBadRequest)
	}
}

func serveLuaScripts(w http.ResponseWriter, r *http.Request) {
	script := urlToLuaScript(r.URL.Path)
	if !fileExistsInApp(script) {
		http.NotFound(w, r)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpError(w, "Error reading request body", err)
		return
	}
	defer r.Body.Close()

	output, err := runScript(script, r.Method, string(body))
	if err != nil {
		httpError(w, "Error executing script", err)
		return
	}

	fmt.Fprint(w, output)
}

func urlToLuaScript(urlPath string) string {
	return urlPath + LuaExtension
}

func fileExistsInApp(path string) bool {
	_, err := os.Stat(AppDir + path)
	return !os.IsNotExist(err)
}

func runScript(script, method, body string) (string, error) {
	cmd := exec.Command(LuaJIT, "."+script)
	cmd.Dir = AppDir
	cmd.Env = []string{"REQUEST_METHOD=" + method, "REQUEST_BODY=" + strings.TrimSpace(body)}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing script '%s': %v, output: %s", script, err, output)
	}
	return string(output), nil
}

func runStartupScript() {
	runAndLogOutput(exec.Command(LuaJIT, StartupScript), "startup script")
}

func printFileTree() {
	runAndLogOutput(exec.Command("tree"), "file tree")
}

func printCWD() {
	cwd, _ := os.Getwd()
	logger.Printf("Serving files from CWD: %s", cwd)
}

func runAndLogOutput(cmd *exec.Cmd, context string) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Printf("Error running %s: %v", context, err)
	}
	logger.Printf("%s output: %s", context, strings.TrimSpace(string(output)))
}

func httpError(w http.ResponseWriter, msg string, err error) {
	logger.Printf("%s: %v", msg, err)
	http.Error(w, msg, http.StatusInternalServerError)
}
