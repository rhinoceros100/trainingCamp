问题：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


应该Wrap这个error抛给上层。因为:
1.如果直接抛sql.ErrNoRows错误给上层，那么上层就会直接依赖这个sentinel err，将来如果想要扩展这个error并添加一些上下文的话就会有问题；
2.需要有文档记录，会增加API的表面积；
3.在两个包之间创建了源代码依赖。
所以我们应该Wrap这个error给上层，这样上层每次调用的时候都需要Unwrap。

code:
func OpDb() error {
	var err error = sql.ErrNoRows
	return errors.Wrap(err, "DAO op db")
}
