## Workflow of site

==========

### Directories

----------

- ./content/ contains all the blog posts in markdown.
- ./templates/ contains scaffold of html pages to which we write our blog post.

### Server

----------

- main.go entry point for all code.
- main will bind endpoints to their respective handler functions.

### Pages

----------

- Home
> description: feed of posts, sorted data wise. also introduction.  
> implementation: Walk dir and send list to template.  
> template: home.html

- Post
> description: A post, which will be bound to an md file object.  
> implementation: bind content file name to an endpoint of same name.  
> template: post.html  

- About
> description: about my self  
> implementation: static  
> template: about.html

- Resume
> description: My resume in detail.  
> implementation: static  
> template: resume.html

- Contact
> description: contains contact info about me.  
> implementation: static  
> template: contact.html
