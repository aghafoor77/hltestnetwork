package chaincode
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Asset defines the structure of the asset
type EmailToDid struct {
	Email       string   `json:"email"`
	Did []string `json:"dids"`
}

// SmartContract is the chaincode contract
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the ledger with sample EmailToDids
func (c *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	/*emailtodids := []EmailToDid{
		{
			Email:       "john@example.com",
			Did: []string{"1234567890"},
		},
		{
			Email:       "jane@example.com",
			Did: []string{"12345"},
		},
	}

	for _, emailtodid := range emailtodids {
		emailtodidJSON, err := json.Marshal(emailtodid)
		if err != nil {
			return fmt.Errorf("failed to marshal asset JSON: %v", err)
		}

		err = ctx.GetStub().PutState(emailtodid.Email, emailtodidJSON)
		if err != nil {
			return fmt.Errorf("failed to put asset on ledger: %v", err)
		}
	}*/

	return nil
}

// GetAllAssets returns all assets on the ledger
func (c *SmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]EmailToDid, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get assets: %v", err)
	}
	defer resultsIterator.Close()

	var emailtodids []EmailToDid

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over assets: %v", err)
		}

		emailtodid := EmailToDid{}
		err = json.Unmarshal(queryResponse.Value, &emailtodid)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal asset: %v", err)
		}

		emailtodids = append(emailtodids, emailtodid)
	}

	return emailtodids, nil
}

// WriteAsset writes a new asset to the ledger
func (c *SmartContract) WriteAsset(ctx contractapi.TransactionContextInterface, email string, did []string) error {
	asset := EmailToDid{
		Email:       email,
		Did: 	     did,
	}
	
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(email, assetJSON)
}

// UpdateAsset updates an existing asset on the ledger by replacing old value to the new value
func (c *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, email string, did string) error {
	assetJSON, err := ctx.GetStub().GetState(email)
	if err != nil {
		return fmt.Errorf("failed to read email-did asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("email-did asset with email %s does not exist", email)
	}

	asset := EmailToDid{}
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return fmt.Errorf("failed to unmarshal email-did asset JSON: %v", err)
	}
	asset.Did = append(asset.Did, did)
	//asset.Did = did

	updatedAssetJSON, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal updated email-did asset JSON: %v", err)
	}

	err = ctx.GetStub().PutState(email, updatedAssetJSON)
	if err != nil {
		return fmt.Errorf("failed to put updated email-did asset on ledger: %v", err)
	}

	return nil
}

// SearchAsset searches for an existing asset on the ledger
func (c *SmartContract) SearchAsset(ctx contractapi.TransactionContextInterface, email string) (*EmailToDid, error) {
	assetJSON, err := ctx.GetStub().GetState(email)
	if err != nil {
		return nil, fmt.Errorf("failed to read email-did asset from ledger: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("email-did asset with email %s does not exist", email)
	}

	asset := &EmailToDid{}
	err = json.Unmarshal(assetJSON, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal email-did asset JSON: %v", err)
	}

	return asset, nil
}
