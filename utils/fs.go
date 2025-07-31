package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/goccy/go-json"
	cp "github.com/otiai10/copy"
)


func ReadJson(filePath string, jsonData any) (any, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil

}

// dir create
func DirCreate(fp string) (bool, error) {
	err := os.MkdirAll(fp, os.ModePerm)
	if err != nil {
		return false, err
	}
	return true, nil
}

// dir delete
func DirRemove(fp string) (bool, error) {
	err := os.RemoveAll(fp)
	if err != nil {
		return false, err
	}
	return true, nil
}

// dir move
func DirMove(src string, dst string) (bool, error) {
	return true, nil
}

// dir copy
func DirCopy(src string, dst string) (bool, error) {
	err := cp.Copy(src, dst)
	if err != nil {
		return false, err
	}
	return true, nil
}


// dir exists
func IsDirExists(fp string) bool {
   info, err := os.Stat(fp)
   if os.IsNotExist(err) {
      return false
   }
   return info.IsDir()
}

// file exists
func IsFileExists(fp string) bool {
   info, err := os.Stat(fp)
   if os.IsNotExist(err) {
      return false
   }
   return !info.IsDir()
}

// file delete
func FileDelete(fp string) (bool, error) {
	err := os.Remove(fp)
	if err != nil {
		return false, err
	}
	return true, nil
}


// file copy to another place
func FileCopy(srcFile string, dstFile string) (bool, error) {
	// Provide the path of source file
	src, err := os.Open(srcFile)
	if err != nil {
		return false, fmt.Errorf("failed to open %s @ %s", srcFile, err) 
	}
	defer src.Close()

	// check srcFile stats
	_, err = os.Stat(srcFile)
	if err != nil {
		return false, fmt.Errorf("failed to check stats for %s @ %s", srcFile, err) 
	}

	// print srcFile stats
	// perm := fileStat.Mode().Perm()
	// fmt.Printf("File permission before copying %v \n", perm)

	// Create the destination file with default permission
	dst, err := os.Create(dstFile)
	if err != nil {
		return false, fmt.Errorf("failed to create %s @ %s", dstFile, err) 
	}
	defer dst.Close()

	// preserve permissions from srcFile to dstFile
	srcStat, _ := src.Stat()
	// fmt.Println("Changing permission of ", dstFile)
	os.Chmod(dstFile, srcStat.Mode())

	// check dstFile stats
	_, err = os.Stat(dstFile)
	if err != nil {
		return false, fmt.Errorf("failed to check stats for %s @ %s", dstFile, err) 
	}

	// print dstFile stats
	// perm2 := newFileStats.Mode().Perm()
	// fmt.Printf("File permission After copying %v \n", perm2)

	// Copy the content of srcFile to dstFile
	if _, err := io.Copy(dst, src); err != nil {
		return false, fmt.Errorf("failed copy `%s` to `%s` @ %s",    srcFile, dstFile, err) 
	}

	return true, nil
}

