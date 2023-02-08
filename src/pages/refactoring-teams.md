---
layout: ../layouts/Post.astro
title: Refactoring Teams
subtitle: Aligning people for fast and safe flow
date: 2023-02-05
abstract: |
  The idea of refactoring is well known by most software engineers.
  Because many software managers are former engineers,
  they are usually also familiar with this term.
  However, with time, when managers become “the business people,”
  they tend to treat refactoring as a necessary evil,
  only slowing down the development of new features.
  The idea of refactoring can be applied not only to code,
  but also to teams organization,
  which can considerably speed up the development.
  I hope that it becomes a valuable tool
  for every software engineering manager
  who treats their job seriously.
imgUrl: /refactoring-teams/tour-de-france.jpg
imgDescription: A professional cycling team in Tour de France
keywords:
  - Team Topologies
  - Management
  - Software Engineering
  - Refactoring
---

A quick note about the cover image:
In cycling, a leader needs to spend the most energy
to overcome the air resistance.
If the team is well aligned in the slip stream,
it experiences far less resistance than when cycling alone.
In the middle of the peloton individual cyclists
experience only 5–10% of air resistance
compared to cycling alone.[^peloton-and-air-resistance]
I believe a good leader in a company should strive
to align their team in such a slip stream
in order to enable fast flow of value delivery.

