#!/bin/bash

docker run -p 8080:8080 -v $(pwd)/stub/mappings:/home/wiremock/mappings -v $(pwd)/stub/__files:/home/wiremock/__files wiremock/wiremock --verbose
