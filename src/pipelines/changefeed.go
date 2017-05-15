package pipelines

import (
	"github.com/astaxie/beego/logs"
	r "unicontract/src/core/db/rethinkdb"
)

type ChangeFeed struct {
	node      Node
	db        string
	table     string
	operation []string
}

func (c *ChangeFeed) runForever() {
	c.runChangeFeed()
}

func (c *ChangeFeed) runChangeFeed() {
	logs.Info("change feed run")
	var value interface{}
	res := r.Changefeed(c.db, c.table)
	for res.Next(&value) {
		m := value.(map[string]interface{})
		logs.Info(c.table, "Changefeed result : %s", m["new_val"])
		c.node.output <- m["new_val"]
	}
	logs.Info("change feed out")
}