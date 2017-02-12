package main

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetPermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	assetOwnerID := args[0]
	assetPermissionedID := args[1]
	assetID := args[2]

	if assetOwnerID == assetPermissionedID {
		return nil, errors.New("assetOwnerID and assetPermissionedID cannot be the same")
	}

	ap := &AssetPermission{
		AssetOwnerID:        assetOwnerID,
		AssetPermissionedID: assetPermissionedID,
		AssetID:             assetID,
		PermissionedRevoked: false,
	}

	pRes, _ := json.Marshal(ap)
	fmt.Print("INFO Invoke - CreatePermission: ")
	fmt.Println(string(pRes))

	fmt.Printf("New Asset permission is [%s] [%s] [%s] [%t]", ap.AssetOwnerID, ap.AssetPermissionedID, ap.AssetID, ap.PermissionedRevoked)

	ok, err := stub.InsertRow("AssetPermissionTable", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: ap.AssetOwnerID}},
			&shim.Column{Value: &shim.Column_String_{String_: ap.AssetPermissionedID}},
			&shim.Column{Value: &shim.Column_String_{String_: ap.AssetID}},
			&shim.Column{Value: &shim.Column_Bool{Bool: ap.PermissionedRevoked}},
		},
	})
	if !ok && err == nil {
		fmt.Println("Error inserting new row in AssetPermissionTable: ", err)
		return nil, errors.New("Asset was already assigned.")
	}

	AssetPermissionJSON, _ := json.Marshal(ap)
	return AssetPermissionJSON, nil

}

func (t *SimpleChaincode) revokeAssetPermission(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil

}
