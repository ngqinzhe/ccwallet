# Wallet App
Simple wallet app to handle deposit, withdraw, transfer, get_balance, get_transactions.

### Design Decisions

#### 1. Data Storage
User & wallet information is stored in PostgreSql database. To simplify things, we don't have information about the user, hence a two tables are sufficient in this case: Wallet & Transactions. Wallet will store the respective wallet of the user and Transaction will store the historical transactions made by the user.

Assumptions: 

1. There is only 1 single currency
2. User has only 1 wallet

```
CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL UNIQUE,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    type INT
    information TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES wallets(user_id) ON DELETE CASCADE
);
```

Deposit / Withdraw / Transfer functions require atomicity, hence row level locking will be implemented to ensure that every wallet balance update is atomic. Upon successfully executing these functions, they will each create a new record in the `transaction` table.

#### 2. Data retrieval
Reads should ideally be as fast and as performant as possible. Records should be stored in the redis cache most of the time for fast retrieval. 

1. Initial simple idea was to write a write-through cache, but it is not as straight forward since we need to rollback cache updates on error. 

2. Best solution is provide a event-streaming service where the redis cache can consume the binlog of the db updates asynchronously, and users can just fetch data from the cache.

Of which both I do not consider simple tasks and are rather time consuming, so I will not implement here.

#### 3. Repository Structure
The repository structure is segregated into its own components
```bash
ccwallet/
│── main.go            # main server entry point
│── internal/          # Private application code
│    │── api/          # API layer
│        │── handler/  # Handles HTTP requests/responses validation
│        │── router/   # Defines API endpoints and methods
│        │── service/  # Business logic
│    │── cache/        # Handles caching logic
│    │── db/           # Database interactions
│    │── mocks/        # Mock interfaces for unit testing
│    │── model/        # Database models
│    │── util          # Utility functions
│── config/            # Configuration files to set up postgres / redis
│── go.mod             # Go module file
│── go.sum             # Go sum file
│── README.md          # Documentation
```

### 4. Things to improve

1. Caching logic. Not implemented due to time constraint. For wallet balances we can try to stream binlog events to redis to get the latest balances for users, since we expect this function to have high read QPS. As for the transaction history, if we expect a large number of transactions, we can consider caching weekly transactions and have a cap of maybe 1000? Beyond that we can implement LRU eviction logic in our Redis. If we expect to scale to millions and our apps runs for years, we can consider only storing latest 1 year data in hot databases and older data into cold databases to reduce memory.

2. Memory storage. Scaling to millions require large storage, and large storage can lead to performance issues. We can implement sharding and shard by user_id, where each query can efficient query and retrieve using the logical shard, improving performance. 
Database indexes can also be created based on the user_id in `wallet` and `created_at` in transactions, where transactions may need to sorted from latest to oldest.

3. Unit testing. Unit testing has to provide 100% coverage, however due to constraints in time, I will cover happycases and obvious fail cases. There are obviously many edge cases to cover.

4. Setting up to run this repo effectively. Again, due to time constraints, I will not be setting up scripts to allow the user to run postgresql db and redis on their docker image. This can be something to be improved on but I feel is not the important point of this test.

Time Taken: ~5 hours