package servers

import (
	"github.com/kgoval/valueid-membership/protogen"
	"github.com/micro/go-micro/client"
)

type Server struct{
	Member member.MemberService
}

var Servers Server


func Register(c client.Client){

	Servers.Member = member.NewMemberService("msa.valueid.membership", c)
}