---
layout: ../layouts/Post.astro
title: Refactoring Management
date: 2023-02-05
draft: true
draftDescription: |
  This article is incomplete and lacks proper references.
  It will be extended and completed soon.
imgUrl: /refactoring-management/team-rafting-on-a-waterfall.jpg
imgDescription: A team rafting on a waterfall
---

The term "refactoring" has become frown upon between some managers.
Putting aside the validity of such aversion,
I think some refactoring techniques can be used
to realign the team with its desired value stream.
They can help manage the relative domain
complexity[^relative-domain-complexity] of a team,
bring development closer to business needs,
and mitigate the risks of management debt[^management-debt].

[^relative-domain-complexity]: [[relative-domain-complexity]]
[^management-debt]: Similar to technical debt being suboptimal code architecture, which is harder to maintain, management debt is suboptimal organization architecture, which makes fast flow harder. More about this term in [this section](#management-smells).

From my experience, software engineering managers
are usually former software engineers.
I, therefore, find it helpful to present
the ideas that can improve organization architecture
in terms understood by their past self.
When we think about it,
teams are more complex and less malleable analog to software.
People in teams have to communicate with other teams,
and do useful work for business,
just like every modern software has to communicate with other software,
and do useful work for its users.

That's why software engineers do refactors of their software.
This tool is crucial for managing complexity[^experts-in-managing-complexity],
which makes introducing changes useful for business easier,
improves testability,
and makes observability more robust[^effective-software-development].
When done correctly, it dramatically speeds up development,
makes application support easy,
and enables the business to make decisions based on facts,
not opinions[^decisions-based-on-facts],
empowering us to become experts in learning[^experts-in-learning].

[^decisions-based-on-facts]: [[farley-2022]] [[apply-scientific-methods-in-engineering]]
[^effective-software-development]: [[farley-2022]] [[tools-for-effective-strategy-of-software-development]]
[^experts-in-managing-complexity]: [[farley-2022]] [[become-an-expert-in-managing-complexity]]
[^experts-in-learning]: [[farley-2022]] [[become-an-expert-in-learning]]

It's all well, every developer knows it.
But how does it translate to management?
Let's take a look at a couple of code smells[^code-smells]
and see if there are any analogous _management smells_.

[^code-smells]: [[fowler-2019]]

## Management smells

The most obvious code smell which has its counterpart in management
is Large Class[^large-class].
It isn't only too big teams, though.
This code smell indicates that a class is **trying to do too much**,
so a better analog will be a team that has too big of a domain,
or even a couple of them.
Then, the relative domain complexity[^relative-domain-complexity]
overloads the team's cognitive load,
and such a team can't do its work efficiently anymore.

We have static code analyzers that tell us when a class
has too many fields, dependencies, or lines of code.
We don't have such easy metrics for measuring the
team's cognitive load.
We can, however, gather feedback from people,
and just ask them:

> Do you feel like you are effective
> and able to respond in a timely fashion
> to the work you are asked to do?[^cognitive-load-feedback]

Clearly negative response should bring our attention.

A similar, but arguably more problematic smell
is when a person is the only one who knows some area of team's work.
Just like a long function[^long-function],
or a brain method in other words,
most work in a team will inevitably
end up having such a person as a dependency.

An all-knowing person is a massive risk for the organization.
When they are at work, they become a bottleneck,
and all work gets slowed down.
A bigger problem arises when they are out.
Then, the work just doesn't get finished.
If the knowledge is not spread across team members,
this area of work has a bus factor of one.
Managers have to mitigate this risk
by encouraging people to share information between one another.
It can be done with pair programming.

[^large-class]: [[fowler-2019]], p. 82
[^cognitive-load-feedback]: [[relative-domain-complexity]]
[^long-function]: [[fowler-2019]], p. 73
