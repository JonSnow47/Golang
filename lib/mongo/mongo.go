/*
 * Revision History:
 *     Initial: 2018/05/30        Chen Yanchen
 */
package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Connection connect to MongoDB.
type Connection struct {
	session  *mgo.Session
	C        *mgo.Collection
	database string
	cname    string
}

// NewConnection.
func NewConnection(s *mgo.Session, database, cname string) (con *Connection) {
	return &Connection{
		session:  s,
		C:        s.DB(database).C(cname),
		database: database,
		cname:    cname,
	}
}

// dial to MongoDB and get a session.
func Dial(url, database string) *mgo.Session {
	s, err := mgo.DialWithTimeout(url+"/"+database, 5*time.Second)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	return s
}

// Connect represents clone a new Connection.
func (con *Connection) Connect() *Connection {
	s := con.session.Clone()
	c := con.session.DB(con.database).C(con.cname)
	return &Connection{
		session: s,
		C:       c,
	}
}

// Disconnect a clone session.
func (con *Connection) Disconnect() {
	con.session.Close()
}
