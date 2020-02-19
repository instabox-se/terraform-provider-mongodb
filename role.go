package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Role struct {
	Role       string      `bson:"role"`
	Db         string      `bson:"db"`
	Roles      []RoleRef   `bson:"roles"`
	Privileges []Privilege `bson:"privileges"`
}

type RoleInfoResult struct {
	Ok    int    `bson:"ok"`
	Roles []Role `bson:"roles"`
}

func (client *Client) GetRole(dbName string, role string) (*Role, error) {
	db := client.client.Database(dbName)

	command := bson.D{{"rolesInfo", role}, {"showPrivileges", true}}
	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result RoleInfoResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Roles) == 1 {
		return &result.Roles[0], nil
	}

	return nil, nil
}

func (client *Client) CreateRole(role Role) error {
	db := client.client.Database(role.Db)

	command := bson.D{
		{"createRole", role.Role},
		{"roles", bsonRoleRefs(role.Roles)},
		{"privileges", bsonPrivileges(role.Privileges)},
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result CreateResult

	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (client *Client) UpdateRole(role Role) error {
	db := client.client.Database(role.Db)

	command := bson.D{
		{"updateRole", role.Role},
		{"roles", bsonRoleRefs(role.Roles)},
		{"privileges", bsonPrivileges(role.Privileges)},
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result bson.M
	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}

func (client *Client) DeleteRole(role Role) error {
	db := client.client.Database(role.Db)

	command := bson.D{
		{"dropRole", role.Role},
	}

	opts := options.RunCmd().SetReadPreference(readpref.Primary())

	var result bson.M
	if err := db.RunCommand(*client.ctx, command, opts).Decode(&result); err != nil {
		return err
	}

	return nil
}
