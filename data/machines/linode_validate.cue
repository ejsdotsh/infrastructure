// SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
//
// SPDX-License-Identifier: Apache-2.0

package machines

import "encoding/yaml"

_linodeRaw: string @tag(linodeFile, type=string)
_linodeData: yaml.Unmarshal(_linodeRaw)
_linodeData: [...#LinodeMachine]
