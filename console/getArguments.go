package console

import "flag"

type UserArgument struct {
	Name string
	Value int
}

type UserInputGetter interface {
	GetArgument(name string, defaultValue int) UserArgument
}

type UserInput struct {}

func (u *UserInput) GetArgument(name string, defaultValue int) UserArgument {
	return UserArgument{
		Name: name,
		Value: *flag.Int(name, defaultValue, "A string variable"),
	}
}

func GetArguments(u UserInputGetter) []UserArgument {
	var arguments []UserArgument
	arguments = append(arguments, u.GetArgument("hours", 24))
	arguments = append(arguments, u.GetArgument("maxrecords", 100000))
	return arguments
}