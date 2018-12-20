package tests

import (
	. "gopkg.in/check.v1"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/network"
	"github.com/polyswarm/perigord/testing"

	"math/big"
	"perigord/bindings"
)

type MonsterCoreSuite struct {
	network *network.Network
}

var _ = Suite(&MonsterCoreSuite{})

func (s *MonsterCoreSuite) SetUpTest(c *C) {
	nw, err := testing.SetUpTest()
	if err != nil {
		c.Fatal(err)
	}

	s.network = nw
}

// func (s *MonsterCoreSuite) TearDownTest(c *C) {
// 	testing.TearDownTest()
// }

// USER TESTS GO HERE
func (s *MonsterCoreSuite) TestTotalSupply(c *C) {
	session := contract.Session("MonsterCore")
	c.Assert(session, NotNil)
	token_session, ok := session.(*bindings.MonsterCoreSession)
	c.Assert(ok, Equals, true)
	c.Assert(token_session, NotNil)
	ret, _ := token_session.TotalSupply()
	c.Assert(ret.Text(10), Equals, "0")
}
