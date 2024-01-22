# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning].
The format is based on [Keep a Changelog].

## [Unreleased]
- Added badges to README
  
## [v1.1.0] - 2024-01-22

- Embedded SVG images as a file system
- Improved test coverage to 100%

## [v1.0.0] - 2023-02-16

A Go library for applications that use playing cards.
Contains basic types for:
- `Suit` - SPADES, HEARTS, DIAMONDS, and CLUBS.
- `Rank` - 2 through Ace, with ordering functions to handle regular and pinochle decks.
- `Card` - A combination of `Rank` and `Suit`. The library contains SVG images for each card.
- `Deck` - A collection of `Card`s for a regular 52-card deck.
- `PinochleDeck` - A collection of `Card`s for a 48-card Pinochle deck.

## [v0.1.0] - 2023-02-08

Incomplete development version

[Semantic Versioning]: http://semver.org
[Keep a Changelog]: http://keepachangelog.com
[Unreleased]: https://github.com/philhanna/cards/compare/v1.1.0..HEAD 
[v1.1.0]: https://github.com/philhanna/cards/compare/v1.0.0..v1.1.0
[v1.0.0]: https://github.com/philhanna/cards/compare/v0.1.0..v1.0.0
[v0.1.0]: https://github.com/philhanna/cards/compare/ef5db44..v0.1.0
