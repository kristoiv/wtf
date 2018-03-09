An attempt at structuring a piece of go software according to the "WTF Dial" blog series / collaborative project.

- Original blog post: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- Blog implementation example series: https://medium.com/wtf-dial
- Blog implementation example github repo: https://github.com/benbjohnson/wtf

# Project conclusion

This project was never intended to be optimal or complete. It started with me reading the post linked above about structuring of golang projects around a root pkg containing only the domain model (and potentially features that only depend on the domain model, eg. cache.go). The upsides are many, including extreme decoupling which simplifies testing of whole modules. It also seems to result in a very neat code base. There is not complete test coverage, not a point.

I also got to have some fun, testing the gocui project to make a "Console User Interface" (cui), which I hadn't really done before. We are talking multi-view, keybindings, etc. Fun!

The main problem I found that I'm unsure how to deal with "correctly" is use of protobuf / grpc in multiple sub-pkgs (sorted by dependency) since protobuf has a global namespace. Thus having protobuf generated code in an "internal" sub-sub-pkg of two different sub-pkgs which both have proto-files with an "message Item" results in a crash ("duplicate models internal.Item"). I first tried to pull all proto-generated code out in a seperate pkg but that goes against the idea of keeping my dependencies self-contained (arguably the proto-generated code is a part of the domain model, but it contains implementation details for example for grpc, which I wanted to contain in the grpc sub-pkg). I resolved this by renaming "Item" in one of the proto files. Having this duplicate is also a result of having the "Item" type duplicated in proto-files, however that is more "okay" since the database might store more info than say the grpc transfers to the client.
