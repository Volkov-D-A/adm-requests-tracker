package api

import (
	"encoding/json"
	"fmt"

	"github.com/volkov-d-a/adm-requests-tracker/internal/models"
	"github.com/volkov-d-a/adm-requests-tracker/pkg/utils"
)

func getTokenData(token string, key string) (*models.UserToken, error) {
	var ut models.UserToken
	str, err := utils.DecryptToken(key, token)
	if err != nil {
		return nil, fmt.Errorf("error decrypting token: %v", err)
	}

	err = json.Unmarshal([]byte(str), &ut)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling userdata: %v", err)
	}

	return &ut, nil
}

func getUserToken(ur *models.UserToken, key string) (string, error) {
	roleJson, err := json.Marshal(ur)
	if err != nil {
		return "", fmt.Errorf("error marshalling data: %v", err)
	}

	token, err := utils.EncryptToken(key, string(roleJson))
	if err != nil {
		return "", fmt.Errorf("error encrypting token: %v", err)
	}

	return token, nil
}
