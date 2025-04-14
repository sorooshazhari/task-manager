package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"task-manager/entity"
)

type UserFileStore struct {
	filepath string
}

func NewUserFileStorage(path string) UserFileStore {
	return UserFileStore{
		filepath: path,
	}
}

func (fStore UserFileStore) Save(user entity.User) error {
	data, mErr := json.Marshal(user)
	if mErr != nil {
		return fmt.Errorf("error in marshalling %v", mErr)
	}
	file, oErr := os.OpenFile(fStore.filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if oErr != nil {
		return fmt.Errorf("error in oppening file %v", oErr)
	}
	defer file.Close()
	data = append(data, []byte("\n")...)
	_, wErr := file.Write(data)
	if wErr != nil {
		return fmt.Errorf("error in writing file %v", wErr)
	}
	return nil
}

func (fStore UserFileStore) Load() ([]entity.User, error) {
	var uStore = []entity.User{}

	if _, err := os.Stat(fStore.filepath); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", fStore.filepath)
		return []entity.User{}, nil
	}
	file, oErr := os.Open(fStore.filepath)
	if oErr != nil {
		return []entity.User{}, fmt.Errorf("error in oppening file %v", oErr)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, rErr := reader.ReadString('\n')
		if len(line) > 0 {
			var u entity.User
			uErr := json.Unmarshal([]byte(line), &u)
			if uErr != nil {
				return []entity.User{}, fmt.Errorf("error in unmarshalling line %v", uErr)
			}
			uStore = append(uStore, u)
		}
		if rErr != nil {
			if rErr == io.EOF {
				break
			} else {
				return []entity.User{}, fmt.Errorf("error in reading line from file %v", rErr)
			}
		}
	}
	return uStore, nil
}
