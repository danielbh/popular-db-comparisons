## DynamoDB

### Demo

Terminal 1

* `docker pull amazon/dynamodb-local`
* `docker run -p 8000:8000 amazon/dynamodb-local`

Terminal 2 

* `go run dynamodb.go`

### Features

* Amazing free-tier pricing. Great for quick MVP.
* Provisioned/On-Demand throughput with burst capacity
* Document/KV store
* auto-scale
* automatic data replication over three AZ in single region
* Integrated with AWS EMR, Data Pipeline, Kinesis
* easy access control
* DynamoDB stream, before and after hooks time-ordered within 24 hr.
* Built-in TTL
* Caching for improved latency (DAX)
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
  * Items: like rows in RD. DynamoDB requires a Primary Key. PK supports Hash or Hash + Range.
* Two Index Types
  * Local Secondary Index (LSI): range key is mandatory. Limit item collection size to 10 GB.
  * Global Secondary Index (GSI): hash key or a hash+range key. GSIs span multiple partitions and are placed in separate tables. DynamoDB supports up to five GSIs. Will be used for partitioning. Offer only eventual consistency.
* Amazon DynamoDB JavaScript Web Shell
* Adaptive Capacity for Hot Partition Keys [more info](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html#bp-partition-key-partitions-adaptive)
* Table encryption with custom keys

### Cost Analysis
 * Can be provisioned or on-demand billing [ref](https://dynobase.dev/dynamodb-pricing-calculator/)
 * Billed on [ref](https://dynobase.dev/dynamodb-pricing-calculator/)
    * Amount of data stored
    * Amount of data written and read
    * Data transfer
    * Backups and restore operations
    * DynamoDB Streams
    * Replicated write request units when using Global Tables
    * Consistency of Reads
* If possible it's helpful to manage cost with good balance between provisioned and on-demand capacity. Ex: spikey predictable workloads.
### Deployment/Operations

* Fully managed, but need to track costs carefully.
* Partitioning is transparent, but if using GSI need to be aware. Partitioning depends on table size and throughput

<p align="center">
  <img src="https://cloudacademy.com/wp-content/uploads/2015/11/amazon-dynamodb-cacl.png">
  <br/>
  <i><a href=https://cloudacademy.com/wp-content/uploads/2015/11/amazon-dynamodb-cacl.png>Image curteosy CloudAcademy, amazon-dynamodb-calc</a></i>
</p>

### Use-Cases

* Duolingo: Supports high throughput with ease
* GE Healthcare: Supports enterprise workloads on cloud
* Docomo: marketing data (user events, profiles, clicks). Ad targeting, attribution, real-time bidding
* Hess Coporation: Used to separate buyer data

[ref](https://www.geeksforgeeks.org/dynamodb-understand-the-benefits-with-real-life-use-cases/)

### Drawbacks

* Althought adaptive capacity can alleviate hot partition key issues, be sure to have random partion key values
* Item size limited to 400 KB which includes: item size in table, size of LSI corresponding to that item.
* Partition max size: 10 GB
* Partition max 3000 RCU and max 1000 WCU
* Strongly consistent reads more expensive, and will predictable prefer CP over AP (CAP Theorem) which can impact if app is multi-region.
* Transactions with following restrictions [see dynamodb limits page](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
* Don't pretend this is relational database, you will get burned
* DynamoDB tables with DynamoDB Streams enabled has limits. 40k wcu or 10k rc depending on region
### References:

* [AWS DynamoDB Limits](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
* [AWS best practices in partition key design](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html#bp-partition-key-partitions-adaptive)
* [AWS DynamoDB Best Practices for Large Items](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-use-s3-too.html)
* [11 things you know before starting with dynamodb (Yugabyte competitor)](https://blog.yugabyte.com/11-things-you-wish-you-knew-before-starting-with-dynamodb)
* [DynamoDB vs MongoDB vs Cassandra for Fast Growing Geo-Distributed Apps (Yugabyte competitor)](https://blog.yugabyte.com/dynamodb-vs-mongodb-vs-cassandra-for-fast-growing-geo-distributed-apps)
* [DynamoDB pricing and tips to reduce cost](https://dynobase.dev/dynamodb-pricing-calculator/)
* [How DynamoDB’s Pricing Works, Gets Expensive Quickly and the Best Alternatives
(yugabyte)](https://blog.yugabyte.com/dynamodb-pricing-calculator-expensive-vs-alternatives/)
* [3 cost-cutting tips for Amazon DynamoDB
](https://rockset.com/blog/3-cost-cutting-tips-for-amazon-dynamodb/)
* [The million dollar engineering problem -- appears to be outdated due to adaptive capacity feature a year later (Segment)](https://segment.com/blog/the-million-dollar-eng-problem/)
* [dynamoDB (trustradius)](https://www.trustradius.com/products/amazon-dynamodb/reviews?qs=pros-and-cons)
* [2020 year in review AWS DynamoDB](https://aws.amazon.com/blogs/database/2020-the-year-in-review-for-amazon-dynamodb/)
* [10 thinks to know about DynamoDB(Cloud Academy)](https://cloudacademy.com/blog/amazon-dynamodb-ten-things/)
* [DynamoDB: Understand The Benefits With Real Life Use Cases(GeeksForGeeks)](https://www.geeksforgeeks.org/dynamodb-understand-the-benefits-with-real-life-use-cases/)
* [db-engines: dynamodb](https://db-engines.com/en/system/Amazon+DynamoDB)
* [reducing dynamoDB cost (Nike)](https://medium.com/nikeengineering/reducing-dynamodb-costs-in-aws-5047cbf726c9)
