## 问题
我们在数据库操作的时候，比如dao层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层，为什么，应该怎么做？
