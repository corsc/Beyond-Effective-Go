# Skip

This package intends to make it easy to skip tests based on environment variables.

Common use-cases include:
* Enabling/Disabling tests that require a database
* Disabling long-running tests (like UAT or capacity tests) during the dev cycle

These use-cases can be achieved using build flags but I have found that that build flags are cumbersome and sometimes lead to broken code 
(as the compiler does not check the code unless that tag is used)
