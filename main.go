package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
}

func main() {

	var templatePath string
	var jsonPath string
	var outputDir string
	var outputFileNameProp string

	flag.StringVar(&templatePath, "tpl", "", "template file path")
	flag.StringVar(&jsonPath, "json", "", "json file path")
	flag.StringVar(&outputDir, "out", "", "output dir")
	flag.StringVar(&outputFileNameProp, "np", "", "output file name property")

	flag.Parse()

	if templatePath == "" || jsonPath == "" || outputDir == "" || outputFileNameProp == "" {
		fmt.Println("One or more required input is missing. Use -help to see supported options")
		os.Exit(0)
	}

	jsonData, err := readJSON(jsonPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for _, item := range jsonData {

		result, err := generateOutput(templatePath, item)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fileName, err := generateFileName(item, outputFileNameProp)
		err = writeResultToFile(result, outputDir+"/"+fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func readJSON(jsonPath string) ([]map[string]interface{}, error) {
	jsonData, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		return nil, err
	}

	var listOfMaps []map[string]interface{}

	err = json.Unmarshal(jsonData, &listOfMaps)

	if err != nil {
		return nil, err
	}

	return listOfMaps, nil

}

func generateOutput(templatePath string, data map[string]interface{}) (string, error) {

	templateBytes, err := ioutil.ReadFile(templatePath)

	if err != nil {
		return "", err
	}
	tpl := template.New("result")
	tpl.Funcs(funcMap)
	templateText := string(templateBytes)
	tpl, err = tpl.Parse(string(templateText))

	if err != nil {
		return "", err
	}

	w := bytes.Buffer{}
	err = tpl.Execute(&w, data)

	if err != nil {
		return "", err
	}

	return w.String(), nil
}

func writeResultToFile(result, fileName string) error {
	return ioutil.WriteFile(fileName, []byte(result), 777)
}

func generateFileName(item map[string]interface{}, fileNameProp string) (string, error) {

	tpl := template.New("filename")
	tpl.Funcs(funcMap)
	fileName := item[fileNameProp].(string)
	tpl, err := tpl.Parse(fileName)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	err = tpl.Execute(&buf, item)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
