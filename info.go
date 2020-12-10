package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	name        = "tree"
	author      = "Michael Treanor"
	description = `Tree is a recursive directory listing command that produces a
depth indented listing of files, which is colorized ala dircolors
if color tty output is available.`
	license        = "GPLv2"
	startYear      = 2020
	configFileName = "gorepo.json"
)

func GetConfig() (*softwareInfo, error) {

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Error("error opening config file %s: %v", configFileName, err)
		return nil, err
	}

	defer configFile.Close()

	n := new(softwareInfo)

	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(n)
	if err != nil {
		log.Error("error parsing config file %s: %v", configFile, err)
	}
	return n, err
}

type softwareInfo struct {
	Name        string `json:Name`
	Author      string `json:Author`
	Description string `json: Description`
	License     string `json: License`
	StartYear   int    `json: StartYear`
}

func (s *softwareInfo) Copyright() string {
	year := time.Now().Year()
	yearString := fmt.Sprintf("Copyright (C) %d", s.startYear)

	if year != s.startYear {
		yearString += fmt.Sprintf("-%d", year)
	}

	return yearString + " " + s.author
}

func (s *softwareInfo) Header() string {
	sb := strings.Builder{}
}

var (
	b                 []byte
	buf               bytes.Buffer
	sb                strings.Builder
	cliLicenseInfo    string = fmt.Sprintf("%s\n%s %s", Copyright(), name, cliLicenseInfoStub)
	licenseHeaderInfo string = fmt.Sprintf("%s\n%s %s", Copyright(), name, licenseHeaderStuf)

	cliLicenseInfoStub string = `comes with ABSOLUTELY NO WARRANTY.
    This is free software, and you are welcome to redistribute it
    under certain conditions;  for details type 'tree -help'.`

	licenseHeaderStuf string = `
All Rights reserved

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA`
)

const nl = '\n'
