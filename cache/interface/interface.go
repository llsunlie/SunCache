package iu

type ByteViewIU interface {
	Len() (length int)
	Copy() (byteView []byte)
	String() (str string)
}

type MemberIU interface {
	RegisterTeam(team TeamIU)
	GetName() (name string)
	Get(key string) (value ByteViewIU, err error)
}

type TeamIU interface {
	AddMember(members ...MemberIU)
	GetMember(name string) (member MemberIU)
	AddAddress(addresses ...string)
	GetAddress(key string) (address string)
	GetChatServerSocket() (address string)
	GetValueFromRemote(member, key string) (value []byte, err error)
	RunChatServer()
	StartHttpServer(socket string)
}
