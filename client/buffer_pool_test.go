// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"testing"
)

func BenchmarkGetMediaBufferFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			buf := mediaBufferPool.Get().(*bytes.Buffer)
			defer mediaBufferPool.Put(buf)
		}()
	}
}

func BenchmarkGetMediaBufferFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = bytes.NewBuffer(make([]byte, 10<<20))
		}()
	}
}

func BenchmarkGetTextBufferFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			buf := textBufferPool.Get().(*bytes.Buffer)
			defer textBufferPool.Put(buf)
		}()
	}
}

func BenchmarkGetTextBufferFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = bytes.NewBuffer(make([]byte, 64<<10))
		}()
	}
}
