package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateUserParams struct {
	Username string ` bson:"username,omitempty"`
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func (q *Queries) InitUsers() {
	index := mongo.IndexModel{
		Keys: bson.D{{"email", 1}},
		Options: options.Index().
			SetUnique(true).
			SetPartialFilterExpression(bson.D{
				{"email", bson.D{
					{"$exists", true},
				}},
				{"username", bson.D{
					{"$exists", true},
				}},
			}),
	}
	collection := q.db.Collection("users")
	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		panic(err)
	}
}
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	collection := q.db.Collection("users")

	result, err := collection.InsertOne(ctx, arg)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:       result.InsertedID.(primitive.ObjectID),
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
	}, err
}
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	collection := q.db.Collection("users")

	var result User
	filter := bson.D{{"email", email}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

type GetUserByTokenAndIdParams struct {
	AccessToken string
	ID          string
}

func (q *Queries) GetUserByTokenAndId(ctx context.Context, arg GetUserByTokenAndIdParams) (User, error) {
	collection := q.db.Collection("users")
	userId, err := primitive.ObjectIDFromHex(arg.ID)
	if err != nil {
		return User{}, err
	}

	var result User
	filter := bson.D{{"_id", userId}, {"access_token", arg.AccessToken}}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return User{}, err
	}
	return result, nil
}

type SetAccessTokenParams struct {
	AccessToken string
	ID          string
}

func (q *Queries) SetAccessToken(ctx context.Context, arg SetAccessTokenParams) error {
	collection := q.db.Collection("users")
	userId, err := primitive.ObjectIDFromHex(arg.ID)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", userId}}
	update := bson.D{{"$set", bson.D{{"access_token", arg.AccessToken}}}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
