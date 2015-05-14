# mgopath for mgo.v2

Connect to a MongoDB database by using a path like `mongodb://127.0.0.1:27017/mydatabase`:

```go
package main

import "github.com/ory-am/common/mgopath"
import "log"

func main() {
    sess, dbname, err := mgopath.Connect("mongodb://127.0.0.1:27017/mydatabase")
    if err != nil {
        log.Fatal(err)
    }
    db := sess.DB(dbname)
    // ...
}
```
