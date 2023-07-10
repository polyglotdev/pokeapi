# Engineering Take Home Project: Develop an SDK for PokeAPI

## Overview

For this project, envision yourself as a software engineer working with the launch team
for a new API product: PokeAPI, a comprehensive database containing all the
information you'll ever need about Pokémon!
Your assignment involves constructing an SDK (Software Development Kit) for this API,
utilizing any programming language you're comfortable with. The main goal here is to
make the API more user-friendly and easily accessible to developers. Here is an
example of a SDK constructed in Python for [Petstore](https://github.com/speakeasy-sdks/template-sdk)

Feel free to incorporate any tools, including code generators, that may assist you in
building the SDK. You may build this in a language of your choosing

In addition to this, please design and implement an integration test that makes use of your SDK to send an API request and validate the response. Again, you are free to use any tool that you deem fit for this purpose.

## Project Scope

The project involves working with the following endpoints from the [PokeAPI](https://pokeapi.co/docs/v2#pokemon). Specifically, you'll need to cover:

- `GET https://pokeapi.co/api/v2/pokemon/{id or name}/`
- `GET https://pokeapi.co/api/v2/nature/fid or name?/`
- `GET https://pokeapi.co/api/v2/stat/{id or name}/`

## Key Guidelines

> :bulb: Please ensure to include a README file alongside your SDK for user guidance.

This should clearly detail the installation process, usage, and testing instructions for the
SDK.

It's important to note that, **the SDK doesn't necessarily have to replicate the API exactly**.
You're encouraged to add abstractions and/or amalgamate different calls to enhance the
functionality.

Please add a readme file for your SDK users specifying how to install, use and test the
SDK.
