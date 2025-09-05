# proCHK

A powerful, command-line file recovery tool for `.CHK` files, written in Go.

proCHK is a modern, cross-platform utility inspired by classic CHK file recovery tools. It scans `.CHK` files created by the Windows `CHKDSK` or `SCANDISK` utilities and identifies files based on their binary signatures, statistical analysis, and internal structure.

## Features

-   **Extensive File Support**: Recognizes over 70 file formats, including modern images (HEIC, AVIF), RAW photos (CR2, NEF), office documents (DOCX, ODT), archives, and databases.
-   **Deep Scan**: Scans every byte of a `.CHK` file to find embedded files, not just signatures at the beginning.
-   **Recursive Search**: Optionally scans all subdirectories of a source folder for `.CHK` files.
-   **Smart Text/JSON Detection**: Uses statistical analysis to identify plain text files (TXT) and then validates them to check if they are structured JSON files.
-   **Flexible Recovery**: Choose to skip already recovered files or create unique filenames for multiple versions.
-   **Cross-Platform**: Written in Go, it can be compiled and run on Windows, macOS, and Linux.

## Installation

To compile proCHK from the source, you need to have Go (version 1.16 or later) installed on your system.

1.  Clone the repository or download the source code.
2.  Navigate to the project directory in your terminal and run the build command:
    ```bash
    go build prochk.go
    ```
    This will create an executable named `prochk` (or `prochk.exe` on Windows).

## Usage

Run the program from your terminal, providing the source path and any optional flags.

```bash
./prochk -source "/path/to/chks" -dest "/path/to/output" [options]
```

### Parameters

* **`-source`**: (Required) The path to the folder containing your `.CHK` files.
* **`-dest`**: (Optional) The path to the output folder where recovered files will be saved. If this parameter is omitted, you will be prompted to enter it.
* **`-recursive`**: (Optional) Searches for `.CHK` files within all subdirectories of the source folder.
* **`-skip`**: (Optional) Skips the recovery of a file if another file with the same identified name already exists in the destination. By default, the program will create a unique name by adding a number suffix (e.g., `filename-1.jpg`).
* **`-log`**: (Optional) Creates a log file named `prochk.log` detailing the results of the recovery process.

---

## Supported File Types

The program automatically detects and recovers a comprehensive list of file types. Supported formats include:

#### Images & Photos
JPG, PNG, GIF, TIF, BMP, PSD, WEBP, HEIC, AVIF, AI, 3DS, professional RAW formats (CR2, NEF).

#### Documents
PDF, DOC, XLS, PPT, DOCX, XLSX, PPTX, OpenDocument (ODT, ODS, ODP), Outlook PST, RTF, HLP.

#### Audio
MP3, WAV, FLAC, M4A, OGG, MID, RMI.

#### Video
MP4, AVI, MKV, MOV, WMV, FLV, MPG, MPEG, ASF, WebM, SWF.

#### Archives
ZIP, RAR, 7z, GZ, CAB.

#### Other Formats
EXE, DLL, TTF, OTF, MDB, ACCDB, DWG, LNK, URL, HTML, HTM, CDR, TXT, JSON, SQLite3, VCF, ICS, TORRENT, CHM.
