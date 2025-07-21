package web

import (
	"context"
	"encoding/json"
	"fmt"
	"juong/matchbox/matchbox"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	matchbox *matchbox.Matchbox
}

func NewServer(m *matchbox.Matchbox) *Server {
	return &Server{matchbox: m}
}

func (s *Server) Start(ctx context.Context) {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/add_player", s.handleAddPlayer)

	go s.matchbox.RunSimulation(ctx, 10*time.Second)

	fmt.Println("Starting web server on :8080")
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("Status code: %v\n", err)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

type AddPlayerRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Server) handleAddPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AddPlayerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	playerID, err := strconv.ParseUint(req.ID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	player := &matchbox.Player{
		ID:   matchbox.PlayerID(playerID),
		Name: req.Name,
	}
	s.matchbox.AddPlayerToQueue(r.Context(), player)

	w.WriteHeader(http.StatusCreated)
}
