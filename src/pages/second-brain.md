---
layout: ../layouts/Post.astro
title: Software Engineer’s Second Brain
subtitle: Leveraging Engineering Principles for Knowledge Management
abstract: |
  What was that obscure Git command again?
  A quick Google search should do the trick,
  but was it the first, second, or third link?
  We all wish we had a sidekick who remembers
  all the relevant information at all times.
  Problem is, we will annoy people
  if we ask the same question multiple times.
  On the other hand, we can become annoyed
  with the overwhelming number of search results
  that are mostly irrelevant to us.
  What if there was a better way to manage knowledge
  and you, a software engineer, had all the skills
  to make it happen?
  Meet the Second Brain, a note-taking technique
  that not only will improve your Time To Answer,
  but also has a potential to improve your engineering skills
  and make you a happier person.
draft: true
draftDescription: |
  This article is a draft — some references may be missing,
  but the general idea should be complete.
imgUrl: /second-brain/network-switch.jpg
imgDescription: A network switch with patch cords connected to it
date: 2023-03-01
---

## Quick start

Starting with your Second Brain should be like starting
a new software product:
do just enough upfront work to get you started,
but try to make decisions that will speed you up in the future.

You could start a new software project
with only deciding on a language,
but the work will go much more smoothly
if you have a Continuous Delivery pipeline in place
and have a testing library set up with easy access through your IDE.
You are the user of the development process,
so create the best user experience you can.
It's called developer experience for a reason.

The two important factors contributing
to good note-taker experience are note capturing[^capture]
and ease of note organization.[^organize]
It is important not to compromise on either,
because otherwise you will just stop taking notes altogether.

