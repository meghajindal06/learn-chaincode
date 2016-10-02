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
	  "strings"

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
	PossibleActions []string `json:"possibleActions"`
}


type MilestoneHistory struct {
	ID     string  `json:"id"`
	Status string `json:"Status"`
	ActionDate time.Time `json:"actionDate"`
	
}


type Transaction struct {
	ID     string  `json:"id"`
	MilestoneID     string  `json:"milestoneId"`
	FromAccount string `json:"fromAccount"`
	ToAccount string `json:"toAccount"`
	Amount float64 `json:"amount"`
	PaymentDate time.Time `json:"paymentAmount"`
	
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
	
	if function == "init" {
		t.createAccounts(stub);

		t.createMilestomes(stub);
	}else{
		t.createMilestoneLifecycle(stub)
	}
	

	//initialize milestones

//	err := stub.PutState("hello_world", []byte(args[0]))

	return nil, nil
}


func (t *SimpleChaincode) createMilestoneLifecycle(stub *shim.ChaincodeStub )  {
	//createAccounts
	var loanAccount = Account{ID: "loanaccount",AccountNumber: "NL75ABNA0123456789",  Balance: 115000} 
	fmt.Println(" creating account" )

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


               var contractoraccount = Account{ID: "contractoraccount",AccountNumber: "NLINGB053412537",  Balance: 5000.0} 


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

                //create milestones
                var milestones = []Milestone{{ID: "1" ,Name: "FLOOR" , CurrentStatus : "DONE" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "2" ,Name: "WALL" , CurrentStatus : "START" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "3" ,Name: "ROOF" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "4" ,Name: "DOOR" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}}}


	milestonesBytes, err := json.Marshal(&milestones)
    	if err != nil {
        	fmt.Println("error creating milestones")

    	}

	err = stub.PutState("milestones", milestonesBytes)
                
        if err == nil {
            fmt.Println("created milestones successfully" )
        } else {
            fmt.Println("failed to create milestones in init")
        }	

        //create milestone history
        var milestonehistoryArray = []MilestoneHistory{{ID : "1"  , Status : "START" , ActionDate : time.Now()},{ID : "milestonehistory_1"  , Status : "DONE" , ActionDate : time.Now()},{ID : "milestonehistory_1"  , Status : "REJECT" , ActionDate : time.Now()},{ID : "milestonehistory_1"  , Status : "DONE" , ActionDate : time.Now()},{ID : "milestonehistory_1"  , Status : "ACCEPT" , ActionDate : time.Now()}};
        milestonehistoryArrayBytes, err := json.Marshal(&milestonehistoryArray)
    	if err != nil {
        	fmt.Println("error creating milestones")

    	}

	err = stub.PutState("milestonehistory_1", milestonehistoryArrayBytes)
                
        if err == nil {
            fmt.Println("created milestones successfully" )
        } else {
            fmt.Println("failed to create milestones in init")
        }	

}

func (t *SimpleChaincode) createAccounts(stub *shim.ChaincodeStub )  {

	var loanAccount = Account{ID: "loanaccount",AccountNumber: "NL75ABNA0123456789",  Balance: 120000} 
	fmt.Println(" creating account" )

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

	var milestones = []Milestone{{ID: "1" ,Name: "FLOOR" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "2" ,Name: "WALL" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "3" ,Name: "ROOF" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}},
{ID: "4" ,Name: "DOOR" , CurrentStatus : "NOT_INITIATED" , PaymentAmount : 5000.0 ,  PossibleActions : []string{}}}


	milestonesBytes, err := json.Marshal(&milestones)
    	if err != nil {
        	fmt.Println("error creating milestones")

    	}

	err = stub.PutState("milestones", milestonesBytes)
                
        if err == nil {
            fmt.Println("created milestones successfully" )
        } else {
            fmt.Println("failed to create milestones in init")
        }	
}


//Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	}else if function == "init_lifecycle" {
		fmt.Println("init lifecycle")

		return t.Init(stub, "init_lifecycle", args)
	} else if function == "updateStatus" {
		return t.UpdateMilestoneStatus(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) GetAccount(stub *shim.ChaincodeStub , userId string) ([]byte,error){
	if userId == "admin" || userId == "user_type1_61da9cc943" {
		return stub.GetState("loanaccount")
	}else{
		return stub.GetState("contractoraccount")
	}

	
	
}


func (t *SimpleChaincode) UpdateMilestoneStatus(stub *shim.ChaincodeStub , args []string) ([]byte,error){
	fmt.Println("Attempting to Update milestone status")
	var userId = args[0]
	var milestoneId = args[1]
	var action = args[2]
	var err error
	//validate action allowed
	var validate = t.ValidateAction(userId , action)
	if(!validate) {
		return nil, errors.New("Action " + action + " not permitted for user " + userId)
	}


	//update milestone array with status
	err = t.UpdateMilestoneSummary(stub ,milestoneId , action)
	if(err != nil) {
		return nil, errors.New("Error while updating summary")
	}
	//udate milestone history with status
	err = t.UpdateMilestoneHistory(stub ,milestoneId ,action)
	if(err != nil) {
		return nil, errors.New("Error while updating history")
	}

	//if action accept create transaction 

	if action == "ACCEPT" {
		t.CreateTransaction(stub, milestoneId , 1000.0)
	}

	return nil,nil
	
}

func (t *SimpleChaincode) CreateTransaction(stub *shim.ChaincodeStub , milestoneId string , amount float64) (error){
	
	var loanaccount Account
	var contractoraccount Account
	var accountBytes []byte
	var err error
	//debit from loan
	accountBytes,err = stub.GetState("loanaccount")
	if(err != nil){
		// there is no account present already
		 return errors.New("No loan account found")
    }else{
		err := json.Unmarshal(accountBytes , &loanaccount)
		if(err!= nil){
			return errors.New("error unmarshalling loanaccount")
		}else{
			loanaccount.Balance = loanaccount.Balance - amount;
			accountBytes,err := json.Marshal(&loanaccount)
			err = stub.PutState("loanaccount" , accountBytes)
			 if err == nil {
            	fmt.Println("updated loanaccount history" )
        	} else {
           		fmt.Println("failed update loanaccount history ")
            	return errors.New("error updating loanaccount history")
        	}	
		}

	}

	//credit in contractor
	accountBytes,err = stub.GetState("contractoraccount")
	if(err != nil){
		// there is no account present already
		 return errors.New("No contractor account found")
    }else{
		err := json.Unmarshal(accountBytes , &contractoraccount)
		if(err!= nil){
			return errors.New("error unmarshalling contractoraccount")
		}else{
			contractoraccount.Balance = contractoraccount.Balance + amount;
			accountBytes,err := json.Marshal(&contractoraccount)
			err = stub.PutState("contractoraccount" , accountBytes)
			 if err == nil {
            	fmt.Println("updated contractoraccount history" )
        	} else {
           		fmt.Println("failed update contractoraccount history ")
            	return errors.New("error updating contractoraccount history")
        	}	
		}

	}


	//Update transaction asdfasdfssssss
	var transactions []Transaction
	transactionArrayBytes,err := stub.GetState("transactions")
	var transaction = Transaction{ID : "trx_" + milestoneId , MilestoneID : milestoneId , FromAccount : loanaccount.AccountNumber , ToAccount : contractoraccount.AccountNumber , Amount : amount, PaymentDate : time.Now()};
	if(err != nil){
		// there is no history present already
		 if strings.Contains(err.Error(), "unexpected end") {
		 	transactions = []Transaction{transaction}
		}else {
                return errors.New("Error unmarshalling existing transaction history" )
           }
	}else{
		err := json.Unmarshal(transactionArrayBytes , &transactions)
		if(err!= nil){
			return errors.New("error unmarshalling transaction history")
		}else{
			transactions = ExtendTransactionArray(transactions , transaction)
		}

	}

	transactionArrayBytes,err = json.Marshal( &transactions)
		if(err != nil){
			return errors.New("error mashalling milestone history")
		}
	err = stub.PutState("transactions" , transactionArrayBytes)

	 if err == nil {
            fmt.Println("updated milestones history" )
        } else {
            fmt.Println("failed update transaction history ")
            return errors.New("error updating transaction history")
        }	
        return nil

}


func (t *SimpleChaincode) UpdateMilestoneHistory(stub *shim.ChaincodeStub , milestoneId string , action string) (error){

	var milestoneHistoryArray []MilestoneHistory
	milestoneHistoryArrayBytes,err := stub.GetState("milestonehistory_" + milestoneId)
	var milestonehistory = MilestoneHistory{ID : milestoneId , Status : action , ActionDate : time.Now()};
	if(err != nil){
		// there is no history present already
		 if strings.Contains(err.Error(), "unexpected end") {
		 	milestoneHistoryArray = []MilestoneHistory{milestonehistory}
		}else {
                return errors.New("Error unmarshalling existing history for id milestonehistory_" + milestoneId)
           }
	}else{
		err := json.Unmarshal(milestoneHistoryArrayBytes , &milestoneHistoryArray)
		if(err!= nil){
			return errors.New("error unmarshalling milestone history")
		}else{
			milestoneHistoryArray = Extend(milestoneHistoryArray , MilestoneHistory{ID: milestoneId , Status : action , ActionDate : time.Now()})
		}

	}

	milestoneHistoryArrayBytes,err = json.Marshal( &milestoneHistoryArray)
		if(err != nil){
			return errors.New("error mashalling milestone history")
		}
	err = stub.PutState("milestonehistory_" + milestoneId , milestoneHistoryArrayBytes)

	 if err == nil {
            fmt.Println("updated milestones history" )
        } else {
            fmt.Println("failed update milestones history ")
            return errors.New("error updating milestone history")
        }	
        return nil
}

func (t *SimpleChaincode) UpdateMilestoneSummary(stub *shim.ChaincodeStub , milestoneId string , action string) (error){

	var milestones []Milestone
	var err error
	var i int
	milestoneSummaryBytes,err := stub.GetState("milestones")

	 	if err != nil {
	 		fmt.Println("error retrieving milestones ")
           return errors.New("error retrieving milestones for id" + milestoneId)
        } 


       err = json.Unmarshal(milestoneSummaryBytes , &milestones)
       if err != nil {
       	fmt.Println("error unmarshalling milestones for id")
           return errors.New("error unmarshalling milestones for id" + milestoneId)
        } 

        for i= 0; i<4 ;i++ {
        	if(milestones[i].ID == milestoneId){
        		milestones[i].CurrentStatus = action
			}
        }	
        fmt.Println( milestones)
       milestonesBytes, err := json.Marshal(&milestones)
    	if err != nil {
        	fmt.Println("error marshalling milestones")
        	return errors.New("error marshalling milestones" )

    	}

		err = stub.PutState("milestones", milestonesBytes)
                
        if err == nil {
            fmt.Println("updated milestones" )
        } else {
            fmt.Println("failed to update milestones ")
            return errors.New("failed to update milestone id" + milestoneId)
        }	

        return nil

}

func (t *SimpleChaincode) ValidateAction(userId string , action string) (bool){

	if userId == "admin" {
		fmt.Println("admin:" + action == "ACCEPT" || action == "REJECT" )
		return (action == "ACCEPT" || action == "REJECT" )

	}else if userId == "user_type1_5894b6a76a" {
		fmt.Println("user_type1_5894b6a76a:" + action == "START" || action == "DONE" )
		return (action == "START" || action == "DONE")
	}else{
		return false
	}

}

func (t *SimpleChaincode) GetMilestoneHistory(stub *shim.ChaincodeStub , milestoneId string) ([]byte,error){
	
	return stub.GetState("milestonehistory_" + milestoneId)
	
}

func (t *SimpleChaincode) GetMilestones(stub *shim.ChaincodeStub , userId string) ([]byte,error){
	var milestones []Milestone
	fmt.Println("Invoked Get Milestones method")
	milestoneArrayBytes, err := stub.GetState("milestones")
	if err != nil {
		fmt.Println("Error retrieving milestones ")
		return nil, errors.New("Error retrieving milestones ")
	}
	err = json.Unmarshal(milestoneArrayBytes, &milestones)
	if err != nil {
		fmt.Println("Error unmarshalling milestones ")
		return nil, errors.New("Error unmarshalling milestones ")
	}
	
	if userId == "admin" {
		milestones = t.populateActionForCustomer(milestones)
	}else if userId == "user_type1_5894b6a76a"{
		milestones =  t.populateActionForContractor(milestones)
	}

	milestoneArrayBytes, err = json.Marshal(&milestones)
	if err != nil {
		fmt.Println("Error retrieving milestones ")
		return nil, errors.New("Error retrieving milestones ")
	}
		
	return milestoneArrayBytes, nil
}


func (t *SimpleChaincode) populateActionForContractor(milestones []Milestone) ([]Milestone) {
		
	var updatedMileStones = milestones
	var i int
	for i = 0; i < 4; i++ {
		var possibleActions []string
		
		switch(updatedMileStones[i].CurrentStatus) {
      		case "NOT_INITIATED" :
      			if(i ==0 || updatedMileStones[i-1].CurrentStatus == "ACCEPT"){
      				possibleActions = []string{"START"}
      				}else{
      					possibleActions = []string{}
      				}
      				
      		case "START" :
      				possibleActions = []string{"DONE"}
      		case "REJECT" :
      				possibleActions = []string{"DONE"}
      		default :
      				possibleActions = []string{}
      	}
      			updatedMileStones[i].PossibleActions = possibleActions

    }
	return updatedMileStones 
	

}

func Extend(slice []MilestoneHistory, element MilestoneHistory) []MilestoneHistory {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]MilestoneHistory, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}

func ExtendTransactionArray(slice []Transaction, element Transaction) []Transaction {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]Transaction, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}


func (t *SimpleChaincode) populateActionForCustomer( milestones []Milestone) ([]Milestone) {
	
	var updatedMileStones = milestones
	var i int
	for i = 0; i < 4; i++ {
		var possibleActions = []string{}
		switch(updatedMileStones[i].CurrentStatus) {
      		case "DONE" :
      				possibleActions = []string{"ACCEPT","REJECT"}
      		default :
      				possibleActions = []string{}
      	}
      			updatedMileStones[i].PossibleActions = possibleActions

    }
	return updatedMileStones 
}



// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "GetAccountDetails" {
		fmt.Println("Getting account details")
		return t.GetAccount(stub, args[0])
		
	}

	if function == "GetMilestones" {
		fmt.Println("Getting milestones")
		return t.GetMilestones(stub, args[0])
		
	}

	if function == "GetMilestoneHistory" {
		fmt.Println("Getting milestone history")
		return t.GetMilestoneHistory(stub, args[0])
		
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