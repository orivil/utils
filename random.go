// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package utils

import (
	"math/rand"
	"time"
)

// 允许的字符集
var str = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 创建指定位数的随机字符数组
func RandomBytes(bit int) []byte {
	return RandomBytesBySrc(str, bit)
}

// 从src中随机创建bit位的随机字符数组
func RandomBytesBySrc(src []byte, bit int) []byte {
	bts := make([]byte, bit)
	r := make([]byte, bit)
	ln := len(src)
	rand.Read(bts)
	for idx, b := range bts {
		r[idx] = src[int(b)%ln]
	}
	return r
}

// 随机生成一个 min - max 之间的数字, 用于生成邮箱或手机验证码
// Example:
//  100000, 999999 => 658459
func GenerateCode(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
