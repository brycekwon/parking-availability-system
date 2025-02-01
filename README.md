# Parking Availability System

The Parking Availability System (PAS) is a proof-of-concept solution for real-time monitoring of
parking space availability, designed to be low-cost, low-maintenance, and scalable. The system
employs a distributed network of time-of-flight sensors and magnetic field sensors to detect the
presence of vehicles on each parking space, transmitting that data wirelessly to an interactive
website for real-time monitoring. This design aims to enhance parking efficiency by providing
accurate, up-to-date information on parking space occupancy, reducing both time and stress.

# Setup and Deployment

> [!Note]
> This is a monorepository containing the source code of multiple independent components used for
> assembling the system. Please consult the specific documentation of each service for proper setup
> and deployment instructions.

- [LoRa End-Devices]()
- [LoRaWAN Gateway]()
- [LoRaWAN Server]()
- [Web Application]()

# System Architecture

This system can be broken down into 4 distinct components: 1) The LoRa End-Devices, which are
responsible for recording live-data from each parking spot; 2) The LoRaWAN Gateway, which is
responsible for collecting data from the end-devices and forwarding them to the network server; 3)
The LoRaWAN server, which is responsible for managing gateways and devices, as well as forwarding
processed information on to the user-facing web application; 4) An interactive web application for
users to quickly view for up-to-date information.

## LoRa End-Devices

## LoRaWAN Gateway

## LoRaWAN Server

## Web Application

# References

- https://github.com/chirpstack
- https://github.com/TheThingsNetwork
- https://www.chirpstack.io/docs/index.html
- https://github.com/brocaar/chirpstack-simulator
- https://github.com/lorabasics/basicstation
- https://www.chirpstack.io/docs/gateway-configuration/index.html
- https://www.thethingsindustries.com/docs/gateways/models/raspberry-pi/
- https://github.com/TheThingsNetwork/lorawan-devices
- https://github.com/chirpstack/lorawan-device-profiles