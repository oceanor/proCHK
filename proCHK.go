// proCHK - A file recovery tool for .CHK files
// Copyright (C) 2025 Federico Fallico
// Licensed under the MIT License. See LICENSE file for details.

package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Signature struct {
	Extension string
	Header    []byte
	Contains  []byte
}

var signatures []Signature

const programName = "proCHK"
const version = "1.02"

func initializeSignatures() {
	addSignature("JPG", "ffd8", "4a464946")
	addSignature("EXIF.JPG", "ffd8", "45786966")
	addSignature("PNG", "89504e470d0a1a", "")
	addSignature("GIF", "47494638", "")
	addSignature("TIF", "49492a00", "")
	addSignature("TIF", "4d4d002a", "")
	addSignature("BMP", "424d", "")
	addSignature("PSD", "38425053", "")
	addSignature("WEBP", "52494646", "57454250")
	addSignature("HEIC", "000000", "6674797068656963")
	addSignature("AVIF", "000000", "6674797061766966")
	addSignature("CR2", "49492a00", "4352")
	addSignature("NEF", "4d4d002a", "4e696b6f6e")
	addSignature("AI", "25215053", "")
	addSignature("EPS", "c5d0d3c6", "")
	addSignature("3DS", "4d4d", "")
	addSignature("FPX", "d0cf11e0", "49006d006100670065")
	addSignature("PSP", "5061696e742053686f702050726f", "")
	addSignature("WPG", "ff575043", "")
	addSignature("WAV", "52494646", "57415645")
	addSignature("MP3", "494433", "")
	addSignature("MID", "4d546864", "4d54726b")
	addSignature("FLAC", "664C6143", "")
	addSignature("M4A", "000000", "667479704d344120")
	addSignature("OGG", "4f676753", "")
	addSignature("RMI", "52494646", "524d4944")
	addSignature("AVI", "52494646", "415649")
	addSignature("MP4", "000000", "66747970")
	addSignature("WebM", "1a45dfa3", "7765626d")
	addSignature("MKV", "1a45dfa3", "")
	addSignature("MOV", "000000", "6674797071742020")
	addSignature("WMV", "3026b2758e66cf11", "415346")
	addSignature("FLV", "464c5601", "")
	addSignature("MPG", "000001b3", "")
	addSignature("MPEG", "000001ba", "")
	addSignature("ASF", "3026b2758e66cf11a6d900aa0062ce6c", "")
	addSignature("SWF", "465753", "")
	addSignature("3GP", "000000", "667479703367")
	addSignature("DOC", "d0cf11e0a1b11ae1", "4d6963726f736f667420576f7264")
	addSignature("XLS", "d0cf11e0a1b11ae1", "4d6963726f736f667420457863656c")
	addSignature("PPT", "d0cf11e0a1b11ae1", "4d6963726f736f667420506f776572506f696e74")
	addSignature("PST", "2142444e", "")
	addSignature("WRI", "31be", "00002e0d0a")
	addSignature("DOCX", "504B0304", "776F72642F")
	addSignature("XLSX", "504B0304", "786C2F")
	addSignature("PPTX", "504B0304", "7070742F")
	addSignature("ODT", "504B0304", "6d696d65747970656170706c69636174696f6e2f766e642e6f617369732e6f70656e646f63756d656e742e74657874")
	addSignature("ODS", "504B0304", "6d696d65747970656170706c69636174696f6e2f766e642e6f617369732e6f70656e646f63756d656e742e7370726561647368656574")
	addSignature("ODP", "504B0304", "6d696d65747970656170706c69636174696f6e2f766e642e6f617369732e6f70656e646f63756d656e742e70726573656e746174696f6e")
	addSignature("ODG", "504B0304", "6d696d65747970656170706c69636174696f6e2f766e642e6f617369732e6f70656e646f63756d656e742e6772617068696373")
	addSignature("PDF", "25504446", "")
	addSignature("EPUB", "504B0304", "6d696d65747970656170706c69636174696f6e2f657075622b7a6970")
	addSignature("TTF", "0001000000", "")
	addSignature("OTF", "4f5454f0", "")
	addSignature("RAR", "52617221", "")
	addSignature("7z", "377abcaf271c", "")
	addSignature("GZ", "1f8b", "")
	addSignature("ACE", "2a2a4143452a2a", "")
	addSignature("ZIP", "504B0304", "")
	addSignature("CAB", "4d534346", "")
	addSignature("MDB", "000100005374616e64617264204a6574204442", "")
	addSignature("ACCDB", "000100005374616e6461726420414345204442", "")
	addSignature("SQLite3", "53514c69746520666f726d6174203300", "")
	addSignature("DWG", "41433130", "")
	addSignature("NC", "434446", "")
	addSignature("PY", "2321", "707974686f6e")
	addSignature("SH", "2321", "2f62696e2f")
	addSignature("CLASS", "cafebabe", "")
	addSignature("BDSPROJ", "3c3f786d6c", "426f726c616e6450726f6a656374")
	addSignature("VCF", "424547494e3a5643415244", "")
	addSignature("ICS", "424547494e3a5643414c454e444152", "")
	addSignature("TORRENT", "64383a616e6e6f756e6365", "")
	addSignature("EXE", "4d5a", "")
	addSignature("DLL", "4d5a", "")
	addSignature("OCX", "4d5a", "446c6c5265676973746572536572766572")
	addSignature("RTF", "7b5c727466", "7b5c666f6e7474626c")
	addSignature("CHM", "49545346", "")
	addSignature("HLP", "3f5f0300", "")
	addSignature("LNK", "4c0000000114020000000000c0000000", "")
	addSignature("URL", "5b496e7465726e657453686f72746375745d", "")
	addSignature("CDR", "52494646", "434452")
	addSignature("HTML", "3c21444f4354595045", "")
	addSignature("HTM", "3c68746d6c", "")
	addSignature("CLP", "50c30100", "")
}

