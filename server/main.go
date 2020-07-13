package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"../config"
	"../models"
	"../protos"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

type taskServer struct{}

//ProjectsColl in not  exported and contains projects collection
var ProjectsColl *mongo.Collection

//TasksColl in not  exported and contains tasks collection
var TasksColl *mongo.Collection

//TeamsColl in not  exported and contains teams collection
var TeamsColl *mongo.Collection

//UsersColl in not  exported and contains users collection
var UsersColl *mongo.Collection

//Login logs the user in
func (taskServer) Login(ctx context.Context, in *protos.UserLogin) (*protos.RegisterResponse, error) {
	var user models.Users

	if err := UsersColl.FindOne(ctx, &models.Users{Email: in.GetEmail()}).Decode(&user); err != nil {
		return &protos.RegisterResponse{ID: "", Message: "Invalid Email OR Password"}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(in.GetPassword()), []byte(user.Password)); err != nil {
		return &protos.RegisterResponse{ID: "", Message: "Invalid Email OR Password"}, err
	}
	return &protos.RegisterResponse{ID: user.ID.Hex(), Message: ""}, nil
}

//Signup creates an account
func (taskServer) Signup(ctx context.Context, in *protos.UserSignup) (*protos.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), 10)
	if err != nil {
		return &protos.RegisterResponse{ID: "", Message: "Sorry, something went wrong"}, err
	}

	newUser := models.Users{
		UserName: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: string(hash),
	}

	user, err := UsersColl.InsertOne(ctx, newUser)
	if err != nil {
		return &protos.RegisterResponse{ID: "", Message: "Account already exists"}, err
	}
	if oid, ok := user.InsertedID.(primitive.ObjectID); ok {
		return &protos.RegisterResponse{ID: oid.Hex(), Message: ""}, nil
	}
	return &protos.RegisterResponse{ID: "", Message: "Sorry, something went wrong"}, errors.New("Sorry, something went wrong")
}

//GetUsers returns all the users in th collection
func (taskServer) GetUsers(ctx context.Context, _ *protos.Empty) (*protos.AllUsers, error) {
	userCursor, err := UsersColl.Find(ctx, bson.M{})
	if err != nil {
		return &protos.AllUsers{}, err
	}
	var allUsers protos.AllUsers
	if err = userCursor.All(ctx, &allUsers); err != nil {
		return &protos.AllUsers{}, err
	}

	defer userCursor.Close(ctx)
	return &allUsers, nil

}

//GetTeamMember returns all the members under a team leader
func (taskServer) GetTeamMembers(ctx context.Context, in *protos.Empty) (*protos.AllUsers, error) {
	matchWith := bson.D{{Key: "$match", Value: bson.D{{Key: "leaderID", Value: in.GetID()}}}}
	groupBy := bson.D{{
		Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$leaderID"},
			{Key: "members", Value: bson.D{
				{Key: "$push", Value: "$memberID"},
			}},
		},
	}}

	lookUp := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "members"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "memberData"},
		},
	}}

	teamMemberCursor, err := TeamsColl.Aggregate(ctx, mongo.Pipeline{matchWith, groupBy, lookUp})
	if err != nil {
		return &protos.AllUsers{}, err
	}
	var teamMembers protos.AllUsers
	if err = teamMemberCursor.All(ctx, &teamMembers); err != nil {
		return &protos.AllUsers{}, err
	}

	return &teamMembers, nil
}

//AddTask inserts a new task
func (taskServer) AddTask(ctx context.Context, in *protos.TaskRequest) (*protos.TaskResponse, error) {
	startD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetStartDate())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	dueD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetDueDate())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	assigneeID, err := primitive.ObjectIDFromHex(in.GetAssigneeID())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	projectID, err := primitive.ObjectIDFromHex(in.GetProjectID())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}

	newTask := models.Tasks{
		Subject:      in.GetSubject(),
		Description:  in.GetDescription(),
		Status:       in.GetStatus(),
		Priority:     in.GetPriority(),
		Category:     in.GetCategory(),
		DateCreated:  time.Now(),
		DateModified: time.Now(),
		StartDate:    startD,
		DueDate:      dueD,
		AssigneeID:   assigneeID,
		ProjectID:    projectID,
	}

	_, err = TasksColl.InsertOne(ctx, newTask)
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	return &protos.TaskResponse{Ok: true, Message: ""}, nil
}

