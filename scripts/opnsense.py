#!/usr/bin/env python3
#
#


import json
import requests

from os import getenv
from requests.auth import HTTPBasicAuth


# using some global variables for testing
api_key = getenv("OPN_KEY")
api_secret = getenv("OPN_SECRET")
base_url = getenv("OPN_URL")


def get_upgrade_available() -> None:
    # check for updates
    session = requests.Session()
    session.auth = HTTPBasicAuth(api_key, api_secret)
    session.headers.update({"Accept": "application/json"})
    r = session.get(
        f"{base_url}/core/system/status",
        timeout=15,
        verify=False,
    )
    r.raise_for_status()

    print(json.dumps(r.json(), indent=2))

    r = session.get(
        f"{base_url}/core/firmware/status",
        timeout=15,
        verify=False,
    )
    r.raise_for_status()

    print(json.dumps(r.json(), indent=2))


def get_interface_stats() -> None:
    session = requests.Session()
    r = session.get(
        f"{base_url}/diagnostics/interface/get_interface_statistics",
        timeout=15,
        verify=False,
    )
    r.raise_for_status()
    iface_stats = r.json()
    print(json.dumps(r.json(), indent=2))

    rows = iface_stats.get("rows") or iface_stats.get("interfaces") or []
    if isinstance(rows, list) and rows:
        print("summary")
        for it in rows:
            name = it.get("name") or it.get("interface") or "?"
            rx = it.get("ibytes") or it.get("rxbytes") or it.get("rx") or "?"
            tx = it.get("obytes") or it.get("txbytes") or it.get("tx") or "?"
            print(f"{name}: RX={rx} TX={tx}")


if __name__ == "__main__":
    get_upgrade_available()
