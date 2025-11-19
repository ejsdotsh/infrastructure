<!--
SPDX-FileCopyrightText: 2025 e.j. sahala <ej@saha.la>

SPDX-License-Identifier: Apache-2.0
-->

# Reanimating Notes

reanimating my personal infrastructure...as code

## Design Decisions

- [CUE](https://cuelang.org)
  - Reduce the number of languages in the build pipeline
  - Typing and validation
  - Packages...because code reuse is good
- [pulumi](https://https://www.pulumi.com/)
  - SDKs
  - Licensing
  - IaC
- *FUTURE* [netbox](https://netboxlabs.com/)
  - Source-of-Truth (SoT)
    ~~- [nickel](https://https://nickel-lang.org)~~
    ~~- Records~~
    ~~- Typing~~
    ~~- Single install; no packages/libraries~~
- "Why not HCL?"
  - It was created by/for Hashicorp, and their licensing is undesirable.
  - As a configuration language, HCL lacks features that facilitate higher levels of abstraction and reduce repetition

## Import existing infrastructure as CUE

TBD

## References

- DigitalOcean [API docs](https://docs.digitalocean.com/reference/api/digitalocean)
