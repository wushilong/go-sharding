/*
 * Copyright 2021. Go-Sharding Author All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  File author: Anders Xiao
 */

//配置参考：https://shardingsphere.apache.org/document/legacy/4.x/document/cn/manual/sharding-jdbc/configuration/config-yaml/

package core

type ShardingTable struct {
	Name             string
	Resources        map[string][]string
	ShardingColumns  []string
	TableStrategy    ShardingStrategy
	DatabaseStrategy ShardingStrategy
}

func NoShardingTable(tableName string) *ShardingTable {
	return &ShardingTable{
		Name:          tableName,
		TableStrategy: NoneShardingStrategy,
	}
}

func (t *ShardingTable) HasColumn(column string) bool {
	for _, column := range t.ShardingColumns {
		if TrimAndLower(column) == column {

		}
	}
	return false
}

//指示是否需要进行分片
func (t *ShardingTable) IsSharding() bool {
	return t.TableStrategy != NoneShardingStrategy && t.DatabaseStrategy != NoneShardingStrategy
}
