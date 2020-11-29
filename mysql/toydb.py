#!/usr/bin/env python
# coding=utf8

# Create time:2017-5-4
# Update time:2017-5-4
# Version: 1
# sudo apt-get install python-mysqldb

import MySQLdb
import sys

reload(sys)
sys.setdefaultencoding('utf8')


class ToyDB:
    conn = None
    cur = None

    def connect(self, user, passwd, ip, db, port=3306):
        self.conn = MySQLdb.connect(
            host=ip, port=port,
            user=user, passwd=passwd,
            db=db, charset="utf8")
        self.cur = self.conn.cursor()
        self.cur.execute("SET NAMES utf8;")
        self.conn.commit()

    def run_sql(self, sql):
        self.cur.execute(sql)
        self.conn.commit()

    def select_sql(self, sql):
        aa = self.cur.execute(sql)
        ret = self.cur.fetchmany(aa)
        return ret

    def close(self):
        self.cur.close()
        self.conn.close()


def test_mysql_db():
    db = ToyDB()
    db.connect("test", "123456", "127.0.0.1", "test", 3306)

    ret = []
    sql = "SELECT * FROM test.test_tables where id>=0 order by id asc;"
    rs = db.select_sql(sql)
    for itor in rs:
        if itor is None:
            break
        if ret is not None:
            ret.append({"id": itor[0],
                        "name": itor[1],
                        "age": itor[2]})
    print ret
    db.close()


if __name__ == "__main__":
    test_mysql_db()
