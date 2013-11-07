library(ggplot2)
library(ggmap)
args <- commandArgs(TRUE)
yelp <- read.delim(args[1], na.strings="N/A")
az <- get_map("Pheonix, Arizona, United States", maptype="roadmap", zoom = 10)
ggmap(az, extent='device', legend='none') + 
      geom_point(aes(x=lon, y=lat), data=subset(yelp, cat == "Food"), na.rm=TRUE) 