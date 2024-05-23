package pkg

import (
	"dijktras/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstra(t *testing.T) {
	asserT := assert.New(t)
	linkGraph := NewLinkGraph()
	configFile := config.ConfigFile{
		Links: []struct {
			From string `json:"from"`
			To   string `json:"to"`
			Cost int    `json:"cost"`
		}(make([]struct {
			From string
			To   string
			Cost int
		}, 4)),
	}
	configFile.Links[0].From = "A"
	configFile.Links[0].To = "B"
	configFile.Links[0].Cost = 1
	configFile.Links[1].From = "B"
	configFile.Links[1].To = "C"
	configFile.Links[1].Cost = 2
	configFile.Links[2].From = "A"
	configFile.Links[2].To = "C"
	configFile.Links[2].Cost = 3
	configFile.Links[3].From = "C"
	configFile.Links[3].To = "D"
	configFile.Links[3].Cost = 1
	for _, i := range configFile.Links {
		linkGraph.AddLinks(config.Config{
			NodeStart: i.From,
			NodeEnd:   i.To,
			Cost:      i.Cost,
		})
	}
	dijkstra := linkGraph.Dijkstra("A", "D")
	asserT.Equal(dijkstra.Path, []string{"A", "C", "D"})
	asserT.Equal(dijkstra.TotalCost, 4)
}
