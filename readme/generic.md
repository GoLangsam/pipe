# Generic

As of now, the 'generics' of this project is powered by [genny](https://github.com/cheekybits/genny).

In `pipe/s` [precedessor project](https://github.com/GoLangsam/AnyType/)
a (self-made) agnostic template-driven all-purpose generator was used:
[dotgo](https://github.com/GoLangsam/dotgo/).

Such a generator can be useful in lot's of situations.

Just: To generate program code is not among such.
The maintenance effort grows considerably as a project evolves.
This does not come as a surprise, this is common knowledge - just look around the community.

Thus: A dedicated generator to the rescue.

Different approaches have been taken.
Many implementations (or attempts thereof) can be found -
from overly simple 'proof-of-conept'
up to highly sophisticated
'bells and whistles' such as traits.

For the project at hand - the most suitable choice is the middle way.

For the project at hand - what is important?

- Real code:

  Real code can be compiled & tested right away.
  The efficiency thus gained is considered more important than irrelevant 'bells and whistles',

  Tools which embed the code in templates, typewriters, strings-in-code or wherever are avoided.

- Documentation:

  Names & ID's must be adjusted in comments as well.
  (Some 'proof-of-conept' implementations fail do do so.)

- Ease of use!

  (We're lazy, are we not?)

- AST-based transformation:

  appreciated, but not mandatory.

- Popularity and a good userbase:

  favourable.

Thus: [genny](https://github.com/cheekybits/genny) became chosen.

Note: Later on we've learned:
maintenance and improvement stall.
Other forks progress...

That's why the leading sentence has to begin with "As of now".
A future major version might use another tool... time will tell.

## Generators

Tools we looked at - an incomplete list:

- [Using code generation to survive without generics in Go](https://dev.to/joncalhoun/using-code-generation-to-survive-without-generics-in-go)
Uses simple {{.Name}} & {{.Type}} templates (no multi-types). Leaves 'import' to `goimport`.

- [gen](http://clipperhouse.github.io/gen/) uses typewriters ...

- [genny](https://github.com/cheekybits/genny) uses `generic.Type`

- [gengen](https://github.com/joeshaw/gengen)

- [StaticTemplate](https://github.com/bouk/statictemplate)
is a code generator for Go's text/template and html/template packages.
  - It works by reading in the template files, and generating the needed functions based on the combination of requested function names and type signatures.
  - Please read [my blogpost](http://bouk.co/blog/code-generating-code/) about this project for some background.

---
## Related materials

- [Experience Reports - Generics](https://github.com/golang/go/wiki/ExperienceReports#generics)

- [The problem with interfaces](https://deedlefake.com/2017/07/the-problem-with-interfaces/)

  "Generics, specifically type parameters, are the standard solution to these issues.
  Though it may not be the best solution, examples of the need for a solution are so rampant that they can even be found throughout the standard library, and particularly in the various `container` subpackages, and most recently in the form of `sync.Map`.
  Just the fact that there are so few packages under `container` could be considered an example of these problems."

---
[Back to overview](overview.md)
