package chatTool

import (
	iu "SunCache/cache/interface"
	"SunCache/cache/log"
	"SunCache/cache/member"
	context "context"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type Server struct {
	socket string
	team   iu.TeamIU

	UnimplementedChatToolServer
}

func NewServer(socket string) (server *Server) {
	server = &Server{
		socket: socket,
	}
	return
}

func (c *Server) GetSocket() (socket string) {
	return c.socket
}

func (c *Server) RegisterTeam(team iu.TeamIU) {
	if c.team != nil {
		panic("ChatToolServer RegisterTeam() called more than once")
	}
	c.team = team
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.socket)
	if err != nil {
		log.Error("[%v] net listen err: %v", s.socket, err)
		return
	}
	server := grpc.NewServer()
	RegisterChatToolServer(server, s)
	if err := server.Serve(listener); err != nil {
		log.Error("[%v] gRPC server start err: %v", s.socket, err)
		return
	}
	log.Info("[%v] gRPC server listen on [%v]", s.socket, strings.Split(s.socket, ":")[1])
}

func (s *Server) Get(ctx context.Context, req *Request) (res *Response, err error) {
	memberName := req.Member
	key := req.Key

	m := s.team.GetMember(memberName)
	var vIU iu.ByteViewIU
	vIU, err = m.Get(key)
	v := vIU.(member.ByteView)
	if err != nil {
		return
	}
	res = &Response{Value: v.Copy()}
	return
}
