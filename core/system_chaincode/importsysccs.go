/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package system_chaincode

import (
	"github.com/hyperledger/fabric/core/system_chaincode/api"
	//import system chain codes here
	"github.com/hyperledger/fabric/bddtests/syschaincode/noop"
	"github.com/hyperledger/fabric/core/system_chaincode/escc"
	"github.com/hyperledger/fabric/core/system_chaincode/lccc"
	"github.com/hyperledger/fabric/core/system_chaincode/vscc"
)

//see systemchaincode_test.go for an example using "sample_syscc"
var systemChaincodes = []*api.SystemChaincode{
	{
		Enabled:   true,
		Name:      "noop",
		Path:      "github.com/hyperledger/fabric/bddtests/syschaincode/noop",
		InitArgs:  [][]byte{},
		Chaincode: &noop.SystemChaincode{},
	},
	{
		Enabled:   true,
		Name:      "lccc",
		Path:      "github.com/hyperledger/fabric/core/system_chaincode/lccc",
		InitArgs:  [][]byte{[]byte("")},
		Chaincode: &lccc.LifeCycleSysCC{},
	},
	{
		Enabled:   true,
		Name:      "escc",
		Path:      "github.com/hyperledger/fabric/core/system_chaincode/escc",
		InitArgs:  [][]byte{[]byte("")},
		Chaincode: &escc.EndorserOneValidSignature{},
	},
	{
		Enabled:   true,
		Name:      "vscc",
		Path:      "github.com/hyperledger/fabric/core/system_chaincode/vscc",
		InitArgs:  [][]byte{[]byte("")},
		Chaincode: &vscc.ValidatorOneValidSignature{},
	}}

//RegisterSysCCs is the hook for system chaincodes where system chaincodes are registered with the fabric
//note the chaincode must still be deployed and launched like a user chaincode will be
func RegisterSysCCs() {
	for _, sysCC := range systemChaincodes {
		api.RegisterSysCC(sysCC)
	}
}
