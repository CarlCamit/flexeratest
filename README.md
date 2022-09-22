# Flexera

To use the app please install Docker on your local machine and run the command below to build the image

`docker build . -t flexeratest`

Once the image is built you can run the app by running

`docker run flexeratest -file <filepath/to/csv>`

## Concerns

The approach used is not as optimal timewise as it does need to loop through the csv and then once again through a map, it is however architected in a way that is easy to change, add validations or increase scope
