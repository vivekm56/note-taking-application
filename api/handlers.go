package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/noteTakingApplication/model"
)

var (
	accounts    map[string]model.Signup
	userSID     map[string]string
	notes       []model.Note
	noteCounter uint32
)

func init() {
	accounts = make(map[string]model.Signup)
	userSID = make(map[string]string)
	notes = make([]model.Note, 0)
	noteCounter = 0
}

// ------------------------------Signup-----------------------//
func CreateSignup(w http.ResponseWriter, r *http.Request) {

	var newSignup model.Signup
	err := json.NewDecoder(r.Body).Decode(&newSignup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newSignup.Name == "" || newSignup.Email == "" || newSignup.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if !model.IsValidEmail(newSignup.Email) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !model.IsValidPassword(newSignup.Password) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if _, ok := accounts[newSignup.Email]; ok {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		accounts[newSignup.Email] = newSignup
	}
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// -----------Login------------------------//
func CreateLogin(w http.ResponseWriter, r *http.Request) {

	var newLogin model.Login
	err := json.NewDecoder(r.Body).Decode(&newLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if storedUser, ok := accounts[newLogin.Email]; ok {
		if storedUser.Password == newLogin.Password {
			inputVlues := newLogin.Email + storedUser.Name
			inputBytes := []byte(inputVlues)
			// Calculate the MD5 hash
			hash := md5.Sum(inputBytes)
			// Convert the hash to a hexadecimal string
			sessionID := string(hex.EncodeToString(hash[:]))
			userSID[sessionID] = newLogin.Email
			session := model.SessionID{SID: sessionID}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(session)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else if newLogin.Email == "" || newLogin.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// http.Error(w, err.Error(), http.StatusUnauthorized)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// -----------------------createNotes--------------------------//
func CreateNotes(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SessionID string `json:"sid"`
		Note      string `json:"note"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userSID[requestData.SessionID]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if requestData.SessionID == "" || requestData.Note == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		note := model.Note{
			ID:      getNextNoteID(),
			Note:    requestData.Note,
			Session: requestData.SessionID,
		}

		notes = append(notes, note)

		responseData := struct {
			ID uint32 `json:"id"`
		}{
			ID: note.ID,
		}

		response, err := json.Marshal(responseData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
func getNextNoteID() uint32 {
	noteCounter++
	return noteCounter
}

// ------------------------readNotes---------------//
func GetNotes(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SessionID string `json:"sid"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.SessionID == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if _, ok := userSID[requestData.SessionID]; ok {
		userNotes := make([]model.Note, 0)

		for _, note := range notes {
			if note.Session == requestData.SessionID {
				userNotes = append(userNotes, note)
			}
		}

		responseData := struct {
			Notes []model.Note `json:"notes"`
		}{
			Notes: userNotes,
		}

		response, err := json.Marshal(responseData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}

// -----------------------deleteNote------------------------//
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SessionID string `json:"sid"`
		ID        uint32 `json:"id"`
	}
	index := -1

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.ID == 0 || requestData.SessionID == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if _, ok := userSID[requestData.SessionID]; ok {
		for i, note := range notes {
			if note.ID == requestData.ID && note.Session == requestData.SessionID {
				index = i
				break
			}
		}
		if index == -1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		notes = append(notes[:index], notes[index+1:]...)
		// w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
