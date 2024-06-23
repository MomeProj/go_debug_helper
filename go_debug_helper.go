package go_debug_helper

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func D_FILE() string {
	var buf bytes.Buffer
	log.New(&buf, "", log.Lshortfile).Output(2, "") //2=>CALLER
	dbgOutStr := strings.TrimSpace(buf.String())
	dbgOutStrTrimed := strings.Split(dbgOutStr, ":")
	return dbgOutStrTrimed[0]
}

func D_LINE() string {
	var buf bytes.Buffer
	log.New(&buf, "", log.Lshortfile).Output(2, "") //2=>CALLER
	dbgOutStr := strings.TrimSpace(buf.String())
	dbgOutStrTrimed := strings.Split(dbgOutStr, ":")
	return dbgOutStrTrimed[1]
}

func D_DIR() string {
	var buf bytes.Buffer
	log.New(&buf, "", log.Llongfile).Output(2, "") //2=>CALLER
	dbgOutStr := strings.TrimSpace(buf.String())
	_file := strings.TrimRight(dbgOutStr, ":1234567890")
	_dir := filepath.Dir(_file)
	return _dir
}

func D_FILE_LINE() string {
	var buf bytes.Buffer
	log.New(&buf, "", log.Lshortfile).Output(2, "") //2=>CALLER
	dbgOutStr := strings.TrimSpace(buf.String())
	dbgOutStrTrimed := strings.Split(dbgOutStr, ":")
	return fmt.Sprintf("%s %s", dbgOutStrTrimed[0], dbgOutStrTrimed[1])
}
