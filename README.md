# Emo 中文官网 (website-cn)

The Chinese marketing website for [Emo](https://github.com/emo-lang/emo) —
clean, explicit, intuitive. Built with [Airway](https://github.com/daqing/airway)
(Gin + templ), server-rendered with no client-side framework required for the
marketing pages.

## Pages

| Route | View | Content |
|---|---|---|
| `/` | `app/views/home` | Landing: hero, five compilation targets, feature grid, concurrency / interfaces / EmoUI / EmoOS sections |
| `/features` | `app/views/features` | The language's design decisions, section by section |
| `/tour` | `app/views/tour` | The language tour in eight annotated steps |
| `/quickstart` | `app/views/quickstart` | Build the compiler, first program, native builds, packages, examples |

Supporting pieces:

- `app/views/layouts` — page shell (`Base`), brand design system (`SiteStyles`), shared `Nav`/`Footer`
- `app/views/emocode` — a small Emo syntax highlighter plus the terminal-style code block used across the pages
- `app/assets/public` — committed static files (logo), served at `/public/` via `assets.PublicHandler`

The brand palette matches the English site: Yellow `#EFBF6B` · Dark `#202022` ·
Grey `#F9FAFC` · Blue `#7678ED`, with light/dark support through
`prefers-color-scheme`.

## Setup

`.env` is created at scaffold time — set `AIRWAY_ENV` (e.g. `local`) and the
`LISTEN` address (`host:port`, e.g. `:1900`). The marketing pages need no
database; `DSN` is only required if you use models.

```bash
go generate ./...   # regenerate *_templ.go after editing .templ files
go run .            # start the HTTP server
go test ./...       # run the test suite
```

`templ` syntax pitfalls learned on this project (the parser errors are
cryptic, so keep these in mind when editing views):

- Element text cannot start with `if ` / `for ` / `switch ` — templ reads it
  as a control-flow statement. Reword the sentence instead.
- Literal `{` and `${` in text must be written as HTML entities
  (`&#123;`, `&#125;`) — they are expression delimiters in templ.
- Newlines inside text are collapsed; use `<br/>` for pre-formatted command
  blocks.

## Static export

`airway static:build` renders all four pages into `dist/` (registered in
`export.go`) together with the committed frontend bundle.
