# Celestia-Ram-Usage-Analyze
Analyzing RAM Usage on a Server with Celestia Light Node: A Case Study



## Introduction

In today's modern computing landscape, understanding how your server's resources are being utilized is vital to maintaining optimal performance. To that end, we set out to analyze RAM usage on a server with a Celestia light node running, comparing it to an empty server. In this post, we'll dive into the details of our case study and discuss the results obtained using a custom Golang code for testing.

## System Configuration

Our server configuration for this case study was as follows:

16 vCPUs (virtual central processing units)
32 GB of RAM (random access memory)
300 GB SSD (solid-state drive)
Testing Methodology

In order to assess the RAM usage of our server, we conducted two separate tests using a Golang script. The first test was run on an empty server, and the second test was run with a Celestia light node running. Both tests were conducted over a six-hour period.

A Celestia light node is a lightweight node within the Celestia ecosystem, requiring less computational resources compared to a full node. It is designed to work on single-core CPU systems with limited RAM requirements.

## Test Results

### Empty Server Test:
The empty server test showed that the RAM usage percentage hovered around 0.7%. This baseline value indicates the server's RAM consumption when no applications or services are running.

![ram_usage_chart empty](https://user-images.githubusercontent.com/53251494/234844792-9042ae41-9596-4ec8-8aee-0698db9de85a.png)

### Server with Celestia Light Node Test:
When the Celestia light node was running, the server's RAM usage percentage increased to an average of 5%. This increase can be attributed to the light node's operations, which require some RAM to function efficiently.

![ram_usage_chart_with_celestia](https://user-images.githubusercontent.com/53251494/234844889-49a8764c-b81f-4326-b88c-579dffd78f99.png)

## Conclusion

Our case study demonstrates that running a Celestia light node on a server with the given configuration has a noticeable impact on RAM usage. The increase from 0.7% to 5% RAM usage indicates that the Celestia light node is utilizing server resources effectively, while still leaving a significant amount of RAM available for other applications.

For those considering deploying a Celestia light node, our findings suggest that the node can operate efficiently on a server with modest RAM requirements. Moreover, by using a Golang script, it is possible to easily monitor and analyze RAM usage for optimal performance and resource allocation.

# NOTE:
You can see the output of the experiment in graphics and in the png file that are named as **ram_usage_chart_with_celestia** and **ram_usage_chart empty**. 
