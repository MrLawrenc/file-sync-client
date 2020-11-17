package files

import (
	"file-sync-client/service/entity"
	"file-sync-client/service/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ListFiles(parentDir string) []entity.FileInfo {
	infos, e := ioutil.ReadDir(parentDir)
	if e != nil || len(infos) == 0 {
		panic("parent dir is illegal")
		utils.PrintError(e, true)
	}

	files := make([]os.File, 0, 10)
	list(parentDir, &files)

	log.Println("-------------------读取完毕---------------------")
	fileInfos := make([]entity.FileInfo, 0, len(files))
	for i := range files {
		file := files[i]
		defer file.Close()

		info, _ := file.Stat()

		fileInfos = append(fileInfos, entity.FileInfo{FileName: file.Name(), Modify: info.ModTime(), Size: info.Size()})
		//println(file.Name(), "   last mod time:", info.ModTime().Unix(), "ms   file size:", info.Size()/1024, "kb")
	}
	log.Println("target size:", len(fileInfos))
	return fileInfos
}

func list(path string, files *[]os.File) {
	infos, err := ioutil.ReadDir(path)

	utils.PrintError(err, true)

	for _, v := range infos {
		newPath := filepath.Join(path, v.Name())
		if v.IsDir() {
			list(newPath, files)
		} else {
			log.Println("add target file path:", newPath)
			file, err := os.Open(newPath)

			utils.PrintError(err, true)
			if file != nil {
				*files = append(*files, *file)
			}
		}
	}

}
