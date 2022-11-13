package service

import (
	"BeTest-AlexanderBergasov/internal/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Orchestra struct {
	supportedDecoders map[string]Decoder
}

func NewOrchestra() *Orchestra {
	return &Orchestra{
		supportedDecoders: map[string]Decoder{},
	}
}

func (o *Orchestra) RegisterDecoder(ext string, decoder Decoder) {
	log.Println("register decoder for files: ", ext)
	o.supportedDecoders[ext] = decoder
}

func (o *Orchestra) Run(filePath string) error {
	files, err := os.ReadDir(filePath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileExt := filepath.Ext(file.Name())
		decoder, ok := o.supportedDecoders[fileExt]
		if !ok {
			continue
		}

		if err = o.processFile(file.Name(), decoder); err != nil {
			return fmt.Errorf("failed to process file: %w", err)
		}
	}
	return nil
}

func (o *Orchestra) processFile(filePath string, decoder Decoder) error {
	log.Println("parsing file: ", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := decoder.Decode(file)
	if err != nil {
		log.Printf("failed to decode file `%s`: %s\n", filePath, err)
		return err
	}
	return o.saveFile(content, filePath+".html")
}

func (o *Orchestra) saveFile(contentHTML string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	if _, err = file.WriteString(utils.WrapToHTML(contentHTML)); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	log.Println("file saved: ", path)
	return nil
}
