package helper

const (
	CommandAdd   = "add"
	CommandList  = "list"
	CommandGet   = "get"
	CommandClear = "clear"
)

func GetCommandDescription() map[string]string {
	return map[string]string{
		CommandAdd:   "[DIR] - Add a directory to the history (if DIR is not specified, the current directory is used)",
		CommandList:  "- Show directory history",
		CommandGet:   "[INDEX] - Get directory by index",
		CommandClear: "clear - Clear history",
	}
}
