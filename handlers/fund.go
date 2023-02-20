package handlers

import (
	"encoding/json"
	Fundsdto "holyways/dto/fund"
	dto "holyways/dto/result"
	"holyways/models"
	"holyways/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerFund struct {
	FundRepository repositories.FundRepository
}

var path_file = os.Getenv("PATH_FILE")

func HandlerFund(FundRepository repositories.FundRepository) *handlerFund {
	return &handlerFund{FundRepository}
}

func (h *handlerFund) FindFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Funds, err := h.FundRepository.FindFund()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range Funds {
		Funds[i].Thumbnail = path_file + p.Thumbnail
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: Funds}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) GetFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var Fund models.Fund
	Fund, err := h.FundRepository.GetFund(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	Fund.Thumbnail = os.Getenv("PATH_ILE") + Fund.Thumbnail

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: Fund}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) CreateFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("thumbnail")
	filename := dataContex.(string)

	// transaction_id, _ := strconv.Atoi(r.FormValue("transaction_id"))
	request := Fundsdto.FundCreateRequest{
		Title:       r.FormValue("title"),
		Thumbnail:   filename,
		Goal:        r.FormValue("goals"),
		Description: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	Fund := models.Fund{
		Title:       request.Title,
		Thumbnail:   filename,
		Goal:        request.Goal,
		Description: request.Description,
		UserID:      userId,
		// User: models.UserTransactionResponse{},
	}

	data, err := h.FundRepository.CreateFund(Fund)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) UpdateFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(Fundsdto.UpdateFundRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	Fund, err := h.FundRepository.GetFund(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		Fund.Title = request.Title
	}

	if request.Thumbnail != "" {
		Fund.Thumbnail = request.Thumbnail
	}

	if request.Goal != "" {
		Fund.Goal = request.Goal
	}

	if request.Description != "" {
		Fund.Description = request.Description
	}

	data, err := h.FundRepository.UpdateFund(Fund)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFund(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFund) DeleteFund(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	Fund, err := h.FundRepository.GetFund(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FundRepository.DeleteFund(Fund)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseFund(data)}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerFund) GetFundStatus(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var Fund models.Fund
// 	Fund, err := h.FundRepository.GetFund(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 	}

// 	type FundItem struct {
// 		ID           int
// 		DonateAmount string
// 		Status       string
// 		CreatedAt    string
// 		Fund         string
// 	}

// 	var funds []FundItem

// 	for fund_pending, _ := range Fund.Transaction {
// 		fund_status, _ := h.FundRepository.GetFundStatus(Fund.Transaction[fund_pending].ID)
// 		fund_done := FundItem{
// 			ID:           fund_status.ID,
// 			DonateAmount: fund_status.DonateAmount,
// 			Status:       fund_status.Status,
// 		}
// 		if fund_status.Status == "pending" {
// 			funds = append(funds, fund_done)
// 		}
// 	}

// 	type FundItemNumber struct {
// 		Total int
// 	}

// 	fundsNumber := FundItemNumber{
// 		Total: len(funds),
// 	}

// 	Fund.Thumbnail = os.Getenv("PATH_ILE") + Fund.Thumbnail

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: fundsNumber}
// 	json.NewEncoder(w).Encode(response)
// }

func convertResponseFund(u models.Fund) Fundsdto.FundResponse {
	return Fundsdto.FundResponse{
		ID:          u.ID,
		Title:       u.Title,
		Thumbnail:   u.Thumbnail,
		Goal:        u.Goal,
		Description: u.Description,
	}
}
