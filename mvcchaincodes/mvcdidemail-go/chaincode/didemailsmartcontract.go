package chaincode
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset defines the structure of the asset
type DidEmail struct {
	Did       string   `json:"did"`
	Email  string `json:"email"`
}

// SmartContract is the chaincode contract
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger with sample DidEmails
func (c *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	/*didemails := []DidEmail{
		{
			Did:       "123456",
			Email:		"a@gmail.com",
		},
		{
			Did:       "45678",
			Email: 	"b@gmail.com",
		},
	}

	for _, didemail := range didemails {
		statuspubkeyJSON, err := json.Marshal(didemail)
		if err != nil {
			return fmt.Errorf("failed to marshal asset JSON: %v", err)
		}

		err = ctx.GetStub().PutState(didemail.Did, statuspubkeyJSON)
		if err != nil {
			return fmt.Errorf("failed to put asset on ledger: %v", err)
		}
	}*/

	return nil
}

// GetAllAssets returns all assets on the ledger
func (c *SmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]DidEmail, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get assets: %v", err)
	}
	defer resultsIterator.Close()

	var didemails []DidEmail

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over assets: %v", err)
		}

		didemail := DidEmail{}
		err = json.Unmarshal(queryResponse.Value, &didemail)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
		}

		didemails = append(didemails, didemail)
	}

	return didemails, nil
}

// WriteAsset writes a new asset to the ledger
func (c *SmartContract) WriteAsset(ctx contractapi.TransactionContextInterface, did string, email string) error {
	asset := DidEmail{
		Did:       did,
		Email: 		 email,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(did, assetJSON)
}

// UpdateAsset updates an existing asset on the ledger by replacing old value to the new value
/*func (c *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, did string, email string) error {
	assetJSON, err := ctx.GetStub().GetState(did)
	if err != nil {
		return fmt.Errorf("failed to read did-email asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("did-email asset with did %s does not exist", did)
	}

	asset := DidEmail{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal did-email asset JSON: %v", err)
	}

	asset.Email = email

	updatedAssetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated did-email asset JSON: %v", err)
	}

	err = ctx.GetStub().PutState(did, updatedAssetJSON)
	if err != nil {
		return fmt.Errorf("failed to put updated did-email asset on ledger: %v", err)
	}

	return nil
}*/

// SearchAsset searches for an existing asset on the ledger
func (c *SmartContract) SearchAsset(ctx contractapi.TransactionContextInterface, did string) (*DidEmail, error) {
	assetJSON, err := ctx.GetStub().GetState(did)
	if err != nil {
		return nil, fmt.Errorf("failed to read did-email asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("did-email asset with did %s does not exist", did)
	}

	asset := &DidEmail{}
	err = json.Unmarshal(assetJSON, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal did-email asset JSON: %v", err)
	}
	return asset, nil
}
