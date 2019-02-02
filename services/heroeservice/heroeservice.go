package heroeservice

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ronaldadrianruizavila/Curso_Go/API/models/heroes"
)

const APY_KEY = "[APY_KEY]"
const PUBLIC_KEY = "[PUBLIC_KEY]"

var ts = "2022019"

//Devuelve Un heroe
func GetHeroes() {
	var hashcode = (ts + APY_KEY + PUBLIC_KEY)
	// var HASH = md5.Sum(hashcode)

	hasher := md5.New()
	hasher.Write([]byte(hashcode))

	response, err := http.Get("http://gateway.marvel.com/v1/public/characters?orderBy=name&limit=20&ts=" + ts + "&apikey=" + PUBLIC_KEY + "&hash=" + hex.EncodeToString(hasher.Sum(nil)))
	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		results, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(results))
		var responseObject heroes.Response
		json.Unmarshal(results, &responseObject)

		var data = responseObject.Data.Results

		for i := 0; i < len(data); i++ {
			fmt.Println("\n \nPersonaje N: ", i+1)
			fmt.Println("Nombre del Personaje: ", data[i].Name)
			if len(data[i].Description) > 0 {
				fmt.Println("Descripcion del Personaje", data[i].Description, "\n")
			} else {
				fmt.Println("Por el momento no contamos descripcion \n")
			}
		}
	}
}

func GetHeroe(name string) {
	var hashcode = (ts + APY_KEY + PUBLIC_KEY)
	// var HASH = md5.Sum(hashcode)

	hasher := md5.New()
	hasher.Write([]byte(hashcode))

	response, err := http.Get("http://gateway.marvel.com/v1/public/characters?name=" + name + "&orderBy=name&limit=20&ts=" + ts + "&apikey=" + PUBLIC_KEY + "&hash=" + hex.EncodeToString(hasher.Sum(nil)))
	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		results, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(results))
		var responseObject heroes.Response
		json.Unmarshal(results, &responseObject)

		for i := 0; i < len(responseObject.Data.Results); i++ {
			fmt.Println("\n \nPersonaje N: ", i+1)
			fmt.Println("Nombre del Personaje: ", responseObject.Data.Results[i].Name)
			if len(responseObject.Data.Results[i].Description) > 0 {
				fmt.Println("Descripcion del Personaje", responseObject.Data.Results[i].Description, "\n")
			} else {
				fmt.Println("Por el momento no contamos descripcion \n")
			}
		}
	}
}
