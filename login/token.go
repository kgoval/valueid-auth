package login

import (
	"context"
	"fmt"
	"github.com/kgoval/erresp"
	"github.com/kgoval/valueid-auth/protogen"
	"github.com/kgoval/valueid-auth/servers"
	"github.com/pkg/errors"
)

type TokenApp struct{
	servers *servers.Server
}

var token *TokenApp

func Register( server *servers.Server) *TokenApp{
	token = &TokenApp{
		servers:server,
	}
	return token
}

func (c *TokenApp) Token(){

	// just try to return data from membership
	resp, err := c.servers.Member.Detail(context.Background(), &member.MemberDetailRequest{
		Id:"1",
	})

	if err != nil {
		errors.Wrap(erresp.Parse(err.Error()), "")
	}

	fmt.Println(resp)
}