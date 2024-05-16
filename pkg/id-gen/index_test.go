package idgen

import (
	"fmt"
	"sync"
	"testing"

	"project-v/pkg/l"
)

type HashSet struct {
	sync.Mutex

	data map[uint64]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{data: map[uint64]struct{}{}}
}

func (hs *HashSet) Put(id uint64) error {
	hs.Lock()
	defer hs.Unlock()

	if _, isEsists := hs.data[id]; isEsists {
		return fmt.Errorf("exists")
	}
	hs.data[id] = struct{}{}
	return nil
}

func Test_sequenceCount_Inc(t *testing.T) {
	t.SkipNow()
	println("==> ", 11&10)
	n := sequenceCount(0)
	var wg sync.WaitGroup
	b := uint64(0)
	genIDs := NewHashSet()

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		go func() {
			b = b + 1
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			n.Inc()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			id, err := GenID()
			if err != nil {
				t.Fatalf("failed to gen id: %v", l.Error(err))
			}

			if err := genIDs.Put(id); err != nil {
				t.Fatalf("duplicate id: %v", id)
			}
			// t.Log(id)
			wg.Done()
		}()

	}
	wg.Wait()
	println(n)
	println(b)
}

func TestNextID(t *testing.T) {
	t.SkipNow()
	id := int64(642696665231695872) // NextID()
	// t.Log(ExtractType(id))
	t.Log(id)

	typePart := (id >> 59) & 0x1F // 1
	t.Log(typePart)
	timePart := (id >> 18) & 0x1FFFFFFFFFF // 171017496261
	t.Log(uint64(timePart) + epochStart)

	// varPart := (id >> 0) & 0x3FFFF         // 148505

}

func TestNextID2(t *testing.T) {

	id := int64(642698326767738880) // NextID()
	// t.Log(ExtractType(id))
	t.Log(id)

	typePart := (id >> 59) & 0x1F // 1
	t.Log(typePart)
	timePart := (id >> 18) & 0x1FFFFFFFFFF // 171017496261
	t.Log(uint64(timePart) + epochStart)

	// varPart := (id >> 0) & 0x3FFFF         // 148505

}
