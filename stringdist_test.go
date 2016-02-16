package stringdist

import (
	"testing"

	. "gopkg.in/check.v1"
)

type PkgSuite struct{}

var _ = Suite(&PkgSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *PkgSuite) TestJoin(c *C) {

	l := Join([]string{
		"Quentin Tarantino",
	}, []string{
		"Quentin Tarantno",
		"Quentin Tarantino",
		"Quentin Tarntiano",
		"quentin tarantino",
		"Quentin Totenham",
		"Quentin Totanham",
		"Quentinson Tarantinson",
		"Quentinson Tarantinson",
	}, 2)
	c.Assert(len(l), Equals, 3)
	c.Assert(l[0], Equals, "Quentin Tarantino")
	c.Assert(l[1], Equals, "Quentin Totenham")
	c.Assert(l[2], Equals, "Quentinson Tarantinson")

	l = Join([]string{"pl"}, []string{"en"}, 0)
	c.Assert(len(l), Equals, 2)
	c.Assert(l[0], Equals, "pl")
	c.Assert(l[1], Equals, "en")
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{
			"Sprawdz total",
			"Kozak ziomeczek",
			"Quentin Totenham",
			"Quentin Tarantino",
		}, []string{
			"Quentin Tarantno",
			"Quentin Tarantino",
			"Quentin Tarntiano",
			"quentin tarantino",
			"Quentinson Tarantinson",
			"Quentin Totanham",
		}, 2)
	}
}