Without notes, there are no notes,[^without-palace]
so it's crucial to lower the bar of taking them in the first place.
You also have to think about your ideal setup
for organizing text files.
You can read about mine in [My setup section](#my-setup).

[^capture]: [[code-notetaking-strategy]]
[^organize]: [[code-notetaking-strategy]]
[^without-palace]: In a 2002 movie "Astérix & Obélix: Mission Cléopâtre," Amonbofis says "Pas de palais... pas de palais," which translates to "Without palace... there's no palace." [[IMDb entry](https://www.imdb.com/title/tt0250223/characters/nm0201462)] (accessed Feb 22, 2023).

## Connecting ideas

One feature I find particularly useful
when building a Second Brain
is the ability to make and visualize connections
between ideas I would never be able
to make without such system.

For example, in Dave Airey's book "Logo Design Love,"
he found in his professional work that contacting
decision makers directly when getting feedback about
the logo design causes less trouble and disappointment
than contacting a middleman.[^contact-decision-makers-directly]
It is the same problem Extreme Programming mentions:
real customer involvement.[^real-customer-involvement]

It feels obvious and natural that there is
a connection between these two books,
but if you read them one or two years apart,
it is almost impossible to make such connections
without a robust set of notes.
Because of the recency bias,[^recency-bias]
we tend to better recall the information
we acquired recently.
The Second Brain can mitigate that,
opening new possibilities of research.

[^contact-decision-makers-directly]: D. Airey, Logo Design Love: A Guide to Creating Iconic Brand Identities, 2nd ed. New Riders, 2015. (p. 111).
[^real-customer-involvement]: K. Beck and C. Andres, Extreme Programming Explained: Embrace Change, 2nd ed. Addison-Wesley Professional, 2004. (pp. 61–62).
[^recency-bias]: Wikipedia, "Recency bias", 2022. [[Online source](https://en.wikipedia.org/wiki/Recency_bias)] (accessed Mar 2, 2023).

### Go to definition or references

When using [Foam][foam],
we can use WikiLinks to link our notes.
It's a common linking implementation,
which can be also found in different other note-taking applications,
like [Obsidian][obsidian], [Roam Research][roam-research],
[Logseq][logseq], or, well [MediaWiki][mediawiki].
Such links are simply file names of another notes
without the `.md` extension,
wrapped in double square brackets:

```md
# A note

This note links to [[another-note]].
```

One way to navigate through the connected ideas
is simply going through the links we made inside a note.
In Foam, it is simply done by putting a cursor on a WikiLink,
and pressing _Go to definition_ shortcut.
We can also navigate to references
from the _Backlinks_ panel,
similar to what we do in code.

### Visual navigation

Another great feature, which most applications implement
is a graph view of your notes.
Each note is represented by a node,
and each link as an edge.
You can explore your graph,
see how many links are between the ideas,
and discover new, surprising analogs,
which would be much harder to explore otherwise.

You can zoom out and see a general overview
to identify the most profound ideas,
spanning the entirety of your notes.

<figure>
  <img src="/second-brain/whole-second-brain-graph.png">
  <figcaption>An overview of a Second Brain graph</figcaption>
</figure>

You can also zoom in to see the individual notes
and connections between them.

<figure>
  <img class="wide" src="/second-brain/partial-second-brain-graph.png">
  <figcaption>A zoomed-in view of a Second Brain graph</figcaption>
</figure>

## Note-taking as programming practice

One thing every software engineer needs to do
is to practice.[^professionals-practice]
It is especially true for hard or common things
we all have to do every day.
Refactoring code, clearly expressing our ideas,
or giving them a proper name.
It turns out that note-taking
can help us with a lot of those things.

When I read a book or listen to a conference talk,
I make a note of each idea that sticks with me.
That means, I have to understand it well enough
to verbalize it into a comprehensible text.
I also need to keep in mind that I write it for my future self,
so it becomes a bit like explaining an idea to your colleague.

When I have that idea in writing,
I have to come up with a clear, short name,
which can be easily searched for,
and will link nicely with other ideas.

[^professionals-practice]: R. C. Martin, The Clean Coder: A Code of Conduct for Professional Programmers. Prentice Hall, 2011. (Chapter 6).

## Moving on

Let's say you have some unorganized, "legacy" notes
spread across different directories or services.
It would be a good idea to have them organized and properly linked,
but it would require a lot of time and effort
that you probably don't have.

Just like with technical debt in software,
unorganized notes are only bad
if we have to deal with them.
So my advice would be to gather all of the legacy notes
into one place, for example, a directory named "archive"
and just move on with your current notes.

The same goes for the projects you've already finished.
Take a glance if some notes would be useful for some future projects
and organize those notes,
but put all the rest in an archive.
Move on with your projects, work, knowledge, and life.

Moving inactive notes to archive
will make your workspace cleaner,
which in turn will help you focus
on your present goals instead of
going through a bunch of old notes.[^move-to-archive]

Remember, you can always search through all your notes
and if something interesting pops up in your archived ones,
you can always move or copy them to your current set of notes.

[^move-to-archive]: T. Forte, Building a Second Brain: A Proven Method to Organise Your Digital Life and Unlock Your Creative Potential. Profile, 2022. (pp. 102-108).

## Happiness

Externalizing our thoughts through writing has a lot of benefits,
both in psychological and even physical health.[^writing-is-healthy]
It can reduce stress and make us more present in our daily life.
Think about different situations
when you just couldn't get rid of those excruciating thoughts:

> I have to remember to take out the laundry
> before I go to work. Oh, and my friend has birthday today.
> Hmm... I think I've just found a solution to a bug
> I've investigated yesterday, so I have to remember that as well.

You get the picture.

Instead, just set a timer for an hour before the laundry is done,
put your friend's birthday into a calendar,
and a solution to a bug into your note-taking app.
You will immediately feel relieved.
Over time your brain will trust the system more and more
and you will become less stressed,
at least when you will have to remember something.

[^writing-is-healthy]: T. Forte, Building a Second Brain: A Proven Method to Organise Your Digital Life and Unlock Your Creative Potential. Profile, 2022. (p. 77).

## My setup

Although my setup is in [VSCode][vscode],
which I use every day for development anyway,
unless you have a similar situation,
I'd advise you to start with [Obsidian][obsidian],
which will provide a wonderful note-taking experience
out of the box.

My ideal setup for organizing text files
has to be version controlled with Git,
use Markdown format,
have search with regex,
and support Vim keybindings.
That's why I've set it up like any other Git repository,
and I organize it in [VSCode][vscode] with [Foam][foam] plugin.
I've written a [Github Action][issue-to-note-action] that automatically
converts a Github issue to a Markdown file in an "inbox" directory.

[foam]: https://github.com/foambubble/foam
[issue-to-note-action]: https://gist.github.com/sewera/f9241f4bbc7f73f27b6dde99cc896672
[logseq]: https://logseq.com
[mediawiki]: https://www.mediawiki.org
[obsidian]: https://obsidian.md
[roam-research]: https://roamresearch.com
[vscode]: https://code.visualstudio.com
