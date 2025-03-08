package main

import (
	"bytes"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.parquetenisclube.com.br/booking/srvc.aspx/ObtenerCuadro"
	payload := []byte(`{"idCuadro":"5","fecha":"07/03/2025","key":"eNEe29kXfZB5XujQIJ1PSycBzUl5pQROoHBe07jvCDjglvloaKL4JA=="}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Origin", "http://www.parquetenisclube.com.br")
	req.Header.Set("Referer", "http://www.parquetenisclube.com.br/Booking/Grid.aspx")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Cookie", "MPOpcionCookie=todos; _ga=GA1.3.1350129106.1739580587; ASP.NET_SessionId=orpacyyanaqz2huzjenkhv55; _gid=GA1.3.17902871.1741389152; i18next=pt-BR")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))



	type Column struct {
		Id              string `json:"Id"`
		TextoPrincipal  string `json:"TextoPrincipal"`
		TextoSecundario string `json:"TextoSecundario"`
		IdImagen        string `json:"IdImagen"`
	}

	type PtcResponse struct {
		D struct {
			Columnas []Column `json:"Columnas"`
		} `json:"d"`
	}

	var response PtcResponse

	jsonerr := json.Unmarshal(body, &response)

	if jsonerr != nil {
		fmt.Print("error", jsonerr)
	}




	for _, column := range response.D.Columnas {
		fmt.Printf("Id: %s\n", column.Id)
		fmt.Printf("TextoPrincipal: %s\n", column.TextoPrincipal)
		fmt.Printf("TextoSecundario: %s\n", column.TextoSecundario)
		fmt.Printf("IdImagen: %s\n", column.IdImagen)
	}


}
