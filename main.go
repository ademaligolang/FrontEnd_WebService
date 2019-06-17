package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	song "github.com/ademaligolang/Song_Definitions"
	"github.com/gorilla/mux"
)

// Port to listen for this webservice
const port = ":8080"

// Main function
func main() {

	// Using gorilla mux for http routing
	router := mux.NewRouter().StrictSlash(true)

	// Set up the routing options for this webservice
	router.HandleFunc("/AddSongs", AddSongs)
	router.HandleFunc("/GetSongGroups", GetSongGroups)

	// Print any errors to the command line
	error := http.ListenAndServe(port, router)
	fmt.Println(error)
}

// Function to open the json file and add songs to the song webservice
func AddSongs(w http.ResponseWriter, r *http.Request) {

	// Temporary list of songs from the json
	var songs song.Songs

	// Open the json file
	jsonFile, err := os.Open("song_list.json")

	// Make sure the file opened correctly
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		panic(err)
	}

	// Wait til we've closed the json file
	defer jsonFile.Close()

	// Read the json
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Convert the json into objects
	if err := json.Unmarshal(byteValue, &songs); err != nil {
		// We've hit an error so prepare an appropriate response
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Prepare some error variables
	errorFound := false
	errorText := ""

	// Iterate all of the songs and send them to the song webservice
	for _, song := range songs.Songs {
		jsonValue, _ := json.Marshal(song)
		response, err := http.Post("http://localhost:8081/AddSong", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			// We have an error so lets send an appropriate response
			errorText += fmt.Sprintf("%s", err)
			fmt.Printf("The HTTP request failed with error %s\n", err)
			errorFound = true
		} else {
			// No errors so print some info to the command line
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}

	// We found an error so prepare an appropriate response
	if errorFound {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error! %s", errorText)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		// All done!
		fmt.Fprintf(w, "Completed!")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}

}

// Function to get the song groups from the song webservice and print them to the browser
func GetSongGroups(w http.ResponseWriter, r *http.Request) {
	// Send the request on to the Song_WebService
	response, err := http.Get("http://localhost:8081/GetSongGroups")
	if err != nil {
		// Print the error to the console
		fmt.Printf("The HTTP request failed with error %s\n", err)

		// Prepare the response
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error!")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		// Write the get response into the body of this response
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintf(w, "%s", data)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
}
