package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Username            string    `bson:"user"`
	Password            string    `bson:"-"`
	AllowPasswordUpdate bool      `bson:"-"`
	Name                string    `bson:"customData.name"`
	Db                  string    `bson:"db"`
	Roles               []RoleRef `bson:"roles"`
}

type UsersInfoResult struct {
	Ok    int    `bson:"ok"`
	Users []User `bson:"users"`
}

func (client *Client) GetUser(databaseName string, name string) (*User, error) {
	db := client.client.Database(databaseName)

	command := bson.D{{"usersInfo", name}}
	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result UsersInfoResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Users) == 1 {
		return &result.Users[0], nil
	}

	return nil, nil
}

func (client *Client) CreateUser(user User) error {
	db := client.client.Database(user.Db)

	command := bson.D{
		{"createUser", user.Username},
		{"pwd", user.Password},
		{"customData", bson.D{{"name", user.Name}, {"__managedByTerraform", true}}},
		{"roles", bsonRoleRefs(user.Roles)},
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result CreateResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (client *Client) UpdateUser(user User) error {
	db := client.client.Database(user.Db)

	command := bson.D{
		{"updateUser", user.Username},
		{"customData", bson.D{{"name", user.Name}, {"__managedByTerraform", true}}},
		{"roles", bsonRoleRefs(user.Roles)},
	}

	if user.AllowPasswordUpdate {
		command = append(command, bson.E{"pwd", user.Password})
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result CreateResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (client *Client) DeleteUser(user User) error {
	db := client.client.Database(user.Db)

	command := bson.D{
		{"dropUser", user.Username},
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result CreateResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}
