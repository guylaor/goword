// package goword parses docx files to get its containing text
package goword

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseText(filename string) (string, error) {

	doc, err := openWordFile(filename)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error opening file %s - %s", filename, err))
	}

	docx, err := Parse(doc)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing %s - %s", filename, err))
	}

	return docx.AsText(), nil

}

// func Parse(doc string) (WordDocument, error) {

// 	docx := WordDocument{}
// 	err := xml.Unmarshal([]byte(doc), &docx)
// 	if err != nil {
// 		return docx, err
// 	}
// 	fmt.Printf("\n %-v \n", docx)
// 	return docx, nil
// }

func Parse(doc string) (WordDocument, error) {

	docx := WordDocument{}
	r := strings.NewReader(string(doc))
	decoder := xml.NewDecoder(r)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "p" {
				var p Paragraph
				decoder.DecodeElement(&p, &se)
				docx.Paragraphs = append(docx.Paragraphs, p)
			}
		}
	}
	return docx, nil
}

func openWordFile(filename string) (string, error) {

	// Open a zip archive for reading. word files are zip archives
	r, err := zip.OpenReader(filename)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// find document.xml
	for _, f := range r.File {

		//fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			return "", err
		}
		if f.Name == "word/document.xml" {
			doc, err := ioutil.ReadAll(rc)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s", doc), nil
		}
		rc.Close()
	}

	return "", nil
}
