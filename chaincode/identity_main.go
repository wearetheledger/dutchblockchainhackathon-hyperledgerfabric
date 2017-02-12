package main

import (
	"errors"
	"fmt"
	_ "go/types"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//AssetPermission is something
type AssetPermission struct {
	AssetOwnerID        string `json:"assetOwnerID"`
	AssetPermissionedID string `json:"assetPermissionedID"`
	AssetID             string `json:"assetID"`
	PermissionedRevoked bool   `json:"permissionedRevoked"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ##### MAIN ##### //
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init is executed 1 time when the chaincode is deployed
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	t.initTables(stub)

	//return nothing for the moment
	return nil, nil
}

// Invoke is called
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Create a Permission
	if function == "create_assetPermission" {
		if len(args) < 3 {
			return nil, errors.New("Incorrect number of arguments. Expecting > 3")
		}
		return t.createAssetPermission(stub, args)
	}

	// Revoke a Permission
	if function == "revoke_assetPermission" {
		if len(args) < 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting > 2")
		}
		return t.revokeAssetPermission(stub, args)
	}

	//return nothing for the moment
	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Query a Permission
	if function == "retrieve_assetPermission" {
		if len(args) < 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting > 2")
		}
		return t.retrieveAssetPermission(stub, args)
	}

	//return nothing for the moment
	return nil, nil

}
