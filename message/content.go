package message

type Content int

const (
	DataContent Content = iota
	TextContent
	TextAndDataContent
)
