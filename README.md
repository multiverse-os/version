[<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">](https://github.com/multiverse-os)

## Multiverse: Semantic Versioning & Checksum Validation Library 
**URL** [multiverse-os.org](https://multiverse-os.org)

#### Introduction
A semantic versioning library that combines the features built across 5+
impelmentations built individually in separate libraries. 

The library is often just re-implemented to avoid dependencies but the 
`portalgun` design requires checksum validation; the validation utilizes a  
merkle tree for each version in a development stream, and the same tree holds 
all known development streams to provide greater security and allows for
validation of any version using the current root. 


#### Development 
Development is currently working on merging in a more generic and broadly used
implementation that will also support the comprasion method needed for Gemfile's
allowing us to implement a RubyGem implementation in Go.
