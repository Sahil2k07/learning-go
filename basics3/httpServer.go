package basics3

import (
	"encoding/json"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		response := map[string]interface{}{
			"statusCode": 405,
			"success":    false,
			"message":    "Method not Allowed",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"statusCode": 200,
		"success":    true,
		"message":    "Server is running Successfully",
	}

	json.NewEncoder(w).Encode(response)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		response := map[string]interface{}{
			"statusCode": 405,
			"success":    false,
			"message":    "Method not Allowed",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	var user struct {
		Name   string   `json:"name"`
		Email  string   `json:"email"`
		Skills []string `json:"skills"`
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.Email == "" || user.Name == "" || (user.Skills == nil || len(user.Skills) <= 0) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		response := map[string]interface{}{
			"statusCode": 400,
			"success":    false,
			"message":    "Invalid request payload",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"statusCode": 200,
		"success":    true,
		"data":       user,
	}

	json.NewEncoder(w).Encode(response)

}

func httpServer() {
	const port = "localhost:3000"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/user", userHandler)

	log.Printf("Server is running on port %v", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
