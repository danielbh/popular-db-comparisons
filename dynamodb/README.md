## DynamoDB

### Demo

Terminal 1

* `docker pull amazon/dynamodb-local`
* `docker run -p 8000:8000 amazon/dynamodb-local`

Terminal 2 

* `go run dynamodb.go`

### Features

* Provisioned throughput with burst capacity
* Document/KV store
* auto-scale
* automatic data replication over three AZ in single region
* Integrated with AWS EMR, Data Pipeline, Kinesis
* easy access control
* DynamoDB stream, before and after hooks time-ordered within 24 hr.
* Built-in TTL
* Good for un-structured data
* Easy back-ups
* This is AWS hosted Cassandra
* ACID transactions with restrictions
* Eventual Consistency/Strong Consistency - User Controlled
* Secondary Indexes
* Enterprise ReadyL SLA, monitoring, private VPN
* Supports the following data-types
  * Scalar: Number, String, Binary, Boolean, Null
  * Multi-valued: String, Number Set, Binary Set
  * Document: List and Map
* Data Model Units: Tables > Items > Attributes
  * Attributes: kv pairs
  * Tables: like RD w/o fixed schemas
  * Items: like rows in RD. DynamoDB requires a Primary Key. PK supports Hash or Range.
* Two Index Types: Local Secondary Index (LSI) and Global Secondary Index (GSI).
  * LSI: 
  * GSI: 

### Cost Analysis

* Cost is based on provisioned throughput
* Hot partitions lead to massive cost

### Deployment/Operations

* Fully managed, but need to track costs carefully.
### Use-Cases

* Duolingo: Supports high throughput with ease
* GE Healthcare: Supports enterprise workloads on cloud
* Docomo: marketing data (user events, profiles, clicks). Ad targeting, attribution, real-time bidding
* Hess Coporation: Used to separate buyer data

### Drawbacks

* Need to be careful about partition keys, hot partitions can lead to massive cost.
### References:

* [AWS DynamoDB Limits](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
* [AWS DynamoDB Best Practices for Large Items](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-use-s3-too.html)
* [11 things you know before starting with dynamodb (Yugabyte competitor)](https://blog.yugabyte.com/11-things-you-wish-you-knew-before-starting-with-dynamodb)
* [DynamoDB vs MongoDB vs Cassandra for Fast Growing Geo-Distributed Apps (Yugabyte competitor)](https://blog.yugabyte.com/dynamodb-vs-mongodb-vs-cassandra-for-fast-growing-geo-distributed-apps)
* [DynamoDB pricing](https://dynobase.dev/dynamodb-pricing-calculator/)
* [How DynamoDB’s Pricing Works, Gets Expensive Quickly and the Best Alternatives
(yugabyte)](https://blog.yugabyte.com/dynamodb-pricing-calculator-expensive-vs-alternatives/)
* [3 cost-cutting tips for Amazon DynamoDB
](https://rockset.com/blog/3-cost-cutting-tips-for-amazon-dynamodb/)
* [The million dollar engineering problem (Segment)](https://segment.com/blog/the-million-dollar-eng-problem/)
* [dynamoDB (trustradius)](https://www.trustradius.com/products/amazon-dynamodb/reviews?qs=pros-and-cons)
* [2020 year in review AWS DynamoDB](https://aws.amazon.com/blogs/database/2020-the-year-in-review-for-amazon-dynamodb/)
* [10 thinks to know about DynamoDB(Cloud Academy)](https://cloudacademy.com/blog/amazon-dynamodb-ten-things/)
* [DynamoDB: Understand The Benefits With Real Life Use Cases(GeeksForGeeks)](https://www.geeksforgeeks.org/dynamodb-understand-the-benefits-with-real-life-use-cases/)
* [db-engines: dynamodb](https://db-engines.com/en/system/Amazon+DynamoDB)