---
title: Hello
author_profile: true
permalink: /talks/
---

## Invite me

I deliver talks on:

* taster university engineering lectures
* teacher training
* engineering ethics
* more

## Past talks

* list coming soon

{% for post in site.talks reversed %}
  {% include archive-single-talk.html %}
{% endfor %}
