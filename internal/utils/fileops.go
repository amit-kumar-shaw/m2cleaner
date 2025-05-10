package utils

import (
    "fmt"
    "os"
    "path/filepath"
)

func PathExists(path string) bool {
    _, err := os.Stat(path)
    return !os.IsNotExist(err)
}

func IsDirectory(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return info.IsDir()
}

func ListFiles(path string) ([]string, error) {
    var files []string
    items, err := os.ReadDir(path)
    if err != nil {
        return nil, err
    }

    for _, item := range items {
        if !item.IsDir() {
            files = append(files, filepath.Join(path, item.Name()))
        }
    }
    return files, nil
}

func DeleteFile(path string) error {
    if err := os.Remove(path); err != nil {
        return fmt.Errorf("failed to delete file %s: %w", path, err)
    }
    fmt.Println("Deleted file:", path)
    return nil
}

func DeleteDirectory(path string) error {
    if err := os.RemoveAll(path); err != nil {
        return fmt.Errorf("failed to delete directory %s: %w", path, err)
    }
    fmt.Println("Deleted directory:", path)
    return nil
}
