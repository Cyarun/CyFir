/*
Velociraptor - Dig Deeper
Copyright (C) 2019-2025 Rapid7 Inc.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package plugins

// This is a do nothing package which just forces an import of the
// various VQL plugin directories. The plugins will register
// themselves.

import (
	_ "github.com/Cyarun/CyFir/vql/aggregates"
	_ "github.com/Cyarun/CyFir/vql/common"
	_ "github.com/Cyarun/CyFir/vql/efi"
	_ "github.com/Cyarun/CyFir/vql/filesystem"
	_ "github.com/Cyarun/CyFir/vql/functions"
	_ "github.com/Cyarun/CyFir/vql/golang"
	_ "github.com/Cyarun/CyFir/vql/networking"
	_ "github.com/Cyarun/CyFir/vql/parsers"
	_ "github.com/Cyarun/CyFir/vql/parsers/authenticode"
	_ "github.com/Cyarun/CyFir/vql/parsers/crypto"
	_ "github.com/Cyarun/CyFir/vql/parsers/csv"
	_ "github.com/Cyarun/CyFir/vql/parsers/ese"
	_ "github.com/Cyarun/CyFir/vql/parsers/event_logs"
	_ "github.com/Cyarun/CyFir/vql/parsers/journald"
	_ "github.com/Cyarun/CyFir/vql/parsers/sql"
	_ "github.com/Cyarun/CyFir/vql/parsers/syslog"
	_ "github.com/Cyarun/CyFir/vql/parsers/usn"
	_ "github.com/Cyarun/CyFir/vql/protocols"
	_ "github.com/Cyarun/CyFir/vql/sigma"
	_ "github.com/Cyarun/CyFir/vql/tools"
	_ "github.com/Cyarun/CyFir/vql/tools/collector"
	_ "github.com/Cyarun/CyFir/vql/tools/dns"
	_ "github.com/Cyarun/CyFir/vql/tools/logscale"
	_ "github.com/Cyarun/CyFir/vql/tools/packaging"
	_ "github.com/Cyarun/CyFir/vql/tools/process"
	_ "github.com/Cyarun/CyFir/vql/tools/rsyslog"
)
