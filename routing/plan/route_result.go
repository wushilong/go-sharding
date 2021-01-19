// Copyright 2019 The Gaea Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"fmt"
	"github.com/XiaoMi/Gaea/core"
	"github.com/scylladb/go-set/strset"
)

// RouteResult is the route result of a statement
// 遍历AST之后得到的路由结果
// db, table唯一确定了一个路由, 这里只记录分片表的db和table, 如果是关联表, 必须关联到同一个父表
type RouteResult struct {
	Targets      map[string]*strset.Set // key = Physical database name, value = Physical table name
	ignoreCase   bool
	currentDb    string
	currentTable string
}

// NewRouteResult constructor of RouteResult
func NewRouteResult() *RouteResult {
	return &RouteResult{
		Targets:    make(map[string]*strset.Set, 0),
		ignoreCase: true,
	}
}

func (r *RouteResult) normalizeStr(db string) string {
	d := db
	if r.ignoreCase {
		d = core.TrimAndLower(db)
	}
	return d
}

func (r *RouteResult) normalizeSet(slice []string) *strset.Set {
	set := strset.NewWithSize(len(slice))
	for _, t := range slice {
		set.Add(r.normalizeStr(t))
	}
	return set
}

func (r *RouteResult) normalizeSlice(slice []string) []string {
	set := strset.NewWithSize(len(slice))
	for _, t := range slice {
		set.Add(r.normalizeStr(t))
	}
	return set.List()
}

// Inter inter indexes with origin indexes in RouteResult
// 如果是关联表, db, table需要用父表的db和table
func (r *RouteResult) Intersect(databases []string, tables []string) {
	if len(databases) == 0 {
		r.Targets = make(map[string]*strset.Set, 0)
	} else {
		dbSet := r.normalizeSlice(databases)
		set := r.normalizeSet(tables)
		contains := strset.New()
		for _, db := range dbSet {
			for dbName, tableSet := range r.Targets {
				if dbName == db {
					r.Targets[db] = strset.Intersection(tableSet, set)
					contains.Add(dbName)
				}
			}
		}

		merged := make(map[string]*strset.Set, contains.Size())
		contains.Each(func(item string) bool {
			merged[item] = r.Targets[item]
			return true
		})

		r.Targets = merged
	}
}

// Union union indexes with origin indexes in RouteResult
// 如果是关联表, db, table需要用父表的db和table
func (r *RouteResult) Union(databases []string, tables []string) {
	if len(databases) > 0 {
		tableSet := r.normalizeSet(tables)
		dbSet := r.normalizeSlice(databases)
		for _, db := range dbSet {
			if set, ok := r.Targets[db]; ok {
				r.Targets[db] = strset.Union(set, tableSet)
			} else {
				r.Targets[db] = tableSet.Copy()
			}
		}
	}
}

// GetShardIndexes get shard indexes
//func (r *RouteResult) GetShardIndexes() []int {
//	return r.indexes
//}

// GetCurrentTableIndex get current table index
func (r *RouteResult) GetCurrentTable() (string, error) {
	if r.currentDb == "" {
		return r.currentDb, fmt.Errorf("table index out of range in route result")
	}
	return r.currentDb, nil
}

// Next get next table index
func (r *RouteResult) Next() bool {
	db := r.currentDb
	table := ""
	if db == "" {
		db = r.nextDatabase(r.currentDb)
		if db != "" {
			table = r.nextTable(db, "")
		}
	} else {
		table = r.nextTable(db, r.currentTable)
	}
	hasNext := db != "" && table != ""
	if hasNext {
		r.currentDb = db
		r.currentTable = table
	}
	return hasNext
}

func (r *RouteResult) nextTable(db string, table string) string {
	tables, ok := r.Targets[db]
	if ok {
		found := table == ""
		current := ""
		tables.Each(func(item string) bool {
			if found {
				current = item
				return false //跳出玄幻
			} else {
				found = item == table
				return true
			}
		})
		return current
	}
	return ""
}

func (r *RouteResult) nextDatabase(db string) string {
	if db == "" {
		for t, _ := range r.Targets {
			return t
		}
	} else {
		found := false
		for t, _ := range r.Targets {
			if found {
				return t
			}
			if db == t {
				found = true
			}
		}
	}
	return ""
}

// Reset reset the cursor of index
func (r *RouteResult) Reset() {
	r.currentDb = ""
	r.currentTable = ""
}
