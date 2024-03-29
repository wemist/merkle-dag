package merkledag

import (
	"errors"
)

// Hash2File 根据哈希值和路径从KVStore中检索文件内容。

func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 假设hash是文件内容的直接键
	data, err := store.Get(hash)
	if err != nil {

		return nil
	}

	// 假设数据已经是文件内容，直接返回
	return data
}

var (
	ErrNotFound = errors.New("not found")
)

func Hash2FileRobust(store KVStore, hash []byte, path string, hp HashPool) ([]byte, error) {
	// 这里应该添加逻辑来解析Merkle DAG结构，根据hash递归检索节点内容
	// ...

	// 假设找到了文件内容
	data := []byte("file content")

	// 假设没有错误发生
	return data, nil
}

// 以下是一个可能的Merkle DAG节点实现的例子
type MerkleNode struct {
	Size uint64
	Name string
	Type int
	Data []byte // 对于文件节点，这里存储文件内容
	// 对于目录节点，这里可能存储子节点的哈希列表
}

// 实现Node接口
func (n *MerkleNode) Size() uint64 {
	return n.Size
}

func (n *MerkleNode) Name() string {
	return n.Name
}

func (n *MerkleNode) Type() int {
	return n.Type
}

// 实现File接口（如果MerkleNode代表文件）
func (n *MerkleNode) Bytes() []byte {
	return n.Data
}
