// This file is part of the GOfax.IP project - https://github.com/gonicus/gofaxip
// Copyright (C) 2014 GONICUS GmbH, Germany - http://www.gonicus.de
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; version 2
// of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.

package gofaxlib

import (
	"log"

	"github.com/anis-bedhiafi/gofaxip-mysql/gofaxlib/logger"

	gcfg "gopkg.in/gcfg.v1"
)

var (
	// Config is the global configuration struct
	Config config
)

type config struct {
	Freeswitch struct {
		Socket            string
		Password          string
		Gateway           []string
		Ident             string
		Header            string
		Verbose           bool
		SoftmodemFallback bool
	}
	Hylafax struct {
		Spooldir string
		Modems   uint
	}
	Log struct {
		Xferfaxlog string
	}
	Gofaxd struct {
		EnableT38              bool
		RequestT38             bool
		Socket                 string
		Answerafter            uint64
		Waittime               uint64
		FaxRcvdCmd             string
		DynamicConfig          string
		AllocateInboundDevices bool
	}
	Gofaxsend struct {
		EnableT38            bool
		RequestT38           bool
		FaxNumber            string
		CallPrefix           string
		DynamicConfig        string
		DisableV17AfterRetry string
		DisableECMAfterRetry string
		CidName              string
		FailedResponse       []string
		FailedResponseMap    map[string]bool
	}
	MySQL struct {
		Host     string
		Port     string
		User     string
		Pass     string
		Database string
		Charset  string
	}
}

// LoadConfig loads the configuration from given file path
func LoadConfig(filename string) {
	err := gcfg.ReadFileInto(&Config, filename)
	if err != nil {
		logger.Logger.Print("Config: ", err)
		log.Fatal("Config: ", err)
	} else {
		Config.Gofaxsend.FailedResponseMap = make(map[string]bool)
		for _, i := range Config.Gofaxsend.FailedResponse {
			logger.Logger.Print("adding failed response: ", i)
			Config.Gofaxsend.FailedResponseMap[i] = true
		}
	}
}

func FailedHangupcause(hangupcause string) bool {
	if Config.Gofaxsend.FailedResponseMap[hangupcause] {
		logger.Logger.Print(hangupcause, " is a final failed response!")
		return true
	} else {
		logger.Logger.Print(hangupcause, " is not a final failed response!")
		return false
	}
}
