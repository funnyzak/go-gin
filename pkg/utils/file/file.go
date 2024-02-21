package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

func FilExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func WriteToFile(filePath string, content string, filemode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	if err := os.WriteFile(filePath, []byte(content), filemode); err != nil {
		return err
	}
	return nil
}

func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)
	md5Hash := hex.EncodeToString(hashInBytes)

	return md5Hash, nil
}

func MkdirAllIfNotExists(pathname string, perm os.FileMode) error {
	if _, err := os.Stat(pathname); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(pathname, perm); err != nil {
				return err
			}
		}
	}
	return nil
}
