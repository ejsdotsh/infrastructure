#!/usr/bin/env bash
# SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>
#
# SPDX-License-Identifier: Apache-2.0

# Validate all data files against their CUE schemas.
# Usage: ./data/validate.sh

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "${SCRIPT_DIR}"

echo "=== Validating data definitions ==="
cue vet -t "machinesFile=$(cat machines.yaml)" -t "dnsFile=$(cat dns.yaml)" .
echo "  machines.yaml: OK"
echo "  dns.yaml: OK"

echo "=== All validations passed ==="
