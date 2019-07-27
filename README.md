# Romaji2Kana

[![Build Status](https://travis-ci.com/kinbiko/romaji2kana.svg?branch=master)](https://travis-ci.com/kinbiko/romaji2kana)
[![Go Report Card](https://goreportcard.com/badge/github.com/kinbiko/romaji2kana)](https://goreportcard.com/report/github.com/kinbiko/romaji2kana)
[![Coverage Status](https://coveralls.io/repos/github/kinbiko/romaji2kana/badge.svg)](https://coveralls.io/github/kinbiko/romaji2kana)
[![Latest version](https://img.shields.io/github/tag/kinbiko/romaji2kana.svg?label=latest%20version&style=flat)](https://github.com/kinbiko/romaji2kana/releases)
[![Go Documentation](http://img.shields.io/badge/godoc-documentation-blue.svg?style=flat)](http://godoc.org/github.com/kinbiko/romaji2kana)
[![License](https://img.shields.io/github/license/kinbiko/romaji2kana.svg?style=flat)](https://github.com/kinbiko/romaji2kana/blob/master/LICENSE)

> **This repository is now deprecated in favour of [nihon-go](https://github.com/kinbiko/nihon-go), which contains a function to achieve the same behaviour, and more.**

This package exposes a `ToKana` function for converting romaji strings into kana.

## Restrictions

This package was written in order to allow myself to write quick notes without switching keyboard layout when taking notes (but rather process these later).
Therefore, there's very little 'intelligence' in this package.
For example, there's no grammatical inference of whether 'wa' is「は」or 「わ」.

Moreover the form in which the kana should be written is very strict.
In particular, only the following form is supported:

```
toukyou - とうきょう
```

The following romanization forms will not work as one might expect:

```
tōkyō - error
tokyo - ときょ
tohkyoh - error
tookyoo - とおきょお
```

The package expects **hiragana to be all lower case** and **katakana to be all upper case**.

```
Toukyou - error
TOUKYOU - トーキョー
```

## Usage

See [the official documentation on godoc](https://godoc.org/github.com/kinbiko/romaji2kana) for usage.
