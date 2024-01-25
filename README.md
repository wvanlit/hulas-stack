# HULAS Stack

![Icon](./docs/hula-hoop.png)

A "modern" web stack based on **H**tmx, **U**nix, **L**ua, **A**i (Chat-GPT) and **S**qlite

## How to run 

First time? Run `build.sh && run.sh`

Docker already built? Then only execute `run.sh`

This initializes the docker container and launches the HULAS stack application. 

Your application should now be accessible at http://localhost:8080

## TODO List
- [x] Add some styling using DaisyUI
    - [x] Setup using Play CDN
    - [ ] Setup CLI tool with `--watch`
- [ ] Build out a library of utilities for apps
    - [ ] Possibly HTML builder / components
- [ ] Allow requesting `.html` files without specifing `[file].html` for better MPAs
- [ ] Build an example application using HULAS
- [ ] Develop hulas CLI tool
    - [ ] hulas create
    - [ ] hulas build
    - [ ] hulas run

## Attribution

<a href="https://www.flaticon.com/free-icons/hula-hoop" title="hula hoop icons">Hula hoop icons created by Freepik - Flaticon</a>