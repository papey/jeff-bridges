# Jeff-Bridges

Multiply bridges, for fun and profit

# Build

Get deps using dep

    dep ensure

Build

    cd cmd
    go build jb.go

# Usage

Jeff-Bridges can add or delete bridge

## Add

For example :

    jb add test 127.17.0.1/16

Will add a new bridge name test with next available ip. In this case, ip will be 127.18.0.1 if not used by another iterface.

## Delete

For example :

    jb delete test

Will delete bridge named test

# Licence

See LICENCE file

# Misc

Yes, this project is build using the gitflow workflow, yes it's overkill for a small project like this but I mainly do this just for fun.
