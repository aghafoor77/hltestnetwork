package chaincode
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset defines the structure of the asset
type KeyidPubkey struct {
	Keyid       string   `json:"keyid"`
	Pubkey  string `json:"pubkey"`
	Algo  string `json:"algo"`
	Status  int `json:"status"`
}


// SmartContract is the chaincode contract
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger with sample KeyidPubkey
func (c *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	keyidpubkey := []KeyidPubkey{
		{
			Keyid:       "1",
			Pubkey:		"0x34343434354",
			Algo: "RSA", 
			Status: 0,
		},
		{
			Keyid:       "2",
			Pubkey: 	"0x4568092754237043",
			Algo: "RSA", 
			Status: 0,
		},
	}

	for _, keyidpub := range keyidpubkey {
		statuspubkeyJSON, err := json.Marshal(keyidpub)
		if err != nil {
			return fmt.Errorf("failed to marshal asset JSON: %v", err)
		}

		err = ctx.GetStub().PutState(keyidpub.Keyid, statuspubkeyJSON)
		if err != nil {
			return fmt.Errorf("failed to put asset on ledger: %v", err)
		}
	}

	return nil
}

// GetAllAssets returns all assets on the ledger
func (c *SmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]KeyidPubkey, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get assets: %v", err)
	}
	defer resultsIterator.Close()

	var keyidpubkey []KeyidPubkey

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over assets: %v", err)
		}

		keyidpub := KeyidPubkey{}
		err = json.Unmarshal(queryResponse.Value, &keyidpub)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
		}

		keyidpubkey = append(keyidpubkey, keyidpub)
	}

	return keyidpubkey, nil
}

// WriteAsset writes a new asset to the ledger
func (c *SmartContract) WriteAsset(ctx contractapi.TransactionContextInterface, keyid string, pubkey string, algo string, 
			status int) error {
	asset := KeyidPubkey{
		Keyid:      keyid,
		Pubkey: 	pubkey,
		Algo: 		algo, 
		Status: 		status,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(keyid, assetJSON)
}

// SearchAsset searches for an existing asset on the ledger
func (c *SmartContract) SearchAsset(ctx contractapi.TransactionContextInterface, keyid string) (*KeyidPubkey, error) {
	assetJSON, err := ctx.GetStub().GetState(keyid)
	if err != nil {
		return nil, fmt.Errorf("failed to read keyid-pubkey asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("keyid-pubkey asset with keyid %s does not exist", keyid)
	}

	asset := &KeyidPubkey{}
	err = json.Unmarshal(assetJSON, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal keyid-pubkey asset JSON: %v", err)
	}
	return asset, nil
}

// UpdateAsset updates an existing asset on the ledger by replacing old value to the new value
func (c *SmartContract) ChangeStatus(ctx contractapi.TransactionContextInterface, keyid string, status int) error {
	assetJSON, err := ctx.GetStub().GetState(keyid)
	if status < 0 || status >1 {
		return fmt.Errorf("Status should be either 0 or 1 but it is : %v", status)
	} 

	if err != nil {
		return fmt.Errorf("failed to read keyid-pup asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("keyid-pup asset with keyid %s does not exist", keyid)
	}

	asset := KeyidPubkey{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal keyid-pup asset JSON: %v", err)
	}

	asset.Status = status

	updatedAssetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated keyid-pup asset JSON: %v", err)
	}

	err = ctx.GetStub().PutState(keyid, updatedAssetJSON)
	if err != nil {
		return fmt.Errorf("failed to put updated keyid-pup asset on ledger: %v", err)
	}

	return nil
}