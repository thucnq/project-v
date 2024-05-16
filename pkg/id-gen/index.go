package idgen

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	objectType = uint64(1)             // using 15 for workflow
	epochStart = uint64(1420070400000) // 2015-01-01 00:00:00 GMT

	// typeBits     = 5
	epochBits    = 41
	shardBits    = 8
	sequenceBits = 10

	maximumRandomValue = 262143
)

// func hash(s string) uint64 {
// 	h := fnv.New64a()
// 	_, _ = h.Write([]byte(s))
// 	return h.Sum64()
// }

func readMachineID() []byte {
	id := make([]byte, 2)
	hid, err := readPlatformMachineID()
	if err != nil || len(hid) == 0 {
		hid, err = os.Hostname()
	}
	if err == nil && len(hid) != 0 {
		hw := md5.New()
		_, _ = hw.Write([]byte(hid))
		copy(id, hw.Sum(nil))
	} else {
		if _, randErr := crand.Reader.Read(id); randErr != nil {
			panic(
				fmt.Errorf(
					"cannot get hostname nor generate a random number: %v; %v",
					err, randErr,
				),
			)
		}
	}
	return id
}

var mutex sync.Mutex

type sequenceCount uint64

func (s *sequenceCount) Inc() uint64 {
	mutex.Lock()
	defer mutex.Unlock()
	a := (*uint64)(s)
	*s = sequenceCount((*a + 1) & maxSequence)
	return uint64(*s)
}
func (s *sequenceCount) Reset() uint64 {
	mutex.Lock()
	defer mutex.Unlock()
	*s = 0
	return uint64(*s)
}

var (
	lastTs      = uint64(0)
	lastTsPass  = uint64(0)
	sequence    = sequenceCount(0)
	maxSequence = uint64(math.Pow(2, sequenceBits) - 1)
	maxShardID  = uint64(math.Pow(2, shardBits) - 1)

	machineID = readMachineID()
	// idCount   = randInt()
	// pid       = os.Getpid()
)

// func randInt() uint32 {
// 	b := make([]byte, 3)
// 	if _, err := crand.Reader.Read(b); err != nil {
// 		panic(fmt.Errorf("cannot generate random number: %v;", err))
// 	}
// 	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
// }

func toMilliseconds(t time.Time) uint64 {
	return uint64(t.UnixMilli())
}

func getTimestamp() uint64 {
	return toMilliseconds(time.Now()) - epochStart
}

func getTimestampFromTimestampMillisecond(timestamp uint64) uint64 {
	return timestamp - epochStart
}

func waitNextMillis(currentTs uint64) uint64 {
	for currentTs == lastTs {
		// currentTs = getTimestamp()
		currentTs = toMilliseconds(time.Now())
	}
	return currentTs
}

func waitNextMillisPass(currentTs uint64) uint64 {
	for currentTs == lastTsPass {
		tmpDt := time.Unix(0, int64(currentTs)*int64(time.Millisecond))
		currentTs = uint64(tmpDt.Add(time.Millisecond).UnixMilli())
	}
	return currentTs
}

var lastTime uint64
var lastSeq uint64

// AtomicResolver define as atomic sequence resolver, base on standard sync/atomic.
func AtomicResolver(ms uint64) (uint64, error) {
	var last uint64
	var seq, localSeq uint64

	for {
		last = atomic.LoadUint64(&lastTime)
		localSeq = atomic.LoadUint64(&lastSeq)
		if last > ms {
			return maxSequence, nil
		}

		if last == ms {
			seq = maxSequence & (localSeq + 1)
			if seq == 0 {
				return maxSequence, nil
			}
		}

		if atomic.CompareAndSwapUint64(
			&lastTime, last, ms,
		) && atomic.CompareAndSwapUint64(&lastSeq, localSeq, seq) {
			return seq, nil
		}
	}
}

// NextID ...DISCARD
// This function have bug generate duplicate id
func NextID() uint64 {
	currentTS := getTimestamp()
	if currentTS == lastTs {
		seq := sequence.Inc()
		if seq == 0 {
			currentTS = waitNextMillis(currentTS)
		}
	} else {
		sequence.Reset()
	}

	lastTs = currentTS
	shardID := uint64(binary.BigEndian.Uint16(machineID)) & maxShardID
	// shardID := uint64(binary.BigEndian.Uint16(machineID)) & uint64(os.Getpid()) & createNodeID() & maxShardID
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | (shardID << sequenceBits) | uint64(sequence)
}

// createNodeID create the shard id.
// the problem is when we scale by run multiple instance with same docker image => mechineID is same each instance.
// pid may be is same because have no diff process
// func createNodeID() uint64 {
// 	ifas, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Printf("createNodeID: fail to get net interface: %v \n", err.Error())
// 		return 0
// 	}
// 	var as []string
// 	for _, ifa := range ifas {
// 		a := ifa.HardwareAddr.String()
// 		if a != "" {
// 			as = append(as, a)
// 		}
// 	}
//
// 	return hash(strings.Join(as, ""))
// }

func NextIDWithRandom() uint64 {
	currentTS := getTimestamp()
	randomizedValue := uint64(rand.Int63n(maximumRandomValue))
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | randomizedValue
}

// NewIDFromTime ...DISCARD
func NewIDFromTime(created uint64) uint64 {
	currentTS := getTimestampFromTimestampMillisecond(created)
	if currentTS == lastTs {
		// sequence = (sequence + 1) & maxSequence
		sequence.Inc()
	} else {
		// sequence = 0
		sequence.Reset()
	}

	lastTs = currentTS
	// shardID := uint64(binary.BigEndian.Uint16(machineID)) & maxShardID
	shardID := uint64(binary.BigEndian.Uint16(machineID)) & uint64(os.Getpid()) & maxShardID
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | (shardID << sequenceBits) | uint64(sequence)
}

// NewIDFromTime2 ...DISCARD
func NewIDFromTime2(created uint64, sequence uint64) uint64 {
	currentTS := getTimestampFromTimestampMillisecond(created)
	shardID := uint64(binary.BigEndian.Uint16(machineID)) & uint64(os.Getpid()) & maxShardID
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | (shardID << sequenceBits) | sequence
}

func GenIDFromTime(createdMS uint64) (uint64, error) {
	c := createdMS
	seq, err := AtomicResolver(c)
	if err != nil {
		return 0, err
	}

	for seq >= maxSequence {
		c = waitNextMillis(c)
		seq, err = AtomicResolver(c)
		if err != nil {
			return 0, err
		}
	}

	currentTS := getTimestampFromTimestampMillisecond(c)

	shardID := uint64(binary.BigEndian.Uint16(machineID)) & maxShardID
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | (shardID << sequenceBits) | seq, nil
}

func GenIDFromTimePass(createdMS uint64) (uint64, error) {
	c := createdMS
	// note: trust caller with limit 1 req/ms
	seq := uint64(0)
	currentTS := getTimestampFromTimestampMillisecond(c)

	shardID := uint64(binary.BigEndian.Uint16(machineID)) & maxShardID
	return (objectType << (epochBits + shardBits + sequenceBits)) | (currentTS << (shardBits + sequenceBits)) | (shardID << sequenceBits) | seq, nil
}

// GenID generate sequence id
func GenID() (uint64, error) {
	c := toMilliseconds(time.Now())
	return GenIDFromTime(c)
}

// NewID generate id as string
func NewID() (string, error) {
	id, err := GenID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}
