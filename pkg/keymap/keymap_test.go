// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package keymap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const keyMaps = `
/*
 * 主要切分依据是 title.production_year
 */
title.id <=> aka_title.movie_id
		 <=> cast_info.movie_id
		 <=> complete_cast.movie_id
		 <=> movie_companies.movie_id
		 <=> movie_info.movie_id
		 <=> movie_keyword.movie_id
		 <=> movie_link.movie_id;

char_name.id <=> cast_info.person_role_id; // char_name 不做切分

comp_cast_type.id <=> complete_cast.subject_id
                  <=> complete_cast.status_id; // comp_cast_type 不做切分

company_name.id <=> movie_companies.company_id; // company_name 不做切分

company_type.id <=> movie_companies.company_type_id; // company_type 不做切分

info_type.id <=> movie_info.info_type_id; // info_type 不做切分

keyword.id <=> movie_keyword.keyword_id; // keyword 不做切分

kind_type.id <=> title.kind_id; // kind_type 不做切分

link_type.id <=> movie_link.link_type_id; // link_type 不做切分

name.id <=> cast_info.person_id // 忽略
		<=> aka_name.person_id
		<=> person_info.person_id;

role_type.id <=> cast_info.role_id; // role_type 不做切分

`

func TestParse(t *testing.T) {
	maps, err := Parse(keyMaps)
	assert.Nil(t, err)
	assert.Len(t, maps, 11)
	assert.Len(t, maps[0].ForeignKeys, 7)
	assert.Len(t, maps[9].ForeignKeys, 3)
	assert.Equal(t, Key{Table: "title", Column: "id"}, *maps[0].PrimaryKey)
	assert.Equal(t, Key{Table: "aka_title", Column: "movie_id"}, *maps[0].ForeignKeys[0])
	assert.Equal(t, Key{Table: "movie_link", Column: "movie_id"}, *maps[0].ForeignKeys[6])
	assert.Equal(t, Key{Table: "name", Column: "id"}, *maps[9].PrimaryKey)
	assert.Equal(t, Key{Table: "cast_info", Column: "person_id"}, *maps[9].ForeignKeys[0])
	assert.Equal(t, Key{Table: "person_info", Column: "person_id"}, *maps[9].ForeignKeys[2])
}
