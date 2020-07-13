package routers

import (
	"encoding/json"
	"net/http"

	"../../config"
	"../../protos"
)

//Users is the json user model
type Users struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

//Projects is the json project model
type Projects struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

//Tasks is the json task model
type Tasks struct {
	ID           string `json:"id"`
	Subject      string `json:"subject"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Priority     int32  `json:"priority"`
	Category     string `json:"category"`
	DateCreated  string `json:"dateCreated"`
	DateModified string `json:"dateModified"`
	StartDate    string `json:"startDate"`
	DueDate      string `json:"dueDate"`
	ProjectID    string `json:"productID"`
	Assignee     Users  `json:"assignee"`
}

func getAllProjects(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	var projects []Projects
	stream, err := client.GetAllProjects(ctx, &protos.Empty{})
	if err != nil {
		json.NewEncoder(res).Encode(projects)
		return
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			break
		}
		p := Projects{
			ID:    result.GetProject().GetID(),
			Title: result.GetProject().GetTitle(),
		}
		projects = append(projects, p)
	}
	json.NewEncoder(res).Encode(projects)
}

func getAllTasks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()

	var tasks []Tasks
	stream, err := client.GetAllTasks(ctx, &protos.Empty{})
	if err != nil {
		json.NewEncoder(res).Encode(tasks)
		return
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			break
		}
		tt := result.GetTask()
		t := Tasks{
			ID:           tt.GetID(),
			Subject:      tt.GetSubject(),
			Description:  tt.GetDescription(),
			Status:       tt.GetStatus(),
			Category:     tt.GetCategory(),
			Priority:     tt.GetPriority(),
			DateCreated:  tt.GetDateCreated(),
			DateModified: tt.GetDateModified(),
			StartDate:    tt.GetStartDate(),
			DueDate:      tt.GetDueDate(),
			ProjectID:    tt.GetProjectID(),
			Assignee: Users{
				ID:       tt.GetAssigneeObj().GetID(),
				Username: tt.GetAssigneeObj().GetUsername(),
			},
		}
		tasks = append(tasks, t)
	}
	json.NewEncoder(res).Encode(tasks)

}

func addTaskForProject(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	newTask := protos.TaskModel{}
	json.NewDecoder(req.Body).Decode(&newTask)
	if result, err := client.AddTask(ctx, &newTask); err != nil {
		json.NewEncoder(res).Encode(map[string]bool{"Ok": false})
	} else {
		json.NewEncoder(res).Encode(result)
	}
}

func getTaskStatusCount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	response := make(map[string]int32)

	pID, _ := req.URL.Query()["pId"]
	stream, err := client.GetStatusCount(ctx, &protos.RegisterID{ID: pID[0]})
	if err != nil {
		json.NewEncoder(res).Encode(response)
		return
	}
	for {
		result, err := stream.Recv()
		if err != nil {
			break
		}
		for k, v := range result.GetStatus() {
			response[k] = v
		}
	}
	json.NewEncoder(res).Encode(response)
}

func getAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	var users []Users

	stream, err := client.GetUsers(ctx, &protos.Empty{})
	if err != nil {
		json.NewEncoder(res).Encode(users)
		return
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			break
		}
		u := Users{
			ID:       result.GetUser().GetID(),
			Username: result.GetUser().GetUsername(),
		}
		users = append(users, u)
	}
	json.NewEncoder(res).Encode(users)
}

func updateTask(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	tID, _ := req.URL.Query()["tId"]
	newTask := protos.TaskModel{ID: tID[0]}
	json.NewDecoder(req.Body).Decode(&newTask)
	if result, err := client.UpdateTask(ctx, &newTask); err != nil {
		json.NewEncoder(res).Encode(map[string]bool{"Ok": false})
	} else {
		json.NewEncoder(res).Encode(result)
	}
}

func createUserAccount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	newUser := protos.UserSignup{}
	json.NewDecoder(req.Body).Decode(&newUser)
	if result, err := client.Signup(ctx, &newUser); err != nil {
		json.NewEncoder(res).Encode(map[string]string{"Message": "Account already exists"})
	} else {
		json.NewEncoder(res).Encode(result)
	}
}

func logIntoUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()
	user := protos.UserLogin{}
	json.NewDecoder(req.Body).Decode(&user)
	if result, err := client.Login(ctx, &user); err != nil {
		json.NewEncoder(res).Encode(map[string]string{"Message": "Invalid Email OR Password"})
	} else {
		json.NewEncoder(res).Encode(result)
	}
}

func getTeamMembers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	ctx, cancel := config.GetCtx()
	defer cancel()

	var members []Users
	cID, _ := req.URL.Query()["cId"]
	stream, err := client.GetTeamMembers(ctx, &protos.RegisterID{ID: cID[0]})
	if err != nil {
		json.NewEncoder(res).Encode(members)
		return
	}

	for {
		result, err := stream.Recv()
		if err != nil {
			break
		}
		m := Users{
			ID:       result.GetUser().GetID(),
			Username: result.GetUser().GetUsername(),
		}
		members = append(members, m)
	}
	json.NewEncoder(res).Encode(members)
}
