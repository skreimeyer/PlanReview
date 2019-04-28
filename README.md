# Plan Review

Plan Review is a tool meant to assist with writing staff comments on building permit applications and subdivision proposals. Reviews tend to be highly formulaic, ie, staff looks up the property on several GIS web applications and then tailors boilerplate responses to reflect that information. Only small portions of the review, such as the accuracy of supporting hydraulic studies, actually require human intelligence. (Even the necessity of that is not a certainty)

Plan Review is also a tool very specific to the City of Little Rock. If you would benefit from automation of a similar rule set (flood ordinances, for example, tend to be fairly uniform), then the source code may have teachable information. Documentation should be adequate to understand the basic functioning of the program. Unfortunately, generalizing this application sufficiently to apply turn-key to different municipalities is not within the scope of the project.

If you have specific questions or want help, then you should contact me at the City of Little Rock Public Works during normal business hours. I would be happy to assist.

If your interest is interacting with Esri REST APIs, which tend to vary greatly, then you might find this tool very useful [json2gostruct](https://github.com/skreimeyer/json2gostruct), which I wrote to generate the structs I used in this application.

## SITREP

Essential helper scripts (see json2gostruct) are complete and have made handling the many different map services possible. Core functionality exists in pkg/planreview. Fetching useful data from the REST APIs is at MVP stage

Currently, the implementation is very weak. Error handling and edge cases are basically not considered. The web front end does not exist yet.

## TODO
[X] - Implement an API to fetch useful information from GIS servers
[ ] - Write the logic to template public works review comments
[ ] - Provide a simple web interface to interact with the API and view templated comments
[ ] - Make the code less bad
[ ] - Deploy to Heroku
