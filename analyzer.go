package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	// "math/rand"
	// "net/http"
	"os"
	// "strings"

	"github.com/jrallison/go-workers"
)

type Analyzer struct {
	exercismHost   string
	analysseurHost string
	Logger         *log.Logger
}

func NewAnalyzer(exercism, analysseur string) *Analyzer {
	return &Analyzer{
		exercismHost:   exercism,
		analysseurHost: analysseur,
		Logger:         log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

type codePayload struct {
	TrackID       string            `json:"track_id"`
	SolutionFiles map[string]string `json:"solution_files"`
	Error         string            `json:"error"`
}

type analysisResult struct {
	Type string   `json:"type"`
	Keys []string `json:"keys"`
}

type analysisPayload struct {
	Results []analysisResult `json:"results"`
	Error   string           `json:"error"`
}

type commentBody struct {
	Comment string `json:"comment"`
}

func getCodeAndLanguageFromAPI(submissionKey string) *codePayload {
	// Step 1: get code + language from exercism.io api
	// url := fmt.Sprintf("%s/api/v1/submissions/%s", analyzer.exercismHost, submissionKey)
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	lgr.Printf("cannot prepare request to %s - %s\n", url, err)
	// 	return
	// }
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	lgr.Printf("request to %s failed - %s\n", url, err)
	// 	return
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	lgr.Printf("cannot read response - %s\n", err)
	// }
	// resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	lgr.Printf("%s responded with status %d - %s\n", url, resp.StatusCode, string(body))
	// 	return
	// }

	// var cp codePayload
	// if err := json.Unmarshal(body, &cp); err != nil {
	// 	lgr.Printf("%s - %s\n", submissionKey, err)
	// 	return
	// }

	// if cp.TrackID != "ruby" {
	// 	lgr.Printf("Skipping code in %s\n", cp.TrackID)
	// 	return
	// }

	// var solutions []string
	// for _, solution := range cp.SolutionFiles {
	// 	solutions = append(solutions, solution)
	// }

	return &codePayload{}
}

func readParamsFromMessage(msg *workers.Msg) (string, string, error) {
	submissionKey, err := msg.Args().GetIndex(0).String()
	submissionType, err := msg.Args().GetIndex(1).String()

	return submissionKey, submissionType, err
}

func (analyzer *Analyzer) process(msg *workers.Msg) {
	lgr := analyzer.Logger

	submissionKey, submissionType, err := readParamsFromMessage(msg)
	if err != nil {
		lgr.Printf("unable to extract params from message: %s\n", err)
		return
	}

	// The different cases will populate this
	commentToSubmit := &commentBody{}

	// check if this is a special submission that does not need analysis
	switch submissionType {
	case "hello_world":
		fmt.Println("selecting hello world comment")
		c := NewComment("hello_world", "hello_world", "")
		commentToSubmit.Comment = c.String()

	default:
		fmt.Println("grabbing submission info from API")
		fmt.Println("analyzing code")

		// _ = getCodeAndLanguageFromAPI(submissionKey)

		fmt.Println("select comment based on analysis result")
		commentToSubmit.Comment = "Some other comment"

		// Step 2: submit code to analysseur
		// url = fmt.Sprintf("%s/analyze/%s", analyzer.analysseurHost, cp.TrackID)
		// codeBody := struct {
		// 	Code string `json:"code"`
		// }{
		// 	strings.Join(solutions, "\n"),
		// }
		// codeBodyJSON, err := json.Marshal(codeBody)
		// if err != nil {
		// 	lgr.Printf("%s - %s\n", submissionKey, err)
		// 	return
		// }

		// req, err = http.NewRequest("POST", url, bytes.NewReader(codeBodyJSON))
		// if err != nil {
		// 	lgr.Printf("%s - cannot prepare request to %s - %s\n", submissionKey, url, err)
		// 	return
		// }

		// resp, err = http.DefaultClient.Do(req)
		// if err != nil {
		// 	lgr.Printf("%s - request to %s failed - %s\n", submissionKey, url, err)
		// 	return
		// }

		// body, err = ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	lgr.Printf("%s - failed to fetch submission - %s\n", submissionKey, err)
		// 	return
		// }
		// resp.Body.Close()
		// if resp.StatusCode != http.StatusOK {
		// 	lgr.Printf("%s - %s responded with status %d - %s\n", submissionKey, url, resp.StatusCode, string(body))
		// 	return
		// }

		// var ap analysisPayload
		// err = json.Unmarshal(body, &ap)
		// if err != nil {
		// 	lgr.Printf("%s - %s\n", submissionKey, err)
		// 	return
		// }

		// if ap.Error != "" {
		// 	lgr.Printf("analysis api is complaining about %s - %s\n", submissionKey, ap.Error)
		// 	return
		// }

		// if len(ap.Results) == 0 {
		// 	// no feedback, bailing
		// 	return
		// }

		// sanity := log.New(os.Stdout, "SANITY: ", log.Ldate|log.Ltime|log.Lshortfile)
		// for _, result := range ap.Results {
		// 	for _, key := range result.Keys {
		// 		sanity.Printf("%s : %s - %s\n", submissionKey, result.Type, key)
		// 	}
		// }

		// // Step 3: Find comments based on analysis result
		// for _, result := range ap.Results {
		// 	for _, key := range result.Keys {
		// 		c := NewComment(cp.TrackID, result.Type, key)
		// 		b, err := c.Bytes()
		// 		if err != nil {
		// 			lgr.Printf("We probably need to add a comment at %s - %s\n", c.Path(), err)
		// 		}
		// 		if len(b) > 0 {
		// 			comments = append(comments, b)
		// 		}
		// 	}
		// }
	}

	// if len(comments) == 0 {
	// 	// no comments, bailing
	// 	return
	// }

	// comment := comments[rand.Intn(len(comments))]
	// commentBody := &commentBody{Comment: "Comment Text"}

	submitComment(submissionKey, commentToSubmit)
}

// add info about the comment
func formatComment(commentBody *commentBody) *commentBody {
	experiment := "_This is an automated nitpick. [Read more](http://exercism.io/rikki) about this experiment._"
	s := fmt.Sprintf("%s\n----\n%s", string(commentBody.Comment), experiment)

	commentBody.Comment = s
	return commentBody
}

func submitComment(submissionKey string, commentBody *commentBody) {
	formattedComment := formatComment(commentBody)

	commentBodyJSON, err := json.Marshal(formattedComment)
	if err != nil {
		log.Println(err)
		return
	}

	url := fmt.Sprintf(
		"%s/api/v1/submissions/%s/comments?shared_key=%s",
		*exercismFlag,
		submissionKey,
		authentication.Key())

	fmt.Printf("%s\n", commentBodyJSON)
	fmt.Println(url)

	// req, err = http.NewRequest("POST", url, bytes.NewReader(commentBodyJSON))
	// if err != nil {
	// 	lgr.Printf("cannot prepare request to %s - %s\n", url, err)
	// 	return
	// }
	// resp, err = http.DefaultClient.Do(req)
	// if err != nil {
	// 	lgr.Printf("request to %s failed - %s\n", url, err)
	// 	return
	// }

	// body, err = ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if resp.StatusCode != http.StatusNoContent {
	// 	lgr.Printf("%s responded with status %d - %s\n", url, resp.StatusCode, string(body))
	// 	return
	// }
}
