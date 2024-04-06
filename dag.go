package merkledag

import "hash"

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

func Add(store KVStore, node Node, hp HashPool) []byte {
	// TODO 将分片写入到KVStore中，并返回Merkle Root
	switch node.Type() {
	case FILE:
		file := node.(File) // 将node转换为File类型
		// 将文件内容写入到KVStore中
		err := store.Put([]byte("file"), file.Bytes())
		if err != nil {
			panic(err)
		}
	case DIR:
		dir := node.(Dir) // 将node转换为Dir类型
		it := dir.It()    // 获取DirIterator
		for it.Next() {
			childNode := it.Node()
			// 递归调用Add函数处理子节点
			Add(store, childNode, h)
		}
	}
	
	
	// 计算Merkle Root
	root := h.Sum(nil)

	// 将Merkle Root写入到KVStore中
	err := store.Put([]byte("root"), root)
	if err != nil {
		panic(err)
	}

	return root
}