[^peloton-and-air-resistance]: Eindhoven University of Technology, "A cyclist in a peloton experiences considerably less air resistance than previously assumed", Jun. 29, 2018. [[Online source](https://www.tue.nl/en/news-and-events/news-overview/a-cyclist-in-a-peloton-experiences-considerably-less-air-resistance-than-previously-assumed/)] (accessed Feb. 05, 2023).

The term "refactoring" has become frowned upon by some managers.
Putting aside the validity of such aversion,
I think some refactoring techniques can be used
to realign the team with its desired value stream.
They can help manage the cognitive load[^team-cognitive-load] of a team,
bring development closer to business needs,
and mitigate the risks of management debt.[^management-debt]

[^team-cognitive-load]: M. Skelton and M. Pais, Team Topologies: Organizing Business and Technology Teams for Fast Flow. IT Revolution, 2019. (p. 11).
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
This tool is crucial for managing complexity,[^experts-in-managing-complexity]
which makes introducing changes useful for business easier,
improves testability,
and makes observability more robust.[^effective-software-development]
When done correctly, it dramatically speeds up development,
makes application support easy,
and enables the business to make decisions based on facts,
not opinions,[^decisions-based-on-facts]
empowering us to become experts in learning.[^experts-in-learning]

[^decisions-based-on-facts]: D. Farley, Modern Software Engineering: Doing What Works to Build Better Software Faster. Addison-Wesley, 2022. (p. 3, 6).
[^effective-software-development]: D. Farley, Modern Software Engineering: Doing What Works to Build Better Software Faster. Addison-Wesley, 2022. (p. 5).
[^experts-in-managing-complexity]: D. Farley, Modern Software Engineering: Doing What Works to Build Better Software Faster. Addison-Wesley, 2022. (p. 5, 38).
[^experts-in-learning]: D. Farley, Modern Software Engineering: Doing What Works to Build Better Software Faster. Addison-Wesley, 2022. (p. 4, 36).

It's all well, every developer knows it.
But how does it translate to management?
Let's take a look at a couple of code smells[^code-smells]
and see if there are any analogous _management smells_.

[^code-smells]: M. Fowler, Refactoring: Improving the Design of Existing Code. Addison-Wesley, 2019. (Chapter 3).

## Management smells

A management smell, similar to code smell is a pattern indicating
a situation that is potentially risky to keep.
In code, it is duplication, low cohesion, high coupling, and so on.
If we don't refactor those smells,
they have a potential to cause much bigger problems,
or simply more expensive development.

### Cognitive overload

The most obvious code smell which has its counterpart in management
is Large Class.[^large-class]
The first idea that comes to our mind
is too big of a team.
It isn't only big teams, though.
This code smell indicates that a class is **trying to do too much**,
so a better analog is a team that has too big of a domain,
or even a couple of them.

Then, the relative domain complexity[^relative-domain-complexity]
overloads the team's cognitive load,
and such a team can't do its work efficiently anymore.
People on such team spend most of their time
switching between contexts,
trying to orchestrate their work around
constraints imposed by some middle management, like
"up to 20% of team's time should be
dedicated for maintenance of the old projects."

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
The problematic thing about assigning a certain percentage
of time to different projects is that it comes from the assumption
that there is zero context-switching overhead.
It also takes away the mastery[^mastery] of the domain,
making the engineers feel less motivated.

[^large-class]: M. Fowler, Refactoring: Improving the Design of Existing Code. Addison-Wesley, 2019. (p. 82).
[^relative-domain-complexity]: Relative domain complexity is an internal metric to assess the cognitive load on a team working on all assigned domains.
[^cognitive-load-feedback]: M. Skelton and M. Pais, Team Topologies: Organizing Business and Technology Teams for Fast Flow. IT Revolution, 2019. (p. 41).
[^mastery]: One of the three elements of intrinsic motivation. Based on: M. Skelton and M. Pais, Team Topologies: Organizing Business and Technology Teams for Fast Flow. IT Revolution, 2019. (p. 11).

### One key person

A similar, but arguably more problematic smell
is when a person is the only one who knows some area of team's work.
Just like a long function,[^long-function]
or a brain method in other words,
most work in a team will inevitably
end up having such a person as a dependency.

The one key person in a team may become such
by a number of reasons:
they never feel the need to explain anything to their colleagues,
they don't know how to clearly communicate, or
their manager actively prevents them from any communication,
treating it as a "distraction."
Such people fix everything,
yet never explain how they did it.
A good example in literature is Brent,[^brent]
the most experienced IT engineer in Parts Unlimited,
a fictional company from an amazing book — The Phoenix Project.

When the management started measuring what impact
he has on the work of the entire IT department,
it became clear that such person
is a massive risk for the organization.
When they are at work, they become a bottleneck,
and all work gets slowed down.
A bigger problem arises when they are out.
Then, the work just doesn't get finished.

If the knowledge is not spread across team members,
this area of work has a bus factor[^bus-factor] of one.
Managers have to mitigate this risk
by encouraging people to share information between one another.
It can be done with pair (or mob) programming,[^pair-programming]
one of the key practices of Extreme Programming.

[^brent]: G. Kim, K. Behr, and G. Spafford, The Phoenix Project: A Novel about IT, DevOps, and Helping Your Business Win, 3rd ed. IT Revolution, 2018. (pp. 115–116).
[^long-function]: M. Fowler, Refactoring: Improving the Design of Existing Code. Addison-Wesley, 2019. (p. 73).
[^bus-factor]: Wikipedia, "Bus factor", 2023. [[Online source](https://en.wikipedia.org/wiki/Bus_factor)] (accessed Feb. 08, 2023).
[^pair-programming]: K. Beck and C. Andres, Extreme Programming Explained: Embrace Change, 2nd ed. Addison-Wesley Professional, 2004. (pp. 42–44).

### Blame game

Another code smell is the "log and throw" pattern.[^log-and-throw]
In code, it can indicate that the person writing the implementation
didn't know how to deal with the unhappy path,
and just passed the buck to some other part of the code.
It's the same mindset that drives us to play the blame game.

A blame game is a common phenomenon in low-trust teams
and organizations, which focuses mostly on finding
the person responsible for a failure that occurred,
but not how to fix the issue,
or prevent such failure in the future.[^blame-game]

People playing it aim at avoiding solving the problem themselves,
just like with the "log and throw" pattern.
Playing the blame game is very counterproductive,
and is not in the organization's interest, team interest,
nor should be in the individual's interest.

A good manager needs to encourage their team
to find the best solution to the problem at hand,
not a guilty person.
To achieve this, a team must be of high trust,
and a safe environment has to be formed within a team,
Such safe-for-failure environment fosters innovation.

[^log-and-throw]: H. Kevlin. "Clean Coders Hate What Happens to Your Code When You Use These Enterprise Programming Tricks", Feb. 27, 2017. [[Online video](https://youtu.be/FyCYva9DhsI?t=3306)]. (55:06–55:12) (accessed: Feb. 08, 2023).
[^blame-game]: G. Kim, K. Behr, and G. Spafford, The Phoenix Project: A Novel about IT, DevOps, and Helping Your Business Win, 3rd ed. IT Revolution, 2018. (pp. 104–106).

To be continued.
