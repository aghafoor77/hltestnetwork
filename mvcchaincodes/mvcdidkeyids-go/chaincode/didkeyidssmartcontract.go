package chaincode
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset defines the structure of the asset
type DidToKeyids struct {
	Did       string   `json:"did"`
	Keyids []string `json:"keyids"`
}

// SmartContract is the chaincode contract
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger with sample EmailToDids
func (c *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	/*didtokeyids := []DidToKeyids{
		{
			Did:       "john@example.com",
			Keyids: []string{"1234567890"},
		},
		{
			Did:       "jane@example.com",
			Keyids: []string{"12345"},
		},
	}

	for _, didtokeyid := range didtokeyids {
		didtokeyidJSON, err := json.Marshal(didtokeyid)
		if err != nil {
			return fmt.Errorf("failed to marshal asset JSON: %v", err)
		}

		err = ctx.GetStub().PutState(didtokeyid.Did, didtokeyidJSON)
		if err != nil {
			return fmt.Errorf("failed to put asset on ledger: %v", err)
		}
	}*/

	return nil
}

// GetAllAssets returns all assets on the ledger
func (c *SmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]DidToKeyids, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get assets: %v", err)
	}
	defer resultsIterator.Close()

	var didtokeyids []DidToKeyids

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over assets: %v", err)
		}

		didtokeyid := DidToKeyids{}
		err = json.Unmarshal(queryResponse.Value, &didtokeyid)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
		}

		didtokeyids = append(didtokeyids, didtokeyid)
	}

	return didtokeyids, nil
}

// WriteAsset writes a new asset to the ledger
func (c *SmartContract) WriteAsset(ctx contractapi.TransactionContextInterface, did string, keyids []string) error {
	asset := DidToKeyids{
		Did:       did,
		Keyids: 	keyids,
	}
	
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(did, assetJSON)
}

// UpdateAsset updates an existing asset on the ledger by replacing old value to the new value
func (c *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, did string, keyids string) error {
	assetJSON, err := ctx.GetStub().GetState(did)
	if err != nil {
		return fmt.Errorf("failed to read did-keyids asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("did-keyids asset with did %s does not exist", did)
	}

	asset := DidToKeyids{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal did-keyids asset JSON: %v", err)
	}
	asset.Keyids = append(asset.Keyids, keyids)
	//asset.Keyids = keyids

	updatedAssetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated did-keyids asset JSON: %v", err)
	}

	err = ctx.GetStub().PutState(did, updatedAssetJSON)
	if err != nil {
		return fmt.Errorf("failed to put updated did-keyids asset on ledger: %v", err)
	}

	return nil
}

// SearchAsset searches for an existing asset on the ledger
func (c *SmartContract) SearchAsset(ctx contractapi.TransactionContextInterface, did string) (*DidToKeyids, error) {
	assetJSON, err := ctx.GetStub().GetState(did)
	if err != nil {
		return nil, fmt.Errorf("failed to read did-keyids asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("did-keyids asset with did %s does not exist", did)
	}

	asset := &DidToKeyids{}
	err = json.Unmarshal(assetJSON, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal did-keyids asset JSON: %v", err)
	}

	return asset, nil
}