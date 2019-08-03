FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD aromahousekpg aromahousekpg
EXPOSE 8080
EXPOSE 9090
CMD ["/aromahousekpg"]
