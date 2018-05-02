/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except 
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and 
 * limitations under the License.
 */
 
package api

import (
	"configcenter/src/framework/core/input"
	"configcenter/src/framework/core/output"
	"time"
)

// RegisterInputer execute a non-blocking inputer, only execute once
func RegisterInputer(inputer input.Inputer, putter output.Puter, exceptionFunc input.ExceptionFunc) (input.InputerKey, error) {

	return mgr.InputerMgr.AddInputer(input.InputerParams{
		Target:    inputer,
		Kind:      input.ExecuteOnce,
		Putter:    putter,
		Exception: exceptionFunc,
	}), nil
}

// RegisterTimingInputer execute a non-blocking timing inputer, only execute once
func RegisterTimingInputer(inputer input.Inputer, putter output.Puter, frequency time.Duration, exceptionFunc input.ExceptionFunc) (input.InputerKey, error) {

	return mgr.InputerMgr.AddInputer(input.InputerParams{
		IsTiming:  true,
		Frequency: frequency,
		Target:    inputer,
		Kind:      input.ExecuteOnce,
		Putter:    putter,
		Exception: exceptionFunc,
	}), nil
}