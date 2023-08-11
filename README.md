# openfga-demo
Openfga demo for self-learning

[![issue](https://img.shields.io/github/issues/MGMCN/openfga-demo?logo=github)](https://github.com/MGMCN/openfga-demo/issues?logo=github)
[![license](https://img.shields.io/github/license/MGMCN/openfga-demo)](https://github.com/MGMCN/openfga-demo/blob/main/LICENSE)
![last_commit](https://img.shields.io/github/last-commit/MGMCN/openfga-demo?color=red&logo=github)

## Usage
```Bash
$ docker pull openfga/openfga
$ docker run -p 8080:8080 -p 3000:3000 openfga/openfga run
```
Access [http://localhost:3000/playground](http://localhost:3000/playground) and create a store called ```demo```.  
Modify the AUTHORIZATION MODEL to 👇🏻
```
model
  schema 1.1
type user
type admin
type group
  relations
    define member: [admin,user]
type folder
  relations
    define can_create_file: owner
    define owner: [admin]
    define parent: [folder]
    define viewer: [admin,user,user:*,group#member] or owner or viewer from parent
type doc
  relations
    define can_change_owner: owner
    define can_read: viewer or owner or viewer from parent
    define can_share: owner or owner from parent
    define can_write: owner or owner from parent
    define owner: [admin]
    define parent: [folder]
    define viewer: [admin,user,user:*,group#member]
```
Run the demo code
```Bash
$ go get -u github.com/openfga/go-sdk
$ go mod tidy
$ go build -o demo ./main
$ ./demo
The permission for gaoshan to read the 'doc1' file is true
The permission for gaoshan to write the 'doc1' file is true
The permission for normal_user to read the 'doc1' file is true
The permission for normal_user to write the 'doc1' file is false
```

## How to create relation tuples
We used the 'gaoshan' user under 'admin' to create the 'docs' folder, and then established a relationship with the 'doc1' file located within the 'docs' folder. We also configured normal users to be able to view the 'docs' folder.
```Go
if err := client.CreateRelationTuple(ctx, "folder:docs", "owner", "admin:gaoshan"); err != nil {
  fmt.Println(err)
}
if err := client.CreateRelationTuple(ctx, "folder:docs", "viewer", "user:normal_user"); err != nil {
  fmt.Println(err)
}
if err := client.CreateRelationTuple(ctx, "doc:doc1", "parent", "folder:docs"); err != nil {
  fmt.Println(err)
}
```

## Test the access permissions for gaoshan and normal_user
It's important to note that gaoshan is the owner of the 'docs' folder and has read-write permissions for all the files within that folder. In the AUTHORIZATION MODEL, we specify that normal_user is only allowed to read.
```Go
if permission, err := client.GetCheck(ctx, "doc:doc1", "can_read", "admin:gaoshan"); err != nil {
  fmt.Println(err)
} else {
  fmt.Println("The permission for gaoshan to read the 'doc1' file is", permission)
}
if permission, err := client.GetCheck(ctx, "doc:doc1", "can_write", "admin:gaoshan"); err != nil {
  fmt.Println(err)
} else {
  fmt.Println("The permission for gaoshan to write the 'doc1' file is", permission)
}
if permission, err := client.GetCheck(ctx, "doc:doc1", "can_read", "user:normal_user"); err != nil {
  fmt.Println(err)
} else {
  fmt.Println("The permission for normal_user to read the 'doc1' file is", permission)
}
if permission, err := client.GetCheck(ctx, "doc:doc1", "can_write", "user:normal_user"); err != nil {
  fmt.Println(err)
} else {
  fmt.Println("The permission for normal_user to write the 'doc1' file is", permission)
}
```

### I'm still a beginner, so please feel free to point out any mistakes. 🙏🏻 (You can create an issue for that.)
