## Golang app template 


### Project structure
```
├── cmd                   # Command line entrypoints
├── db                    # Database related operations and migrations
├── internal              # Internal packages
├── models                # Business models
├── routes                # HTTP handlers 
│   ├── api               # HTTP application/json
│   └── app               # HTTP [x-www-form-urlencoded,text/html]
└── views                 # HTML templates
    ├── components        # Reusable components
    ├── layouts           # Layouts
    └── locations         # Your domain specific views
```