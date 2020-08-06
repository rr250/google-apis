/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [START sheets_quickstart]
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets "+drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	// srv, err := drive.New(client)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Drive client: %v", err)
	// }

	// r, err := srv.Files.List().PageSize(1).
	// 	Fields("nextPageToken, files(id, name)").Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve files: %v", err)
	// }
	// fmt.Println("Files:")
	// if len(r.Files) == 0 {
	// 	fmt.Println("No files found.")
	// } else {
	// 	for _, i := range r.Files {
	// 		fmt.Printf("%s (%s)\n", i.Name, i.Id)
	// 	}
	// }

	// resp, _ := srv.Files.Get("1lWCjxxur2oUZqQb_Q6X00ufw4cq3pzA7gjHwtv2QIXA").Do()
	// fmt.Println(resp)
	// permission := &drive.Permission{
	// 	Type: "anyone",
	// 	Role: "reader",
	// }
	// resp1, err := srv.Permissions.Create("1lWCjxxur2oUZqQb_Q6X00ufw4cq3pzA7gjHwtv2QIXA", permission).Do()
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// fmt.Println(resp1)

	// srv1, err := sheets.New(client)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve Sheets client: %v", err)
	// }

	// // Prints the names and majors of students in a sample spreadsheet:
	// // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	// spreadsheetId := "1lWCjxxur2oUZqQb_Q6X00ufw4cq3pzA7gjHwtv2QIXA"
	// readRange := "Sheet1!A:Z"
	// resp, err := srv1.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve data from sheet: %v", err)
	// }
	// //fmt.Println(resp)
	// resp1, _ := srv1.Spreadsheets.Get(spreadsheetId).Do()
	// fmt.Println(resp1.SpreadsheetUrl)
	// if len(resp.Values) == 0 {
	// 	fmt.Println("No data found.")
	// } else {
	// 	fmt.Println("Name, Major:")
	// 	for _, row := range resp.Values {
	// 		// Print columns A and E, which correspond to indices 0 and 4.
	// 		fmt.Printf("%s, %s\n", row[0], row[1])
	// 	}
	// }
	// valueRange := &sheets.ValueRange{
	// 	Values: [][]interface{}{
	// 		{
	// 			"Name",
	// 			"Email",
	// 			"Resume",
	// 			"Experience",
	// 			"Reason",
	// 			"FilledAt",
	// 		},
	// 		{
	// 			"Rishabh Ranjan",
	// 			"rrrishabh7@gmail.com",
	// 			"google.com",
	// 			1,
	// 			"test",
	// 			"2020-07-24",
	// 		},
	// 	},
	// }
	// resp2, err := srv1.Spreadsheets.Values.Append(spreadsheetId, readRange, valueRange).ValueInputOption("USER_ENTERED").Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve data from sheet: %v", err)
	// }
	// fmt.Println(resp2)
	// ss := &sheets.Spreadsheet{
	// 	Properties: &sheets.SpreadsheetProperties{
	// 		Title: "123",
	// 	},
	// }
	// resp3, err3 := srv1.Spreadsheets.Create(ss).Do()
	// if err3 != nil {
	// 	log.Fatalf("Unable to retrieve data from sheet: %v", err3)
	// }
	// fmt.Println(resp3)
}

// [END sheets_quickstart]
