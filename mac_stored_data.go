package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	did "github.com/jlinclabs/go-jlinc-did"
)

func MacStoredData(service string) (string, error) {
	const datadirname string = "JLINCAuthorizer"
	const datafilename string = "diddata.json"
	dataFile := make(map[string]DidData)
	var zcap string

	home, _ := os.UserHomeDir()
	datadir := home + "/Library/Application Support/" + datadirname
	files, err := ioutil.ReadDir(datadir)
	if err != nil { // there's no data directory
		if _, patherr := err.(*os.PathError); patherr {
			// create the directory and the data file
			err := os.Mkdir(datadir, 0700)
			if err != nil && !os.IsExist(err) {
				return zcap, err
			}
			data, err := newServiceData(service, dataFile, datadir+"/"+datafilename)
			if err != nil {
				return zcap, err
			}
			log.Printf("data: %v, error: %v\n", data, err)
			// return invocation JWT here
		} else {
			return zcap, err //generic ioutil.ReadDir error
		}
	} // else found the data directory

	for _, file := range files {
		if file.Name() == "diddata.json" {
			// read diddata.json and look for didData[service]
			jsonData, err := ioutil.ReadFile(datadir + "/diddata.json")
			if err != nil {
				return zcap, err
			}
			err = json.Unmarshal(jsonData, &dataFile)
			if data, ok := dataFile[service]; ok {
				log.Printf("service: %s, data: %v\n", service, data)
				// return invocation JWT here
				return zcap, nil
			}
		}
	}

	// did not find service data in data file
	data, err := newServiceData(service, dataFile, datadir+"/"+datafilename)
	log.Printf("data: %v, error: %v\n", data, err)
	// return invocation JWT here
	return zcap, nil
}

func newServiceData(service string, datafile map[string]DidData, datafileaddr string) (DidData, error) {
	// registers a new DID and writes it to the diddata.json file, creating the file if needed
	var forSaving DidData
	jsonForSaving, err := did.RegisterDID(didServerUrl)
	if err != nil {
		return DidData{}, err
	}

	err = json.Unmarshal([]byte(jsonForSaving), &forSaving)
	if err != nil {
		return DidData{}, err
	}

	datafile[service] = forSaving
	//marshal dataFile and write it to file
	dataFileJson, err := json.Marshal(datafile)
	if err != nil {
		return DidData{}, err
	}
	err = os.WriteFile(datafileaddr, dataFileJson, 0600)
	if err != nil {
		return DidData{}, err
	}

	return forSaving, nil
}
