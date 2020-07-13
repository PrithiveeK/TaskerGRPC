package main

import (
	"context"
	"time"

	"../config"
	"../models"
	"../protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskModel Collections model
type TaskModel struct {
	ID           primitive.ObjectID `bson:"_id"`
	Subject      string             `bson:"subject"`
	Description  string             `bson:"description"`
	Status       string             `bson:"status"`
	Priority     int32              `bson:"priority"`
	Category     string             `bson:"category"`
	DateCreated  time.Time          `bson:"dateCreated"`
	DateModified time.Time          `bson:"dateModified"`
	StartDate    time.Time          `bson:"startDate"`
	DueDate      time.Time          `bson:"dueDate"`
	ProjectID    primitive.ObjectID `bson:"projectID"`
	Assignee     []struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
	} `bson:"assignee"`
}

//AddTask inserts a new task
func (taskServer) AddTask(ctx context.Context, in *protos.TaskModel) (*protos.TaskResponse, error) {
	startD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetStartDate())
	if err != nil {
		return &protos.TaskResponse{}, err
	}
	dueD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetDueDate())
	if err != nil {
		return &protos.TaskResponse{}, err
	}
	assigneeID, err := primitive.ObjectIDFromHex(in.GetAssigneeID())
	if err != nil {
		return &protos.TaskResponse{}, err
	}
	projectID, err := primitive.ObjectIDFromHex(in.GetProjectID())
	if err != nil {
		return &protos.TaskResponse{}, err
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
		return &protos.TaskResponse{}, err
	}
	return &protos.TaskResponse{Ok: true}, nil
}

//UpdateTask updates an existing task
func (taskServer) UpdateTask(ctx context.Context, in *protos.TaskModel) (*protos.TaskResponse, error) {
	taskID, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return &protos.TaskResponse{}, err
	}
	startD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetStartDate())
	if err != nil {
		return &protos.TaskResponse{}, err
	}

	dueD, err := time.Parse("2006-01-02T15:04:05.000Z", in.GetDueDate())
	if err != nil {
		return &protos.TaskResponse{}, err
	}

	assigneeID, err := primitive.ObjectIDFromHex(in.GetAssigneeID())
	if err != nil {
		return &protos.TaskResponse{}, err
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
		return &protos.TaskResponse{}, err
	}
	return &protos.TaskResponse{Ok: true}, nil
}

//GetAllTasks returns details about all the tasks
func (taskServer) GetAllTasks(_ *protos.Empty, stream protos.TaskerApp_GetAllTasksServer) error {
	ctx, cancel := config.GetCtx()
	defer cancel()

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
		return err
	}
	for taskCursor.Next(ctx) {
		var task TaskModel
		if err := taskCursor.Decode(&task); err != nil {
			return err
		}
		stream.Send(&protos.AllTaskInfo{
			Task: &protos.TaskModel{
				ID:           task.ID.Hex(),
				Subject:      task.Subject,
				Description:  task.Description,
				Status:       task.Status,
				Priority:     task.Priority,
				Category:     task.Category,
				DateCreated:  task.DateCreated.String(),
				DateModified: task.DateModified.String(),
				StartDate:    task.StartDate.String(),
				DueDate:      task.DueDate.String(),
				ProjectID:    task.ProjectID.Hex(),
				Assignee: &protos.TaskModel_AssigneeObj{
					AssigneeObj: &protos.UserInfo{
						ID:       task.Assignee[0].ID.Hex(),
						Username: task.Assignee[0].Username,
					},
				},
			},
		})
	}

	if err := taskCursor.Err(); err != nil {
		return err
	}

	defer taskCursor.Close(ctx)
	return nil
}

//GetStatusCount returns a map of status and count
func (taskServer) GetStatusCount(in *protos.RegisterID, stream protos.TaskerApp_GetStatusCountServer) error {
	ctx, cancel := config.GetCtx()
	defer cancel()

	projectID, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return err
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
		return err
	}
	type Result struct {
		ID    string `bson:"_id"`
		Count int32  `bson:"count"`
	}
	for resultCursor.Next(ctx) {
		var result Result
		if err := resultCursor.Decode(&result); err != nil {
			return err
		}
		stream.Send(&protos.StatusResponse{
			Status: map[string]int32{
				result.ID: result.Count,
			},
		})
	}

	if err = resultCursor.Err(); err != nil {
		return err
	}

	defer resultCursor.Close(ctx)
	return nil
}

//UpdateTaskStatus updates the status of the tasks
func (taskServer) UpdateTaskStatus(ctx context.Context, _ *protos.Empty) (*protos.Empty, error) {
	if _, err := TasksColl.UpdateMany(ctx, bson.M{"status": "Open", "startDate": bson.M{"$lte": time.Now()}}, bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "status", Value: "In Progress"},
			{Key: "dateModified", Value: time.Now()},
		},
	}}); err != nil {
		return &protos.Empty{}, err
	}
	if _, err := TasksColl.UpdateMany(ctx, bson.M{"status": "In Progress", "dueDate": bson.M{"$lte": time.Now()}}, bson.D{{
		Key: "$set", Value: bson.D{
			{Key: "status", Value: "Closed"},
			{Key: "dateModified", Value: time.Now()},
		},
	}}); err != nil {
		return &protos.Empty{}, err
	}
	return &protos.Empty{}, nil
}

//GetAllProjects returns details about all the projects
func (taskServer) GetAllProjects(_ *protos.Empty, stream protos.TaskerApp_GetAllProjectsServer) error {
	ctx, cancel := config.GetCtx()
	defer cancel()

	projectCursor, err := ProjectsColl.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	for projectCursor.Next(ctx) {
		var project models.Projects
		if err := projectCursor.Decode(&project); err != nil {
			return err
		}
		stream.Send(&protos.ProjectResponse{
			Project: &protos.ProjectOne{
				ID:    project.ID.Hex(),
				Title: project.Title,
			},
		})
	}

	if err := projectCursor.Err(); err != nil {
		return err
	}

	defer projectCursor.Close(ctx)
	return nil
}
