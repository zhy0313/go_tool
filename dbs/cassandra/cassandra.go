package cassandra
/*
cassandra操作工具类
*/
import (
	"github.com/gocql/gocql"
)

//配置特别简单，主机名和keyspace
type Config struct {
	Host     []string
	Keyspace string
}

//这是一个重要的结构体，cassandra连接包装于此
type Cdb struct {
	*Config
	baseSession *gocql.Session
}

//配置初始化连接
func NewCdbWithConf(c *Config) (cdb *Cdb) {
	cdb = &Cdb{
		Config: c,
	}
	cdb.Connect()
	return
}

//一般初始化连接
func NewCdb(host []string, keyspace string) (cdb *Cdb) {
	cdb = &Cdb{
		Config: &Config{
			Host:     host,
			Keyspace: keyspace,
		},
	}
	cdb.Connect()
	return
}

//连接，出错panic
func (self *Cdb) Connect() {
	cluster := gocql.NewCluster(self.Host...)
	cluster.Keyspace = self.Keyspace
	cluster.Consistency = gocql.Quorum
	baseSession, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	self.baseSession = baseSession

}

//别名，为了容易使用
func New(c *Config) (cdb *Cdb) {
	cdb = &Cdb{
		Config: c,
	}
	cdb.Connect()
	return
}

//构造查询语句，包括插入，查找，不仅仅是查询
func (c *Cdb) Query(stmt string, values ...interface{}) *gocql.Query {
	if c.baseSession == nil {
		return nil
	}
	return c.baseSession.Query(stmt, values...)
}

//使用上面的查询语句，开始执行,一般是插入操作
func (c *Cdb) Exec(q *gocql.Query) error {
	return q.Exec()
}

//使用上面的查询语句，开始执行,一般是查找操作
func (c *Cdb) Iter(q *gocql.Query) *gocql.Iter {
	return q.Iter()
}

