// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package dns

import "encoding/yaml"

_doRaw: string @tag(doFile, type=string)
_doData: yaml.Unmarshal(_doRaw)
_doData: [...#DODomain]
