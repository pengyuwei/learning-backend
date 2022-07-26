#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from colorama import Fore, init


def main() -> None:
    init(autoreset=True)
    print(Fore.RED + "Hello ", end='')
    print(Fore.YELLOW + "Color ", end='')
    print(Fore.GREEN + "World!")
    print("Do it.")


if __name__ == "__main__":
    main()
