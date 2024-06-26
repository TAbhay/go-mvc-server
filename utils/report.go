package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type TestResult struct {
	Type        string `json:"type"`
	Api         string `json:"api"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Passed      bool   `json:"passed"`
}

type TestGroup struct {
	Name  string       `json:"name"`
	Tests []TestResult `json:"tests"`
}

type TestResults struct {
	Tests []TestGroup `json:"tests"`
}

type ReportData struct {
	TestGroups []TestGroup
}

func (tg *TestGroup) UnmarshalJSON(data []byte) error {
	type Alias TestGroup
	aux := &struct {
		Tests interface{} `json:"tests"`
		*Alias
	}{
		Alias: (*Alias)(tg),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.Tests.(type) {
	case []interface{}:
		for _, item := range v {
			switch item := item.(type) {
			case []interface{}:
				for _, subitem := range item {
					var tr TestResult
					b, _ := json.Marshal(subitem)
					if err := json.Unmarshal(b, &tr); err != nil {
						return err
					}
					tg.Tests = append(tg.Tests, tr)
				}
			case map[string]interface{}:
				var tr TestResult
				b, _ := json.Marshal(item)
				if err := json.Unmarshal(b, &tr); err != nil {
					return err
				}
				tg.Tests = append(tg.Tests, tr)
			default:
				return fmt.Errorf("unexpected type in tests field")
			}
		}
	default:
		return fmt.Errorf("unexpected type in tests field")
	}

	return nil
}

func GenerateReport() {
	dir := "reports"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var allTestGroups []TestGroup

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
				continue
			}

			var testResults TestResults
			if err := json.Unmarshal(data, &testResults); err != nil {
				log.Printf("Failed to parse JSON in file %s: %v", file.Name(), err)
				continue
			}

			allTestGroups = append(allTestGroups, testResults.Tests...)
		}
	}
	fmt.Println(allTestGroups)

	reportData := ReportData{TestGroups: allTestGroups}
	htmlTemplate := "templates/template.html"

	tmpl, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	reportFile, err := os.Create("report.html")
	if err != nil {
		log.Fatalf("Failed to create report file: %v", err)
	}
	defer reportFile.Close()

	if err := tmpl.Execute(reportFile, reportData); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	log.Println("Report generated successfully as report.html")
}
