# Free 5GC "aper" project

Encoding and decoding of ASN.1 structs to APER encoding format

This is copied from `v1.0.0` of https://github.com/free5gc/aper

Changes:

* logrus has been replaced by own logging
* errors are generated through our own package
* can handle structs that have one or more private fields (lowercase)