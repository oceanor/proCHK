# proCHK v1.02

A powerful, command-line file recovery tool for `.CHK` files, written in Go. Licensed under the MIT License.

proCHK is a modern, cross-platform utility inspired by classic CHK file recovery tools. It scans `.CHK` files created by the Windows `CHKDSK` or `SCANDISK` utilities and identifies files based on their binary signatures, statistical analysis, and internal structure.

## Features

-   **Extensive File Support**: Recognizes a huge range of file formats, including modern images (HEIC, AVIF), RAW photos (CR2, NEF), office documents (DOCX, ODT), archives, and databases.
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

* **`-source`**: (Required) The path to the folder containing your .CHK files.
* **`-dest`**: (Optional) The path to the output folder where recovered files will be saved. If this parameter is omitted, you will be prompted to enter it.
* **`-recursive`**: (Optional) Searches for .CHK files within all subdirectories of the source folder.
* **`-skip`**: (Optional) Skips the recovery of a file if another file with the same identified name already exists in the destination. By default, the program will create a unique name by adding a number suffix (e.g., `filename-1.jpg`).
* **`-log`**: (Optional) Creates a log file named `prochk.log` detailing the results of the recovery process.

---

## Supported File Types

proCHK automatically detects and recovers a comprehensive list of file types. Supported formats include:

#### Images & Graphics
3DS, AI, AVIF, BMP, CDR, CR2, EPS, FPX, GIF, HEIC, JPG, NEF, PNG, PSD, PSP, TIF, WEBP, WPG.

#### Documents & Office
DOC, DOCX, HLP, ODP, ODS, ODT, PDF, PPT, PPTX, PST, RTF, WRI, XLS, XLSX.

#### Audio & Video
3GP, ASF, AVI, FLAC, FLV, M4A, MID, MKV, MOV, MP3, MP4, MPG, MPEG, OGG, RMI, SWF, WAV, WebM, WMV.

#### Archives
7z, ACE, CAB, GZ, RAR, ZIP.

#### Code, Scripts & Projects
BDSPROJ, CLASS (Java), HTM, HTML, PY (Python), SH (Shell Script).

#### System & Other
ACCDB, CHM, CLP, DLL, DWG, EXE, ICS, LNK, MDB, NC, OCX, OTF, SQLite3, TORRENT, TTF, URL, VCF, TXT, JSON.

---

## License

This project is licensed under the **MIT License**.

See the `LICENSE` file in this repository for the full license text.