//UpdateTask updates an existing task
func (taskServer) UpdateTask(ctx context.Context, in *protos.TaskRequest) (*protos.TaskResponse, error) {
	taskID, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	startD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetStartDate())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}

	dueD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetDueDate())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}

	assigneeID, err := primitive.ObjectIDFromHex(in.GetAssigneeID())
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}

	_, err = TasksColl.UpdateOne(ctx, bson.M{"_id": taskID}, bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "subject", Value: in.GetSubject()},
			{Key: "description", Value: in.GetDescription()},
			{Key: "status", Value: in.GetStatus()},
			{Key: "priority", Value: in.GetPriority()},
			{Key: "category", Value: in.GetCategory()},
			{Key: "startDate", Value: startD},
			{Key: "dueDate", Value: dueD},
			{Key: "assigneeID", Value: assigneeID},
			{Key: "dateModified", Value: time.Now()},
		}},
	})
	if err != nil {
		return &protos.TaskResponse{Ok: false, Message: "Sorry, something went wrong"}, err
	}
	return &protos.TaskResponse{Ok: true, Message: ""}, nil
}

//GetAllTasks returns details about all the tasks
func (taskServer) GetAllTasks(ctx context.Context, _ *protos.Empty) (*protos.AllTaskInfo, error) {
	lookUp := bson.D{{
		Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "assigneeID"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "assignee"},
		},
	}}

	taskCursor, err := TasksColl.Aggregate(ctx, mongo.Pipeline{lookUp})
	if err != nil {
		return &protos.AllTaskInfo{}, err
	}

	var tasks protos.AllTaskInfo

	if err = taskCursor.All(ctx, &tasks); err != nil {
		return &protos.AllTaskInfo{}, err
	}

	defer taskCursor.Close(ctx)
	return &tasks, nil
}

//GetStatusCount returns a map of status and count
func (taskServer) GetStatusCount(ctx context.Context, in *protos.Empty) (*protos.StatusResponse, error) {
	projectID, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return &protos.StatusResponse{}, err
	}

	matchWith := bson.D{{
		Key: "$match", Value: bson.D{
			{Key: "projectID", Value: projectID},
		},
	}}
	groupBy := bson.D{{
		Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$status"},
			{Key: "count", Value: bson.D{
				{Key: "$sum", Value: 1},
			}},
		},
	}}

	resultCursor, err := TasksColl.Aggregate(ctx, mongo.Pipeline{matchWith, groupBy})
	if err != nil {
		return &protos.StatusResponse{}, err
	}

	var result protos.StatusResponse

	if err = resultCursor.All(ctx, &result); err != nil {
		return &protos.StatusResponse{}, err
	}

	defer resultCursor.Close(ctx)
	return &result, nil
}

//UpdateTaskStatus updates the status of the tasks
func (taskServer) UpdateTaskStatus(ctx context.Context, _ *protos.Empty) (*protos.Empty, error) {
	if _, err := TasksColl.UpdateMany(ctx, bson.M{"status": "Open", "startDate": bson.M{"$lte": time.Now()}}, bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "status", Value: "In Progress"},
		},
	}}); err != nil {
		return &protos.Empty{}, err
	}
	if _, err := TasksColl.UpdateMany(ctx, bson.M{"status": "In Progress", "dueDate": bson.M{"$lte": time.Now()}}, bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "status", Value: "Closed"},
		},
	}}); err != nil {
		return &protos.Empty{}, err
	}
	return &protos.Empty{ID: "success"}, nil
}

//GetAllProjects returns details about all the projects
func (taskServer) GetAllProjects(ctx context.Context, _ *protos.Empty) (*protos.ProjectResponse, error) {
	projectCursor, err := ProjectsColl.Find(ctx, bson.M{})
	if err != nil {
		return &protos.ProjectResponse{}, err
	}

	var projects protos.ProjectResponse

	if err = projectCursor.All(ctx, &projects); err != nil {
		return &protos.ProjectResponse{}, err
	}

	defer projectCursor.Close(ctx)
	return &projects, nil
}

func main() {
	client := config.Client
	ctx, cancel := config.GetCtx()
	taskerDB := client.Database("taskerapp")

	defer client.Disconnect(ctx)
	defer cancel()

	ProjectsColl = taskerDB.Collection("projects")
	TasksColl = taskerDB.Collection("tasks")
	TeamsColl = taskerDB.Collection("teams")
	UsersColl = taskerDB.Collection("users")

	server := grpc.NewServer()
	protos.RegisterTaskerAppServer(server, taskServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal("Server starting...", server.Serve(listener).Error())
	}()
	grpcWebServer := grpcweb.WrapServer(server)
	httpServer := &http.Server{
		Addr: ":9001",
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	log.Fatal("Serving Proxy: ", httpServer.ListenAndServe().Error())
}
