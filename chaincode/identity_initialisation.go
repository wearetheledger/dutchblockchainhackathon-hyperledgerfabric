package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) initTables(stub shim.ChaincodeStubInterface) ([]byte, error) {

	// AssetPermissionTable
	err := stub.CreateTable("AssetPermissionTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "AssetOwnerID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "AssetPermissionedID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "AssetID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "PermissionedRevoked", Type: shim.ColumnDefinition_BOOL, Key: true},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed creating Permission table, [%v]", err)
	}
	fmt.Println("INIT - AssetPermissionTable created")

	return nil, nil

}
