package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const APP_DIR = "./app"
const PORT = ":8080"

var logger *log.Logger

func serveLuaScripts(w http.ResponseWriter, req *http.Request) {
	script := urlToLuaScript(req)

	if !fileExists(script) {
		logger.Println("File not found:", script)
		http.NotFound(w, req)
	}

	output := runScript(script)

	if output == "" {
		http.Error(w, "Error executing script", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", output)
}

func urlToLuaScript(req *http.Request) string {
	// /api/[script] -> ./app/[script].lua
	path := req.URL.Path
	path = strings.Replace(path, "/api/", "", 1)
	return APP_DIR + "/" + path + ".lua"
}

func fileExists(path string) bool {
	_, err := exec.LookPath(path)
	return err == nil
}

func runScript(script string) string {
	cmd := exec.Command("luajit", script)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Println("Error executing script '", script, "':", err)
		logger.Println("Output:", string(output))
		return ""
	}
	return string(output)
}

func printFileTree() {
	cmd := exec.Command("tree")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Println("Error printing file tree:", err)
	}
	logger.Println("File tree:")
	logger.Println(string(output))
}

func main() {
	logger = log.New(log.Writer(), "[HULAS server] ", log.LstdFlags)

	logger.Println("Starting HULAS...")
	logger.Println("Running on ", "http://localhost"+PORT)

	cwd, _ := os.Getwd()
	logger.Println("Serving files from CWD", cwd)

	printFileTree()

	fileServer := http.FileServer(http.Dir(APP_DIR))

	http.Handle("/", http.StripPrefix("/", fileServer))

	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path != "/api/" {
			serveLuaScripts(w, r)
		} else {
			http.Error(w, "Invalid path", http.StatusBadRequest)
		}
	})

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}