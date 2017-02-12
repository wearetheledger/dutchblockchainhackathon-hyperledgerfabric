package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) retrieveAssetPermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("INFO query - Retrieving permissions")

	assetOwnerID := args[0]
	assetPermissionedID := args[1]

	var columns []shim.Column

	rowChannel, err := stub.GetRows("AssetPermissionTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed retrieving AssetPermissions of assetOwnerID " + assetOwnerID + " and assetPermissionedID" + assetPermissionedID + ". Error " + err.Error() + ". \"}"
		return nil, errors.New(jsonResp)
	}
	var rows []shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, row)
			}
		}
		if rowChannel == nil {
			break
		}
	}

	var filteredRows []AssetPermission
	for _, row := range rows {
		var ap AssetPermission
		if row.Columns[0].GetString_() == assetOwnerID && row.Columns[1].GetString_() == assetPermissionedID {
			ap.AssetOwnerID = row.Columns[0].GetString_()
			ap.AssetPermissionedID = row.Columns[1].GetString_()
			ap.AssetID = row.Columns[2].GetString_()
			ap.PermissionedRevoked = row.Columns[3].GetBool()
			filteredRows = append(filteredRows, ap)
		}

	}
	AssetPermissionJSON, _ := json.Marshal(filteredRows)
	if filteredRows == nil {
		fmt.Printf("ERROR query")
		jsonResp := "[]"
		AssetPermissionJSON, _ = json.Marshal(jsonResp)
		fmt.Println(AssetPermissionJSON)
		return AssetPermissionJSON, nil
	}
	fmt.Println("FILTEREDROWS")
	fmt.Println(filteredRows)

	fmt.Println("ASSETPERMISSION")
	fmt.Println(AssetPermissionJSON)
	return AssetPermissionJSON, nil

}
