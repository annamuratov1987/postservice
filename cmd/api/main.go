package main

import (
	"flag"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"postservice/delivery"
	"postservice/internal/config"
	"postservice/internal/usecase"
	"postservice/pkg/logger"
	"strconv"
)

var (
	configPath string
	logPath    string
)

func init() {
	flag.StringVar(&configPath, `config-path`, `./configs`, `Configuration file path`)
	flag.StringVar(&logPath, `log-path`, ``, `Configuration log path`)
	flag.Parse()
}

func main() {

	// Initial logger
	lg := newLogger()
	defer func() { _ = lg.Sync() }()

	lg.Info(`Start`)
	defer lg.Info(`Stop`)

	// Initial config
	cfg := newConfig(lg)

	// Start listen port
	address := cfg.Api().GetHost() + ":" + strconv.Itoa(cfg.Api().GetPort())
	listen, err := net.Listen("tcp", address)
	if err != nil {
		lg.Fatal(
			`TcpListener.Start.Error`,
			zap.String(`address`, address),
			zap.Error(err),
		)
	}
	defer listen.Close()

	apiHandler := delivery.NewApiHandler(
		cfg,
		lg,
		usecase.New(cfg, lg, nil),
	)

	// Start server
	lg.Info(`Main.Listening`, zap.String(`address`, address))
	if err := http.Serve(listen, apiHandler.GetHttpHandler()); err != nil {
		log.Fatal(err)
	}
}

// Initial config
func newConfig(lg logger.ILogger) config.IConfig {
	config, err := config.New(configPath)
	if err != nil {
		lg.Fatal(`Config.Error`, zap.Error(err))
	}

	return config
}

// Initial logger
func newLogger() logger.ILogger {
	lg, err := logger.New(logPath)
	if err != nil {
		log.Fatalf(`Start logger: %s\n`, err)
	}

	return lg
}

/*
func LoaderStart(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(configuration.Get("loader.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := protos.NewLoaderServiceClient(conn)

	startRequest := &protos.LoaderRequest{}
	res, err := client.LoaderStart(context.Background(), startRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func LoaderCheck(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(configuration.Get("loader.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := protos.NewLoaderServiceClient(conn)

	checkRequest := &protos.LoaderRequest{}
	res, err := client.LoaderCheck(context.Background(), checkRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(configuration.Get("grud.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := protos.NewGrudServiceClient(conn)
	grudRequest := &protos.GrudRequest{}
	res, err := client.GetAll(context.Background(), grudRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  res.GetStatus().GetName(),
		Message: res.GetStatus().GetMessage(),
		Data:    res.GetData(),
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func Get(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := Request{
		Id:     -1,
		UserId: 0,
		Title:  "",
		Body:   "",
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Id < 0 {
		log.Println("Request field 'id' is not set")
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(configuration.Get("grud.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	grudRequest := &protos.GrudRequest{
		Id:     int64(req.Id),
		UserId: 0,
		Title:  "",
		Body:   "",
	}

	client := protos.NewGrudServiceClient(conn)
	res, err := client.Get(context.Background(), grudRequest)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  res.GetStatus().GetName(),
		Message: res.GetStatus().GetMessage(),
		Data:    res.GetData(),
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := Request{
		Id:     -1,
		UserId: 0,
		Title:  "",
		Body:   "",
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Id < 0 {
		log.Println("Request field 'id' is not set")
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(configuration.Get("grud.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := protos.NewGrudServiceClient(conn)

	grudRequest := &protos.GrudRequest{
		Id:     int64(req.Id),
		UserId: 0,
		Title:  "",
		Body:   "",
	}
	res, err := client.Update(context.Background(), grudRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  res.GetName(),
		Message: res.GetMessage(),
		Data:    nil,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := Request{
		Id:     -1,
		UserId: 0,
		Title:  "",
		Body:   "",
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Id < 0 {
		log.Println("Request field 'id' is not set")
		http.Error(w, "Request field 'id' is not set", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(configuration.Get("grud.address"), grpc.WithInsecure())
	if err != nil {
		log.Println("Dial:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := protos.NewGrudServiceClient(conn)

	grudRequest := &protos.GrudRequest{
		Id:     int64(req.Id),
		UserId: 0,
		Title:  "",
		Body:   "",
	}
	res, err := client.Delete(context.Background(), grudRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  res.GetName(),
		Message: res.GetMessage(),
		Data:    nil,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
*/
