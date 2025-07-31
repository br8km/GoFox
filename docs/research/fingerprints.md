# Fingerpints

---

## References

- [BrowserLeaks](https://browserleaks.com/)

## Details

BrowserLeaks is a suite of tools that offers a range of tests to evaluate the security and privacy of your web browser. These tests focus on identifying ways in which websites may leak your real IP address, collect information about your device, and perform a browser fingerprinting.

By understanding these risks, you can take appropriate steps to protect your online privacy and minimize your exposure to potential threats
IP Address

The main tools for checking IP address privacy showcase server-side abilities to uncover a user's identity. These include displaying your IP address, reverse IP lookup, and HTTP request headers, your country, state, city, ISP/ASN, and local time. The tool also includes features such as IP address whois lookup, TCP/IP OS fingerprinting, WebRTC, DNS, and IPv6 leak tests.
JavaScript

By utilizing the basic functionality of JavaScript and modern Web APIs, it's possible to extract a wealth of data about the user's system. This includes information such as User-Agent, screen resolution, system language, local time, CPU architecture, the number of logical cores, the battery status, network information, installed plugins, and more.
WebRTC Leak Test

The WebRTC Leak Test is a critical tool for anyone using a VPN, as it leverages the WebRTC API to communicate with a STUN server and potentially reveal the user's real local and public IP addresses, even when using a VPN, proxy server, or behind a NAT. This tool can help verify whether a real public IP is being leaked.
Canvas Fingerprinting

A tracking method known as Canvas Fingerprinting uses HTML5 Canvas code to generate a unique identifier for each individual user. The method is based on the fact that the unique pixels generated through Canvas code can vary depending on the system and browser used, making it possible to identify users.
WebGL Report

The WebGL Report is a diagnostic tool to analyze your browser's WebGL support and create a unique WebGL Fingerprint that can potentially identify your web browser. This tool exposes information about your graphics card and other WebGL and GPU capabilities, which can be used to differentiate your browser from others.
Font Fingerprinting

Font fingerprinting is a technique used to track online activity by analyzing the unique characteristics of a user's system fonts. By measuring the dimensions of text or individual Unicode glyphs, enumerating fonts and finding rendering differences, this method can create a unique fingerprint that is difficult to spoof or alter.
Geolocation API

The Geolocation API allows websites to retrieve geographical location information from the user's device, and this HTML5 Geolocation API testing tool provides a detailed analysis of your geolocation and browser permissions.
Features Detection

The Web Browser's Features Detection tool provides a detailed list of HTML5 feature detectors, allowing you to determine which features your web browser supports or lacks, and how modifying them may impact your browser's digital footprint.
SSL/TLS Client Test

Check your browser's supported TLS protocols, cipher suites, TLS extensions, and key exchange groups. Identify weak or insecure options, generate a JA3 TLS fingerprint, and test how the browser handles insecure mixed content.
Content Filters

This page provides detectors to identify the usage of content filters that manipulate the connection and content between the browser and the visited web page. Examples of such filters include Tor Browser and AdBlockers.

More Tools

Here is a list of new, experimental, controversial, broken, and deprecated tools:

    HTTP/2 Fingerprinting – reading HTTP/2 frames and creating an imprint in Akamai format
    Chrome Extensions Detection – web accessible resources scanner
    WebGPU Browser Report – the successor to WebGL
    Client Hints Test – HTTP and JavaScript Client Hints test page
    DNS Leak Test – standalone page for DNS Leak Test
    CSS Media Queries – brute-forcing Media Queries in pure CSS
    ClientRects Fingerprinting – measuring the size and position of rendered HTML elements
    Do Not Track – checking whether the DNT or GPC is enabled in your browser
    Flash Player – showing device info using Flash Player
    Silverlight – showing device info using Silverlight plugin
    Java Applet – showing device info using Java applet
    Firefox Resources Reader – reading Resource URLs in pre-Quantum Firefox
