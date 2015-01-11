# Today, an experiment in better days.

_This is an excuse for me to improve proficiency with [Go](https://golang.org).
If something useful happens, that is on accident._

## Important things through the day

This utility helps to proactively nudge the user into acknowledging important
things, and if needed, write them down.

## Review

Reviewing what you have left to do, and what you set out to do yesterday is
pretty important. The default operation of `today` is to show the review view.
If all items are listed as done, then there is helpful prompts to plan for your
upcoming day.

When you want to mark stuff off, use the `did` command.

## Planning with `must`, `should`, and `want`

Prioritization is for Six Sigma Blackbelts. What matters most to me is focusing
on important things. Urgent things get taken care of automatically, or they
really aren't that urgent.

I categorize things to do each day as a `must`, in that it is very important
but probably not urgent. If I do this item, I will feel very good about my day
and feel like a superhero.

The `should` items are about increasing your capacity to produce. Learning new
tools and techniques fall under this category.

And we humans, with desires and emotions that often sabotage our progress just
as often as they motivate us. By listing out things we simply `want` to do, we
ensure we're taking care of our gentler side. The `wants` are there to ensure
we enjoy our lives.

## Recording with `did`

### Marking planned activities as done

Running `today did` will display a view where you can quickly mark off what you
got done, and add notes to each item as necessary. The notes here are formatted
in Markdown.

Anything that is not marked as done within 24 hours will disappear from view
unless you explicitly ask for it.

### Free form notes

Sometimes you just want to jot down that something happened. That's pretty easy
to do with the `did` command!

`today did something neat that happened`

But wait, what if you want a big long message? Perhaps in Markdown? With tags?

Can you do that, too? Yes, yes you can! Just use `-m` and your `$EDITOR` will
fire up and you can write out some nice Markdown notes.

## Summary Displays

## The Today Server

A full featured HTML-based GUI view is packaged up, so you can get great
visibility into your data as you accumulate.

Run `today server` to start this up (or keep it running via `upstart` or
`launchtl`!)
