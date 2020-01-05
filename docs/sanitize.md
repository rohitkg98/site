## Parsing and Sanitizing markdown to html for rendering.

- Struct to store a line.

struct Line {  
  content string  
  type    string  
}

- content  
    contents of the line.

- type
    > use the start of the line to determine type.  
    > apply css tags accordingly.  
    > type are like: heading, bullet points etc.  
