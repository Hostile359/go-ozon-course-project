package commandhandler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	startCmd = "start"
	helpCmd = "help"
	addCmd  = "add"
	updateCmd = "update"
	getCmd = "get"
	listCmd = "list"
	deleteCmd = "delete"
)

type CommandFunc func(*CommandHandler, string) string

type CommandHandler struct {
	route map[string]CommandFunc
	userApp userapp.App
	lastId user.UserId
}

func New(userApp userapp.App) *CommandHandler {
	route := make(map[string]CommandFunc)
	route[startCmd] = startFunc
	route[helpCmd] = helpFunc
	route[addCmd] = addFunc
	route[updateCmd] = updateFunc
	route[getCmd] = getFunc
	route[listCmd] = listFunc
	route[deleteCmd] = deleteFunc

	return &CommandHandler{
		userApp: userApp,
		route: route,
		lastId: 1,
	}
}

func (c *CommandHandler) HandleCommand(cmd, args string) string {
	cmdFunc, ok := c.route[cmd];
	if !ok {
		return "Unknown command, use /help to get info about available commands"
	}

	return cmdFunc(c, args)
}

func startFunc(_ *CommandHandler, s string) string {
	return "Bot started, use /help to get more info"
}

func helpFunc(_ *CommandHandler, s string) string {
	return "/help - list commands\n" +
		"/add <name> <password> - add new user with name and password\n" +
		"/update <used_id> <new_name> <new_password> - update user's name and password\n" +
		"/get <used_id> - get user info\n" +
		"/list - get users list\n" +
		"/delete <used_id> - delete user"
}

func addFunc(c *CommandHandler, data string) string {
	args := strings.Split(data, " ")
	if len(args) != 2 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}
	name, password := args[0], args[1]
	u := user.NewUser(c.lastId, name, password)

	if err := c.userApp.Add(u); err != nil {
		if errors.Is(err, userapp.ErrValidationArgs) || errors.Is(err, userapp.ErrUserExists) {
			return err.Error()
		}
		return "internal error"
	}

	c.lastId += 1

	return "user added"
}

func updateFunc(c *CommandHandler, data string) string {
	args := strings.Split(data, " ")
	if len(args) != 3 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	name, password := args[1], args[2]
	u := user.NewUser(userId, name, password)

	if err := c.userApp.Update(u); err != nil {
		if errors.Is(err, userapp.ErrValidationArgs) || errors.Is(err, userapp.ErrUserNotExists) {
			return err.Error()
		}
		return "internal error"
	}

	return "user updated"
}

func getFunc(c *CommandHandler, data string) string {
	args := strings.Split(data, " ")
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	u, err := c.userApp.Get(userId)
	if err != nil {
		if errors.Is(err, userapp.ErrUserNotExists) {
			return err.Error()
		}
		return "internal error"
	}
	return fmt.Sprintf("%v", u.String())
}

func listFunc(c *CommandHandler, data string) string {
	usersList := c.userApp.List()
	res := make([]string, 0, len(usersList)+1)
	res = append(res, "Users list:")
	for _, u := range usersList {
		res = append(res, u.String())
	}
	return strings.Join(res, "\n")
}

func deleteFunc(c *CommandHandler, data string) string {
	args := strings.Split(data, " ")
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	if err := c.userApp.Delete(userId); err != nil {
		if errors.Is(err, userapp.ErrUserNotExists) {
			return err.Error()
		}
		return "internal error"
	}
	return "user deleted"
}

func checkId(id string) (user.UserId, error) {
	parsedId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0, errors.Wrapf(userapp.ErrValidationArgs, "<%v>, id must be number", id)
	}
	return user.UserId(parsedId), nil
}
