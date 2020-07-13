package main

import (
	"context"
	"errors"

	"../config"
	"../models"
	"../protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//Login logs the user in
func (taskServer) Login(ctx context.Context, in *protos.UserLogin) (*protos.RegisterID, error) {
	var user models.Users

	if err := UsersColl.FindOne(ctx, &models.Users{Email: in.GetEmail()}).Decode(&user); err != nil {
		return &protos.RegisterID{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.GetPassword())); err != nil {
		return &protos.RegisterID{}, err
	}
	return &protos.RegisterID{ID: user.ID.Hex()}, nil
}

//Signup creates an account
func (taskServer) Signup(ctx context.Context, in *protos.UserSignup) (*protos.RegisterID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), 10)
	if err != nil {
		return &protos.RegisterID{}, err
	}

	newUser := models.Users{
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
		Password: string(hash),
	}

	user, err := UsersColl.InsertOne(ctx, newUser)
	if err != nil {
		return &protos.RegisterID{}, err
	}
	if oid, ok := user.InsertedID.(primitive.ObjectID); ok {
		return &protos.RegisterID{ID: oid.Hex()}, nil
	}
	return &protos.RegisterID{}, errors.New("Sorry, something went wrong")
}

//GetUsers returns all the users in th collection
func (taskServer) GetUsers(_ *protos.Empty, stream protos.TaskerApp_GetUsersServer) error {
	ctx, cancel := config.GetCtx()
	defer cancel()
	userCursor, err := UsersColl.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	for userCursor.Next(ctx) {
		var uC models.Users
		if err := userCursor.Decode(&uC); err != nil {
			return err
		}
		stream.Send(&protos.AllUsers{
			User: &protos.UserInfo{
				ID:       uC.ID.Hex(),
				Username: uC.Username,
			},
		})
	}

	if err := userCursor.Err(); err != nil {
		return err
	}
	defer userCursor.Close(ctx)
	return nil

}

//GetTeamMember returns all the members under a team leader
func (taskServer) GetTeamMembers(in *protos.RegisterID, stream protos.TaskerApp_GetTeamMembersServer) error {
	ctx, cancel := config.GetCtx()
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(in.GetID())
	if err != nil {
		return err
	}

	matchWith := bson.D{{Key: "$match", Value: bson.D{{Key: "leaderID", Value: userID}}}}
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
		return err
	}

	type TeamMem struct {
		MemberData []struct {
			ID       primitive.ObjectID `bson:"_id"`
			Username string             `bson:"username"`
		} `bson:"memberData"`
	}
	var teamMems []TeamMem
	if err := teamMemberCursor.All(ctx, &teamMems); err != nil {
		return err
	}

	for _, user := range teamMems[0].MemberData {
		stream.Send(&protos.AllUsers{
			User: &protos.UserInfo{
				ID:       user.ID.Hex(),
				Username: user.Username,
			},
		})
	}

	defer teamMemberCursor.Close(ctx)
	return nil

}
