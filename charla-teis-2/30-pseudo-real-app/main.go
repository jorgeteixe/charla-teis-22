package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jorgeteixe/charla-teis-22/charla-teis-2/30-pseudo-real-app/config"
)

type Status struct {
	Status string `json:"status"`
}

type Alumno struct {
	ID     int    `json:"id"`
	Login  string `json:"login"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT `id`, `login`, `nombre`, `email` FROM `alumno`")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	alumnos := make([]Alumno, 0)

	for rows.Next() {
		alum := Alumno{}
		err := rows.Scan(&alum.ID, &alum.Login, &alum.Nombre, &alum.Email)

		if err != nil {
			panic(err)
		}

		alumnos = append(alumnos, alum)
	}

	json.NewEncoder(w).Encode(alumnos)
}

func getByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	row := config.DB.QueryRow("SELECT `id`, `login`, `nombre`, `email` FROM `alumno` WHERE `id` = ?", ID)

	alum := Alumno{}

	err = row.Scan(&alum.ID, &alum.Login, &alum.Nombre, &alum.Email)

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(alum)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	alum := Alumno{}

	err := json.NewDecoder(r.Body).Decode(&alum)

	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec("INSERT INTO `alumno` (`login`, `nombre`, `email`) VALUES (?, ?, ?)", alum.Login, alum.Nombre, alum.Email)

	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	alum.ID = int(lastId)

	json.NewEncoder(w).Encode(alum)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// _, err := config.DB.Query("SELECT 1")

	// if err != nil {
	// 	panic(err)
	// }

	json.NewEncoder(w).Encode(Status{Status: "ok"})
}

func workHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 2)
	close(done)

	json.NewEncoder(w).Encode(Status{Status: "ok"})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ok", okHandler).Methods("GET")

	r.HandleFunc("/work", workHandler).Methods("GET")

	r.HandleFunc("/{id:[0-9]+}", getByIdHandler).Methods("GET")

	r.HandleFunc("/", getHandler).Methods("GET")

	r.HandleFunc("/", postHandler).Methods("POST")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	println("Started app...")
	log.Fatal(http.ListenAndServe(":8090", logRequest(r)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", requestGetRemoteAddress(r), r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// requestGetRemoteAddress returns ip address of the client making the request,
// taking into account http proxies
func requestGetRemoteAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		return strings.Join(parts, ",")
	}
	return hdrRealIP
}
