---
title: From over 200 MB of dependencies to less than 4 MB
subtitle:  Migrating my blog from Astro to Go
date: 2024-02-24
abstract: |
  How much work maintaining the dependencies is too much?
  How big must be the node_modules directory to start bothering you?
  Do you like spending an hour every now and then
  to adjust your code to the APIs you depend on,
  just to have up-to-date libraries?
  Are those breaking changes enough improvements
  to warrant you going back to your side project from a year ago
  just because a new vulnerability surfaced?
  Many of us don't realize that other people's code
  in our dependencies is still our liability.
  Library maintainers sure do improve them,
  but nobody will help us maintain the integration
  of those libraries with our own projects.
imgUrl: /from-over-200-mb-of-dependencies-to-less-than-4/simple-desk.jpg
imgDescription: A simple and clean desk
keywords:
  - JavaScript
  - Dependencies
  - Go
  - Blog
---

## An easy way in

Astro[^astro] is an amazing framework for static site generation.
It supports multiple UI libraries for defining components,
as well as its own component templating language.
When I started developing the second iteration[^second-iteration] of my blog,
I decided to use it, as I was somewhat familiar with React.

I obviously didn't choose React for my side project,
that would be too uncool.
I chose Solid[^solid].
The Developer Experience™ was not bad at all,
and I made myself a nice blog renderer,
together with a Continuous Delivery™ pipeline,
so the only thing left was to write great blog posts in Markdown.
Easy enough.

## Prepare to move with your framework

Library authors have every right to break all the APIs they wish.
Especially with open-source licenses,
like the one Astro uses, MIT,
the license literally screams at you with the following statement:

```
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

I'm very grateful for their work,
but that means if I want to be up-to-date with this framework,
I have to be prepared to make changes.
Changes probably leading into a better DX™,
but changes nonetheless.
And that means time.
Worse, that means switching my context in my free time
from reading, writing, or watching ThePrimeagen's hot takes,
to updating my dependencies.

A recent, third major release of Astro
brought many amazing things,[^astro-3]
but unfortunately I won't benefit from any of them in my blog.
What I will get instead,
is a couple of hours spent reading their guide
on deprecations and breaking changes[^astro-3-deprecations-and-breaking-changes]
to make sure my blog still renders with Astro 4.

I didn't have to wait that long for the next major release of Astro —
version 4[^astro-4] came out on December 5, 2023.
Because Astro 3 came out on August 30, 2023,
it is only 3 months between major updates.
I consider my hypothetical maximum writing capacity of one article per month.
That would make me do a major framework update every third article!
I already have enough reasons to procrastinate on writing articles,
and another one would probably grind my writing to a halt.

[^astro]: Astro. The web framework for content-driven websites. [[Online source](https://astro.build)] (accessed Feb 24, 2024).
[^second-iteration]: The first iteration was in Python and Django,
    and somehow had an Internal Server Error
    whenever I had a non-ASCII character.
    I was stupid back then, so I never figured out why.
    I still don't know why, so not much changed.

[^solid]: Solid. A declarative, efficient, and flexible JavaScript library for building user interfaces. [[Online source](https://www.solidjs.com)] (accessed Feb 24, 2024).
[^astro-3]: Third major release of Astro, Aug 30, 2023. [[Online source](https://astro.build/blog/astro-3)] (accessed Feb 24, 2024).
[^astro-3-deprecations-and-breaking-changes]: Astro 3 deprecations and breaking changes, 2023. [[Online source](https://docs.astro.build/en/guides/upgrade-to/v3)] (accessed Feb 24, 2024).
[^astro-4]: Fourth major release of Astro, Dec 05, 2023. [[Online source](https://astro.build/blog/astro-4)] (accessed Feb 24, 2024).

## Unplugging from the JavaScript Matrix

Frontend development has became synonymous with JavaScript
to the point that I chose Astro instead of Hugo,[^hugo]
because I didn't write any frontend code
that didn't involve JavaScript in my job.

Only when I learned about HTMX[^htmx]
I realized that I can generate my static site in whatever language I wanted.
Even though I didn't need to use HTMX for my blog,
through their essays and memes[^htmx-essays-and-memes]
I convinced myself that the JavaScript ecosystem
is not what I want and need.
I wanted a break from it all.

I was already familiar with Go and had used it for many side projects
and I always loved how stable and extensive the standard library is.
It has an HTML templating language built-in,
so it seems to be a perfect solution for generating a simple static website.

[^hugo]: Hugo The world’s fastest framework for building websites. [[Online source](https://gohugo.io)] (accessed Feb 24, 2024).
[^htmx]: HTMX. High power tools for HTML. [[Online source](https://htmx.org)] (accessed Feb 24, 2024).
[^htmx-essays-and-memes]: HTMX. Essays and memes. [[Online source](https://htmx.org/essays)] (accessed Feb 24, 2024).

## Minimal dependencies

When implementing the custom blog renderer,
I wanted to depend only on the standard library
and as little external libraries as possible.
Additionally, I wrapped each library into its own module,
so that it wouldn't spread into my application,
and it would be less painful to change if I needed to.
A thing that was extremely hard to do when using Astro
or any JavaScript framework like React or Solid.

I used about 3.5 megabytes of dependencies:
6 direct and 4 indirect, making it 10 in total.
In contrast, when using Astro,
I used over 200 megabytes of dependencies:
12 direct, and a whopping number of 768 indirect,
making it 780 in total.
That is 58 times less code and 78 times fewer dependencies.

Let me say that again: from over 200 MB to about 3.5!
If you wanted proof that the JavaScript ecosystem
makes it extremely easy to just add new layers of abstraction ad infinitum,
here you go.

### But there is a catch

Giving up those many layers of abstraction
lead to more _application_ code, though.
With Astro, I needed 2209 lines of code,
complete with all the JSON configuration files,
to implement my blog.
With Go, I needed 3309 lines of code,
so quite a lot more!

The bigger number isn't at all strange given that I needed to implement
a myriad of features that were already ready to use in Astro, like:

- Parsing Frontmatter for all articles,
- Sorting the articles based on the date in the Frontmatter,
- Copying static assets to the dist directory,
- Formatting the generated HTML,
- Re-rendering the changed pages.

I also didn't include some of the Astro niceties,
because the implementation was too complex for me,
like WebSocket-based hot module replacement,
or I simply didn't need them, like image optimization.

Obviously, I couldn't use the main selling point of Astro,
which is interoperability with different JavaScript component frameworks.
Fortunately, my blog isn't a very big project,
so rewriting everything in Go templates wasn't very hard.

The last annoyance was the lack of customizability
of the footnote section's heading.
I don't think it was customizable in Astro, either,
but the default behavior already satisfied me,
because the default heading for the footnote section was "Footnotes."
The Goldmark library[^goldmark] did not put any heading,
only a horizontal rule to mark the start of the footnotes.
So I needed to reimplement some internal methods
to put the heading back in its place.

[^goldmark]: Goldmark. A markdown parser written in Go. [[Online source](https://github.com/yuin/goldmark)] (accessed Feb 24, 2024).

## Conclusion

You are probably familiar with many such rewrite stories,
where a project moves away from JavaScript.
Less dependencies, faster development, better application performance.
My story was a bit different.
Yes, the number of dependencies decreased significantly,
but my biggest goal was to be able to forget about my project
for a couple of months and be able to just write an article,
without studying a migration guide for half a day.
I think I managed to do it,
because I've started writing this article November 24, 2023,
and went back to it on and off for three months.
Now it's February 24, 2024 and I didn't need to read any migration guide —
my blog renderer Just Works™.
