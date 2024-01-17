# Golb

Golb is a simple static site generator for simple blog. It aims to reduce the
overwhelming complexity of existing blog generators. 

### Basic usage
Create a folder anywhere named after your blog and inside that folder, create
another folder named posts that will contain all of your blog posts.
```
mkdir -p ~/my_blog/posts
``` 

All blog post must be Markdown files. Naviguate to the posts directory and
create a new blog post:
```
cd ~/my_blog/posts
touch first_post.md
```

Once you finished to write your new post, simply run the following command:
```
golb ~/my_blog
```

The above command will interpret the all the Markdown files in your posts
directory into html files and your append a new link on your blog's landing page
to your new post


### Basic blog post elements
golb supports most of the Mardown elements.
```
# First Title
## Second Title
### Third Title
#### Fourth Title

<image file name> Images

[Link text](Link Url)

Regular Text

**Bold**
__Bold__

*Italic*
_Italic_
```

### Images
To add an image to a blog post, add the image into the `image` folder and add
an element to your blog post like so:
`<image file name>`