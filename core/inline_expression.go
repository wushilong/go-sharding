/*
 *
 *  * Copyright 2021. Go-Sharding Author All Rights Reserved.
 *  *
 *  *  Licensed under the Apache License, Version 2.0 (the "License");
 *  *  you may not use this file except in compliance with the License.
 *  *  You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  *  Unless required by applicable law or agreed to in writing, software
 *  *  distributed under the License is distributed on an "AS IS" BASIS,
 *  *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  *  See the License for the specific language governing permissions and
 *  *  limitations under the License.
 *  *
 *  *  File author: Anders Xiao
 *
 */

package core

type FlatChecker func(string) error

type InlineExpression interface {
	Flat(checker FlatChecker) ([]string, error)
}

type inlineExpr struct {
	expression string
	segments   []*inlineSegment
}

func (i *inlineExpr) Flat(checker FlatChecker) ([]string, error) {
	for _, s := range i.segments {

	}
	return nil, nil
}

func NewInlineExpression(expression string) (InlineExpression, error) {
	expr := &inlineExpr{expression: expression}
	if segments, err := splitInlineExpression(expression); err != nil {
		return nil, err
	} else {
		expr.segments = segments
	}
	return expr, nil
}