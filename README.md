# Philatelist

Philatelist is a RESTful server, that allows searching images in web by location or Google placeid.

## the task
> for the next project, i need something quite interesting =) i need a REST API that i can provide an address or google placeid to, and then get back a list of images in that area 


> to start with, i would like to see if we can find some free images on the internet to fill this, perhaps using google? or picasa? or something like that
then, later, we can perhaps add a street view option

> but for now lets stick to just images from the region
the real purpose of this project is for a deal comparison website, and the images will be used as a background in the site

> so when the customer enters their address into the site, the site will have images from the customer's area on the site
to make the site feel more local and more customised for the user

> the API must be very easy to use.. for example:

`GET /api/v1/images/address-text?address=20+kelmarna+avenue,+ponsonby`

> or 


`GET /api/v1/images/google-place-id?placeid=xxxxxxxx`

> and the response should be something like:

```
[
{ url: "http://picasa.com/aaaa1" },
{ url: "http://picasa.com/aaaa2" },
{ url: "http://picasa.com/aaaa3" }
]
```
