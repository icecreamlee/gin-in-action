package core

import (
	"time"
)

// Token Bucket 令牌桶
type TokenBucket struct {
	PutInterval time.Duration // 放令牌间隔
	Tokens      chan int64    // 桶内令牌
}

// take取走一个令牌
func (bucket *TokenBucket) Take() bool {
	select {
	case <-bucket.Tokens:
		return true
	default:
		return false
	}
}

// putInterval 每隔一段时间放入一个令牌
func (bucket *TokenBucket) putInterval() {
	ticker := time.NewTicker(bucket.PutInterval)
	for {
		select {
		case <-ticker.C:
			select {
			case bucket.Tokens <- time.Now().UnixNano():
			default:
			}
		}
	}
}

func NewTokenBucket(putInterval time.Duration, maxToken int) *TokenBucket {
	bucket := &TokenBucket{PutInterval: putInterval, Tokens: make(chan int64, maxToken)}
	go bucket.putInterval()
	return bucket
}
