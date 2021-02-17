This is a sandbox project created as a means of exploring and learning about the consumer first, contract testing tool "[pact](https://docs.pact.io/)".

The project is comprised of 2 would-be microservices:

### videos service

This is a microservice which can be used to store and retrieve videos.

#### example request: store a video

```
curl --location --request POST 'localhost:8080/video' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "123",
    "name": "The Matrix",
    "decription": "The most overrated film of all time"
}'
```

#### example request: retrieve a video:

```
curl --location --request GET 'localhost:8080/videos?id=123'
```

When retrieving a video, the video service suppliments the data with a list of user submitted ratings. For this, it calls upon the second microservice, the "ratings service".



### ratings service

This is a microservice which can be used to submit a rating for a video and also retrieve all of the submitted ratings for a given video.


### architecture

The video service is a client of the ratings service. The contract between the two will be established and protected with a pact, created in the consumer (the video service) and then used in the producer (the ratings service).

In a real world scenario, these 2 microservices may be built by separate teams, and estabvlishing and utilising this contract (using pact in this case) would allow team autonomy with safety.