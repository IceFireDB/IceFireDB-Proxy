/*
 *
 *  * Licensed to the Apache Software Foundation (ASF) under one or more
 *  * contributor license agreements.  See the NOTICE file distributed with
 *  * this work for additional information regarding copyright ownership.
 *  * The ASF licenses this file to You under the Apache License, Version 2.0
 *  * (the "License"); you may not use this file except in compliance with
 *  * the License.  You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package redisNode

import (
	"log"
	"testing"
	"time"

	redisclient "github.com/gomodule/redigo/redis"
)

func TestSync(t *testing.T) {
	cli := &redisclient.Pool{
		MaxIdle:     2,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redisclient.Conn, error) {
			c, err := redisclient.Dial("tcp", "127.0.0.1:6379",
				redisclient.DialConnectTimeout(time.Duration(5)*time.Second),
				redisclient.DialReadTimeout(time.Duration(1)*time.Second),
				redisclient.DialWriteTimeout(time.Duration(1)*time.Second))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redisclient.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	r := NewRouter(cli)

	err := r.Sync([]interface{}{"set", "aa"})
	log.Println(err)
}
