package servers

import (
	"github.com/kgoval/valueid-auth/protogen"
	"github.com/micro/go-micro/client"
)

type Server struct{
	Member member.MemberService
}

var Servers Server


func Register(c client.Client){

	Servers.Member = member.NewMemberService("msa.valueid.membership", c)
}