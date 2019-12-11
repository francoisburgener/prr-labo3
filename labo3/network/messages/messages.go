package messages

type Message struct {
	Id uint16
}

type MessageResult struct {
	Map map[uint16]bool
}

type MessageNotif struct {
	Map map[uint16]uint16
}
