(bundle exec jekyll serve --incremental --drafts) | while read line; do 
    [[ $line =~ "Server running" ]] && open http://127.0.0.1:4000
done
