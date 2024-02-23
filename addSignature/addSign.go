package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
파일 내 문서 정보와 copylight를 추가하는 애플리케이션
*/
func main() {
	// 원본 파일 및 변환 파일 저장 경로
	srcFile := "C:\\Downloads\\imsi\\HATEOAS.md"
	dstFile := "C:\\Downloads\\imsi\\HATEOAS_signed.md"

	// 텍스트 파일 열기
	file, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 스캐너 생성
	scanner := bufio.NewScanner(file)

	// 수정된 내용을 담을 배열
	var modifiedLines []string

	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLines = append(modifiedLines, line)
		lineNumber++

		// add doc-info at 2'nd line
		if lineNumber == 2 {
			modifiedLines = append(modifiedLines, getDocInfo(getFileInfo(srcFile)))
		}
	}
	// add copylight at last line
	modifiedLines = append(modifiedLines, getDocCopylight())

	// write to new file
	err = ioutil.WriteFile(dstFile, []byte(strings.Join(modifiedLines, "\n")), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File converting finish!!")

}

func getFileInfo(filePath string) (string, string, string) {
	// 파일 정보 가져오기
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		//return
	}
	fileName := fileInfo.Name()
	createDate := fileInfo.ModTime().Format("2006년 01월 02일 15:04:05")
	modifyDate := fileInfo.ModTime().Format("2006년 01월 02일 15:04:05")

	return fileName, createDate, modifyDate
}

func getDocInfo(fileName string, createDate string, modifyDate string) string {
	var docInfo bytes.Buffer
	docInfo.WriteString("\n\n**문서정보**")
	docInfo.WriteString("\n* 문서명 : ")
	docInfo.WriteString(fileName)
	docInfo.WriteString("\n* 최초 작성일 : ")
	docInfo.WriteString(createDate)
	docInfo.WriteString("\n* 최종 업데이트일 : ")
	docInfo.WriteString(modifyDate)
	docInfo.WriteString("\n* 작성자 : 오픈서비스사업팀 / 이대현 \n\n")

	return docInfo.String()
}

func getDocCopylight() string {
	return "\n\n<center><b>Copyright Gaeasoft co., Ltd., All right reserved.</b></center>"
}
