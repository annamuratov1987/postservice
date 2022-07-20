package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"postservice/internal/config"
	postRepo "postservice/internal/repository/post"
	"postservice/internal/usecase"
	"postservice/pkg/logger"
)

type IApiHandler interface {
	LoaderStart(w http.ResponseWriter, r *http.Request)
	LoaderCheck(w http.ResponseWriter, r *http.Request)
	GetHttpHandler() http.Handler
}

type apiHandler struct {
	cfg              config.IConfig
	lg               logger.ILogger
	loaderConnection *grpc.ClientConn
	usecase          usecase.IUseCase
}

func NewApiHandler(cfg config.IConfig, lg logger.ILogger, apiUseCase usecase.IUseCase) IApiHandler {
	return &apiHandler{
		cfg:     cfg,
		lg:      lg,
		usecase: apiUseCase,
	}
}

func (a *apiHandler) GetHttpHandler() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/loader/start", a.LoaderStart).Methods(http.MethodGet)
	router.HandleFunc("/loader/check", a.LoaderCheck).Methods(http.MethodGet)

	router.HandleFunc("/grud/get", a.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/grud/get", a.Get).Methods(http.MethodPost)
	router.HandleFunc("/grud/update", a.Update).Methods(http.MethodPut)
	router.HandleFunc("/grud/delete", a.Delete).Methods(http.MethodDelete)
	return router
}

func (a *apiHandler) writeResponse(w http.ResponseWriter, res ApiResponse) {
	jsonRes, err := json.Marshal(res)
	if err != nil {
		a.lg.Error("ApiHandler.WriteResponse.JsonEncoding.Error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (a *apiHandler) getApiRequest(r *http.Request) (*ApiRequest, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.lg.Error("ApiHandler.getApiRequest.ReadRequestBodyError", zap.Error(err))
		return nil, err
	}

	req := ApiRequest{
		Id:     -1,
		UserId: 0,
		Title:  "",
		Body:   "",
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		a.lg.Error("ApiHandler.getApiRequest.RequestBodyUnmarshalError", zap.Error(err))
		return nil, err
	}

	return &req, nil
}
func (a *apiHandler) LoaderStart(w http.ResponseWriter, r *http.Request) {
	res, err := a.usecase.Api().LoaderStart(r.Context())
	if err != nil {
		a.lg.Error("ApiHandler.LoaderStart.Error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetStatus(),
		Message: "",
		Data:    nil,
	}
	a.writeResponse(w, response)
}

func (a *apiHandler) LoaderCheck(w http.ResponseWriter, r *http.Request) {
	res, err := a.usecase.Api().LoaderCheck(r.Context())
	if err != nil {
		a.lg.Error("ApiHandler.LoaderCheck.Error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetStatus(),
		Message: "",
		Data:    nil,
	}
	a.writeResponse(w, response)
}

func (a *apiHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := a.usecase.Api().GetAll(r.Context())
	if err != nil {
		a.lg.Error("ApiHandler.GetAll.UseCaseGetAllError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetStatus().GetName(),
		Message: res.GetStatus().GetMessage(),
		Data:    res.GetData(),
	}
	a.writeResponse(w, response)
}

func (a *apiHandler) Get(w http.ResponseWriter, r *http.Request) {
	req, err := a.getApiRequest(r)
	if err != nil {
		a.lg.Error("ApiHandler.Get.getApiRequestError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Id < 0 {
		a.lg.Error("ApiHandler.Get.RequestFieldIdNotSetError", zap.Error(err))
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	res, err := a.usecase.Api().Get(r.Context(), req.Id)
	if err != nil {
		a.lg.Error("ApiHandler.Get.UseCaseGetError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetStatus().GetName(),
		Message: res.GetStatus().GetMessage(),
		Data:    res.GetData(),
	}
	a.writeResponse(w, response)
}

func (a *apiHandler) Update(w http.ResponseWriter, r *http.Request) {
	req, err := a.getApiRequest(r)
	if err != nil {
		a.lg.Error("ApiHandler.Update.getApiRequestError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Id < 0 {
		a.lg.Error("ApiHandler.Update.RequestFieldIdNotSetError", zap.Error(err))
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	post := postRepo.Post{
		Id:     req.Id,
		UserId: req.UserId,
		Title:  req.Title,
		Body:   req.Body,
	}
	res, err := a.usecase.Api().Update(r.Context(), &post)
	if err != nil {
		a.lg.Error("ApiHandler.Update.UseCaseGetError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetName(),
		Message: res.GetMessage(),
		Data:    nil,
	}
	a.writeResponse(w, response)
}

func (a *apiHandler) Delete(w http.ResponseWriter, r *http.Request) {
	req, err := a.getApiRequest(r)
	if err != nil {
		a.lg.Error("ApiHandler.Delete.getApiRequestError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Id < 0 {
		a.lg.Error("ApiHandler.Delete.RequestFieldIdNotSetError", zap.Error(err))
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	res, err := a.usecase.Api().Delete(r.Context(), req.Id)
	if err != nil {
		a.lg.Error("ApiHandler.Delete.UseCaseGetError", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ApiResponse{
		Status:  res.GetName(),
		Message: res.GetMessage(),
		Data:    nil,
	}
	a.writeResponse(w, response)
}
