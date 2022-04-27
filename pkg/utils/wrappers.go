/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package utils

import "github.com/golang/protobuf/ptypes/wrappers"

// WrapperUint32 uint32
func WrapperUint32(value uint32) *wrappers.UInt32Value {
	return &wrappers.UInt32Value{Value: value}
}

// WrapperString string
func WrapperString(value string) *wrappers.StringValue {
	return &wrappers.StringValue{Value: value}
}

// WrapperInt64 int64
func WrapperInt64(value int64) *wrappers.Int64Value {
	return &wrappers.Int64Value{Value: value}
}

// WrapperBool bool
func WrapperBool(value bool) *wrappers.BoolValue {
	return &wrappers.BoolValue{Value: value}
}
