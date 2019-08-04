FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD aromahousekpg aromahousekpg
ADD static static
ADD views views
EXPOSE 8080
EXPOSE 9090
CMD ["/aromahousekpg"]
