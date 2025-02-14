package det

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Word struct {
	Word          string     `json:"word"`
	Vocabulary    Vocabulary `json:"vocabulary"`
	Vocabulary_id string     `json:"vocabulary_id"`
}
type Vocabulary struct {
	Items []Item `json:"items"`
}
type Item struct {
	PartOfSpeech     string           `json:"part_of_speech"`
	UsageTranslation UsageTranslation `json:"usage_translation"`
}

type UsageTranslation struct {
	Vietnamese string `json:"Vietnamese"`
}

type ResItem struct {
	ID  int    `json:"id"`
	En1 string `json:"en1"`
	En2 string `json:"en2"`
	En3 string `json:"en3"`
	En4 string `json:"en4"`
	Vn1 string `json:"vn1"`
}

func Convert() {
	apiURL := "https://core.goarno.io/flashcards"
	authToken := "eyJhbGciOiJIUzI1NiIsImtpZCI6InNiZ1RrN1dVbGhMdXN1MG4iLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3BuZmtpa2xsZG9pc2hxcW5sb2hlLnN1cGFiYXNlLmNvL2F1dGgvdjEiLCJzdWIiOiIyNzEyZDY5ZC1kNTAwLTQ5YzAtYWM0ZS1kYmE4OGYxNDYwODMiLCJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNzM5MTAxMDM1LCJpYXQiOjE3MzkwOTc0MzUsImVtYWlsIjoiYW5oa2hvaS5ucDk0QGdtYWlsLmNvbSIsInBob25lIjoiIiwiYXBwX21ldGFkYXRhIjp7InByb3ZpZGVyIjoiZ29vZ2xlIiwicHJvdmlkZXJzIjpbImdvb2dsZSJdfSwidXNlcl9tZXRhZGF0YSI6eyJhdmF0YXJfdXJsIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jSXBuTkxSdDA5b0J6UFNBRll3Q2x1YWx6WThEejczMW8ydktzT2dadm5BTGZUUkVBPXM5Ni1jIiwiZW1haWwiOiJhbmhraG9pLm5wOTRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZ1bGxfbmFtZSI6Iktow7RpIE5ndXnhu4VuIiwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50cy5nb29nbGUuY29tIiwibmFtZSI6Iktow7RpIE5ndXnhu4VuIiwicGhvbmVfdmVyaWZpZWQiOmZhbHNlLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jSXBuTkxSdDA5b0J6UFNBRll3Q2x1YWx6WThEejczMW8ydktzT2dadm5BTGZUUkVBPXM5Ni1jIiwicHJvdmlkZXJfaWQiOiIxMTY5NjE2NDI1NzQ4NTQ0OTkwNzIiLCJzdWIiOiIxMTY5NjE2NDI1NzQ4NTQ0OTkwNzIifSwicm9sZSI6ImF1dGhlbnRpY2F0ZWQiLCJhYWwiOiJhYWwxIiwiYW1yIjpbeyJtZXRob2QiOiJwYXNzd29yZCIsInRpbWVzdGFtcCI6MTczODExMjA5NH1dLCJzZXNzaW9uX2lkIjoiMmVmODVhNzMtZTYwNS00OTFhLTllY2EtOTEyOGYwODRkNjFlIiwiaXNfYW5vbnltb3VzIjpmYWxzZX0.fgQUUndbsnQvzmhivkMFYNxgTOQHysFRQHm4OXx_sAg"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add the Authentication header
	req.Header.Add("Authentication", authToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: received status code %d", resp.StatusCode)
	}

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var items []Word
	if err := json.Unmarshal(body, &items); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Stored words
	var storedWords []ResItem
	data, err := ioutil.ReadFile("output.json")
	if err != nil {
		fmt.Println("Failed")
	}
	if err := json.Unmarshal(data, &storedWords); err != nil {
		fmt.Println("Failed")
	}

	outputFile := "output.json"
	addCount := 0
	currentId := len(storedWords) - 1

	var resItem []ResItem
	for i := 0; i < len(items); i++ {
		saved := false
		for y := 0; y < len(storedWords); y++ {
			if storedWords[y].En1 == items[i].Word {
				saved = true
				break
			}
		}

		if !saved {
			vn := ""
			re := regexp.MustCompile(`\((.*?)\)`)
			match := re.FindStringSubmatch(items[i].Vocabulary.Items[0].UsageTranslation.Vietnamese)
			if len(match) > 1 {
				//fmt.Println(match[1]) // Output: sự tinh lọc
				vn = match[1]
			}
			typeW := items[i].Vocabulary.Items[0].PartOfSpeech
			resItem = append(resItem,
				ResItem{
					ID:  currentId + 1,
					En1: items[i].Word,
					En2: "",
					En3: "",
					En4: "",
					Vn1: "(" + typeW + ") " + vn,
				},
			)
			currentId += 1
			addCount += 1
		}

	}

	// Convert items back to JSON
	processedData, err := json.MarshalIndent(append(storedWords, resItem...), "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Write processed JSON to another file
	err = ioutil.WriteFile(outputFile, processedData, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Add %v words\n", addCount)
	fmt.Printf("Processed data written to %s\n", outputFile)
}
