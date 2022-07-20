package commandhandler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/memoryuserstore"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/userstore"
)

const (
	startCmd = "start"
	helpCmd = "help"
	addCmd  = "add"
	updateCmd = "update"
	getCmd = "get"
	getallCmd = "getall"
	deleteCmd = "delete"
)

type CommandHandler struct {
	userStore userstore.UserStore
}

func (commandHandler *CommandHandler) Init() {
	commandHandler.userStore = &memoryuserstore.MemoryUserStore{}
	commandHandler.userStore.Init()
}

func (commandHandler *CommandHandler) HandleCommand(cmd, args string) string {
	var res string
	switch cmd {
	case startCmd:
		res = commandHandler.startFunc(args)
	case helpCmd:
		res = commandHandler.helpFunc(args)
	case addCmd:
		res = commandHandler.addFunc(args)
	case updateCmd:
		res = commandHandler.updateFunc(args)
	case getCmd:
		res = commandHandler.getFunc(args)
	case getallCmd:
		res = commandHandler.getallFunc(args)
	case deleteCmd:
		res = commandHandler.deleteFunc(args)
	default:
		res = "Unknown command, use /help to get info about available commands"
	}

	return res
}

func (CommandHandler) startFunc(s string) string {
	return "Bot started, use /help to get more info"
}

func (CommandHandler) helpFunc(s string) string {
	return "/help - list commands\n" +
		"/add <name> <password> - add new user with name and password\n" +
		"/update <used_id> <new_name> <new_password> - update user's name and password\n" +
		"/get <used_id> - get user info\n" +
		"/getall - get users list\n" +
		"/delete <used_id> - delete user"
}

func (commandHandler *CommandHandler) addFunc(data string) string {
	args := strings.Split(data, " ")
	if len(args) != 2 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	name, password := args[0], args[1]

	if err := checkName(name); err != nil {
		return err.Error()
	}

	if err := checkPassword(password); err != nil {
		return err.Error()
	}

	err := commandHandler.userStore.AddUser(name, password)
	if err != nil {
		return err.Error()
	}
	return "user added"
}

func (commandHandler *CommandHandler) updateFunc(data string) string {
	args := strings.Split(data, " ")
	if len(args) != 3 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	name, password := args[1], args[2]

	if err := checkName(name); err != nil {
		return err.Error()
	}

	if err := checkPassword(password); err != nil {
		return err.Error()
	}

	err = commandHandler.userStore.UpdateUser(userId, name, password)
	if err != nil {
		return err.Error()
	}
	return "user updated"
}

func (commandHandler CommandHandler) getFunc(data string) string {
	args := strings.Split(data, " ")
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	u, err := commandHandler.userStore.GetUser(userId)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%v", u.String())
}

func (commandHandler CommandHandler) getallFunc(s string) string {
	usersList := commandHandler.userStore.GetAllUsers()
	res := make([]string, 0, len(usersList)+1)
	res = append(res, "Users list:")
	for _, u := range usersList {
		res = append(res, u.String())
	}
	return strings.Join(res, "\n")
}

func (commandHandler *CommandHandler)deleteFunc(data string) string {
	args := strings.Split(data, " ")
	if len(args) != 1 {
		return fmt.Sprintf("bad arguments <%v>", args)
	}

	userId, err := checkId(args[0])
	if err != nil {
		return err.Error()
	}

	err = commandHandler.userStore.DeleteUser(userId)
	if err != nil {
		return err.Error()
	}
	return "user deleted"
}

func checkId(id string) (uint, error) {
	parsedId, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return 0, errors.Errorf("bad id <%v>, id must be number", id)
	}
	return uint(parsedId), nil
}

func checkName(name string) error {
	if len(name) == 0 || len(name) > 10 {
		return errors.Errorf("bad name <%v>, len should be from 1 to 10", name)
	}
	return nil
}

func checkPassword(password string) error {
	if len(password) < 6 || len(password) > 10 {
		return errors.Errorf("bad password <%v>, len should be from 6 to 10", password)
	}	
	return nil
}
