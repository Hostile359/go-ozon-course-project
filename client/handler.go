package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
)

const (
	helpCmd = "help"
	addCmd  = "add"
	updateCmd = "update"
	getCmd = "get"
	listCmd = "list"
	deleteCmd = "delete"
)

type CommandFunc func([]string, context.Context) string

type CommandHandler struct {
	route map[string]CommandFunc
	client pb.UserClient
}

func NewCommandHandler(client pb.UserClient) *CommandHandler {
	route := make(map[string]CommandFunc)
	cH := CommandHandler{
		route: route,
		client: client,
	}

	route[helpCmd] = cH.helpFunc
	route[addCmd] = cH.addFunc
	route[updateCmd] = cH.updateFunc
	route[getCmd] = cH.getFunc
	route[listCmd] = cH.listFunc
	route[deleteCmd] = cH.deleteFunc

	return &cH
}

func (c *CommandHandler) HandleCommand(cmd string, ctx context.Context) string {
	args := strings.Fields(cmd)
	cmd = args[0]

	cmdFunc, ok := c.route[cmd];
	if !ok {
		return "Unknown command, use /help to get info about available commands"
	}

	return cmdFunc(args[1:], ctx)
}

func (*CommandHandler) helpFunc(s []string, _ context.Context) string {
	return "/help - list commands\n" +
		"/add <name> <password> - add new user with name and password\n" +
		"/update <used_id> <new_name> <new_password> - update user's name and password\n" +
		"/get <used_id> - get user info\n" +
		"/list - get users list\n" +
		"/delete <used_id> - delete user"
}

func (c *CommandHandler) addFunc(args []string, ctx context.Context) string {
	if len(args) != 2 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	name, password := args[0], args[1]

	req := pb.UserAddRequest{
		Name: name,
		Password: password,
	}

	response, err := c.client.UserAdd(ctx, &req)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprint(response)
}

func (c *CommandHandler) updateFunc(args []string, ctx context.Context) string {
	if len(args) != 3 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	name, password := args[1], args[2]

	req := pb.UserUpdateRequest{
		Id: userId,
		Name: name,
		Password: password,
	}

	response, err := c.client.UserUpdate(ctx, &req)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprint(response)
}

func (c *CommandHandler) getFunc(args []string, ctx context.Context) string {
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	req := pb.UserGetRequest{
		Id: userId,
	}

	response, err := c.client.UserGet(ctx, &req)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprint(response)
}

func (c *CommandHandler) listFunc(_ []string, ctx context.Context) string {
	req := pb.UserListRequest{}

	response, err := c.client.UserList(ctx, &req)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprint(response)
}

func (c *CommandHandler) deleteFunc(args []string, ctx context.Context) string {
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	req := pb.UserDeleteRequest{
		Id: userId,
	}

	response, err := c.client.UserDelete(ctx, &req)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprint(response)
}

func checkId(id string) (uint64, error) {
	parsedId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0, errors.Errorf("<%v>, id must be number", id)
	}
	return parsedId, nil
}