func addSignature(ext, headerHex, containsHex string) {
	header, err := hex.DecodeString(headerHex)
	if err != nil {
		log.Fatalf("Error decoding header for %s: %v", ext, err)
	}
	var contains []byte
	if containsHex != "" {
		contains, err = hex.DecodeString(containsHex)
		if err != nil {
			log.Fatalf("Error decoding 'contains' for %s: %v", ext, err)
		}
	}
	signatures = append(signatures, Signature{Extension: ext, Header: header, Contains: contains})
}

func isLikelyTextFile(data []byte) bool {
	totalBytes := len(data)
	if totalBytes == 0 || totalBytes > 1048576 {
		return false
	}
	textByteCount := 0
	for _, b := range data {
		if (b >= 32 && b <= 126) || b == 9 || b == 10 || b == 13 {
			textByteCount++
		}
	}
	if textByteCount == 0 {
		return false
	}
	ratio := float64(totalBytes) / float64(textByteCount)
	return ratio < 1.25
}

func main() {
	fmt.Printf("%s v%s\n\n", programName, version)
	initializeSignatures()

	sourceDir := flag.String("source", "", "The folder where the .CHK files are located (required)")
	destDir := flag.String("dest", "", "The folder where recovered files will be saved (optional)")
	logFile := flag.Bool("log", false, "Create a log file (prochk.log)")
	skipExisting := flag.Bool("skip", false, "Skip recovery if a file with the same name already exists in the destination")
	recursive := flag.Bool("recursive", false, "Search for .CHK files in subdirectories as well")
	flag.Parse()

	if *sourceDir == "" {
		fmt.Printf("Error: The source folder is required.\n")
		fmt.Printf("Usage: %s -source <path_to_chk_folder>\n", programName)
		flag.PrintDefaults()
		return
	}
	if *destDir == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the path to the folder where recovered files will be saved: ")
		path, _ := reader.ReadString('\n')
		*destDir = strings.TrimSpace(path)
	}
	if _, err := os.Stat(*sourceDir); os.IsNotExist(err) {
		log.Fatalf("Source folder '%s' does not exist.", *sourceDir)
	}
	if err := os.MkdirAll(*destDir, 0755); err != nil {
		log.Fatalf("Could not create destination folder '%s': %v", *destDir, err)
	}
	fmt.Printf("Source: %s\n", *sourceDir)
	fmt.Printf("Destination: %s\n", *destDir)

	var logger *log.Logger
	if *logFile {
		logFilePath := fmt.Sprintf("%s.log", programName)
		f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Could not create log file: %v", err)
		}
		defer f.Close()
		logger = log.New(f, "", log.LstdFlags)
		fmt.Printf("Log will be saved to %s\n", logFilePath)
	}

	var chkFiles []string
	var err error
	if *recursive {
		fmt.Println("Recursive search enabled...")
		err = filepath.Walk(*sourceDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.EqualFold(filepath.Ext(path), ".chk") {
				chkFiles = append(chkFiles, path)
			}
			return nil
		})
	} else {
		globPattern := filepath.Join(*sourceDir, "*.[cC][hH][kK]")
		chkFiles, err = filepath.Glob(globPattern)
	}
	if err != nil {
		log.Fatalf("Error finding .chk files: %v", err)
	}
	if len(chkFiles) == 0 {
		fmt.Println("No .CHK files found in the specified folder.")
		return
	}
	fmt.Printf("Found %d .CHK files. Starting scan...\n", len(chkFiles))

	for i, chkPath := range chkFiles {
		processFile(chkPath, *destDir, i+1, len(chkFiles), logger, *skipExisting)
	}
	fmt.Println("\nProcess complete.")
}

