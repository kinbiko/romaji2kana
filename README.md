# Romaji2Kana

This package exposes a function for converting [traditional hepburn romaji strings](https://en.wikipedia.org/wiki/Hepburn_romanization) into kana.

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
```

## Usage

See [the official documentation on godoc](https://godoc.org/github.com/kinbiko/romaji2kana) for usage.
