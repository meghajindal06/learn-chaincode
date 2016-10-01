/*
Copyright IBM Corp 2016 All Rights Reserved.

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

package main

import (
"encoding/json"	
"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}



type Account struct {
	ID          string  `json:"id"`
	AccountNumber string `json:"iban"`
	Balance float64 `json:"balance"`
}

type Milestone struct {
	ID          string  `json:"id"`
	Name string `json:"name"`
	CurrentStatus string `json:"currentStatus"`
	PaymentAmount float64 `json:"paymentAmount"`
	PaymentDate *time.Time `json:"paymentAmount"`
	PossibleActions []string `json:"possibleActions"`
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	//create  accounts list
	t.createAccounts(stub);

	t.milestones(stub);

	//initialize milestones

//	err := stub.PutState("hello_world", []byte(args[0]))

	return nil, nil
}

func (t *SimpleChaincode) createAccounts(stub *shim.ChaincodeStub )  {

	var loanAccount = Account{ID: "loanaccount",AccountNumber: "NL75ABNA0123456789",  Balance: 10000000.0} 


	loanAccountBytes, err := json.Marshal(&loanAccount)
    	if err != nil {
        	fmt.Println("error creating account" + loanAccount.ID)

    	}

	err = stub.PutState("loanaccount", loanAccountBytes)
                
                if err == nil {
                    fmt.Println("created account" + loanAccount.ID)
                } else {
                    fmt.Println("failed to create initialize account for " + loanAccount.ID)
                }	


               var contractoraccount = Account{ID: "contractoraccount",AccountNumber: "NLINGB053412537",  Balance: 0.0} 


	contractoraccountBytes, err := json.Marshal(&contractoraccount)
    	if err != nil {
        	fmt.Println("error creating account" + "contractoraccount")

    	}
err = stub.PutState("contractoraccount", contractoraccountBytes)
                
                if err == nil {
                    fmt.Println("created account" + "contractoraccount")
                } else {
                    fmt.Println("failed to create initialize account for " + "contractoraccount")
                }	


 
}

func (t *SimpleChaincode) createMilestomes(stub *shim.ChaincodeStub )  {

	var milestones = []Milestone{{ID: "1" ,Name: "FLOOR" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 , PaymentDate : "" , PossibleActions : []string{"STARTED"}},
{ID: "2" ,Name: "WALL" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 , PaymentDate : nil , PossibleActions : []string{"STARTED"}},
{ID: "3" ,Name: "ROOF" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 , PaymentDate : nil , PossibleActions : []string{"STARTED"}},
{ID: "4" ,Name: "DOOR" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 , PaymentDate : nil , PossibleActions : []string{"STARTED"}}}


	milestonesBytes, err := json.Marshal(&milestones)
    	if err != nil {
        	fmt.Println("error creating milestones")

    	}

	err = stub.PutState("milestones", milestonesBytes)
                
        if err == nil {
            fmt.Println("created milestones" )
        } else {
            fmt.Println("failed to create milestones ")
        }	
}


//Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

func GetAccount(stub *shim.ChaincodeStub , accountid string) (Account,error){
	var account Account

	accountBytes, err := stub.GetState(accountid)
	if err != nil {
		fmt.Println("Error retrieving account " + accountid)
		return account, errors.New("Error retrieving account " + accountid)
	}
		
	err = json.Unmarshal(accountBytes, &account)
	if err != nil {
		fmt.Println("Error unmarshalling account " + accountid)
		return account, errors.New("Error unmarshalling account " + accountid)
	}
		
	return account, nil
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

if args[0] == "GetAccountDetails" {
		fmt.Println("Getting account details")
		account, err := GetAccount(stub, args[1])
		if err != nil {
			fmt.Println("Error Getting particular account")
			return nil, err
		} else {
			accountBytes, err1 := json.Marshal(&account)
			if err1 != nil {
				fmt.Println("Error marshalling account")
				return nil, err1
			}	
			return accountBytes, nil		 
		}
	}

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
