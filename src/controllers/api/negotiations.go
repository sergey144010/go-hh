package api

import (
	"encoding/json"
	// AppTokenService "go-hh/src/services/appToken"
	"net/http"
	// VacancyService "go-hh/src/services/vacancy"

	"fmt"
	"io"
	// "log"
	"os"
	"bytes"
	"mime/multipart"
)

type NegotiationsRequestData struct {
	Message string `json:"message"`
	ResumeId string `json:"resume_id"`
	VacancyIdList []string `json:"vacancy_id_list"`
	AccessToken string `json:"access_token"`
}

func Send(w http.ResponseWriter, r *http.Request) {
	var request NegotiationsRequestData
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// data, err := json.Marshal(request)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// reader := bytes.NewReader(data)
	// io.Copy(os.Stdout, reader)

	// client := &http.Client{}
	// req, _ := request()
	// resp, _ := client.Do(req)

	for _, vacancyId := range request.VacancyIdList {
		buffer := &bytes.Buffer{}
		mpw := multipart.NewWriter(buffer)
	
		resumeIdWriter, err := mpw.CreateFormField("resume_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = resumeIdWriter.Write([]byte(request.ResumeId))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		// messageWriter, err := mpw.CreateFormField("message")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
		// _, err = messageWriter.Write([]byte(request.Message))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	
		vacancyIdWriter, errV := mpw.CreateFormField("vacancy_id")
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
		}
		_, errV = vacancyIdWriter.Write([]byte(vacancyId))
		if errV != nil {
			http.Error(w, errV.Error(), http.StatusInternalServerError)
		}
	
		err = mpw.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		client := &http.Client{}
	
		req, err := http.NewRequest("POST", "https://api.hh.ru/negotiations", buffer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		req.Header.Add("Content-Type", mpw.FormDataContentType())
		req.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
		req.Header.Add("Authorization", "Bearer " + request.AccessToken)
	
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
		defer resp.Body.Close()
	
		fmt.Println("Vacancy Id:", vacancyId)
		fmt.Println("Status Code:", resp.StatusCode)
	}

	// io.Copy(os.Stdout, resp.Body)

	// body, errR := io.ReadAll(resp.Body)
	// if errR != nil {
	// 	http.Error(w, errR.Error(), http.StatusInternalServerError)
	// }
}