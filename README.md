# go-oasapi

Based on **gph03n1x** API Documentation. I only implemented the endpoints that seemed to work properly.

```
(OK) url: "/api/?act=webGetLinesWithMLInfo",
(OK) url: '/api/?act=webGetLangs',
(OK) url: "/api/?act=webGetLines",
(OK) url: '/api/?act=getLineName&p1='+lineCode,
(OK) url: '/api/?act=webGetRoutes&p1='+lineCode,
(OK) url: '/api/?act=getRoutesForLine&p1='+linecode,
(OK) url: '/api/?act=webRouteDetails&p1='+routeCode,
(OK) url: '/api/?act=getRouteName&p1='+routeCode,
(OK) url: '/api/?act=webRoutesForStop&p1='+stopCode,
(OK) url: '/api/?act=webGetRoutesDetailsAndStops&p1='+routeCode, JSON
(OK) url: '/api/?act=webGetStops&p1='+routeCode,
(OK) url: '/api/?act=getStopArrivals&p1='+stopCode,
(OK) url: '/api/?act=getStopNameAndXY&p1='+stopCode,
(OK) url: '/api/?act=getClosestStops&p1='+x+'&p2='+y,
(OK) url: '/api/?act=getBusLocation&p1='+routeCode,
(OK) url: '/api/?act=getScheduleDaysMasterline&p1='+lineCode,
(OK) url: '/api/?act=getMLName&p1='+mlCode,
     url: '/api/?act=getDailySchedule&line_code='+lineCode, JSON

----------Not implemented: Return Null Responses----------
url: "/api/?act=getPlacesTypes",
url: '/api/?act=getPlaces&p1='+catCode,
url: '/api/?act=getLinesAndRoutesForMl&p1='+mlCode,
url: '/api/?act=getLinesAndRoutesForMlandLCode&p1='+mlCode+'&p2='+lid,
url: '/api/?act=getSchedLines&p1='+mlCode+'&p2='+sdcCode+'&p3='+lineCode,
url: '/api/?act=getEsavlDirections&p1=4&p2=3&p3=300&p4='+from[1]+'&p5='+from[0]+'&p6='+to[1]+'&p7='+to[0],
```

## References

[Unofficial Documentation](https://oasa-telematics-api.readthedocs.io/en/latest/)

[OASA-Telematics](http://telematics.oasa.gr/#main)

[OASA Telematics Script](http://telematics.oasa.gr/js/script.js)

**TODO**

1. Better Tests
2. Example Application
