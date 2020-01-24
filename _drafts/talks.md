---
title: Talks
author_profile: true
permalink: /talks/
draft: true
---

## Invite me

I deliver talks on:

* taster university engineering lectures
* teacher training
* engineering ethics
* more

## Past talks

{% for post in site.talks reversed %}
  {% include archive-single-talk.html %}
{% endfor %}