func processFile(chkPath, destDir string, currentFile, totalFiles int, logger *log.Logger, skipExisting bool) {
	fileName := filepath.Base(chkPath)
	fileRoot := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fmt.Printf("\rScanning [%d/%d]: %s...", currentFile, totalFiles, fileName)
	data, err := ioutil.ReadFile(chkPath)
	if err != nil {
		fmt.Printf("\nError reading file %s: %v\n", fileName, err)
		if logger != nil {
			logger.Printf("%s\tERROR READING FILE", fileName)
		}
		return
	}
	if len(data) == 0 {
		return
	}
	fileRecovered := false
	for i := 0; i < len(data); i++ {
		for _, sig := range signatures {
			if (i+len(sig.Header) <= len(data)) && bytes.HasPrefix(data[i:], sig.Header) {
				containsMatch := true
				if len(sig.Contains) > 0 {
					if !bytes.Contains(data[i+len(sig.Header):], sig.Contains) {
						containsMatch = false
					}
				}
				if containsMatch {
					saveRecoveredFile(destDir, fileRoot, fileName, sig.Extension, data[i:], currentFile, totalFiles, skipExisting, logger)
					fileRecovered = true
				}
			}
		}
	}

	if !fileRecovered {
		if isLikelyTextFile(data) {
			if json.Valid(data) {
				saveRecoveredFile(destDir, fileRoot, fileName, "json", data, currentFile, totalFiles, skipExisting, logger)
			} else {
				saveRecoveredFile(destDir, fileRoot, fileName, "txt", data, currentFile, totalFiles, skipExisting, logger)
			}
			fileRecovered = true
		}
	}

	if !fileRecovered && logger != nil {
		logger.Printf("%s\t(not recovered)", fileName)
	}
}

func saveRecoveredFile(destDir, fileRoot, fileName, extension string, data []byte, currentFile, totalFiles int, skipExisting bool, logger *log.Logger) {
	recoveredFileName := fmt.Sprintf("%s.%s", fileRoot, extension)
	var recoveredFilePath string
	if skipExisting {
		recoveredFilePath = filepath.Join(destDir, recoveredFileName)
		if _, err := os.Stat(recoveredFilePath); !os.IsNotExist(err) {
			fmt.Printf("\rSkipped! [%d/%d]: %s -> %s (already exists)\n", currentFile, totalFiles, fileName, filepath.Base(recoveredFilePath))
			if logger != nil {
				logger.Printf("%s\tSKIPPED (exists) -> %s", fileName, extension)
			}
			return
		}
	} else {
		recoveredFilePath = getUniquePath(filepath.Join(destDir, recoveredFileName))
	}
	err := ioutil.WriteFile(recoveredFilePath, data, 0644)
	if err != nil {
		fmt.Printf("\nError saving recovered file %s: %v\n", recoveredFilePath, err)
	} else {
		fmt.Printf("\rFound! [%d/%d]: %s -> %s\n", currentFile, totalFiles, fileName, filepath.Base(recoveredFilePath))
		if logger != nil {
			logger.Printf("%s\t%s", fileName, extension)
		}
	}
}

func getUniquePath(path string) string {
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(filepath.Base(path), ext)
	newPath := path
	counter := 1
	for {
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
		newPath = filepath.Join(dir, fmt.Sprintf("%s-%d%s", base, counter, ext))
		counter++
	}
}