// Copyright 2020 Dolthub, Inc.
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

package set

import "sort"

type Uint64Set struct {
	uints map[uint64]bool
}

func NewUint64Set(uints []uint64) *Uint64Set {
	l := len(uints)
	if l < 16 {
		l = 16
	}

	s := &Uint64Set{make(map[uint64]bool, l)}

	for _, b := range uints {
		s.uints[b] = true
	}

	return s
}

func (us *Uint64Set) Contains(i uint64) bool {
	_, present := us.uints[i]
	return present
}

func (us *Uint64Set) ContainsAll(uints []uint64) bool {
	for _, b := range uints {
		if _, present := us.uints[b]; !present {
			return false
		}
	}

	return true
}

func (us *Uint64Set) Add(vals ...uint64) {
	for _, val := range vals {
		us.uints[val] = true
	}
}

func (us *Uint64Set) Remove(i uint64) {
	delete(us.uints, i)
}

func (us *Uint64Set) Intersection(other *Uint64Set) *Uint64Set {
	inter := &Uint64Set{uints: make(map[uint64]bool)}
	for member := range us.uints {
		if other.Contains(member) {
			inter.Add(member)
		}
	}
	return inter
}

func (us *Uint64Set) AsSlice() []uint64 {
	sl := make([]uint64, 0, us.Size())
	for k := range us.uints {
		sl = append(sl, k)
	}
	sort.Slice(sl, func(i, j int) bool { return sl[i] < sl[j] })
	return sl
}

func (us *Uint64Set) Size() int {
	return len(us.uints)
}

func (us *Uint64Set) Iter(fn func(uint64)) {
	for n := range us.uints {
		fn(n)
	}
}
