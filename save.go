package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type person struct {
	ID    string `json:"ID"`
	Score int    `json:"Score"`
}

func SaveDecodingPerson(file string) interface{} {
	data, _ := ioutil.ReadFile(file) //file = ./savePerson.json se trouve dans le dossier de base

	var P person

	/* data = les valeurs contenus dans le fichier save.json
	ici un exemple de ce qu'il pourrait contenir mais on doit
	mtn savoir append le texte du .json dans la variable data.
	data := []byte(`{
		"ID":"abc",
		"Score": 1000,
	}`)*/

	err := json.Unmarshal(data, &P) //recup les valeurs voulu du fichier
	if err != nil {
		fmt.Println("error Unmarshal json : ", err)
	}
	//creer list avec elemennt different type
	Utilisateur := []interface{}(nil)
	Utilisateur = append(Utilisateur, P.ID)
	Utilisateur = append(Utilisateur, P.Score)
	fmt.Println("ID = ", Utilisateur[0])   // index 0 = ID
	fmt.Println("Score =", Utilisateur[1]) // index 1 = Score
	fmt.Println("json : ", P)
	fmt.Printf("ID : %s\nSCORE : %d\n", P.ID, P.Score)

	return Utilisateur // liste contenant caractéristique ID & Score d'un utilisateur
}

func SaveEncodingPerson(id string, score int) {
	save := person{ID: id, Score: score}                    // ce qu'on veut dans le format json
	jsonFromMarshal, _ := json.MarshalIndent(save, "", "	") //transforme au format json
	f, _ := os.Create("./savePerson.json")                  //
	defer f.Close()
	_, err2 := f.WriteString(string(jsonFromMarshal)) //ecrit dans le fichier
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Game Saved in savePerson.json")
	return
}
