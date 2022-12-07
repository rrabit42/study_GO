package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block // pointer들의 slice
}

// Singleton 패턴: 단 하나의 instance만을 공유하는 방법
// 이 변수의 instance를 직접 공유하지 않고, 그 대신 우릴 대신해서 이 변수의 instance를 드러내주는 function 생성
// 이 말인 즉슨, 다른 패키지에서 우리의 blockchain이 어떻게 드러날 지를 제어할 수 있음

var b *blockchain // 소문자로 시작. -> main에서는 접근 못함. blockchain package에서만 접근 가능
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash)) // string도 []byte(byte slice)의 일종. for loop로 문자 하나씩 돌리면 binary이기 떄문. But string은 immutable, array는 길이가 정해져 있음. byte slice는 mutable. 따라서 sha 함수 안에서 변할 수 있는 데이터([]byte)가 필요함을 추측할 수 있음
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		// genesis block일 경우
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	// blockchain이 어떻게 초기화되고 공유될지 제어, 어떻게 생성될지
	if b == nil {
		once.Do(func() {
			// 단 한번만 실행, 이 이후에는 b가 nil이 아니기 때문
			// := 아님. 업데이트의 의미로 =
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

// 둘 중 하나 선택: 어디에 붙일 건지
// blockchain 모듈에서 AllBlocks 호출 -> singleton에 더 적합
// func AllBlocks() []*block {
// 	return GetBlockchain().blocks
// }

// chain에서 AllBlocks 호출
func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
