package merkledag

import (
	"crypto/sha256"
	"errors"
	"hash"
)

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

// splitData 将数据切分为指定大小的块  
func splitData(data []byte, chunkSize int) [][]byte {  
	var chunks [][]byte  
	for len(data) > 0 {  
		if len(data) >= chunkSize {  
			chunk := data[:chunkSize]  
			data = data[chunkSize:]  
			chunks = append(chunks, chunk)  
		} else {  
			chunks = append(chunks, data)  
			data = nil  
		}  
	}  
	return chunks  
}  
  
// calculateHash 使用给定的哈希函数计算数据的哈希值  
func calculateHash(h hash.Hash, data []byte) []byte {  
	h.Write(data)  
	return h.Sum(nil)  
}  
  
// buildMerkleTree 构建Merkle树并返回Merkle Root  
func buildMerkleTree(hashes [][]byte) []byte {  
	if len(hashes) == 0 {  
		return nil  
	}  
	for len(hashes) > 1 {  
		if len(hashes)%2 != 0 {  
			hashes = append(hashes, hashes[len(hashes)-1]) // 如果数量是奇数，复制最后一个哈希值  
		}  
		newHashes := make([][]byte, 0, len(hashes)/2)  
		for i := 0; i < len(hashes); i += 2 {  
			left := hashes[i]  
			right := hashes[i+1]  
			combined := append(left, right...)  
			hash := calculateHash(sha256.New(), combined)  
			newHashes = append(newHashes, hash)  
		}  
		hashes = newHashes  
	}  
	return hashes[0] // 返回Merkle Root  
}  
  
// Add 函数实现  
func Add(store KVStore, node Node, h hash.Hash) ([]byte, error) {  
	var hashes [][]byte  
  
	switch t := node.(type) {  
	case File:  
		data := t.Bytes()  
		chunks := splitData(data, 1024) // 假设每个分片的大小为1024字节  
  
		for i, chunk := range chunks {  
			key := []byte("chunk-" + string(rune(i))) // 使用简单的键来存储分片  
			if err := store.Put(key, chunk); err != nil {  
				return nil, err // 返回错误  
			}  
			hash := calculateHash(h, chunk)  
			hashes = append(hashes, hash)  
		}  
	case Dir:  
		it := t.It()  
		for it.Next() {  
			subdirHash, err := Add(store, it.Node(), h)  
			if err != nil {  
				return nil, err // 返回错误  
			}  
			hashes = append(hashes, subdirHash)  
		}  
	default:  
		return nil, ErrUnsupportedNodeType // 返回不支持的节点类型错误  
	}  
  
	merkleRoot := buildMerkleTree(hashes)  
	return merkleRoot, nil  
}  
var ErrUnsupportedNodeType = errors.New("unsupported node type")
