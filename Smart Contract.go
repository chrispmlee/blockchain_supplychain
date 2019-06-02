/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Order struct {
	OrderId   string `json:"orderid"`
	Owner  string `json:"owner"`
	Distributor  string `json:"distributor"`
	DrugId string `json:"drugid"`
	Quantity int `json:"quantity"`
	Price float `json:"price"`
	Status string `json:"status"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract 
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
    if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "orderTransaction" {
		return s.orderTransaction(APIstub, args)
	}  else if function == "createOrder" {
		return s.createOrder(APIstub, args)
	} 
	return shim.Error("Invalid Smart Contract function name.")
}


func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	orders := []Orders{
		Order{OrderId: "1", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "117", Quantity: 100, Price: 13.4, Status: 'Ordered'},
	 	Order{OrderId: "2", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "777", Quantity: 100, Price: 144, Status: 'Ordered'},
	 	Order{OrderId: "3", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "631", Quantity: 100, Price: 11.4, Status: 'Ordered'},
	 	Order{OrderId: "4", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "222", Quantity: 111, Price: 133.4, Status: 'Ordered'},
	 	Order{OrderId: "5", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "41", Quantity: 111, Price: 1313.4, Status: 'Ordered'},
	 	Order{OrderId: "6", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "17", Quantity: 111, Price: 1153.4, Status: 'Ordered'},
	 	Order{OrderId: "7", Owner:"Kaiser Permanente", Distributor:"Pfizer", DrugId: "1", Quantity: 111, Price: 1361.4, Status: 'Ordered'}
	 }

	i := 0
	for i < len(orders) {
		fmt.Println("i is ", i)
		ordersAsBytes, _ := json.Marshal(orders[i])
		APIstub.PutState("ORDER"+strconv.Itoa(i), orders)
		fmt.Println("Added", order)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createOrder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	if(RegisteredHospitals[args[2]]){
		return shim.Error('Invalid Hospitals cannot create Order')
	}

	if(RegisteredDistributors[args[3]]){
		return shim.Error('Invalid Distributor cannot create Order')
	}
	var order = Order{OrderId: args[1], Owner: args[2], Distributor: args[3], DrugId: args[4], Quantity: args[4], Price: args[5], Status: args[6]}

	orderAsBytes, _ := json.Marshal(order)
	APIstub.PutState(args[0], orderAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) orderTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if(args[2].to != Order(OrderId).Owner)
		return shim.Error('Invalid Owner trying to make changes')

	if(args[3].to != Order(OrderId).Distributor)
		return shim.Error('Invalid Distributor trying to make changes')
 
 	if(args[3].to != Order(OrderId).Distributor)
		return shim.Error('Invalid Distributor trying to make changes')

	orderAsBytes, _ := APIstub.GetState(args[0])
	order := Order{}

	json.Unmarshal(ordersAsBytes, &order)
	order.Status = args[4]

	orderAsBytes, _ = json.Marshal(order)
	APIstub.PutState(args[0], orderAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}