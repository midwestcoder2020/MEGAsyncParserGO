package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type mega_row struct {
	nodehandle         string `json:"Nodehandle,omitempty"`
	parenthandle       string `json:"Parenthandle,omitempty"`
	name               string `json:"Name,omitempty"`
	fingerprint        string `json:"Fingerprint,omitempty"`
	origFingerprint    string `json:"OrigFingerprint,omitempty"`
	Type               string `json:"Type,omitempty"`
	mimetypeVirtual    string `json:"MimetypeVirtual,omitempty"`
	fingerprintVirtual string `json:"FingerprintVirtual,omitempty"`
	sizeVirtual        string `json:"SizeVirtual,omitempty"`
	share              string `json:"Share,omitempty"`
	fav                string `json:"Fav,omitempty"`
	ctime              string `json:"Ctime,omitempty"`
	mtime              string `json:"Mtime,omitempty"`
	flags              string `json:"Flags,omitempty"`
	counter            string `json:"Counter,omitempty"`
	node               string `json:"Node,omitempty"`
	label              string `json:"Label,omitempty"`
	description        string `json:"Description,omitempty"`
	tags               string `json:"Tags,omitempty"`
}

// Schema has changed as of 2026
const mega_query = `
        SELECT
            *
        FROM
            nodes;`

func parse_mega_files(db_path string, output_path string) {
	db_handle, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatalf("Unable to open db file: %s", db_path)
	}

	defer db_handle.Close()

	log.Printf("[+] Opened db %s\n", db_path)
	log.Println("[+] Executing MegaFileDB Parsing")

	rows, err := db_handle.Query(mega_query)
	if err != nil {
		log.Fatalln(err)
	}

	var header = []string{"nodehandle", "parenthandle", "name", "fingerprint", "origFingerprint", "type", "mimetypeVirtual", "fingerprintVirtual", "sizeVirtual", "share", "fav", "ctime", "mtime", "flags", "counter", "node", "label", "description", "tags"}

	var output_entries [][]string = make([][]string, 0)
	output_entries = append(output_entries, header)

	for rows.Next() {
		var row mega_row

		err = rows.Scan(
			&row,
		)

		if err != nil {
			log.Println(err)
		}

		output_entries = append(output_entries, []string{
			row.nodehandle,
			row.parenthandle,
			row.name,
			row.fingerprint,
			row.origFingerprint,
			row.Type,
			row.mimetypeVirtual,
			row.fingerprintVirtual,
			row.sizeVirtual,
			row.share,
			row.fav,
			row.ctime,
			row.mtime,
			row.flags,
			row.counter,
			row.node,
			row.label,
			row.description,
			row.tags,
		})
	}

	output_csv(output_entries, output_path, "mega_files")

}

func output_csv(entries [][]string, output string, filename string) {
	log.Printf("[+] Saving %s output to %s.csv\n", filename, filename)
	var csv_f, err = os.Create(fmt.Sprintf("%s/%s.csv", output, filename))
	if err != nil {
		log.Fatalln(err)
	}

	var writer = csv.NewWriter(csv_f)

	for _, v := range entries {
		we := writer.Write(v)
		if we != nil {
			log.Println(we)
		}
	}

	writer.Flush()
	if ferr := writer.Error(); ferr != nil {
		log.Println(ferr)
	}

	log.Printf("[+] File: %s.csv Saved!\n", filename)
}

func main() {
	log.Println("Go MEGASyncParser")
	var args = os.Args

	if len(args) < 3 {
		log.Println("[-] Invalid arguments")
		log.Fatalln("[-] usage: <program> <input_file> <output_folder>")
	}

	var target_file = args[1]
	var output_dir = args[2]

	log.Println(target_file)
	log.Println(output_dir)

	parse_mega_files(target_file, output_dir)

}
