# Golb

Golb is a simple static site generator for simple blogs. It aims to reduce the
overwhelming complexity of existing blog generators.

### Requirements
The `go` toolchain is required in order to build the project.

### Basic usage
Init your new blog by using the init command following your new blog name
```
golb --init [BLOG NAME]
```

All blog post must be Markdown files. Naviguate to the posts directory and
create a new blog post:
```
cd ~/[BLOG NAME]/posts
touch first_post.md
```

Once you finished to write your new post, simply run the following command:
```
golb --build
```

The above command will interpret the Markdown files in your posts
directory into html files and your append a new link on your blog's landing page
to your new post

To run your website, simply run the following command at the root of your blog:
```
golb --serve
```

### Basic blog post elements
golb supports most of the Mardown elements.

**Titles**
```
# First Title
## Second Title
### Third Title
#### Fourth Title
```

**Images**
```
<image file name>
```
**Text**
```
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
