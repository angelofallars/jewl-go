package jewl

import "os"

type RecorderConfigFile struct{
    path string
}

func (f RecorderConfigFile) Write(data []byte) error{
    if _, err := os.Stat(f.path); err != nil{
        _, err := os.Create(f.path)
        if err != nil{
            return err
        }
    }
    err := os.WriteFile(f.path, data, 0666)
    if err != nil{
        return err
    }
    return nil
}

func (f RecorderConfigFile) Load() ([]byte, error){
    if _, err := os.Stat(f.path); err != nil{
        _, err := os.Create(f.path)
        if err != nil{
            return []byte{}, err
        }
    }
    data, err := os.ReadFile(f.path)
    if err != nil{
        return []byte{}, err
    }
    return data, nil
}

func FileConfig(path string) RecorderConfigFile{
    return RecorderConfigFile{
        path: path,
    }
}
