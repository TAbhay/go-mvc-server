package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type TestResult struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Passed      bool   `json:"passed"`
}

type TestResultsFile struct {
	Items []TestResult `json:"items"`
}

type ReportData struct {
	TestResults []TestResult
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>QA Test Results</title>
    <style>
        table {
            width: 100%;
            border-collapse: collapse;
        }
        table, th, td {
            border: 1px solid black;
        }
        th, td {
            padding: 10px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
        .pass {
            background-color: #d4edda;
        }
        .fail {
            background-color: #f8d7da;
        }
    </style>
</head>
<body>
    <h1>QA Test Results</h1>
    <table>
        <thead>
            <tr>
                <th>Type</th>
                <th>Name</th>
                <th>Description</th>
                <th>Passed</th>
            </tr>
        </thead>
        <tbody>
            {{ range .TestResults }}
            <tr class="{{ if .Passed }}pass{{ else }}fail{{ end }}">
                <td>{{ .Type }}</td>
                <td>{{ .Name }}</td>
                <td>{{ .Description }}</td>
                <td>{{ .Passed }}</td>
            </tr>
            {{ end }}
        </tbody>
    </table>
</body>
</html>
`

func main() {
	// Specify the directory containing the JSON files
	dir := "../reports"

	files, err := os.ReadDir("../reports")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var allResults []TestResult

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
				continue
			}

			var resultsFile TestResultsFile
			if err := json.Unmarshal(data, &resultsFile); err != nil {
				log.Printf("Failed to parse JSON in file %s: %v", file.Name(), err)
				continue
			}

			allResults = append(allResults, resultsFile.Items...)
		}
	}

	reportData := ReportData{TestResults: allResults}

	tmpl, err := template.New("report").Parse(htmlTemplate)
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
