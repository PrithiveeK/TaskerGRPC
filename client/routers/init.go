package routers

import (
	"log"
	"net/http"

	"../../config"
	"../../protos"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

//Router is exported to create server in main function
var Router *mux.Router
var client protos.TaskerAppClient

func init() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = protos.NewTaskerAppClient(conn)

	Router = mux.NewRouter()

	Router.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		ctx, cancel := config.GetCtx()
		defer cancel()
		client.UpdateTaskStatus(ctx, &protos.Empty{})
	})
	Router.HandleFunc("/api/project/all", getAllProjects).Methods("GET")
	Router.HandleFunc("/api/task/all", getAllTasks).Methods("GET")
	Router.HandleFunc("/api/task/add", addTaskForProject).Methods("POST")
	Router.HandleFunc("/api/task/status", getTaskStatusCount).Methods("GET")
	Router.HandleFunc("/api/user/all", getAllUsers).Methods("GET")
	Router.HandleFunc("/api/task/update", updateTask).Methods("PUT")
	Router.HandleFunc("/api/user/signup", createUserAccount).Methods("POST")
	Router.HandleFunc("/api/user/login", logIntoUser).Methods("POST")
	Router.HandleFunc("/api/team", getTeamMembers).Methods("GET")
}
