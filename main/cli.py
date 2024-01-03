"""CLI interface for bitcoin project.

Be creative! do whatever you want!

- Install click or typer and create a CLI app
- Use builtin argparse
- Start a web application
- Import things from your .base module
"""
import os
# P2PKH Script
from bitcoin import *
from bitcoin.wallet import CBitcoinSecret, P2PKHBitcoinAddress


def main():  # pragma: no cover
    """
    The main function executes on commands:
    `python -m bitcoin` and `$ bitcoin `.

    This is your program's entry point.

    You can change this function to do whatever you want.
    Examples:
        * Run a test suite
        * Run a server
        * Do some other stuff
        * Run a command line application (Click, Typer, ArgParse)
        * List all available tasks
        * Run an application (Flask, FastAPI, Django, etc.)
    """
    

    # Generate a random private key
    private_key = CBitcoinSecret.from_secret_bytes(os.urandom(32))
    # Derive the public key and Bitcoin address
    public_key = private_key.pub
    address = P2PKHBitcoinAddress.from_pubkey(public_key)
    print("Private Key:", private_key)
    print("Public Key:", public_key.hex())
    print("Bitcoin Address:", address)
    print("This will do something")
